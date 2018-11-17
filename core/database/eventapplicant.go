package database

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/ubclaunchpad/pinpoint/core/model"
)

const (
	clubTablePrefix = "ClubData-"
	peidPrefix      = "PEID-"
	applicantPrefix = "Applicant-"
)

type dbEvent struct {
	PeriodEventIDPK string `json:"pk"`
	PeriodEventIDSK string `json:"sk"`
	Name            string `json:"name"`
	Description     string `json:"description"`
}

type dbApplicant struct {
	PeriodEventID string                 `json:"pk"`
	Email         string                 `json:"sk"`
	Name          string                 `json:"name"`
	Info          map[string]interface{} `json:"info"`
}

func newDBEvent(event *model.Event) *dbEvent {
	peID := getPeriodEventID(event.Period, event.EventID)
	return &dbEvent{
		PeriodEventIDPK: peID,
		PeriodEventIDSK: peID,
		Name:            event.Name,
		Description:     event.Description,
	}
}

func newDBApplicant(applicant *model.Applicant) *dbApplicant {
	peID := getPeriodEventID(applicant.Period, applicant.EventID)
	return &dbApplicant{
		PeriodEventID: peID,
		Email:         getApplicantID(applicant.Email),
		Name:          applicant.Name,
		Info:          applicant.Info,
	}
}

func getApplicant(item *dbApplicant) *model.Applicant {
	p, e := getPeriodAndEventID(item.PeriodEventID)
	return &model.Applicant{
		Period:  p,
		EventID: e,
		Email:   removePrefix(item.Email),
		Name:    item.Name,
		Info:    item.Info,
	}
}

func getEvent(item *dbEvent) *model.Event {
	p, e := getPeriodAndEventID(item.PeriodEventIDPK)
	return &model.Event{
		Period:      p,
		EventID:     e,
		Name:        item.Name,
		Description: item.Description,
	}
}

func getClubTable(clubID string) *string {
	return aws.String(clubTablePrefix + clubID)
}

func getPeriodEventID(period string, eventID string) string {
	return peidPrefix + period + "-" + eventID
}

func getApplicantID(email string) string {
	return applicantPrefix + email
}

func getPeriodAndEventID(peid string) (string, string) {
	str := removePrefix(peid)
	s := strings.SplitN(str, "-", 2)
	return s[0], s[1]
}

// AddNewEvent creates a new event in the club table
func (db *Database) AddNewEvent(clubID string, event *model.Event) error {
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

// GetEvent returns event with a application period and eventID
func (db *Database) GetEvent(clubID string, period string, eventID string) (*model.Event, error) {
	peid := aws.String(getPeriodEventID(period, eventID))
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

	var item dbEvent
	if err := dynamodbattribute.UnmarshalMap(result.Item, &item); err != nil {
		return nil, err
	}
	return getEvent(&item), nil
}

// GetEvents returns all the events of an application period
func (db *Database) GetEvents(clubID string, period string) ([]*model.Event, error) {
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
	items := make([]*dbEvent, *result.Count)
	if err := dynamodbattribute.UnmarshalListOfMaps(result.Items, &items); err != nil {
		return nil, err
	}
	events := make([]*model.Event, *result.Count)
	for _, item := range items {
		events = append(events, getEvent(item))
	}
	return events, nil
}

// DeleteEvent deletes an event and all of its applicants
func (db *Database) DeleteEvent(clubID string, event *model.Event) error {
	input := &dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":id": {
				S: aws.String(getPeriodEventID(event.Period, event.EventID)),
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

// AddNewApplicant creates a new applicant in the club table for an event
func (db *Database) AddNewApplicant(clubID string, applicant *model.Applicant) error {
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

// GetEventApplicant returns an applicant for a specific event and application period
func (db *Database) GetEventApplicant(clubID string, period string, eventID string, email string) (*model.Applicant, error) {
	result, err := db.c.GetItem(&dynamodb.GetItemInput{
		TableName: getClubTable(clubID),
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {S: aws.String(getPeriodEventID(period, eventID))},
			"sk": {S: aws.String(getApplicantID(email))},
		},
	})
	if err != nil {
		return nil, err
	}

	var item dbApplicant
	if err := dynamodbattribute.UnmarshalMap(result.Item, &item); err != nil {
		return nil, err
	}
	return getApplicant(&item), nil
}

// GetEventApplicants returns all the applicants for an event and application period
func (db *Database) GetEventApplicants(clubID string, period string, eventID string) ([]*model.Applicant, error) {
	input := &dynamodb.QueryInput{
		TableName: getClubTable(clubID),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":id": {
				S: aws.String(getPeriodEventID(period, eventID)),
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
	items := make([]*dbApplicant, *result.Count)
	if err := dynamodbattribute.UnmarshalListOfMaps(result.Items, &items); err != nil {
		return nil, err
	}
	applicants := make([]*model.Applicant, *result.Count)
	for _, item := range items {
		applicants = append(applicants, getApplicant(item))
	}
	return applicants, nil
}

// DeleteEventApplicant deletes an applicant
func (db *Database) DeleteEventApplicant(clubID string, applicant *model.Applicant) error {
	if _, err := db.c.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: getClubTable(clubID),
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {S: aws.String(getPeriodEventID(applicant.Period, applicant.EventID))},
			"sk": {S: aws.String(getApplicantID(applicant.Email))},
		},
	}); err != nil {
		return err
	}
	return nil
}
