package libcmd

import (
	"fmt"
	"os"

	"go.uber.org/zap"

	"github.com/spf13/cobra"
	"github.com/ubclaunchpad/pinpoint/utils"
)

// Command is the default command-line app for Pinpoint components
type Command struct {
	*zap.SugaredLogger
	cobra.Command
	Host string
	Port string
	Dev  bool
}

// New creates a new CLI app and logger
func New(name, short, long, version, defaultPort string) *Command {
	// initialize app
	app := Command{
		Command: cobra.Command{
			Use:     name,
			Short:   short,
			Long:    long,
			Version: version,
		},
	}

	// register global flags
	app.PersistentFlags().String("host", "127.0.0.1", "service host")
	app.PersistentFlags().String("port", defaultPort, "service port")
	app.PersistentFlags().Bool("dev", os.Getenv("MODE") == "development", "toggle dev mode")

	// set flag and initial config loader
	app.PersistentPreRunE = func(cmd *cobra.Command, args []string) (err error) {
		// set flags
		strFlags, err := GetStringFlags(cmd, "host", "port")
		if err != nil {
			return
		}
		app.Host = strFlags["host"]
		app.Port = strFlags["port"]
		if app.Dev, err = cmd.Flags().GetBool("dev"); err != nil {
			return
		}

		// set logger - TODO: allow more granular config
		if app.SugaredLogger, err = utils.NewLogger(app.Dev); err != nil {
			return fmt.Errorf("failed to init logger: %s", err.Error())
		}

		// report mode
		if app.Dev {
			println("WARNING: dev mode enabled")
		}

		return nil
	}

	return &app
}

// Sync safely calls logger sync
func (c *Command) Sync() error {
	if c.SugaredLogger != nil {
		return c.SugaredLogger.Sync()
	}
	return nil
}
