package config

import (
	"github.com/tkanos/gonfig"
)

var Conf *Configuration

type Configuration struct {
	KUBE_CONFIG_PATH string
}

func init() {
	Conf = &Configuration{}
	fileName := "conf/config.json"
	gonfig.GetConf(fileName, Conf)

}
