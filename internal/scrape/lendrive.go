package scrape

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

func Lendrive(link string) (out ScrapeResponse) {
	c := colly.NewCollector()
	out.Url = link

	// get title
	c.OnHTML("h1.entry-title", func(h *colly.HTMLElement) {
		out.Title = h.Text

	})

	// get description
	c.OnHTML("div.desc", func(h *colly.HTMLElement) {
		desc := strings.ReplaceAll(h.Text, "\t", "")
		desc = strings.ReplaceAll(desc, "\n", "")
		out.Description = desc
	})

	// get urldownload
	c.OnHTML("div.soraddlx", func(h *colly.HTMLElement) {
		if strings.Contains(h.DOM.Find("div.sorattlx").Text(), "Subtitle") {
			return
		}

		h.DOM.Find("div.soraurlx").Each(func(_ int, s *goquery.Selection) {
			resolution := strings.Split(s.Find("strong").Text(), "|")[0]
			resolution = strings.TrimSpace(resolution)

			s.Find("a").Each(func(_ int, s *goquery.Selection) {
				out.Downloads = append(out.Downloads, Download{
					Codec:       "x265",
					Resolution:  resolution,
					FileHosting: strings.ToLower(s.Text()),
					UrlDownload: s.AttrOr("href", ""),
				})
			})
		})
	})

	c.Visit(link)
	return
}
