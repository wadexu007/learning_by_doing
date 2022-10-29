package router

import (
	"net/http"

	"main.go/controller"
	"main.go/lib/logger"
	"main.go/middleware"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitRouter(d *gorm.DB) *gin.Engine {
	db = d
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	health := new(controller.HealthController)
	router.GET("/health", health.Status)
	initAuthentication(router)
	v1 := router.Group("v1")
	initUserRouter(v1)
	return router

}

func initUserRouter(v1 *gin.RouterGroup) {
	v1.Use(middleware.AuthMiddleware.MiddlewareFunc())
	userGroup := v1.Group("account")
	//init repository, service and controller via google wire
	userController := initUserController(db)
	userGroup.GET("/:id", userController.Get)
	userGroup.GET("/search", userController.Search)
	userGroup.PUT("/:id", userController.Update)
	userGroup.DELETE("/:id", userController.Delete)
	userGroup.POST("", userController.Create)

}

func initAuthentication(router *gin.Engine) {
	router.POST("/login", middleware.AuthMiddleware.LoginHandler)
	router.NoRoute(middleware.AuthMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		logger.ErrorF("NoRoute claims: %v", claims)
		c.JSON(http.StatusNotFound, "Not found")
	})
	auth := router.Group("/auth")
	auth.GET("/refresh_token", middleware.AuthMiddleware.RefreshHandler)
}
