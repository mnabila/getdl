package scrape

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

func Samehadaku(link string) Response {
	c := colly.NewCollector()
	result := Response{}
	result.Url = link

	// get title
	c.OnHTML("h1.entry-title", func(h *colly.HTMLElement) {
		result.Title = h.Text

	})

	// get description
	c.OnHTML("div.desc div", func(h *colly.HTMLElement) {
		result.Description = strings.Replace(h.Text, "\n", "", 1)
	})

	// get download links
	c.OnHTML("div.download-eps", func(h *colly.HTMLElement) {
		h.DOM.Each(func(_ int, s *goquery.Selection) {
			ld := ListDownload{}
			ld.Codec = s.Find("p").Text()

			s.Find("li").Each(func(_ int, s *goquery.Selection) {

				d := Download{}
				d.Resolution = s.Find("strong").Text()

				s.Find("a").Each(func(_ int, s *goquery.Selection) {

					d.Links = append(d.Links, FileHosting{strings.ToLower(s.Text()), s.AttrOr("href", "")})
				})

				ld.Downloads = append(ld.Downloads, d)
			})
			result.Downloads = append(result.Downloads, ld)

		})
	})
	c.Visit(link)
	return result
}
