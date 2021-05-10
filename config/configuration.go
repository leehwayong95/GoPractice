package config

import (
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	DB_USER	string
	DB_PW	string
	DB_HOST	string
	DB_Port	string
	DBtype	string
	DB_TARGET	string
	DB_OPTION	string
}

func GetConfiguration() Configuration {
	configuration := Configuration{}
	gonfig.GetConf("config/config.local.json", &configuration)
	return configuration
}

func (c *Configuration) GetValue (Key string) string {
	test := make(map[string]string)
	gonfig.GetConf("config/config.local.json",&test)
	return test[Key]
}