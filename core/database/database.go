package database

import (
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"go.uber.org/zap"
)

// Database handles database connections
type Database struct {
	c *dynamodb.DynamoDB
	l *zap.SugaredLogger
}

// New creates a new database client
func New(awsConfig client.ConfigProvider, logger *zap.SugaredLogger) (*Database, error) {
	db := &Database{dynamodb.New(awsConfig), logger.Named("db")}

	// set up database
	logger.Infow("setting up database",
		"db.client", db.c.ClientInfo)
	if err := db.initTables(); err != nil {
		return nil, err
	}

	// log current tables
	tables, err := db.c.ListTables(nil)
	if err != nil {
		return nil, err
	}
	db.l.Infow("connected to database",
		"tables", tables.TableNames)

	return db, nil
}
