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
	}

	// init all tables, collecting critical errors on the way
	tableErrors := make([]error, 0)
	for _, t := range tables {
		_, err := db.c.CreateTable(t)
		if err != nil {
			// ignore error if table exists
			if strings.Contains(err.Error(), dynamodb.ErrCodeResourceInUseException) {
				db.l.Warnw(fmt.Sprintf("table '%s' already exists - ignoring", *t.TableName),
					"table", *t.TableName,
					"error", err.Error())
			} else {
				tableErrors = append(tableErrors, err)
			}
		}
	}

	if len(tableErrors) > 0 {
		return fmt.Errorf("encountered errors in tables initialization: %v", tableErrors)
	}

	return nil
}
