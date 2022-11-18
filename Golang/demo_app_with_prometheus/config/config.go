package config

import (
	"github.com/tkanos/gonfig"
)

var Conf *Configuration

type Configuration struct {
	FILE_PATH string
	LOG_LEVEL string
	DB_HOST   string
	DB_NAME   string
}

func init() {
	Conf = &Configuration{}
	fileName := "/app/conf/config.json"
	gonfig.GetConf(fileName, Conf)
}
