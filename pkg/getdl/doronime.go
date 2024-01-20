package getdl

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"github.com/mnabila/getdl/entities"
)

func Doronime(link string) (out entities.ScrapeResponse) {
	c := colly.NewCollector()
	out.Url = link

	// get title
	c.OnHTML("h5.Content__title", func(h *colly.HTMLElement) {
		out.Title = h.Text
	})

	// get description
	c.OnHTML("div.Content__description-caption-synopsis", func(h *colly.HTMLElement) {
		p := h.DOM.Find("p")
		out.Description = strings.ReplaceAll(p.Text(), "\n", "")
	})

	// get download links
	c.OnHTML("div.Content__link", func(h *colly.HTMLElement) {
		container := h.DOM.Find("div.Download__container")
		// linsk x264
		container.First().Find("div.Download__group").Each(func(_ int, s *goquery.Selection) {
			resolution := strings.ToLower(s.Find("div.Download__group-title").Text())
			s.Find("div.Download__link").Each(func(_ int, s *goquery.Selection) {
				s.Find("a").Each(func(_ int, s *goquery.Selection) {
					download := entities.Download{
						Codec:       "x264",
						Resolution:  resolution,
						FileHosting: strings.ToLower(s.Find("span").First().Text()),
						UrlDownload: s.AttrOr("href", ""),
					}
					out.Downloads = append(out.Downloads, download)
				})
			})
		})

		// linsk x265
		container.Last().Find("div.Download__group").Each(func(_ int, s *goquery.Selection) {
			resolution := strings.ToLower(s.Find("div.Download__group-title").Text())
			s.Find("div.Download__link").Each(func(_ int, s *goquery.Selection) {
				s.Find("a").Each(func(_ int, s *goquery.Selection) {
					out.Downloads = append(out.Downloads, entities.Download{
						Codec:       "x265",
						Resolution:  resolution,
						FileHosting: strings.ToLower(s.Find("span").First().Text()),
						UrlDownload: s.AttrOr("href", ""),
					})
				})
			})
		})
	})

	c.Visit(link)

	return
}
