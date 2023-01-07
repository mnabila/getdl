package scrape

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

func Samehadaku(link string) (out ScrapeResponse) {
	c := colly.NewCollector()
	out.Url = link

	// get title
	c.OnHTML("h1.entry-title", func(h *colly.HTMLElement) {
		out.Title = h.Text

	})

	// get description
	c.OnHTML("div.desc div", func(h *colly.HTMLElement) {
		out.Description = strings.Replace(h.Text, "\n", "", 1)
	})

	// get download links
	c.OnHTML("div.download-eps", func(h *colly.HTMLElement) {
		h.DOM.Each(func(_ int, s *goquery.Selection) {
			codec := s.Find("p").Text()
			if codec == "mkv" {
				codec = "x264"
			}

			s.Find("li").Each(func(_ int, s *goquery.Selection) {
				resolution := s.Find("strong").Text()
				s.Find("a").Each(func(_ int, s *goquery.Selection) {
					out.Downloads = append(out.Downloads, Download{
						Codec:       codec,
						Resolution:  resolution,
						FileHosting: strings.ToLower(s.Text()),
						UrlDownload: s.AttrOr("href", ""),
					})
				})
			})
		})
	})
	c.Visit(link)
	return
}
