package middleware

import (
	"time"

	"main.go/lib/logger"
	"main.go/lib/utils"
	"main.go/model"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type loginPayload struct {
	UserName string `form:"userName" json:"userName" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var identityKey = "id"
var AuthMiddleware *jwt.GinJWTMiddleware

func InitAuth(adminName string, adminPassword string) {
	var err error
	AuthMiddleware, err = jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "demo",
		Key:         []byte("my key"),
		Timeout:     30 * time.Minute,
		MaxRefresh:  30 * time.Minute,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					identityKey: v.Name,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &model.User{
				Name: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginPayload loginPayload
			if err := c.ShouldBind(&loginPayload); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userName := loginPayload.UserName
			password := loginPayload.Password

			if userName == adminName && password == adminPassword {
				return &model.User{
					Name: userName,
				}, nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*model.User); ok && v.Name == adminName {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": utils.Capitalize(message),
			})
		},
		TokenLookup:   "header: Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		logger.Error(err)
		panic(err)
	}
}
