package main

import (
	"getdl/cmd"

	"github.com/spf13/cobra"
)

func init() {
}

func main() {
	rootCmd := &cobra.Command{
		Short: "getdl",
		Long:  "Get download link from some website",
	}

	rootCmd.AddCommand(cmd.VersionCmd)
	rootCmd.AddCommand(cmd.GetUrlDownload)
	rootCmd.Execute()
}
