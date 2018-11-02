package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/ubclaunchpad/pinpoint/core/cmd"
)

var (
	// Version defines the version of pinpoint-core
	Version string
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		println(err.Error())
	}

	app := cmd.New(Version)
	defer app.Sync()

	if err := app.Execute(); err != nil {
		println(err.Error())
		os.Exit(1)
	}
}
