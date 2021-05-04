package config

import "github.com/tkanos/gonfig"

type Configuration struct {
	DBUrl string
	DBtype string
}

func GetConfiguration() Configuration {
	configuration := Configuration{}
	gonfig.GetConf("config/config.local.json", &configuration)
	return configuration
}