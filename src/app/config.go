package app

import (
    "encoding/json"
    "github.com/CarsonSlovoka/dovego/app/log"
    "io/ioutil"
    "os"
)

var (
    Config Settings
)

type Settings struct {
    Name    string   `json:"name"`
    Version string   `json:"version"`
    Server  Server   `json:"server"`
    Plugins []Plugin `json:"plugins"`
    Debug   `json:"debug"`
}

type Debug struct {
    Enable bool `json:"enable"`
}

type Server struct {
    Port int `json:"port"`
}

type Plugin struct {
    Name string `json:"name"`
    Path string `json:"path"`
}

func LoadConfig() {
    jsonFile, err := os.Open("manifest.json")
    if err != nil {
        log.Error.Printf(`Unable to open the "manifest.json": %v`, err)
    }

    byteValues, _ := ioutil.ReadAll(jsonFile)
    // json.Unmarshal(manifestBytes, &Config)
    if err := json.Unmarshal(byteValues, &Config); err != nil {
        log.Error.Printf("Unable to unmarshal json to struct (Settings): %v", err)
    }
}
