package database

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/ubclaunchpad/pinpoint/utils"
)

// NewTestDB is a utility function to create a connection to a local dynamodb
// instance and expose the underlying client for test purposes
func NewTestDB() (*Database, *dynamodb.DynamoDB) {
	s, err := utils.AWSSession(utils.AWSConfig(true))
	if err != nil {
		return nil, nil
	}
	l, err := utils.NewLogger(true)
	if err != nil {
		return nil, nil
	}

	// attempt to create database with connection
	db, err := New(s, l)
	if err != nil {
		return nil, nil
	}

	return db, db.c
}
