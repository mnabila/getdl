package scrape

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

type MirroredResult FileHosting

func Mirrored(link string) []MirroredResult {
	c := colly.NewCollector()
	result := []MirroredResult{}

	c.OnXML("//a[contains(@href, '://www.mirrored.to/files')]", func(x *colly.XMLElement) {
		fmt.Println(x.Attr("href"))

	})

	c.Visit(link)

	return result
}
