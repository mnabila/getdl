package getdl

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/mnabila/getdl/entities"
)

func Oploverz(link string) (out entities.ScrapeResponse) {
	c := colly.NewCollector()
	out.Url = link

	// get title
	c.OnHTML("h1.entry-title", func(h *colly.HTMLElement) {
		out.Title = h.Text
	})

	// get description
	c.OnHTML("div.desc", func(h *colly.HTMLElement) {
		desc := h.ChildText("div.entry-content-single")
		out.Description = strings.TrimSpace(desc)
	})

	// get download links
	c.OnHTML("div.links_table table", func(h *colly.HTMLElement) {
		h.DOM.Find("tr").Each(func(_ int, s *goquery.Selection) {
			format := strings.Split(s.Find("td strong.quality").Text(), " ")
			out.Downloads = append(out.Downloads, entities.Download{
				Codec:       format[len(format)-1],
				Resolution:  format[0],
				FileHosting: strings.ToLower(s.Find("td b").First().Text()),
				UrlDownload: s.Find("td a").AttrOr("href", ""),
			})
		})
	})

	c.Visit(link)

	return
}
