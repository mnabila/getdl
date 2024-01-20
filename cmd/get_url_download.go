package cmd

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os/exec"
	"strings"

	"github.com/mnabila/getdl/config"
	"github.com/mnabila/getdl/entities"
	"github.com/mnabila/getdl/pkg/getdl"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var UseGetUrlDownload = &cobra.Command{
	Use: "get [url]",
	Short: "Get download links from website",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 || len(args) >= 2 {
			cmd.Help()
			return
		}

		var conf config.Configurations
		if err := viper.Unmarshal(&conf); err != nil {
			fmt.Println(err.Error())
			return
		}

		u, err := url.Parse(args[0])
		if err != nil {
			fmt.Println("Url Tidak Valid")
		}

		response, webConfig := getResponse(u, conf.Website)

		raw, _ := cmd.Flags().GetBool("raw")
		if raw {
			rawJson, _ := json.Marshal(response)
			fmt.Println(string(rawJson))
			return
		}

		urlDownload := getResult(response, webConfig)
		fmt.Println("[ Open URL ] >> ", response.Url)
		fmt.Println("[ Result   ] >> ", urlDownload)

		if urlDownload != "" && conf.OpenInBrowser {
			if exec.Command(conf.Browser, urlDownload).Run(); err != nil {
				fmt.Println(err.Error())
				return
			}
		}
	},
}

func getResponse(u *url.URL, website []config.WebConfig) (response entities.ScrapeResponse, web config.WebConfig) {
	for _, web := range website {
		for _, domain := range web.Domain {
			if strings.Contains(u.Host, domain) {
				switch web.Name {
				case "oploverz":
					return getdl.Oploverz(u.String()), web
				case "doronime":
					return getdl.Doronime(u.String()), web
				case "samehadaku":
					return getdl.Samehadaku(u.String()), web
				case "lendrive":
					return getdl.Lendrive(u.String()), web
				case "animekompi":
					return getdl.Lendrive(u.String()), web
				}
			}
		}
	}
	return
}

func getResult(data entities.ScrapeResponse, conf config.WebConfig) string {
	for _, download := range data.Downloads {
		if download.Codec == conf.Codec {
			if strings.Contains(download.Resolution, conf.Resolution) {
				for _, hosting := range conf.FileHosting {
					if strings.Contains(download.FileHosting, strings.ToLower(hosting)) {
						return download.UrlDownload
					}
				}
			}
		}
	}
	return ""
}

func init() {
	UseGetUrlDownload.PersistentFlags().Bool("raw", false, "Print raw response in json")
}
