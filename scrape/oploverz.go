package scrape

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

func Oploverz(link string) Response {
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

	// get download links
	c.OnHTML("div.soraddlx", func(h *colly.HTMLElement) {
		h.DOM.Each(func(_ int, s *goquery.Selection) {
			ld := ListDownload{}
			ld.Type = s.Find("h3").Text()

			s.Find("div.soraurlx").Each(func(_ int, s *goquery.Selection) {

				d := Download{}
				d.Resolution = s.Find("strong").Text()

				s.Find("a").Each(func(_ int, s *goquery.Selection) {

					d.Links = append(d.Links, Links{strings.ToLower(s.Text()), s.AttrOr("href", "")})
				})

				ld.Downloads = append(ld.Downloads, d)
			})
			result.Downloads = append(result.Downloads, ld)

		})
	})
	c.Visit(link)
	return result
}
