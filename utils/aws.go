package utils

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

// AWSConfig generates an AWS session Configuration
func AWSConfig(dev bool) (cfg *aws.Config) {
	if dev {
		cfg = &aws.Config{
			// dynamodb-local
			Endpoint: aws.String("http://localhost:8000"),

			// static credentials
			Credentials: credentials.NewStaticCredentials("robert", "wow", "launchpad"),

			// arbitrary region
			Region: aws.String("us-west-2"),
		}
		if os.Getenv("AWS_DEBUG") == "true" {
			cfg.LogLevel = aws.LogLevel(aws.LogDebug)
		}
	} else {
		// todo: production aws setup
	}
	return cfg
}

// AWSSession initializes an AWS API session with given configs
func AWSSession(cfg ...*aws.Config) (*session.Session, error) {
	return session.NewSession(cfg...)
}
