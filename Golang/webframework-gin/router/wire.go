//go:build wireinject
// +build wireinject

package router

import (
	"demo-gin/controller/account"
	"demo-gin/repository"
	"demo-gin/service"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func initUserController(db *gorm.DB) account.UserController {
	wire.Build(repository.NewUserRepository, service.NewUserService, account.NewUserController)
	return nil
}
