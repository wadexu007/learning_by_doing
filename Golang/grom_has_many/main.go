/*
Copyright 2022 Wade Xu
*/

package main

import (
	"os"

	"main.go/config"
	"main.go/lib/logger"
	"main.go/middleware"
	"main.go/model"
	"main.go/router"
)

func main() {
	config.InitConf(os.Getenv("RUN_ENV"))
	logger.InitLog(config.Config.GetString("log.config"))
	db := model.ConnectDB()

	middleware.InitAuth(config.Config.GetString("admin.name"), config.Config.GetString("admin.password"))
	router.InitRouter(db).Run(":8080")
}
