package cmd

import (
	"github.com/ubclaunchpad/pinpoint/cmdlib"
)

// CoreCommand is the CLI for pinpoint-core
type CoreCommand struct {
	*cmdlib.Command
}

// New creates a new CoreCommand
func New(version string) *CoreCommand {
	app := &CoreCommand{
		Command: cmdlib.New("pinpoint-core",
			"Pinpoint's core service",
			``,
			version,
			"9111"),
	}

	// register commands
	app.AddCommand(
		app.getRunCommand(),
	)

	return app
}
