package database

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/ubclaunchpad/pinpoint/core/model"
)

const clubsAndUsersTable string = "ClubsAndUsers"

// AddNewUser creates a new user in the database
func (db *Database) AddNewUser(u *model.User) error {
	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"pk": {
				S: prefixUserEmail(u.Email),
			},
			"sk": {
				S: prefixUserEmail(u.Email),
			},
			"name": {
				S: aws.String(u.Name),
			},
			"salt": {
				S: aws.String(u.Salt),
			},
		},
		TableName: aws.String(clubsAndUsersTable),
	}
	if _, err := db.c.PutItem(input); err != nil {
		return err
	}
	return nil
}

// GetUser returns a user from the database with the given email
func (db *Database) GetUser(email string) (*model.User, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(clubsAndUsersTable),
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {
				S: prefixUserEmail(email),
			},
			"sk": {
				S: prefixUserEmail(email),
			},
		},
	}
	result, err := db.c.GetItem(input)
	if err != nil {
		return nil, err
	}
	user := &model.User{}
	if err := dynamodbattribute.UnmarshalMap(result.Item, user); err != nil {
		return nil, err
	}
	user.Email = removePrefix(user.Email) // remove 'User-' prefix
	return user, nil
}

// DeleteUser deletes a user from the database with the given email
func (db *Database) DeleteUser(email string) error {
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(clubsAndUsersTable),
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {
				S: prefixUserEmail(email),
			},
			"sk": {
				S: prefixUserEmail(email),
			},
		},
	}
	if _, err := db.c.DeleteItem(input); err != nil {
		return err
	}
	return nil
}

// AddNewClub creates a new club in the database with a user (creator) associated to it
func (db *Database) AddNewClub(c *model.Club, cu *model.ClubUser) error {
	input := &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]*dynamodb.WriteRequest{
			clubsAndUsersTable: {
				{
					PutRequest: &dynamodb.PutRequest{
						Item: map[string]*dynamodb.AttributeValue{
							"pk": {
								S: prefixClubID(c.ID),
							},
							"sk": {
								S: prefixClubID(c.ID),
							},
							"name": {
								S: aws.String(c.Name),
							},
							"description": {
								S: aws.String(c.Description),
							},
						},
					},
				},
				{
					PutRequest: &dynamodb.PutRequest{
						Item: map[string]*dynamodb.AttributeValue{
							"pk": {
								S: prefixClubID(cu.ClubID),
							},
							"sk": {
								S: prefixUserEmail(cu.Email),
							},
							"name": {
								S: aws.String(cu.UserName),
							},
							"role": {
								S: aws.String(cu.Role),
							},
						},
					},
				},
			},
		},
	}
	if _, err := db.c.BatchWriteItem(input); err != nil {
		return err
	}
	t := &dynamodb.CreateTableInput{
		TableName: aws.String("ClubData-" + c.ID),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("pk"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("sk"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("pk"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("sk"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
	}
	if _, err := db.c.CreateTable(t); err != nil {
		return err
	}
	return nil
}

// GetClub returns a club with the given id
func (db *Database) GetClub(id string) (*model.Club, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(clubsAndUsersTable),
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {
				S: prefixClubID(id),
			},
			"sk": {
				S: prefixClubID(id),
			},
		},
	}
	result, err := db.c.GetItem(input)
	if err != nil {
		return nil, err
	}
	club := &model.Club{}
	if err := dynamodbattribute.UnmarshalMap(result.Item, club); err != nil {
		return nil, err
	}
	club.ID = removePrefix(club.ID) // remove 'Club-' prefix
	return club, nil
}

// GetAllClubUsers returns a list of ClubUsers associated with a club
func (db *Database) GetAllClubUsers(id string) ([]*model.ClubUser, error) {
	input := &dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":id": {
				S: prefixClubID(id),
			},
			":u": {
				S: prefixUserEmail(""),
			},
		},
		KeyConditionExpression: aws.String("pk = :id AND begins_with(sk, :u)"),
		TableName:              aws.String(clubsAndUsersTable),
	}

	result, err := db.c.Query(input)
	if err != nil {
		return nil, err
	}
	cus := make([]*model.ClubUser, *result.Count)
	if err := dynamodbattribute.UnmarshalListOfMaps(result.Items, &cus); err != nil {
		return nil, err
	}
	for _, cu := range cus {
		cu.ClubID = removePrefix(cu.ClubID) // remove 'Club-' prefix
		cu.Email = removePrefix(cu.Email)   // remove 'User-' prefix
	}
	return cus, nil
}

// DeleteClub deletes a club from the database with the given id
func (db *Database) DeleteClub(id string) error {
	input := &dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":id": {
				S: prefixClubID(id),
			},
		},
		KeyConditionExpression: aws.String("pk = :id"),
		ProjectionExpression:   aws.String("pk, sk"),
		TableName:              aws.String(clubsAndUsersTable),
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
			clubsAndUsersTable: batch,
		},
	}
	if _, err = db.c.BatchWriteItem(batchInput); err != nil {
		return err
	}
	t := &dynamodb.DeleteTableInput{
		TableName: aws.String("ClubData-" + id),
	}
	if _, err = db.c.DeleteTable(t); err != nil {
		return err
	}
	return nil
}

func prefixClubID(id string) *string {
	return aws.String("Club-" + id)
}

func prefixUserEmail(email string) *string {
	return aws.String("User-" + email)
}

func removePrefix(str string) string {
	return strings.SplitN(str, "-", 2)[1]
}
