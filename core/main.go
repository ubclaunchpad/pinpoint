package main

import (
	"fmt"
	"net"
	"os"

	"github.com/ubclaunchpad/pinpoint/core/service"
	pinpoint "github.com/ubclaunchpad/pinpoint/grpc"
	"github.com/ubclaunchpad/pinpoint/utils"
	"google.golang.org/grpc"
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

	// set up rpc server
	var (
		s    = grpc.NewServer()
		addr = os.Getenv("CORE_HOST") + ":" + os.Getenv("CORE_PORT")
	)

	// spin up rpc server
	pinpoint.RegisterPinpointCoreServer(s, core)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Fatalw("failed to start service",
			"error", err.Error())
	}
	logger.Infow("spinning up service",
		"host", os.Getenv("CORE_HOST"),
		"core", os.Getenv("CORE_PORT"))
	if err = s.Serve(listener); err != nil {
		logger.Fatalw("error encountered",
			"error", err.Error())
	}
	logger.Info("service shut down")
}
