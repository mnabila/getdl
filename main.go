package main

import (
	"fmt"
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

func getZippyshare(data []scrape.ListDownload) string {
	var result string
	for _, d := range data {
		if strings.Contains(d.Type, "265") {
			for _, urlfile := range d.Downloads {
				if strings.Contains(urlfile.Resolution, "720") {
					for _, u := range urlfile.Links {
						if strings.Contains(u.Label, "zippy") {
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
	u, err := url.Parse(urlDownload)
	if err != nil {
		fmt.Println("Url Tidak Valid")
		return
	}

	result := getResponse(u)
	zippy := getZippyshare(result.Downloads)

	fmt.Println("[ Open URL ] >> ", urlDownload)
	fmt.Println("[ Result   ] >> ", zippy)

	err = exec.Command("xdg-open", zippy).Run()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
