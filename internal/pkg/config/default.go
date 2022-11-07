package config

import (
	"encoding/json"
	"github.com/bilgehannal/harbctl/internal/env"
	"github.com/mittwald/goharbor-client/v5/apiv2/model"
	"log"
	"os"
)

type DefaultConfig struct {
	ActiveContext string    `json:"activeContext"`
	Contexts      []Context `json:"contexts"`
}

func (d DefaultConfig) GetActiveContext() Context {
	activeContextUser := d.ActiveContext
	isContextValid, activeContext := d.GetContextFromUser(activeContextUser)
	if !isContextValid {
		log.Fatal("Active context cannot be found in context list!")
	}
	return activeContext
}

func (d DefaultConfig) GetContextFromUser(userName string) (bool, Context) {
	for _, context := range d.Contexts {
		if context.User == userName {
			return true, context
		}
	}
	return false, Context{}
}

func getDefaultConfigPath() string {
	return os.Getenv(env.HarbctlConfigPath)
}

func getDefaultConfigFilePath() string {
	configFolder := getDefaultConfigPath()
	return configFolder + "/config.json"
}

func GetDefaultConfig() DefaultConfig {
	content, err := os.ReadFile(getDefaultConfigFilePath())
	if err != nil {
		log.Println("There is no default config. Creating default config...")
		return createNewDefualtConfig()
	}
	var payload DefaultConfig
	err = json.Unmarshal(content, &payload)
	if err != nil {
		panic(err)
	}
	return payload
}

func createNewDefualtConfig() DefaultConfig {
	return DefaultConfig{}
}

func (d *DefaultConfig) SetDefaultConfig(ctx Context, userInfo *model.UserResp) error {
	log.Println("girdii")
	if !d.IsThereAnyActiveContext() {
		err := os.MkdirAll(getDefaultConfigPath(), 0700)
		if err != nil {
			return err
		}
		d.ActiveContext = ctx.User
	}
	d.removeContextFromDefaultConfig(ctx.User)
	d.Contexts = append(d.Contexts, ctx)
	file, _ := json.MarshalIndent(d, "", "  ")
	err := os.WriteFile(getDefaultConfigFilePath(), file, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (d *DefaultConfig) removeContextFromDefaultConfig(userName string) {
	for i := 0; i < len(d.Contexts); i++ {
		if d.Contexts[i].User == userName {
			d.Contexts = removeIndex(d.Contexts, i)
			return
		}
	}
}

func removeIndex(s []Context, index int) []Context {
	return append(s[:index], s[index+1:]...)
}

func (d DefaultConfig) IsThereAnyActiveContext() bool {
	if d.ActiveContext == "" {
		return false
	}
	return true
}
