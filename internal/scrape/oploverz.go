package scrape

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

func Oploverz(link string) (out ScrapeResponse) {
	c := colly.NewCollector()
	out.Url = link

	// get title
	c.OnHTML("h1.entry-title", func(h *colly.HTMLElement) {
		out.Title = h.Text

	})

	// get description
	c.OnHTML("div.desc", func(h *colly.HTMLElement) {
		desc := strings.ReplaceAll(h.Text, "\t", "")
		out.Description = strings.ReplaceAll(desc, "\n", "")
	})

	// get download links
	c.OnHTML("div.soraddlx", func(h *colly.HTMLElement) {
		h.DOM.Each(func(_ int, s *goquery.Selection) {
			codec := s.Find("h3").Text()
			if codec == "mkv" {
				codec = "x264"
			}

			s.Find("div.soraurlx").Each(func(_ int, s *goquery.Selection) {
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
