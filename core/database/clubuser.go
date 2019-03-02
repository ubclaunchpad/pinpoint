package database

import (
	"errors"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/ubclaunchpad/pinpoint/protobuf/models"
)

// AddNewUser creates a new user in the database
func (db *Database) AddNewUser(u *models.User, e *models.EmailVerification) error {
	if u.Email == "" || e.Email == "" || e.Hash == "" {
		return errors.New("Keys cannot be empty")
	}
	uItem, err := dynamodbattribute.MarshalMap(newDBUser(u))
	if err != nil {
		return err
	}
	evItem, err := dynamodbattribute.MarshalMap(newDBEmailVerification(e))
	if err != nil {
		return err
	}
	var table = getClubsAndUsersTable()

	if _, err := db.c.BatchWriteItem(&dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]*dynamodb.WriteRequest{
			*table: {
				{PutRequest: &dynamodb.PutRequest{Item: uItem}},
				{PutRequest: &dynamodb.PutRequest{Item: evItem}},
			},
		},
	}); err != nil {
		return err
	}

	return nil
}

// GetUser returns a user from the database with the given email
func (db *Database) GetUser(email string) (*models.User, error) {
	if email == "" {
		return nil, errors.New("email can not be empty")
	}
	var e = aws.String(prefixUserEmail(email))
	result, err := db.c.GetItem(&dynamodb.GetItemInput{
		TableName: getClubsAndUsersTable(),
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {S: e},
			"sk": {S: e},
		},
	})
	if err != nil {
		return nil, err
	}
	var item userItem
	if err := dynamodbattribute.UnmarshalMap(result.Item, &item); err != nil {
		return nil, err
	}
	return newUser(&item), nil
}

// DeleteUser deletes a user from the database with the given email
func (db *Database) DeleteUser(email string) error {
	if email == "" {
		return errors.New("Email can not be empty")
	}
	var e = aws.String(prefixUserEmail(email))
	if _, err := db.c.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: getClubsAndUsersTable(),
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {S: e},
			"sk": {S: e},
		},
	}); err != nil {
		return err
	}
	return nil
}

// GetEmailVerification returns a pending email verification from the database
// with the given email and hash
func (db *Database) GetEmailVerification(email string, hash string) (*models.EmailVerification, error) {
	if email == "" || hash == "" {
		return nil, errors.New("email or hash can not be empty")
	}
	var e = aws.String(prefixUserEmail(email))
	var h = aws.String(prefixVerificationHash(hash))
	var table = getClubsAndUsersTable()
	result, err := db.c.GetItem(&dynamodb.GetItemInput{
		TableName: table,
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {S: e},
			"sk": {S: h},
		},
	})
	if err != nil {
		return nil, err
	}

	var item emailVerificationItem
	if err := dynamodbattribute.UnmarshalMap(result.Item, &item); err != nil {
		return nil, err
	}
	var ev = newEmailVerification(&item)

	if ev.Hash != hash {
		return nil, errors.New("verification code not found")
	}

	// delete verification
	db.c.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: table,
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {S: e},
			"sk": {S: h},
		},
	})

	// check expiry
	if time.Unix(ev.Expiry, 0).Before(time.Now()) {
		return nil, fmt.Errorf("verification expired on %v", item.Expiry)
	}

	u, err := db.GetUser(email)
	if err != nil {
		return nil, fmt.Errorf("user not found in db: %s", err.Error())
	}
	u.Verified = true
	uItem, err := dynamodbattribute.MarshalMap(newDBUser(u))
	if err != nil {
		return nil, fmt.Errorf("failed to marshal user: %s", err.Error())
	}
	if _, err := db.c.PutItem(&dynamodb.PutItemInput{
		Item:      uItem,
		TableName: getClubsAndUsersTable(),
	}); err != nil {
		return nil, fmt.Errorf("failed to update user: %s", err.Error())
	}

	return ev, nil
}

// AddNewClub creates a new club in the database with a user (creator) associated to it
func (db *Database) AddNewClub(c *models.Club, cu *models.ClubUser) error {
	if c.ClubID == "" || cu.ClubID == "" || cu.Email == "" {
		return errors.New("keys cannot be empty")
	}
	cItem, err := dynamodbattribute.MarshalMap(newDBClub(c))
	if err != nil {
		return err
	}
	cuItem, err := dynamodbattribute.MarshalMap(newDBClubUser(cu))
	if err != nil {
		return err
	}
	var table = getClubsAndUsersTable()

	// create new club entry
	if _, err := db.c.BatchWriteItem(&dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]*dynamodb.WriteRequest{
			*table: {
				{PutRequest: &dynamodb.PutRequest{Item: cItem}},
				{PutRequest: &dynamodb.PutRequest{Item: cuItem}},
			},
		},
	}); err != nil {
		return err
	}

	// create table for club
	if _, err := db.c.CreateTable(&dynamodb.CreateTableInput{
		TableName: getClubTable(c.ClubID),
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
func (db *Database) GetClub(id string) (*models.Club, error) {
	if id == "" {
		return nil, errors.New("ID can not be empty")
	}
	var cID = aws.String(prefixClubID(id))
	input := &dynamodb.GetItemInput{
		TableName: getClubsAndUsersTable(),
		Key: map[string]*dynamodb.AttributeValue{
			"pk": {S: cID},
			"sk": {S: cID},
		},
	}
	result, err := db.c.GetItem(input)
	if err != nil {
		return nil, err
	}
	var item clubItem
	if err := dynamodbattribute.UnmarshalMap(result.Item, &item); err != nil {
		return nil, err
	}
	return newClub(&item), nil
}

// GetAllClubUsers returns a list of ClubUsers associated with a club
func (db *Database) GetAllClubUsers(id string) ([]*models.ClubUser, error) {
	if id == "" {
		return nil, errors.New("ID can not be empty")
	}
	input := &dynamodb.QueryInput{
		TableName: getClubsAndUsersTable(),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":id": {S: aws.String(prefixClubID(id))},
			":u":  {S: aws.String(prefixUserEmail(""))},
		},
		KeyConditionExpression: aws.String("pk = :id AND begins_with(sk, :u)"),
	}

	result, err := db.c.Query(input)
	if err != nil {
		return nil, err
	}
	cuItems := make([]*clubUserItem, *result.Count)
	if err := dynamodbattribute.UnmarshalListOfMaps(result.Items, &cuItems); err != nil {
		return nil, err
	}
	cus := make([]*models.ClubUser, *result.Count)
	for i, item := range cuItems {
		cus[i] = newClubUser(item)
	}
	return cus, nil
}

// DeleteClub deletes a club from the database with the given id
func (db *Database) DeleteClub(id string) error {
	if id == "" {
		return errors.New("ID can not be empty")
	}
	var table = getClubsAndUsersTable()
	input := &dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":id": {S: aws.String(prefixClubID(id))},
		},
		KeyConditionExpression: aws.String("pk = :id"),
		ProjectionExpression:   aws.String("pk, sk"),
		TableName:              table,
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
			*table: batch,
		},
	}
	if _, err = db.c.BatchWriteItem(batchInput); err != nil {
		return err
	}
	t := &dynamodb.DeleteTableInput{
		TableName: getClubTable(id),
	}
	if _, err = db.c.DeleteTable(t); err != nil {
		return err
	}
	return nil
}
