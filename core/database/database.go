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
	c := dynamodb.New(awsConfig)
	return &Database{c, logger.Named("db")}, nil
}
