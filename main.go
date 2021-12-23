package main

import (
	"fmt"
	"getdl/config"
	"getdl/scrape"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

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

// func getDummyResponse() scrape.Response {
// 	var result scrape.Response
// 	file, _ := ioutil.ReadFile("result.json")
// 	if err := json.Unmarshal(file, &result); err != nil {
// 		panic(err.Error())
// 	}
// 	return result
// }

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

func main() {
	urlDownload := os.Args[1]
	conf := config.ReadConfig()
	u, err := url.Parse(urlDownload)
	if err != nil {
		fmt.Println("Url Tidak Valid")
		return
	}

	result := getUrlFile(getResponse(u).Downloads)

	fmt.Println("[ Open URL ] >> ", urlDownload)
	fmt.Println("[ Result   ] >> ", result)

	err = exec.Command(conf.Browser, result).Run()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
