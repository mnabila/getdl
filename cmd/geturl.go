package cmd

import (
	"fmt"
	"getdl/config"
	"getdl/scrape"
	"net/url"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

func getUrlFile(data []scrape.ListDownload) string {
	conf := config.ReadConfig()
	var result string
	for _, d := range data {
		if strings.Contains(d.Codec, conf.Codec) {
			for _, urlfile := range d.Downloads {
				if strings.Contains(urlfile.Resolution, conf.Resolution) {
					for _, u := range urlfile.Links {
						if strings.Contains(u.Label, conf.FileHosting) {
							result = u.Link
						}
					}
				}
			}
		}
	}
	return result
}

func getResponse(u *url.URL) scrape.Response {
	var result scrape.Response
	switch {
	case strings.Contains(u.Host, "oploverz"):
		result = scrape.Oploverz(u.String())
	case strings.Contains(u.Host, "doronime"):
		result = scrape.Doronime(u.String())
	case strings.Contains(u.Host, "194.163.183.129"):
		result = scrape.Samehadaku(u.String())
	}
	return result
}

var getUrlCmd = &cobra.Command{
	Use:   "get",
	Short: "Get url download",
	Long:  "Get url download from spesific website",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 || len(args) >= 2 {
			cmd.Help()
			return
		}
		urlWeb := args[0]
		conf := config.ReadConfig()

		u, err := url.Parse(urlWeb)
		if err != nil {
			fmt.Println("Url Tidak Valid")
			return
		}

		result := getUrlFile(getResponse(u).Downloads)

		fmt.Println("[ Open URL ] >> ", urlWeb)
		fmt.Println("[ Result   ] >> ", result)

		err = exec.Command(conf.Browser, result).Run()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

	},
}

func init() {
	RootCmd.AddCommand(getUrlCmd)
}
