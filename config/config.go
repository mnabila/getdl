package config

type Config struct {
	Browser       string      `yaml:"browser"`
	OpenInBrowser bool        `yaml:"open_in_browser"`
	Website       []WebConfig `yaml:"website"`
}

type WebConfig struct {
	Name        string   `yaml:"name"`
	Domain      []string   `yaml:"domain"`
	Codec       string   `yaml:"codec"`
	Resolution  string   `yaml:"resolution"`
	FileHosting []string `yaml:"file_hosting"`
}
