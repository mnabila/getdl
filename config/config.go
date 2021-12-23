package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Configuration struct {
	Codec       string `json:"codec"`
	Resolution  string `json:"resolution"`
	FileHosting string `json:"file_hosting"`
	Browser     string `json:"browser"`
}

func DefaultConfig() Configuration {
	c := Configuration{}
	c.Codec = "265"
	c.Resolution = "720p"
	c.FileHosting = "zippyshare"
	c.Browser = "xdg-open"
	return c
}

func ReadConfig() Configuration {
	configPath := fmt.Sprintf("/home/%s/.config/getdl/config.json", os.Getenv("USER"))
	config := Configuration{}

	if _, err := os.Stat(configPath); err != nil {
		return DefaultConfig()
	} else {
		configFile, err := ioutil.ReadFile(configPath)
		if err != nil {
			return DefaultConfig()
		}
		if err := json.Unmarshal(configFile, &config); err != nil {
			return DefaultConfig()
		}
		return config
	}
}
