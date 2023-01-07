package scrape

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

func Animekompi(link string) (out ScrapeResponse) {
	c := colly.NewCollector()
	out.Url = link

	// get title
	c.OnHTML("div.infolimit", func(h *colly.HTMLElement) {
		out.Title = h.DOM.Find("h2").Text()
	})

	// get description
	c.OnHTML("div.desc", func(h *colly.HTMLElement) {
		desc := strings.ReplaceAll(h.Text, "\n", "")
		out.Description = strings.TrimSpace(desc)
	})

	// get download links
	c.OnHTML("div.soraddlx", func(h *colly.HTMLElement) {
		if strings.Contains(h.DOM.Find("div.sorattlx").Text(), "Subtitle") {
			return
		}

		h.DOM.Find("div.soraurlx").Each(func(_ int, s *goquery.Selection) {
			resolution := s.Find("strong").Text()

			s.Find("a").Each(func(_ int, s *goquery.Selection) {
				out.Downloads = append(out.Downloads, Download{
					Codec:       "x264",
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
