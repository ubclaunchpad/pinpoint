package database

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/ubclaunchpad/pinpoint/utils"
	"go.uber.org/zap/zaptest"
)

// newTestDB is a utility function to create a connection to a local dynamodb
// instance and expose the underlying client for test purposes
func newTestDB(t *testing.T) (*Database, *dynamodb.DynamoDB) {
	s, err := utils.AWSSession(utils.AWSConfig(true))
	if err != nil {
		return nil, nil
	}

	// attempt to create database with connection
	db, err := New(s, zaptest.NewLogger(t).Sugar())
	if err != nil {
		return nil, nil
	}

	return db, db.c
}

func TestNew(t *testing.T) {
	newTestDB(t)
}
