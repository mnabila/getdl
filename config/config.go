package config

type Configuration struct {
	Codec         string       `yaml:"codec"`
	Resolution    string       `yaml:"resolution"`
	FileHosting   string       `yaml:"file_hosting"`
	Browser       string       `yaml:"browser"`
	OpenInBrowser bool        `yaml:"open_in_browser"`
	Domain        DomainScrape `yaml:"domain"`
}

type DomainScrape struct {
	Doronime   string `yaml:"doronime"`
	Samehadaku string `yaml:"samehadaku"`
	Lendrive   string `yaml:"lendrive"`
	Oploverz   string `yaml:"oploverz"`
}
