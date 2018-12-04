package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/ubclaunchpad/pinpoint/core/service"
	"github.com/ubclaunchpad/pinpoint/libcmd"
	"github.com/ubclaunchpad/pinpoint/utils"
)

func (c *CoreCommand) getRunCommand() *cobra.Command {
	run := &cobra.Command{
		Use:   "run",
		Short: "Spin up service",
		Long:  ``,
	}

	// register flags
	run.Flags().String("tls.cert", "", "TLS certificate")
	run.Flags().String("tls.key", "", "TLS key")

	// set run command
	run.RunE = runCommand(c)

	return run
}

func runCommand(c *CoreCommand) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		// retrieve flags
		flags, err := libcmd.GetStringFlags(cmd,
			"tls.cert", "tls.key")
		if err != nil {
			return err
		}

		// Set up AWS credentials
		awsConfig, err := utils.AWSSession(utils.AWSConfig(
			c.Dev, c.SugaredLogger.Named("aws")))
		if err != nil {
			return fmt.Errorf("failed to connect to aws: %s", err.Error())
		}

		// Set up tokens
		coreToken := os.Getenv("PINPOINT_CORE_TOKEN")
		gatewayToken := os.Getenv("PINPOINT_GATEWAY_TOKEN")
		if c.Dev && coreToken == "" && gatewayToken == "" {
			coreToken = "valid_token"
			gatewayToken = "valid_token"
		}

		// Set up service
		core, err := service.New(awsConfig, c.SugaredLogger, service.Opts{
			Token: coreToken,
			TLSOpts: service.TLSOpts{
				CertFile: flags["tls.cert"],
				KeyFile:  flags["tls.key"],
			},
			GatewayOpts: service.GatewayOpts{
				Token: gatewayToken,
			},
		})
		if err != nil {
			return fmt.Errorf("failed to create service: %s", err.Error())
		}

		// handle graceful shutdown
		signals := make(chan os.Signal)
		signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-signals
			core.Stop()
			os.Exit(1)
		}()

		// Serve and block until exit
		if err = core.Run(c.Host, c.Port); err != nil {
			return err
		}

		return nil
	}
}
