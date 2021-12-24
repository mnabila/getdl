package config

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
)

var configPath = fmt.Sprintf("/home/%s/.config/getdl/config.json", os.Getenv("USER"))

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
	config := Configuration{}

	configFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return DefaultConfig()
	}
	if err := json.Unmarshal(configFile, &config); err != nil {
		return DefaultConfig()
	}
	return config
}

func GetConfig(key string) string {
	config := Configuration{}

	configFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		config = DefaultConfig()
	}

	if err := json.Unmarshal(configFile, &config); err != nil {
		config = DefaultConfig()
	}

	switch key {
	case "codec":
		return config.Codec
	case "resolution":
		return config.Resolution
	case "file_hosting":
		return config.FileHosting
	case "browser":
		return config.Browser
	}
	return ""
}

func SetConfig(key string, value string) Configuration {
	config := Configuration{}

	configRead, err := ioutil.ReadFile(configPath)
	if err != nil {
		config = DefaultConfig()
	}

	if err := json.Unmarshal(configRead, &config); err != nil {
		config = DefaultConfig()
	}
	switch key {
	case "codec":
		config.Codec = value
	case "resolution":
		config.Resolution = value
	case "file_hosting":
		config.Codec = value
	case "browser":
		config.Browser = value
	}

	configWrite, err := json.Marshal(config)
	if err != nil {
		fmt.Println("Error Unable Parse Configuration")
	}

	if err := ioutil.WriteFile(configPath, configWrite, fs.FileMode(os.O_RDWR)); err != nil {
		fmt.Println("Error Unable Write Configuration")
	}

	return config
}
