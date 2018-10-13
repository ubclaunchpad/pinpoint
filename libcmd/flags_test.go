package libcmd

import (
	"reflect"
	"testing"

	"github.com/spf13/cobra"
)

func TestGetStringFlags(t *testing.T) {
	type args struct {
		setFlags bool
		flags    []string
	}
	tests := []struct {
		name     string
		args     args
		wantVals map[string]string
		wantErr  bool
	}{
		{"with flags", args{true, []string{"host"}}, map[string]string{"host": "test"}, false},
		{"without flags or input", args{false, []string{}}, map[string]string{}, false},
		{"without flags but with input", args{false, []string{"host"}}, map[string]string{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// create mock command with flags, run tests in the Run command
			cmd := &cobra.Command{Use: "fakecommand", Run: func(cmd *cobra.Command, args []string) {
				gotVals, err := GetStringFlags(cmd, tt.args.flags...)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetStringFlags() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !tt.wantErr && !reflect.DeepEqual(gotVals, tt.wantVals) {
					t.Errorf("GetStringFlags() = %v, want %v", gotVals, tt.wantVals)
				}
			}}
			if tt.args.setFlags {
				cmd.PersistentFlags().String("host", "", "")
				cmd.PersistentFlags().String("port", "", "")
			}

			// set up test run
			args := make([]string, 2*len(tt.args.flags)+1)
			args[0] = "fakecommand"
			i := 1
			if tt.args.setFlags {
				for _, f := range tt.args.flags {
					args[i] = "--" + f
					args[i+i] = "test"
					i++
				}
			}
			cmd.SetArgs(args)

			// execute
			if err := cmd.Execute(); err != nil {
				t.Error(err)
			}
		})
	}
}
