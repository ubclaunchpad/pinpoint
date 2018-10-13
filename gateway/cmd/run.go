package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ubclaunchpad/pinpoint/gateway/api"
	"github.com/ubclaunchpad/pinpoint/libcmd"
	"google.golang.org/grpc"
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
	run.Flags().String("ssl.cert", "", "ssl certificate")
	run.Flags().String("ssl.key", "", "ssl key")

	// set run
	run.RunE = runCommand(g)

	return run
}

func runCommand(g *GatewayCommand) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		// retrieve flags
		flags, err := libcmd.GetStringFlags(cmd,
			"core.host", "core.port", "ssl.cert", "ssl.key")
		if err != nil {
			return err
		}

		// Set up api
		a, err := api.New(g.SugaredLogger)
		if err != nil {
			return fmt.Errorf("failed to create app: %s", err.Error())
		}

		// set connection options
		grpcOpts := make([]grpc.DialOption, 0)
		if g.Dev {
			grpcOpts = append(grpcOpts, grpc.WithInsecure())
		}

		// Let's go!
		if err = a.Run(g.Host, g.Port, api.RunOpts{
			SSLOpts: api.SSLOpts{
				CertFile: flags["ssl.cert"],
				KeyFile:  flags["ssl.key"],
			},
			CoreOpts: api.CoreOpts{
				Host:        flags["core.host"],
				Port:        flags["core.port"],
				DialOptions: grpcOpts,
			},
		}); err != nil {
			return err
		}
		return nil
	}
}
