package database

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/ubclaunchpad/pinpoint/protobuf/models"
)

const (
	clubTablePrefix = "ClubData-"
	peidPrefix      = "PEID-"
	periodPrefix    = "Period-"
	applicantPrefix = "Applicant-"
	tagPrefix       = "Tag-"
)

type eventItem struct {
	PeidPK      string          `json:"pk"`
	PeidSK      string          `json:"sk"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Fields      []*models.Field `json:"fields"`
}

type applicantItem struct {
	Period string   `json:"pk"`
	Email  string   `json:"sk"`
	Name   string   `json:"name"`
	Tags   []string `json:"tags"`
}

type applicationItem struct {
	Peid    string                        `json:"pk"`
	Email   string                        `json:"sk"`
	Name    string                        `json:"name"`
	Entires map[string]*models.FieldEntry `json:"entries"`
}

type tagItem struct {
	Period  string `json:"pk"`
	TagName string `json:"sk"`
}

func newDBEvent(event *models.Event) *eventItem {
	peid := prefixPeriodEventID(event.Period, event.EventID)
	return &eventItem{
		PeidPK:      peid,
		PeidSK:      peid,
		Name:        event.Name,
		Description: event.Description,
		Fields:      event.Fields,
	}
}

func getEvent(item *eventItem) *models.Event {
	p, e := getPeriodAndEventID(item.PeidPK)
	return &models.Event{
		Period:      p,
		EventID:     e,
		Name:        item.Name,
		Description: item.Description,
		Fields:      item.Fields,
	}
}

func newDBApplicant(applicant *models.Applicant) *applicantItem {
	p := prefixPeriodID(applicant.Period)
	return &applicantItem{
		Period: p,
		Email:  prefixApplicantID(applicant.Email),
		Name:   applicant.Name,
		Tags:   applicant.Tags,
	}
}

func getApplicant(item *applicantItem) *models.Applicant {
	return &models.Applicant{
		Period: removePrefix(item.Period),
		Email:  removePrefix(item.Email),
		Name:   item.Name,
		Tags:   item.Tags,
	}
}

func newDBTag(tag *models.Tag) *tagItem {
	return &tagItem{
		Period:  prefixPeriodID(tag.Period),
		TagName: prefixTag(tag.TagName),
	}
}

func getTag(item *tagItem) *models.Tag {
	return &models.Tag{
		Period:  removePrefix(item.Period),
		TagName: removePrefix(item.TagName),
	}
}

func newDBApplication(application *models.Application) *applicationItem {
	peid := prefixPeriodEventID(application.Period, application.EventID)
	return &applicationItem{
		Peid:    peid,
		Email:   prefixApplicantID(application.Email),
		Name:    application.Name,
		Entires: application.Entires,
	}
}

func getApplication(item *applicationItem) *models.Application {
	p, e := getPeriodAndEventID(item.Peid)
	return &models.Application{
		Period:  p,
		EventID: e,
		Email:   removePrefix(item.Email),
		Name:    item.Name,
	}
}

func getClubTable(clubID string) *string {
	return aws.String(clubTablePrefix + clubID)
}

func prefixPeriodEventID(period string, eventID string) string {
	return peidPrefix + period + "-" + eventID
}

func prefixPeriodID(period string) string {
	return periodPrefix + period
}

func prefixApplicantID(email string) string {
	return applicantPrefix + email
}

func prefixTag(tag string) string {
	return tagPrefix + tag
}

func getPeriodAndEventID(peid string) (string, string) {
	str := removePrefix(peid)
	s := strings.SplitN(str, "-", 2)
	return s[0], s[1]
}

// AddNewEvent creates a new event in the club table
func (db *Database) AddNewEvent(clubID string, event *models.Event) error {
	e := newDBEvent(event)
	item, err := dynamodbattribute.MarshalMap(e)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: getClubTable(clubID),
	}
	if _, err := db.c.PutItem(input); err != nil {
		return err
	}
	return nil
}

// GetEvent returns an event from the database
func (db *Database) GetEvent(clubID string, period string, eventID string) (*models.Event, error) {
	peid := aws.String(prefixPeriodEventID(period, eventID))
	result, err := db.c.GetItem(&dynamodb.GetItemInput{
		TableName: getClubTable(clubID),
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {S: peid},
			"sk": {S: peid},
		},
	})
	if err != nil {
		return nil, err
	}

	var item eventItem
	if err := dynamodbattribute.UnmarshalMap(result.Item, &item); err != nil {
		return nil, err
	}
	return getEvent(&item), nil
}

// GetEvents returns all the events of an application period
func (db *Database) GetEvents(clubID string, period string) ([]*models.Event, error) {
	input := &dynamodb.QueryInput{
		TableName: getClubTable(clubID),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":id": {
				S: aws.String(peidPrefix + period + "-"),
			},
		},
		KeyConditionExpression: aws.String("begins_with(pk, :id) AND begins_with(sk, :id)"),
	}

	result, err := db.c.Query(input)
	if err != nil {
		return nil, err
	}
	items := make([]*eventItem, *result.Count)
	if err := dynamodbattribute.UnmarshalListOfMaps(result.Items, &items); err != nil {
		return nil, err
	}
	events := make([]*models.Event, *result.Count)
	for _, item := range items {
		events = append(events, getEvent(item))
	}
	return events, nil
}

// DeleteEvent deletes an event and all of its applications
func (db *Database) DeleteEvent(clubID string, event *models.Event) error {
	input := &dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":id": {
				S: aws.String(prefixPeriodEventID(event.Period, event.EventID)),
			},
		},
		KeyConditionExpression: aws.String("pk = :id"),
		ProjectionExpression:   aws.String("pk, sk"),
		TableName:              getClubTable(clubID),
	}
	result, err := db.c.Query(input)
	if err != nil {
		return err
	}

	var batch []*dynamodb.WriteRequest
	for _, item := range result.Items {
		batch = append(batch, &dynamodb.WriteRequest{
			DeleteRequest: &dynamodb.DeleteRequest{
				Key: item,
			},
		})
	}
	batchInput := &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]*dynamodb.WriteRequest{
			*getClubTable(clubID): batch,
		},
	}
	if _, err = db.c.BatchWriteItem(batchInput); err != nil {
		return err
	}
	return nil
}

// AddNewApplicant creates a new applicant in the club table for an application period
func (db *Database) AddNewApplicant(clubID string, applicant *models.Applicant) error {
	a := newDBApplicant(applicant)
	item, err := dynamodbattribute.MarshalMap(a)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: getClubTable(clubID),
	}
	if _, err := db.c.PutItem(input); err != nil {
		return err
	}
	return nil
}

// GetApplicant returns an applicant for a application period
func (db *Database) GetApplicant(clubID string, period string, email string) (*models.Applicant, error) {
	result, err := db.c.GetItem(&dynamodb.GetItemInput{
		TableName: getClubTable(clubID),
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {S: aws.String(prefixPeriodID(period))},
			"sk": {S: aws.String(prefixApplicantID(email))},
		},
	})
	if err != nil {
		return nil, err
	}

	var item applicantItem
	if err := dynamodbattribute.UnmarshalMap(result.Item, &item); err != nil {
		return nil, err
	}
	return getApplicant(&item), nil
}

// GetApplicants returns all the applicants for an application period
func (db *Database) GetApplicants(clubID string, period string) ([]*models.Applicant, error) {
	input := &dynamodb.QueryInput{
		TableName: getClubTable(clubID),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":id": {
				S: aws.String(prefixPeriodID(period)),
			},
			":a": {
				S: aws.String(applicantPrefix),
			},
		},
		KeyConditionExpression: aws.String("pk = :id AND begins_with(sk, :a)"),
	}

	result, err := db.c.Query(input)
	if err != nil {
		return nil, err
	}
	items := make([]*applicantItem, *result.Count)
	if err := dynamodbattribute.UnmarshalListOfMaps(result.Items, &items); err != nil {
		return nil, err
	}
	applicants := make([]*models.Applicant, *result.Count)
	for _, item := range items {
		applicants = append(applicants, getApplicant(item))
	}
	return applicants, nil
}

// DeleteApplicant deletes an applicant from a application period and their event applications
func (db *Database) DeleteApplicant(clubID string, period string, email string) error {
	input := &dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":id": {
				S: aws.String(prefixApplicantID(email)),
			},
			"peid": {
				S: aws.String(peidPrefix + period),
			},
			"p": {
				S: aws.String(prefixPeriodID(period)),
			},
		},
		KeyConditionExpression: aws.String("sk = :id AND (begins_with(pk, :peid) OR begins_with(pk, :p))"),
		ProjectionExpression:   aws.String("pk, sk"),
		TableName:              getClubTable(clubID),
	}
	result, err := db.c.Query(input)
	if err != nil {
		return err
	}

	var batch []*dynamodb.WriteRequest
	for _, item := range result.Items {
		batch = append(batch, &dynamodb.WriteRequest{
			DeleteRequest: &dynamodb.DeleteRequest{
				Key: item,
			},
		})
	}
	batchInput := &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]*dynamodb.WriteRequest{
			*getClubTable(clubID): batch,
		},
	}
	if _, err = db.c.BatchWriteItem(batchInput); err != nil {
		return err
	}
	return nil
}

// AddNewApplication adds an application to the database
func (db *Database) AddNewApplication(clubID string, application *models.Application) error {
	a := newDBApplication(application)
	item, err := dynamodbattribute.MarshalMap(a)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: getClubTable(clubID),
	}
	if _, err := db.c.PutItem(input); err != nil {
		return err
	}
	return nil
}

// GetApplication returns the application for an event by applicant email
func (db *Database) GetApplication(clubID string, period string, eventID string, email string) (*models.Application, error) {
	result, err := db.c.GetItem(&dynamodb.GetItemInput{
		TableName: getClubTable(clubID),
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {S: aws.String(prefixPeriodEventID(period, eventID))},
			"sk": {S: aws.String(prefixApplicantID(email))},
		},
	})
	if err != nil {
		return nil, err
	}

	var item applicationItem
	if err := dynamodbattribute.UnmarshalMap(result.Item, &item); err != nil {
		return nil, err
	}
	return getApplication(&item), nil
}

// GetApplications returns all the applications for an event
func (db *Database) GetApplications(clubID string, period string, eventID string) ([]*models.Application, error) {
	input := &dynamodb.QueryInput{
		TableName: getClubTable(clubID),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":id": {
				S: aws.String(prefixPeriodEventID(period, eventID)),
			},
			":a": {
				S: aws.String(applicantPrefix),
			},
		},
		KeyConditionExpression: aws.String("pk = :id AND begins_with(sk, :a)"),
	}

	result, err := db.c.Query(input)
	if err != nil {
		return nil, err
	}
	items := make([]*applicationItem, *result.Count)
	if err := dynamodbattribute.UnmarshalListOfMaps(result.Items, &items); err != nil {
		return nil, err
	}
	applications := make([]*models.Application, *result.Count)
	for _, item := range items {
		applications = append(applications, getApplication(item))
	}
	return applications, nil
}

// DeleteApplication deletes the application for an event by applicant email
func (db *Database) DeleteApplication(clubID string, period string, eventID string, email string) error {
	if _, err := db.c.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: getClubTable(clubID),
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {S: aws.String(prefixPeriodEventID(period, eventID))},
			"sk": {S: aws.String(prefixApplicantID(email))},
		},
	}); err != nil {
		return err
	}
	return nil
}

// GetTags returns all the tags for an application period
func (db *Database) GetTags(clubID string, period string) ([]*models.Tag, error) {
	input := &dynamodb.QueryInput{
		TableName: getClubTable(clubID),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":id": {
				S: aws.String(prefixPeriodID(period)),
			},
			":t": {
				S: aws.String(tagPrefix),
			},
		},
		KeyConditionExpression: aws.String("pk = :id AND begins_with(sk, :t)"),
	}

	result, err := db.c.Query(input)
	if err != nil {
		return nil, err
	}
	items := make([]*tagItem, *result.Count)
	if err := dynamodbattribute.UnmarshalListOfMaps(result.Items, &items); err != nil {
		return nil, err
	}
	tags := make([]*models.Tag, *result.Count)
	for _, item := range items {
		tags = append(tags, getTag(item))
	}
	return tags, nil
}
