package config_test

import (
	"encoding/json"
	"fmt"
	"getdl/config"
	"testing"
)

func structTojson(c config.Configuration) string {
	if result, err := json.Marshal(c); err == nil {
		return string(result)
	}
	return ""
}

func TestDefaultConfig(t *testing.T) {
	fmt.Println(structTojson(config.DefaultConfig()))
}

func TestReadConfig(t *testing.T) {
	fmt.Println(structTojson(config.ReadConfig()))

}
