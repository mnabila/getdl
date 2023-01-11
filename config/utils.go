package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"

	"gopkg.in/yaml.v3"
)

func GetConfigPath() string {
	if runtime.GOOS == "linux" {
		return fmt.Sprintf("/home/%s/.config/getdl/config.yaml", os.Getenv("USER"))
	} else {
		cwd, _ := os.Getwd()
		return fmt.Sprintf("%s/config.yaml", cwd)
	}
}

func ReadConfig() (*Config, error) {
	configInByte, err := ioutil.ReadFile(GetConfigPath())
	if err != nil {
		return nil, err
	}

	var output Config
	if err := yaml.Unmarshal(configInByte, &output); err != nil {
		return nil, err
	}

	return &output, nil
}
