package libcmd

import "github.com/spf13/cobra"

// GetStringFlags maps given flags into a map
func GetStringFlags(cmd *cobra.Command, flags ...string) (vals map[string]string, err error) {
	vals = make(map[string]string)
	for _, f := range flags {
		vals[f], err = cmd.Flags().GetString(f)
		if err != nil {
			return
		}
	}
	return
}
