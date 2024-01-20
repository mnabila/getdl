package config

import "github.com/spf13/viper"

type Configurations struct {
	Browser       string      `mapstructure:"browser"`
	OpenInBrowser bool        `mapstructure:"open_in_browser"`
	Website       []WebConfig `mapstructure:"website"`
}

type WebConfig struct {
	Name        string   `mapstructure:"name"`
	Domain      []string `mapstructure:"domain"`
	Codec       string   `mapstructure:"codec"`
	Resolution  string   `mapstructure:"resolution"`
	FileHosting []string `mapstructure:"file_hosting"`
}

func SetupDefaultConfig() error {
	viper.SetDefault("browser", "xdg-open")
	viper.SetDefault("open_in_browser", true)
	viper.SetDefault("website", []WebConfig{
		{
			Name:        "animekompi",
			Domain:      []string{"animekompi.net"},
			Codec:       "x265",
			Resolution:  "720",
			FileHosting: []string{"drive"},
		},
		{

			Name:        "doronime",
			Domain:      []string{"doronime.id"},
			Codec:       "x265",
			Resolution:  "720",
			FileHosting: []string{"drive"},
		},
		{
			Name:        "lendrive",
			Domain:      []string{"lendrive.web.id"},
			Codec:       "x265",
			Resolution:  "720",
			FileHosting: []string{"drive"},
		},
		{
			Name:        "oploverz",
			Domain:      []string{"oploverz.red"},
			Codec:       "x265",
			Resolution:  "720",
			FileHosting: []string{"drive"},
		},
		{
			Name:        "samehadaku",
			Domain:      []string{"samehadaku.help"},
			Codec:       "x265",
			Resolution:  "720",
			FileHosting: []string{"drive"},
		},
	})

	return viper.SafeWriteConfig()
}
