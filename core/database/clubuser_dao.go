package database

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/ubclaunchpad/pinpoint/core/model"
)

const clubuserTable string = "Clubusers"

// AddNewUser creates a new user in the database
func (db *Database) AddNewUser(u *model.User) error {
	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"pk": {
				S: aws.String("User-" + u.Email),
			},
			"sk": {
				S: aws.String("User-" + u.Email),
			},
			"name": {
				S: aws.String(u.Name),
			},
			"salt": {
				S: aws.String(u.Salt),
			},
		},
		TableName: aws.String(clubuserTable),
	}
	_, err := db.c.PutItem(input)
	if err != nil {
		return err
	}
	return nil
}

// GetUser returns a user from the database with the given email
func (db *Database) GetUser(email string) (*model.User, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(clubuserTable),
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {
				S: aws.String("User-" + email),
			},
			"sk": {
				S: aws.String("User-" + email),
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
	user.Email = user.Email[5:] // remove 'User-' prefix
	return user, nil
}

// AddNewClub creates a new club in the database with a user (creator) associated to it
func (db *Database) AddNewClub(c *model.Club, cu *model.Clubuser) error {
	input := &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]*dynamodb.WriteRequest{
			clubuserTable: {
				{
					PutRequest: &dynamodb.PutRequest{
						Item: map[string]*dynamodb.AttributeValue{
							"pk": {
								S: aws.String("Club-" + c.ID),
							},
							"sk": {
								S: aws.String("Club-" + c.ID),
							},
							"name": {
								S: aws.String(c.Name),
							},
							"description": {
								S: aws.String(c.Description),
							},
							"periods": {
								SS: aws.StringSlice(c.Periods),
							},
						},
					},
				},
				{
					PutRequest: &dynamodb.PutRequest{
						Item: map[string]*dynamodb.AttributeValue{
							"pk": {
								S: aws.String("Club-" + cu.ClubID),
							},
							"sk": {
								S: aws.String("User-" + cu.Email),
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
	_, err := db.c.BatchWriteItem(input)
	if err != nil {
		return err
	}
	return nil
}

// GetClub returns a club with the given id
func (db *Database) GetClub(id string) (*model.Club, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(clubuserTable),
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {
				S: aws.String("Club-" + id),
			},
			"sk": {
				S: aws.String("Club-" + id),
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
	club.ID = club.ID[5:] // remove 'Club-' prefix
	return club, nil
}

// GetAllClubusers returns a list of Clubusers associated with a club
func (db *Database) GetAllClubusers(id string) (*[]*model.Clubuser, error) {
	input := &dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":id": {
				S: aws.String("Club-" + id),
			},
			":u": {
				S: aws.String("User-"),
			},
		},
		KeyConditionExpression: aws.String("pk = :id AND begins_with(sk, :u)"),
		TableName:              aws.String(clubuserTable),
	}

	result, err := db.c.Query(input)
	if err != nil {
		return nil, err
	}
	cus := make([]*model.Clubuser, *result.Count)
	if err := dynamodbattribute.UnmarshalListOfMaps(result.Items, &cus); err != nil {
		return nil, err
	}
	for _, cu := range cus {
		cu.ClubID = cu.ClubID[5:] // remove 'Club-' prefix
		cu.Email = cu.Email[5:]   // remove 'User-' prefix
	}
	return &cus, nil
}
