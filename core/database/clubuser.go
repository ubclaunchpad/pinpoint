package database

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/ubclaunchpad/pinpoint/core/model"
)

var (
	tableClubsAndUsers   = aws.String("ClubsAndUsers")
	keyEmailVerification = aws.String("verification")
)

func prefixClubID(id string) *string {
	return aws.String("Club-" + id)
}

func prefixUserEmail(email string) *string {
	return aws.String("User-" + email)
}

func removePrefix(str string) string {
	s := strings.SplitN(str, "-", 2)
	if len(s) > 1 {
		return s[1]
	}
	return str
}

// AddNewUser creates a new user in the database
func (db *Database) AddNewUser(u *model.User, e *model.EmailVerification) error {
	verify, err := dynamodbattribute.MarshalMap(e)
	if err != nil {
		return err
	}
	verify["sk"] = &dynamodb.AttributeValue{S: keyEmailVerification}

	if _, err := db.c.BatchWriteItem(&dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]*dynamodb.WriteRequest{
			*tableClubsAndUsers: {
				{PutRequest: &dynamodb.PutRequest{Item: map[string]*dynamodb.AttributeValue{
					"pk":       {S: prefixUserEmail(u.Email)},
					"sk":       {S: prefixUserEmail(u.Email)},
					"name":     {S: aws.String(u.Name)},
					"salt":     {S: aws.String(u.Salt)},
					"verified": {BOOL: aws.Bool(u.Verified)},
				}}},
				{PutRequest: &dynamodb.PutRequest{Item: verify}},
			},
		},
	}); err != nil {
		return err
	}
	return nil
}

// GetUser returns a user from the database with the given email
func (db *Database) GetUser(email string) (*model.User, error) {
	result, err := db.c.GetItem(&dynamodb.GetItemInput{
		TableName: tableClubsAndUsers,
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {S: prefixUserEmail(email)},
			"sk": {S: prefixUserEmail(email)},
		},
	})
	if err != nil {
		return nil, err
	}

	var user model.User
	if err := dynamodbattribute.UnmarshalMap(result.Item, &user); err != nil {
		return nil, err
	}

	// remove 'User-' prefix for usability
	user.Email = removePrefix(user.Email)
	if email == "" || user.Email != email {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

// DeleteUser deletes a user from the database with the given email
func (db *Database) DeleteUser(email string) error {
	if _, err := db.c.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: tableClubsAndUsers,
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {S: prefixUserEmail(email)},
			"sk": {S: prefixUserEmail(email)},
		},
	}); err != nil {
		return err
	}
	return nil
}

// GetEmailVerification returns a pending email verification from the database
// with the given email
func (db *Database) GetEmailVerification(hash string) (*model.EmailVerification, error) {
	result, err := db.c.GetItem(&dynamodb.GetItemInput{
		TableName: tableClubsAndUsers,
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {S: aws.String(hash)},
			"sk": {S: keyEmailVerification},
		},
	})
	if err != nil {
		return nil, err
	}

	var v model.EmailVerification
	if err := dynamodbattribute.UnmarshalMap(result.Item, &v); err != nil {
		return nil, err
	}

	println(hash, v.Hash)

	if v.Hash != hash {
		return nil, errors.New("verification code not found")
	}

	// delete verification
	db.c.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: tableClubsAndUsers,
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {S: aws.String(hash)},
			"sk": {S: keyEmailVerification},
		},
	})

	// check expiry
	if v.Expiry.Before(time.Now()) {
		return nil, fmt.Errorf("verification expired on %v", v.Expiry)
	}

	return &v, nil
}

// AddNewClub creates a new club in the database with a user (creator) associated to it
func (db *Database) AddNewClub(c *model.Club, cu *model.ClubUser) error {
	// create new club entry
	if _, err := db.c.BatchWriteItem(&dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]*dynamodb.WriteRequest{
			*tableClubsAndUsers: {
				{
					PutRequest: &dynamodb.PutRequest{
						Item: map[string]*dynamodb.AttributeValue{
							"pk":          {S: prefixClubID(c.ID)},
							"sk":          {S: prefixClubID(c.ID)},
							"name":        {S: aws.String(c.Name)},
							"description": {S: aws.String(c.Description)},
						},
					},
				},
				{
					PutRequest: &dynamodb.PutRequest{
						Item: map[string]*dynamodb.AttributeValue{
							"pk":   {S: prefixClubID(cu.ClubID)},
							"sk":   {S: prefixUserEmail(cu.Email)},
							"name": {S: aws.String(cu.UserName)},
							"role": {S: aws.String(cu.Role)},
						},
					},
				},
			},
		},
	}); err != nil {
		return err
	}

	// create table for club
	if _, err := db.c.CreateTable(&dynamodb.CreateTableInput{
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
	}); err != nil {
		return err
	}

	return nil
}

// GetClub returns a club with the given id
func (db *Database) GetClub(id string) (*model.Club, error) {
	input := &dynamodb.GetItemInput{
		TableName: tableClubsAndUsers,
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
		TableName: tableClubsAndUsers,
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":id": {
				S: prefixClubID(id),
			},
			":u": {
				S: prefixUserEmail(""),
			},
		},
		KeyConditionExpression: aws.String("pk = :id AND begins_with(sk, :u)"),
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
		TableName:              tableClubsAndUsers,
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
			*tableClubsAndUsers: batch,
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
