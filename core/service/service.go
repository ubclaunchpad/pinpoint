package service

import (
	"context"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/ubclaunchpad/pinpoint/core/database"
	pinpoint "github.com/ubclaunchpad/pinpoint/protobuf"
	"github.com/ubclaunchpad/pinpoint/protobuf/request"
	"github.com/ubclaunchpad/pinpoint/protobuf/response"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
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

// RunOpts declares configuration for starting up core service
type RunOpts struct {
	TLSOpts
}

// TLSOpts defines TLS configuration
type TLSOpts struct {
	CertFile string
	KeyFile  string
}

// Run starts up the service and blocks until exit
func (s *Service) Run(host, port string, opts RunOpts) error {
	serverOpts := make([]grpc.ServerOption, 0)

	// set up logging params
	grpcLogger := s.l.Desugar().Named("grpc")
	grpc_zap.ReplaceGrpcLogger(grpcLogger)
	zapOpts := []grpc_zap.Option{
		grpc_zap.WithDurationField(func(duration time.Duration) zapcore.Field {
			return zap.Duration("grpc.duration", duration)
		}),
	}
	serverOpts = append(serverOpts,
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.UnaryServerInterceptor(grpcLogger, zapOpts...)),
		grpc_middleware.WithStreamServerChain(
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.StreamServerInterceptor(grpcLogger, zapOpts...)))

	// set up TLS credentials
	if opts.TLSOpts.CertFile != "" {
		creds, err := credentials.NewServerTLSFromFile(opts.TLSOpts.CertFile, opts.TLSOpts.KeyFile)
		if err != nil {
			return fmt.Errorf("could not load TLS keys: %s", err)
		}
		serverOpts = append(serverOpts, grpc.Creds(creds))
	}

	// create server
	grpcServer := grpc.NewServer(serverOpts...)
	pinpoint.RegisterCoreServer(grpcServer, s)

	// let's gooooo
	s.l.Infow("spinning up core service",
		"host", host,
		"core", port)
	listener, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		return err
	}
	if err = grpcServer.Serve(listener); err != nil {
		s.l.Errorf("error encountered - service stopped",
			"error", err)
		return err
	}

	// report shutdown
	s.l.Info("service shut down")
	return nil
}

// GetStatus retrieves status of service
func (s *Service) GetStatus(ctx context.Context, req *request.Status) (*response.Status, error) {
	res := &response.Status{Callback: req.Callback}
	if req.Callback == "I don't like launch pad" {
		return res, errors.New("launch pad is the best and you know it")
	}
	return res, nil
}
