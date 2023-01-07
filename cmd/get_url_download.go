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

		website := args[0]

		u, err := url.Parse(website)
		if err != nil {
			fmt.Println("Url Tidak Valid")
		}
		data := getResponse(u, conf)

		raw, _ := cmd.Flags().GetBool("raw")
		if raw {
			rawJson, _ := json.Marshal(data)
			fmt.Println(string(rawJson))
			return
		}

		urlDownload := getUrlFile(data.Downloads, conf)

		fmt.Println("[ Open URL ] >> ", website)
		fmt.Println("[ Result   ] >> ", urlDownload)

		if urlDownload != "" && conf.OpenInBrowser {
			if exec.Command(conf.Browser, urlDownload).Run(); err != nil {
				fmt.Println(err.Error())
				return
			}
		}
	},
}

func getResponse(u *url.URL, c *config.Configuration) (result scrape.ScrapeResponse) {
	switch u.Host {
	case c.Domain.Oploverz:
		result = scrape.Oploverz(u.String())
	case c.Domain.Doronime:
		result = scrape.Doronime(u.String())
	case c.Domain.Samehadaku:
		result = scrape.Samehadaku(u.String())
	case c.Domain.Lendrive:
		result = scrape.Lendrive(u.String())
	}
	return result
}

func getUrlFile(downloads []scrape.Download, c *config.Configuration) (result string) {
	for _, d := range downloads {
		if strings.Contains(d.Codec, c.Codec) {
			if strings.Contains(d.Resolution, c.Resolution) {
				if strings.Contains(d.FileHosting, c.FileHosting) {
					return d.UrlDownload
				}
			}
		}
	}
	return
}

func init() {
	GetUrlDownload.PersistentFlags().Bool("raw", false, "Print raw response in json")
}
