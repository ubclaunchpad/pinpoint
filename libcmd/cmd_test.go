package libcmd

import (
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/ubclaunchpad/pinpoint/utils"
	"go.uber.org/zap"
)

func TestNew(t *testing.T) {
	type args struct {
		name  string
		flags map[string]string
	}
	tests := []struct {
		name string
		args args
	}{
		{"basic", args{"pinpoint", nil}},
		{"dev enabled", args{"pinpoint", map[string]string{"dev": ""}}},
		{"host/port flags set", args{"pinpoint", map[string]string{"host": "1234", "port": "5678"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := New(tt.args.name, "", "", "", "")
			if app.Use != tt.args.name {
				t.Errorf("use: expected %s, got %s", tt.args.name, app.Use)
				return
			}

			// set mock input
			args := []string{"fakecommand"}
			if tt.args.flags != nil {
				for f, v := range tt.args.flags {
					args = append(args, "--"+f, v)
				}
			}
			app.SetArgs(args)

			// run prerun setup by attaching a mock subcommand and executing it
			app.AddCommand(&cobra.Command{Use: "fakecommand", Run: func(cmd *cobra.Command, args []string) {}})
			if err := app.Execute(); err != nil {
				t.Errorf("unexpected error: %s", err.Error())
				return
			}

			// Check for properties
			if app.SugaredLogger == nil {
				t.Error("no logger found")
			}
			if tt.args.flags != nil {
				if _, ok := tt.args.flags["dev"]; ok && app.Dev == (os.Getenv("mode") == "development") {
					t.Errorf("expected dev to be %v", !(os.Getenv("mode") == "development"))
				}
				if _, ok := tt.args.flags["host"]; ok && app.Host == "" {
					t.Error("expected app.Host to be set")
				}
				if _, ok := tt.args.flags["port"]; ok && app.Port == "" {
					t.Error("expected app.Port to be set")
				}
			}
		})
	}
}

func TestCommand_Sync(t *testing.T) {
	log, err := utils.NewLogger(true, "")
	if err != nil {
		t.Errorf("unexpected error: %s", err.Error())
		return
	}
	type fields struct {
		SugaredLogger *zap.SugaredLogger
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"logger should not throw", fields{log}},
		{"nil logger should not throw", fields{nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Command{SugaredLogger: tt.fields.SugaredLogger}
			c.Sync()
		})
	}
}
