package config

import (
	"encoding/json"
	"github.com/bilgehannal/harbctl/internal/env"
	"io/ioutil"
	"log"
	"os"
)

type DefaultConfig struct {
	ActiveContext string    `json:"activeContext"`
	Contexts      []Context `json:"contexts"`
}

func (d DefaultConfig) GetActiveContext() Context {
	activeContextUrl := d.ActiveContext
	isContextValid, activeContext := d.GetContextFromUrl(activeContextUrl)
	if !isContextValid {
		log.Fatal("Active context cannot be found in context list!")
	}
	return activeContext
}

func (d DefaultConfig) GetContextFromUrl(url string) (bool, Context) {
	for _, context := range d.Contexts {
		if context.Url == url {
			return true, context
		}
	}
	return false, Context{}
}

func GetDefaultConfigFilePath() string {
	configFolder := os.Getenv(env.HarbctlConfigPath)
	return configFolder + "/config.json"
}

func GetDefaultConfig() DefaultConfig {
	content, err := ioutil.ReadFile(GetDefaultConfigFilePath())
	if err != nil {
		log.Fatal("Error when opening config file: ", err)
	}
	var payload DefaultConfig
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	return payload
}

func (d DefaultConfig) SetDefaultConfig() {
	file, _ := json.MarshalIndent(d, "", "  ")
	err := ioutil.WriteFile(GetDefaultConfigFilePath(), file, 0644)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
}
