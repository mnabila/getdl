package config_test

import (
	"encoding/json"
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
	t.Log(structTojson(config.DefaultConfig()))
}

func TestReadConfig(t *testing.T) {
	t.Log(structTojson(config.ReadConfig()))

}

func TestSetConfig(t *testing.T) {
	t.Log(structTojson(config.SetConfig("browserrr", "kivi")))

}

func TestGetConfig(t *testing.T) {
	t.Log(config.GetConfig("browser"))

}
