//go:build wireinject
// +build wireinject

package router

import (
	"cost-analyzer/controller/account"
	"cost-analyzer/repository"
	"cost-analyzer/service"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func initUserController(db *gorm.DB) account.UserController {
	wire.Build(repository.NewUserRepository, service.NewUserService, account.NewUserController)
	return nil
}
