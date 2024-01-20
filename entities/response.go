package entities

type ScrapeResponse struct {
	Url         string     `json:"url"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Downloads   []Download `json:"downloads"`
}

type Download struct {
	Codec       string `json:"codec,omitempty"`
	Resolution  string `json:"resolution,omitempty"`
	FileHosting string `json:"file_hosting,omitempty"`
	UrlDownload string `json:"url_download,omitempty"`
}
