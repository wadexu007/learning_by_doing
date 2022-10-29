// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package router

import (
	"demo-gin/controller/account"
	"demo-gin/repository"
	"demo-gin/service"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func initUserController(db2 *gorm.DB) account.UserController {
	userRepository := repository.NewUserRepository(db2)
	userService := service.NewUserService(userRepository)
	userController := account.NewUserController(userService)
	return userController
}