package service

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/ubclaunchpad/pinpoint/core/crypto"
	"github.com/ubclaunchpad/pinpoint/core/database"
	"github.com/ubclaunchpad/pinpoint/core/mailer"
	"github.com/ubclaunchpad/pinpoint/core/verifier"
	pinpoint "github.com/ubclaunchpad/pinpoint/protobuf"
	"github.com/ubclaunchpad/pinpoint/protobuf/models"
	"github.com/ubclaunchpad/pinpoint/protobuf/request"
	"github.com/ubclaunchpad/pinpoint/protobuf/response"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"

	"google.golang.org/grpc/metadata"
)

// Service provides core application service functionality. It handles most
// logic and connections to various backends. It implements an gRPC interface.
type Service struct {
	l    *zap.SugaredLogger
	grpc *grpc.Server
	mail *mailer.Mailer
	db   database.DBClient

	gateway GatewayOpts
}

// Opts declares configuration for the core service
type Opts struct {
	Token string
	TLSOpts
	GatewayOpts
}

// TLSOpts defines TLS configuration
type TLSOpts struct {
	CertFile string
	KeyFile  string
}

// GatewayOpts declares gateway configuration
type GatewayOpts struct {
	Token string
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
		l:       logger.Named("service"),
		db:      db,
		gateway: opts.GatewayOpts,
	}

	// Construct verification email - TODO: remove from ENV
	s.mail, err = mailer.New(os.Getenv("MAILER_USER"), os.Getenv("MAILER_PASS"))
	if err != nil {
		s.l.Warnw("failed to instantiate mailer",
			"user", os.Getenv("MAILER_USER"))
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

	// set up interceptors
	authUnaryInterceptor, authStreamingInterceptor := newAuthInterceptors(opts.Token)

	// instantiate server options
	serverOpts = append(serverOpts,
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.UnaryServerInterceptor(grpcLogger, zapOpts...),
			authUnaryInterceptor),
		grpc_middleware.WithStreamServerChain(
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.StreamServerInterceptor(grpcLogger, zapOpts...),
			authStreamingInterceptor))

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
	return &response.Status{}, nil
}

// Handshake generates response to a Ping request (For Initial Auth Purpose)
func (s *Service) Handshake(ctx context.Context, req *request.Empty) (*response.Empty, error) {
	s.l.Info("received handshake request from gateway")
	grpc.SendHeader(ctx, metadata.New(map[string]string{
		"authorization": s.gateway.Token,
	}))
	return &response.Empty{}, nil
}

// CreateAccount registers a user and sends an email verification email
func (s *Service) CreateAccount(ctx context.Context, req *request.CreateAccount) (*response.Message, error) {
	if err := crypto.ValidateCredentialValues(req.Email, req.Password); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "unable to validate credentials: %s", err.Error())
	}

	// Generate password salt
	salt, err := crypto.HashAndSalt(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to salt password: %s", err.Error())
	}

	// set up verification code for user
	v := verifier.New(req.Email, s.mail)

	// create user
	if err := s.db.AddNewUser(
		&models.User{Email: req.Email, Name: req.Name, Hash: salt},
		&models.EmailVerification{Email: req.Email, Hash: v.Hash, Expiry: v.Expiry},
	); err != nil {
		return nil, fmt.Errorf("failed to create user: %s", err.Error())
	}

	// send verification email
	if err = v.SendVerification(); err != nil {
		return nil, fmt.Errorf("failed to create email verification: %s", err.Error())
	}

	// If no error, respond success
	return &response.Message{
		Message: "user successfully created - an email verification was sent",
	}, nil
}

// Verify looks up the given hash, and verifies the hash matching email
func (s *Service) Verify(ctx context.Context, req *request.Verify) (*response.Message, error) {
	v, err := s.db.GetEmailVerification(req.GetEmail(), req.GetHash())
	if err != nil {
		return nil, err
	}
	return &response.Message{Message: "successfully verified " + v.Email}, nil
}

// Login looks up the given email and password and attempts to validate the user
func (s *Service) Login(ctx context.Context, req *request.Login) (*response.Message, error) {
	user, err := s.db.GetUser(req.GetEmail())
	if err != nil {
		return nil, fmt.Errorf("failed to authenticate user: %s", err.Error())
	}
	if crypto.ComparePasswords(user.Hash, req.GetPassword()) {
		return &response.Message{Message: "user successfully logged in"}, nil
	}
	return nil, errors.New("user not authenticated")
}
