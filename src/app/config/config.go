package config

import (
	"encoding/json"
	"github.com/CarsonSlovoka/dovego/app/log"
	"io/ioutil"
	"os"
)

type Settings struct {
	Name    string   `json:"name"`
	Version string   `json:"version"`
	Server  Server   `json:"server"`
	Plugins []Plugin `json:"plugins"`
	Debug   `json:"debug"`
}

func LoadConfig(path string, settings *Settings) {
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Error.Printf(`Unable to open the "%s": %v`, path, err)
	}

	manifestBytes, _ := ioutil.ReadAll(jsonFile)
	if err := json.Unmarshal(manifestBytes, &settings); err != nil {
		log.Error.Printf("Unable to unmarshal json to struct (Settings): %v", err)
	}
}
