package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/jellydator/ttlcache/v3"
)

// AppConfig holds the application config
type AppConfig struct {
	ApiTemplateUrl string `json:"api_template_url"`
	UseCache       bool   `json:"use_cache"`
	ApiKey         string `json:"api_key"`
	ServerPort     string `json:"server_port"`
	Cache          *ttlcache.Cache[string, string]
}

// NewConfig setups new AppConfig from config file
func NewConfig(filename string) *AppConfig {
	var appConfig AppConfig

	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Error opening config file", err)
	}
	defer jsonFile.Close()

	// read opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &appConfig)
	if err != nil {
		log.Fatalln("Error unmarshalling config file", err)
	}

	return &appConfig
}
