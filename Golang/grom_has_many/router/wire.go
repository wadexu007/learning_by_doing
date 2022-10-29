//go:build wireinject
// +build wireinject

package router

import (
	"main.go/controller/account"
	"main.go/repository"
	"main.go/service"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func initUserController(db *gorm.DB) account.UserController {
	wire.Build(repository.NewUserRepository, service.NewUserService, account.NewUserController)
	return nil
}
