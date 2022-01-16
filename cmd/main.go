package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Short: "getdl",
	Long: `Get download link from some website
-------------------------
Current Support:
1. Oploverz
2. Samehadaku
3. Doronime
4. Lendrive
-------------------------`,
}

func init() {
	RootCmd.CompletionOptions.DisableDefaultCmd = true
}
