package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"

	"gopkg.in/yaml.v3"
)

func GetConfigPath() (out string) {
	if runtime.GOOS == "linux" {
		out = fmt.Sprintf("/home/%s/.config/getdl/config.yaml", os.Getenv("USER"))
	} else {
		cwd, _ := os.Getwd()
		out = fmt.Sprintf("%s/config.yaml", cwd)
	}
	return
}

func DefaultConfig() *Configuration {
	return &Configuration{
		Codec:         "x264",
		Resolution:    "720p",
		FileHosting:   "google",
		Browser:       "xdg-open",
		OpenInBrowser: false,
	}
}

func ReadConfig() (*Configuration, error) {
	configInByte, err := ioutil.ReadFile(GetConfigPath())
	if err != nil {
		return nil, err
	}

	var output Configuration
	if err := yaml.Unmarshal(configInByte, &output); err != nil {
		return nil, err
	}

	return &output, nil
}
