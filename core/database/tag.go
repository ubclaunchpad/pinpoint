package database

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/ubclaunchpad/pinpoint/core/model"
)

// AddNewTag adds a tag to the database
func (db *Database) AddNewTag(t *model.Tag, c *model.Club) error {
	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"pk":       {S: aws.String(t.ApplicantID)},
			"sk":       {S: aws.String(t.PeriodEventID)},
			"tag_name": {S: aws.String(t.TagName)},
			"type":     {S: aws.String("tag")},
		},
		TableName: aws.String("ClubData-" + c.ID),
	}
	if _, err := db.c.PutItem(input); err != nil {
		return err
	}
	return nil
}

// GetTag gets the tag associated with Applicant_ID, Period_ID, & Event_ID
func (db *Database) GetTag(ApplicantID string, PeriodID string, EventID string, c *model.Club) (*model.Tag, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String("ClubData-" + c.ID),
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {S: aws.String(ApplicantID)},
			"sk": {S: aws.String(PeriodID + "_" + EventID)},
		},
	}
	result, err := db.c.GetItem(input)
	if err != nil {
		return nil, err
	}
	tag := &model.Tag{}
	if err := dynamodbattribute.UnmarshalMap(result.Item, tag); err != nil {
		return nil, err
	}
	return tag, nil
}

// DeleteTag associated with Applicant_ID, Period_ID, & Event_ID
func (db *Database) DeleteTag(ApplicantID string, PeriodID string, EventID string, c *model.Club) error {
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String("ClubData-" + c.ID),
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {S: aws.String(ApplicantID)},
			"sk": {S: aws.String(PeriodID + "_" + EventID)},
		},
	}
	if _, err := db.c.DeleteItem(input); err != nil {
		return err
	}
	return nil
}
