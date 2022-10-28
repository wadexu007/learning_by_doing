/*
Copyright 2022 Wade Xu
*/

package main

import (
	"cost-analyzer/config"
	"cost-analyzer/lib/logger"
	"cost-analyzer/middleware"
	"cost-analyzer/model"
	"cost-analyzer/router"
	"os"
)

func main() {
	config.InitConf(os.Getenv("RUN_ENV"))
	logger.InitLog(config.Config.GetString("log.config"))
	db := model.ConnectDB()

	middleware.InitAuth(config.Config.GetString("admin.name"), config.Config.GetString("admin.password"))
	router.InitRouter(db).Run(":8080")
}
