package service

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/ubclaunchpad/pinpoint/core/database"
	"github.com/ubclaunchpad/pinpoint/grpc/request"
	"github.com/ubclaunchpad/pinpoint/grpc/response"

	"go.uber.org/zap"
)

// Service provides core application service functionality. It handles most
// logic and connections to various backends. It implements an gRPC interface.
type Service struct {
	l  *zap.SugaredLogger
	db *database.Database
}

// New creates a new Service
func New(awsConfig client.ConfigProvider, logger *zap.SugaredLogger) (*Service, error) {
	// set up database
	db, err := database.New(awsConfig, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to init database: %s", err.Error())
	}

	// create service
	return &Service{
		l:  logger.Named("service"),
		db: db,
	}, nil
}

// GetStatus retrieves status of service
func (s *Service) GetStatus(ctx context.Context, req *request.Status) (*response.Status, error) {
	return &response.Status{Callback: req.Callback}, nil
}
