package config

import (
	"encoding/json"
	"os"
)

type ConfigStruct struct {
	GithubOAuth struct {
		ClientID     string `json:"clientID"`
		ClientSecret string `json:"clientSecret"`
	}
	RepoPath    string
	BlihSSH     string
	URL         string
}

var Config ConfigStruct

func LoadConfiguration() {
	configFile, err := os.Open("config/config.json")
	defer configFile.Close()
	if err != nil {
		panic("Can't open config file")
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&Config)
	if err != nil {
		panic("Config file loading failed")
	}
}