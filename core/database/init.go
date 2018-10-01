package database

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func (db *Database) initTables() error {
	// create a test table
	_, err := db.c.CreateTable(&dynamodb.CreateTableInput{
		TableName: aws.String("test-table"),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: aws.String("N"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       aws.String("HASH"),
			},
		},
		// operations per second before throttle
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
	})

	// ignore error if table exists
	if err != nil && !strings.Contains(err.Error(), dynamodb.ErrCodeResourceInUseException) {
		return err
	}

	return nil
}
