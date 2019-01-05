package database

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/ubclaunchpad/pinpoint/protobuf/models"
)

// AddNewEvent creates a new event in the club table
func (db *Database) AddNewEvent(clubID string, event *models.Event) error {
	e := newDBEvent(event)
	item, err := dynamodbattribute.MarshalMap(e)
	if err != nil {
		return fmt.Errorf("Failed to marshal event: %s", err.Error())
	}
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: getClubTable(clubID),
	}
	if _, err := db.c.PutItem(input); err != nil {
		return fmt.Errorf("Failed to put event: %s", err.Error())
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
		return nil, fmt.Errorf("Failed to get event: %s", err.Error())
	}

	var item eventItem
	if err := dynamodbattribute.UnmarshalMap(result.Item, &item); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal event: %s", err.Error())
	}
	return newEvent(&item), nil
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
		return nil, fmt.Errorf("Failed query for all events: %s", err.Error())
	}
	items := make([]*eventItem, *result.Count)
	if err := dynamodbattribute.UnmarshalListOfMaps(result.Items, &items); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal events: %s", err.Error())
	}
	events := make([]*models.Event, *result.Count)
	for _, item := range items {
		events = append(events, newEvent(item))
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
		return fmt.Errorf("Failed query for event: %s", err.Error())
	}

	batch := make([]*dynamodb.WriteRequest, len(result.Items))
	for i, item := range result.Items {
		batch[i] = &dynamodb.WriteRequest{
			DeleteRequest: &dynamodb.DeleteRequest{
				Key: item,
			},
		}
	}
	batchInput := &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]*dynamodb.WriteRequest{
			*getClubTable(clubID): batch,
		},
	}
	if _, err = db.c.BatchWriteItem(batchInput); err != nil {
		return fmt.Errorf("Failed batch delete requests for event: %s", err.Error())
	}
	return nil
}

// AddNewApplicant creates a new applicant in the club table for an application period
func (db *Database) AddNewApplicant(clubID string, applicant *models.Applicant) error {
	item, err := dynamodbattribute.MarshalMap(newDBApplicant(applicant))
	if err != nil {
		return fmt.Errorf("Failed to marshal applicant: %s", err.Error())
	}
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: getClubTable(clubID),
	}
	if _, err := db.c.PutItem(input); err != nil {
		return fmt.Errorf("Failed to put applicant: %s", err.Error())
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
		return nil, fmt.Errorf("Failed to get applicant: %s", err.Error())
	}

	var item applicantItem
	if err := dynamodbattribute.UnmarshalMap(result.Item, &item); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal applicant: %s", err.Error())
	}
	return newApplicant(&item), nil
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
		return nil, fmt.Errorf("Failed query for all applicants: %s", err.Error())
	}
	items := make([]*applicantItem, *result.Count)
	if err := dynamodbattribute.UnmarshalListOfMaps(result.Items, &items); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal all applicants: %s", err.Error())
	}
	applicants := make([]*models.Applicant, *result.Count)
	for _, item := range items {
		applicants = append(applicants, newApplicant(item))
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
		return fmt.Errorf("Failed query for applicant: %s", err.Error())
	}

	batch := make([]*dynamodb.WriteRequest, len(result.Items))
	for i, item := range result.Items {
		batch[i] = &dynamodb.WriteRequest{
			DeleteRequest: &dynamodb.DeleteRequest{
				Key: item,
			},
		}
	}
	batchInput := &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]*dynamodb.WriteRequest{
			*getClubTable(clubID): batch,
		},
	}
	if _, err = db.c.BatchWriteItem(batchInput); err != nil {
		return fmt.Errorf("Failed batch delete request for applicant: %s", err.Error())
	}
	return nil
}

// AddNewApplication adds an application to the database
func (db *Database) AddNewApplication(clubID string, application *models.Application) error {
	a := newDBApplication(application)
	item, err := dynamodbattribute.MarshalMap(a)
	if err != nil {
		return fmt.Errorf("Failed to marshal application: %s", err.Error())
	}
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: getClubTable(clubID),
	}
	if _, err := db.c.PutItem(input); err != nil {
		return fmt.Errorf("Failed to put application: %s", err.Error())
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
		return nil, fmt.Errorf("Failed to get application: %s", err.Error())
	}

	var item applicationItem
	if err := dynamodbattribute.UnmarshalMap(result.Item, &item); err != nil {
		return nil, fmt.Errorf("Failed to put application: %s", err.Error())
	}
	return newApplication(&item), nil
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
		return nil, fmt.Errorf("Failed to get all applications: %s", err.Error())
	}
	items := make([]*applicationItem, *result.Count)
	if err := dynamodbattribute.UnmarshalListOfMaps(result.Items, &items); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal applications: %s", err.Error())
	}
	applications := make([]*models.Application, *result.Count)
	for _, item := range items {
		applications = append(applications, newApplication(item))
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
		return fmt.Errorf("Failed to delete application: %s", err.Error())
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
		return nil, fmt.Errorf("Failed query for all tags: %s", err.Error())
	}
	items := make([]*tagItem, *result.Count)
	if err := dynamodbattribute.UnmarshalListOfMaps(result.Items, &items); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal tags: %s", err.Error())
	}
	tags := make([]*models.Tag, *result.Count)
	for _, item := range items {
		tags = append(tags, newTag(item))
	}
	return tags, nil
}
