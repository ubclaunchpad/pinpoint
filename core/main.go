package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/ubclaunchpad/pinpoint/core/service"
	pinpoint "github.com/ubclaunchpad/pinpoint/grpc"
	"github.com/ubclaunchpad/pinpoint/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
)

var (
	dev = os.Getenv("MODE") == "development"
)

func main() {
	// Set up logger
	logger, err := utils.NewLogger(dev)
	if err != nil {
		fmt.Printf("failed to init logger: %s", err.Error())
		os.Exit(1)
	}
	defer logger.Sync()

	// Set up AWS credentials
	awsConfig, err := utils.AWSSession(utils.AWSConfig(dev))
	if err != nil {
		logger.Fatalw("failed to connect to aws",
			"error", err.Error())
	}

	// Set up service
	core, err := service.New(awsConfig, logger)
	if err != nil {
		logger.Fatalw("failed to create service",
			"error", err.Error())
	}

	// Make sure that log statements internal to gRPC library are logged
	grpcLogger := logger.Desugar().Named("grpc")
	grpc_zap.ReplaceGrpcLogger(grpcLogger)

	// Create a server, make sure we put the grpc_ctxtags context before everything else
	opts := []grpc_zap.Option{
		grpc_zap.WithDurationField(func(duration time.Duration) zapcore.Field {
			return zap.Duration("grpc.duration", duration)
		}),
	}
	grpcServer := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.UnaryServerInterceptor(grpcLogger, opts...)),
		grpc_middleware.WithStreamServerChain(
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.StreamServerInterceptor(grpcLogger, opts...)))

	// Set address from environment
	addr := os.Getenv("CORE_HOST") + ":" + os.Getenv("CORE_PORT")

	// Prep grpc server
	pinpoint.RegisterPinpointCoreServer(grpcServer, core)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Fatalw("failed to start core service",
			"error", err.Error())
	}
	logger.Infow("spinning up core service",
		"host", os.Getenv("CORE_HOST"),
		"core", os.Getenv("CORE_PORT"))

	// Serve and block until exit
	if err = grpcServer.Serve(listener); err != nil {
		logger.Fatalw("error encountered",
			"error", err.Error())
	}

	logger.Info("core service shut down")
}
