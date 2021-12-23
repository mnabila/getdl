package scrape

type Response struct {
	Url         string         `json:"url"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Downloads   []ListDownload `json:"downloads"`
}

type ListDownload struct {
	Type      string     `json:"tipe"`
	Downloads []Download `json:"downloads"`
}

type Download struct {
	Resolution string  `json:"resolution"`
	Links []Links `json:"links"`
}

type Links struct {
	Label string   `json:"label"`
	Link  string `json:"link"`
}
