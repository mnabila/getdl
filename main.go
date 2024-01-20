package main

import (
	"fmt"
	"runtime"

	"github.com/mnabila/getdl/cmd"
	"github.com/mnabila/getdl/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Version string

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if runtime.GOOS == "linux" {
		viper.AddConfigPath("$HOME/.config/getdl")
	} else {
		viper.AddConfigPath(".")
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if err := config.SetupDefaultConfig(); err != nil {
				fmt.Println(err.Error())
				return
			}
		}

		fmt.Println(err.Error())
		return
	}
}

func main() {
	rootCmd := &cobra.Command{
		Use: "getdl",
	}

	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print version",
		Long:  "Print version information and exit ",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println("Version :", Version)
		},
	})

	rootCmd.AddCommand(cmd.UseGetUrlDownload)
	rootCmd.Execute()
}
