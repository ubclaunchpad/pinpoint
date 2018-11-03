package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/ubclaunchpad/pinpoint/gateway/api"
	"github.com/ubclaunchpad/pinpoint/libcmd"
)

func (g *GatewayCommand) getRunCommand() *cobra.Command {
	run := &cobra.Command{
		Use:   "run",
		Short: "Spin up service",
		Long:  ``,
	}

	// register flags
	run.Flags().String("core.host", "127.0.0.1", "pinpoint-core host")
	run.Flags().String("core.port", "9111", "pinpoint-core host")
	run.Flags().String("core.cert", "", "pinpoint-core TLS certificate")
	run.Flags().String("tls.cert", "", "gateway TLS certificate")
	run.Flags().String("tls.key", "", "gateway TLS key")

	// set run
	run.RunE = runCommand(g)

	return run
}

func runCommand(g *GatewayCommand) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		// retrieve flags
		flags, err := libcmd.GetStringFlags(cmd,
			"core.host", "core.port", "core.cert", "tls.cert", "tls.key")
		if err != nil {
			return err
		}

		// Set up api
		a, err := api.New(g.SugaredLogger, api.CoreOpts{
			Host:     flags["core.host"],
			Port:     flags["core.port"],
			CertFile: flags["core.cert"],
		})
		if err != nil {
			return fmt.Errorf("failed to create app: %s", err.Error())
		}

		// handle graceful shutdown
		signals := make(chan os.Signal)
		signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-signals
			a.Stop()
			os.Exit(1)
		}()

		// Let's go!
		if err = a.Run(g.Host, g.Port, api.RunOpts{
			CertFile: flags["tls.cert"],
			KeyFile:  flags["tls.key"],
		},
		); err != nil {
			return err
		}
		return nil
	}
}
