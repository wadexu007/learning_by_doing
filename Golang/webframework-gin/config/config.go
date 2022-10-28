package config

import (
	"github.com/spf13/viper"
)

var Config *viper.Viper

var DBConfig *DatabaseConfig

type DatabaseConfig struct {
	User     string
	Password string
	Host     string
	Name     string
	ConnPool *DatabaseConnPool
}

type DatabaseConnPool struct {
	MaxIdleConns    int64
	MaxOpenConns    int64
	MaxConnLifeTime int64
}

func InitConf(env string) {
	var err error
	Config = viper.New()
	Config.SetConfigType("yaml")
	if len(env) == 0 {
		Config.SetConfigName("app")
	} else {
		Config.SetConfigName("app_" + env)
	}
	Config.AddConfigPath("conf/")
	err = Config.ReadInConfig()
	if err != nil {
		panic(err)
	}

	setDBConfig()
}

/**
Set DB Configuration
*/
func setDBConfig() {
	DBConfig = &DatabaseConfig{
		User: Config.GetString("mysql.user"), Password: Config.GetString("mysql.password"),
		Host: Config.GetString("mysql.host"), Name: Config.GetString("mysql.name"),
		ConnPool: &DatabaseConnPool{
			MaxIdleConns:    Config.GetInt64("mysql.pool.maxIdleConns"),
			MaxOpenConns:    Config.GetInt64("mysql.pool.maxOpenConns"),
			MaxConnLifeTime: Config.GetInt64("mysql.pool.maxConnLifeTime"),
		},
	}
}
