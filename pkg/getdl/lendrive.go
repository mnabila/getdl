package getdl

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/mnabila/getdl/entities"
)

func Lendrive(link string) (out entities.ScrapeResponse) {
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
		out.Description = strings.TrimSpace(desc)
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
				out.Downloads = append(out.Downloads, entities.Download{
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
