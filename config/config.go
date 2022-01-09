package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
)

var configPath = fmt.Sprintf("/home/%s/.config/getdl/config.json", os.Getenv("USER"))

type Configuration struct {
	Codec         string `json:"codec"`
	Resolution    string `json:"resolution"`
	FileHosting   string `json:"file_hosting"`
	Browser       string `json:"browser"`
	OpenInBrowser string `json:"open_in_browser"`
}

func DefaultConfig() Configuration {
	c := Configuration{}
	c.Codec = "265"
	c.Resolution = "720p"
	c.FileHosting = "zippyshare"
	c.Browser = "xdg-open"
	c.OpenInBrowser = "true"
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
	case "codec", "Codec":
		return config.Codec
	case "resolution", "Resolution":
		return config.Resolution
	case "file_hosting", "FileHosting":
		return config.FileHosting
	case "browser", "Browser":
		return config.Browser
	case "open_in_browser", "OpenInBrowser":
		return config.OpenInBrowser
	}
	return ""
}

func SetConfig(key string, value string) error {
	config := Configuration{}

	configRead, err := ioutil.ReadFile(configPath)
	if err != nil {
		config = DefaultConfig()
	}

	if err := json.Unmarshal(configRead, &config); err != nil {
		config = DefaultConfig()
	}

	switch key {
	case "codec", "Codec":
		config.Codec = value
	case "resolution", "Resolution":
		config.Resolution = value
	case "file_hosting", "FileHosting", "filehosting":
		config.FileHosting = value
	case "browser", "Browser":
		config.Browser = value
	case "open_in_browser", "OpenInBrowser", "openinbrowser":
		config.OpenInBrowser = value
	}

	configWrite, err := json.Marshal(config)
	if err != nil {
		return errors.New("error unable parse configuration")
	}

	if err := ioutil.WriteFile(configPath, configWrite, fs.FileMode(os.O_RDWR)); err != nil {
		return errors.New("error unable write configuration")
	}

	return nil
}

func ResetConfig() error {
	config := DefaultConfig()
	configWrite, err := json.Marshal(config)
	if err != nil {
		return errors.New("error unable parse configuration")
	}
	if err := ioutil.WriteFile(configPath, configWrite, fs.FileMode(os.O_RDWR)); err != nil {
		return errors.New("error unable write configuration")
	}
	return nil
}
