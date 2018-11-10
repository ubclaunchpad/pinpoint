package database

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/ubclaunchpad/pinpoint/core/model"
)

/* Adds a new tag */
func (db *Database) AddNewTag(t *model.Tag) error {
	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"pk": {
				S: aws.String(t.Applicant_ID),
			},
			"sk": {
				S: aws.String(t.Period_Event_ID),
			},
			"tag_name": {
				S: aws.String(t.Tag_Name),
			},
		},
		TableName: aws.String("TagTable"),
	}
	if _, err := db.c.PutItem(input); err != nil {
		return err
	}
	return nil
}

/* Gets the tag associated with Applicant_ID, Period_ID, & Event_ID */
func (db *Database) GetTag(Applicant_ID string, Period_ID string, Event_ID string) (*model.Tag, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String("TagTable"),
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {
				S: aws.String(Applicant_ID),
			},
			"sk": {
				S: aws.String(Period_ID + "_" + Event_ID),
			},
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

/* Deletes the tag associated with Applicant_ID, Period_ID, & Event_ID */
func (db *Database) DeleteTag(Applicant_ID string, Period_ID string, Event_ID string) error {
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String("TagTable"),
    Key: map[string]*dynamodb.AttributeValue{
			"pk": {
				S: aws.String(Applicant_ID),
			},
			"sk": {
				S: aws.String(Period_ID + "_" + Event_ID),
			},
		},
	}
	if _, err := db.c.DeleteItem(input); err != nil {
		return err
	}
	return nil
}


/* Changes the tag name associated with Applicant_ID, Period_ID, & Event_ID */
func (db *Database) ChangeTagName(Applicant_ID string, Period_ID string, Event_ID string, New_Name string) (*model.Tag, error)  {
	input := &dynamodb.UpdateItemInput{
    ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
       ":new_name": {
           S: aws.String(New_Name),
       },
   },
   ExpressionAttributeNames: map[string]*string{
        "#attr_name": aws.String("Tag_Name"),
    },
    TableName: aws.String("TagTable"),
    Key: map[string]*dynamodb.AttributeValue{
			"pk": {
				S: aws.String(Applicant_ID),
			},
			"sk": {
				S: aws.String(Period_ID + "_" + Event_ID),
			},
		},
    UpdateExpression: aws.String("SET #attr_name = :new_name"),
    ReturnValues: aws.String("ALL_NEW"),
	}
  result, err := db.c.UpdateItem(input)
	if err != nil {
		return nil, err
	}
	tag := &model.Tag{}
	if err := dynamodbattribute.UnmarshalMap(result.Attributes, tag); err != nil {
		return nil, err
	}
	return tag, nil
}
