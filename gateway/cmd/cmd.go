package cmd

import "github.com/ubclaunchpad/pinpoint/libcmd"

// GatewayCommand is the CLI for pinpoint-gateway
type GatewayCommand struct {
	*libcmd.Command
}

// New creates a new GatewayCommand
func New(version string) *GatewayCommand {
	app := &GatewayCommand{
		Command: libcmd.New("pinpoint-gateway",
			"Pinpoint's RESTful API gateway",
			``,
			version,
			"8081"),
	}

	// register commands
	app.AddCommand(
		app.getRunCommand(),
	)

	return app
}
