package database

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func (db *Database) initTables() error {
	var tables = []*dynamodb.CreateTableInput{
		{
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
		},
		{
			TableName: aws.String("ClubsAndUsers"),
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
		},
		{
			TableName: aws.String("TagTable"),
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
		},
	}

	// init all tables, collecting critical errors on the way
	for _, t := range tables {
		_, err := db.c.CreateTable(t)
		if err != nil {
			// ignore error if table exists
			if strings.Contains(err.Error(), dynamodb.ErrCodeResourceInUseException) {
				db.l.Warnw(fmt.Sprintf("table '%s' already exists - ignoring", *t.TableName),
					"table", *t.TableName,
					"error", err.Error())
			} else {
				return err
			}
		}
	}

	return nil
}
