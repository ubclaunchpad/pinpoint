package service

import (
	"context"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/ubclaunchpad/pinpoint/core/database"
	"github.com/ubclaunchpad/pinpoint/core/mailer"
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
	l    *zap.SugaredLogger
	db   *database.Database
	grpc *grpc.Server
}

// Opts declares configuration for the core service
type Opts struct {
	TLSOpts
}

// TLSOpts defines TLS configuration
type TLSOpts struct {
	CertFile string
	KeyFile  string
}

// New creates a new Service
func New(awsConfig client.ConfigProvider, logger *zap.SugaredLogger, opts Opts) (*Service, error) {
	// set up database
	db, err := database.New(awsConfig, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to init database: %s", err.Error())
	}

	// create service
	s := &Service{
		l:  logger.Named("service"),
		db: db,
	}

	// set up logging params
	serverOpts := make([]grpc.ServerOption, 0)
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
		s.l.Info("setting up TLS")
		creds, err := credentials.NewServerTLSFromFile(opts.TLSOpts.CertFile, opts.TLSOpts.KeyFile)
		if err != nil {
			return nil, fmt.Errorf("could not load TLS keys: %s", err)
		}
		serverOpts = append(serverOpts, grpc.Creds(creds))
	}

	// create server
	s.grpc = grpc.NewServer(serverOpts...)
	pinpoint.RegisterCoreServer(s.grpc, s)

	// create service
	return s, nil
}

// Run starts up the service and blocks until exit
func (s *Service) Run(host, port string) error {
	// let's gooooo
	s.l.Infow("spinning up core service",
		"core.host", host,
		"core.port", port)
	listener, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		return err
	}
	if err = s.grpc.Serve(listener); err != nil {
		s.l.Errorf("error encountered - service stopped",
			"error", err)
		return err
	}

	// report shutdown
	s.l.Info("service shut down")
	return nil
}

// Stop releases resources and shuts down the service
func (s *Service) Stop() {
	if s.grpc != nil {
		s.grpc.GracefulStop()
	}
}

// GetStatus retrieves status of service
func (s *Service) GetStatus(ctx context.Context, req *request.Status) (*response.Status, error) {
	res := &response.Status{Callback: req.Callback}
	if req.Callback == "I don't like launch pad" {
		return res, errors.New("launch pad is the best and you know it")
	}
	return res, nil
}

// Verify sends an email verification email
func (s *Service) Verify(ctx context.Context, req *request.Verify) (*response.Bool, error) {
	res := &response.Bool{Bool: false}

	hash, err := mailer.NewVerifier(req.Email).Init()
	if err != nil {
		return res, err
	}

	// Construct verification email
	// TODO: Change to get email address from user session
	mailer, err := mailer.NewMailer(req.Email, "Title", hash)
	if err != nil {
		return res, err
	}

	// Send email
	if err := mailer.Send(); err != nil {
		return res, err
	}

	// If no error, respond true. TODO: Change this to utilize response codes
	res = &response.Bool{Bool: true}
	return res, nil
}
