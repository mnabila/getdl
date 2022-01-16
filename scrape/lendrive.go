package scrape

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

func Lendrive(link string) Response {
	c := colly.NewCollector()
	result := Response{}
	result.Url = link

	// get title
	c.OnHTML("h1.entry-title", func(h *colly.HTMLElement) {
		result.Title = h.Text

	})

	// get description
	c.OnHTML("div.desc", func(h *colly.HTMLElement) {
		desc := strings.ReplaceAll(h.Text, "\t", "")
		desc = strings.ReplaceAll(desc, "\n", "")
		result.Description = desc
	})

	// get urldownload
	c.OnHTML("div.soraddlx", func(h *colly.HTMLElement) {
		if strings.Contains(h.DOM.Find("div.sorattlx").Text(), "Subtitle") {
			return
		}

		ld := ListDownload{}
		ld.Codec = "x265"

		h.DOM.Find("div.soraurlx").Each(func(_ int, s *goquery.Selection) {
			d := Download{}
			res := strings.Split(s.Find("strong").Text(), "|")[0]
			d.Resolution = strings.TrimSpace(res)

			s.Find("a").Each(func(_ int, s *goquery.Selection) {

				d.Links = append(d.Links, FileHosting{strings.ToLower(s.Text()), s.AttrOr("href", "")})
			})
			ld.Downloads = append(ld.Downloads, d)
		})
		result.Downloads = append(result.Downloads, ld)
	})

	c.Visit(link)
	return result

}
