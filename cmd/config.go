package cmd

import (
	"fmt"
	"getdl/config"
	"os"
	"reflect"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var getConfigCmd = &cobra.Command{
	Use:   "get",
	Short: "Read Configuration",
	Long:  "Read configuration from  ~/.config/getdl",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 && len(args) < 2 {
			fmt.Println(config.GetConfig(args[0]))
			return
		}
		cmd.Help()
	},
}

var setConfigCmd = &cobra.Command{
	Use:   "set",
	Short: "Update Configuration",
	Long:  "Update configuration in ~/.config/getdl",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 && len(args) <= 2 {
			key := args[0]
			value := args[1]
			config.SetConfig(key, value)

			fmt.Println("Update Configuration Success Full")
			return
		}
		cmd.Help()
	},
}

var showConfigCmd = &cobra.Command{
	Use:   "show",
	Short: "Print Configuration",
	Long:  "Print configuration in ~/.config/getdl",
	Run: func(_ *cobra.Command, _ []string) {
		separator := strings.Repeat("-", 25)
		fmt.Println(separator)
		fmt.Println("Getdl Configuration")
		fmt.Println(separator)

		reflectValue := reflect.ValueOf(config.ReadConfig())
		reflectType := reflectValue.Type()
		writer := tabwriter.NewWriter(os.Stdout, 0, 4, 1, '\t', tabwriter.AlignRight)
		for i := 0; i < reflectType.NumField(); i++ {
			fmt.Fprintf(writer, "%s\t: %s\n", reflectType.Field(i).Name, reflectValue.Field(i).Interface())
		}
		writer.Flush()
		fmt.Println()
	},
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configuration",
	Long:  "Update or Read Configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	configCmd.AddCommand(getConfigCmd)
	configCmd.AddCommand(setConfigCmd)
	configCmd.AddCommand(showConfigCmd)
	RootCmd.AddCommand(configCmd)
}
