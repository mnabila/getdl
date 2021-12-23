package scrape

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
)

// func sendPost(urlPost string, data url.Values, csrfToken string) {
// 	client := &http.Client{}
// 	req, err := http.NewRequest(http.MethodPost, urlPost, strings.NewReader(data.Encode()))
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	req.Header.Add("X-CSRF-TOKEN", csrfToken)
//
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	defer resp.Body.Close()
// 	result, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
//
// 	fmt.Println(string(result))
//
// }
//
// func BypassEgao(link string) string {
// 	const linkShortlink = "https://egao.in/safelink"
// 	var csrfToken string
// 	var dataId string
// 	var cookies []*http.Cookie
// 	c := colly.NewCollector()
// 	c.OnXML("/html/head/meta[@name=\"csrf-token\"]", func(x *colly.XMLElement) {
// 		csrfToken = x.Attr("content")
// 	})
//
// 	c.OnHTML("button#SafelinkChecker", func(h *colly.HTMLElement) {
// 		dataId = h.DOM.AttrOr("data-id", "")
// 	})
//
// 	c.OnResponse(func(r *colly.Response) {
//
// 		cookies = c.Cookies(r.Request.URL.String())
// 	})
//
// 	c.Visit(link)
//
// 	data := url.Values{}
// 	data.Add("id", dataId)
//
// 	return "adu"
// }


func Doronime(link string) Response {
	c := colly.NewCollector()
	result := Response{}
	result.Url = link

	// get title
	c.OnHTML("h5.Content__title", func(h *colly.HTMLElement) {
		result.Title = h.Text

	})

	// get description
	c.OnHTML("div.Content__description-caption-synopsis", func(h *colly.HTMLElement) {
		p := h.DOM.Find("p")
		result.Description = strings.ReplaceAll(p.Text(), "\n", "")
	})

	// get download links
	c.OnHTML("div.Content__link", func(h *colly.HTMLElement) {
		h264 := ListDownload{}
		h265 := ListDownload{}

		tipe := h.DOM.Find("div.Download__title")
		h264.Type = tipe.First().Text()
		h265.Type = tipe.Last().Text()

		download := h.DOM.Find("div.Download__container")

		// get data from container mkv
		download.First().Find("div.Download__group").Each(func(_ int, s *goquery.Selection) {
			d := Download{}
			d.Resolution = strings.ToLower(s.Find("div.Download__group-title").Text())

			s.Find("div.Download__link").Each(func(_ int, s *goquery.Selection) {
				s.Find("a").Each(func(_ int, s *goquery.Selection) {
					label := strings.ToLower(s.Find("span").First().Text())
					ahref := s.AttrOr("href", "")

					d.Links = append(d.Links, Links{label, ahref})
				})
			})

			h264.Downloads = append(h264.Downloads, d)

		})

		// get data from container x265
		download.Last().Find("div.Download__group").Each(func(_ int, s *goquery.Selection) {
			d := Download{}
			d.Resolution = strings.ToLower(s.Find("div.Download__group-title").Text())

			s.Find("div.Download__link").Each(func(_ int, s *goquery.Selection) {
				s.Find("a").Each(func(_ int, s *goquery.Selection) {
					label := strings.ToLower(s.Find("span").First().Text())
					ahref := s.AttrOr("href", "")

					d.Links = append(d.Links, Links{label, ahref})
				})
			})

			h265.Downloads = append(h265.Downloads, d)

		})

		result.Downloads = append(result.Downloads, h264)
		result.Downloads = append(result.Downloads, h265)

	})
	c.Visit(link)
	return result
}
