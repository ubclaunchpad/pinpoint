package utils

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

// AWSDebugEnv is the key used for the env variable that toggles debug logs
const AWSDebugEnv = "AWS_DEBUG"

// Logger defines a logger that can be configured for use with the AWS SDK
type Logger interface {
	Info(...interface{})
}

// AWSConfig generates an AWS session Configuration. Only the first provided
// logger is used.
func AWSConfig(dev bool, logger ...Logger) (cfg *aws.Config) {
	if dev {
		cfg = &aws.Config{
			// dynamodb-local
			Endpoint: aws.String("http://localhost:8000"),

			// static credentials
			Credentials: credentials.NewStaticCredentials("robert", "wow", "launchpad"),

			// arbitrary region
			Region: aws.String("us-west-2"),
		}
	} else {
		// todo: more granular production aws setup
		cfg = aws.NewConfig()
	}

	if os.Getenv(AWSDebugEnv) == "true" {
		// log everything
		cfg.LogLevel = aws.LogLevel(aws.LogDebug)
	} else {
		// by default, log error events
		cfg.LogLevel = aws.LogLevel(aws.LogDebugWithRequestErrors)
	}

	// assign logger
	if len(logger) > 0 {
		var l = logger[0]
		cfg.Logger = aws.LoggerFunc(func(args ...interface{}) { l.Info(args...) })
		cfg.Logger.Log("aws logger initialized")
	}

	return cfg
}

// AWSSession initializes an AWS API session with given configs
func AWSSession(cfg ...*aws.Config) (*session.Session, error) {
	return session.NewSession(cfg...)
}
