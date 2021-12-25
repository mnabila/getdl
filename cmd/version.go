package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Long:  "Print version information and exit ",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("Version :", Version)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
