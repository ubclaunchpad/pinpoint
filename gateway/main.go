package main

import (
	"fmt"
	"os"

	"github.com/ubclaunchpad/pinpoint/gateway/api"
	"github.com/ubclaunchpad/pinpoint/utils"
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

	// Set up api
	a, err := api.New(logger, dev)
	if err != nil {
		logger.Fatalw("failed to create app",
			"error", err.Error())
	}

	// Let's go!
	if err = a.Run(os.Getenv("API_HOST"), os.Getenv("API_PORT"), api.RunOpts{
		SSLOpts: &api.SSLOpts{
			CertFile: os.Getenv("API_CERT"),
			KeyFile:  os.Getenv("API_KEY"),
		},
	}); err != nil {
		logger.Fatalw("service shut down",
			"error", err.Error())
	}
}
