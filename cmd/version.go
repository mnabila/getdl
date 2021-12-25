package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version string
var Commit string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Long:  "Print version information and exit ",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("Version : %s (%s)\n", Version, Commit)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
