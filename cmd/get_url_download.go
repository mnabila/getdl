package cmd

import (
	"encoding/json"
	"fmt"
	"getdl/config"
	"getdl/internal/scrape"
	"net/url"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var GetUrlDownload = &cobra.Command{
	Use:   "get",
	Short: "Get url download",
	Long:  "Get url download from spesific website",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 || len(args) >= 2 {
			cmd.Help()
			return
		}

		conf, err := config.ReadConfig()
		if err != nil {
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

func getResponse(u *url.URL, website []config.WebConfig) (response scrape.ScrapeResponse, web config.WebConfig) {
	for _, web := range website {
		for _, domain := range web.Domain {
			if u.Host == domain {
				switch web.Name {
				case "oploverz":
					return scrape.Oploverz(u.String()), web
				case "doronime":
					return scrape.Doronime(u.String()), web
				case "samehadaku":
					return scrape.Samehadaku(u.String()), web
				case "lendrive":
					return scrape.Lendrive(u.String()), web
				case "animekompi":
					return scrape.Lendrive(u.String()), web
				}
			}
		}
	}
	return
}

func getResult(data scrape.ScrapeResponse, conf config.WebConfig) string {
	for _, download := range data.Downloads {
		if download.Codec == conf.Codec {
			if strings.Contains(download.Resolution, conf.Resolution) {
				for _, hosting := range conf.FileHosting {
					if strings.Contains(download.FileHosting, hosting) {
						return download.UrlDownload
					}
				}
			}
		}
	}
	return ""
}

func init() {
	GetUrlDownload.PersistentFlags().Bool("raw", false, "Print raw response in json")
}
