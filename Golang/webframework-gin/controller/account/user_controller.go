package account

import (
	. "cost-analyzer/error"
	"cost-analyzer/lib/logger"
	"cost-analyzer/model"
	. "cost-analyzer/service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type UserController interface {
	Search(c *gin.Context)
	Get(c *gin.Context)
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
}

func NewUserController(userService UserService) UserController {
	return &userController{service: userService}
}

type userController struct {
	service UserService
}

func (u *userController) Search(c *gin.Context) {
	name := c.Request.URL.Query().Get("name")
	if len(name) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	} else {
		users, err := u.service.SearchUser(name, false)
		if err != nil {
			if errors.Is(err, BadRequestError) {
				c.JSON(http.StatusBadRequest, "Failed to search user")
			} else {
				c.JSON(http.StatusInternalServerError, "Failed to search user")
			}

		} else {
			c.JSON(http.StatusOK, users)
		}
	}
	c.Abort()
	return
}

func (u *userController) Get(c *gin.Context) {
	if c.Param("id") != "" {
		logger.DebugF("Start to get %s", c.Param("id"))
		user, err := u.service.GetUser(c.Param("id"), false)
		if err != nil {
			logger.ErrorF("Failed to get user %s %v", c.Param("id"), zap.Error(err))
			if errors.Is(err, BadRequestError) {
				c.JSON(http.StatusBadRequest, "Failed to get user")
			} else {
				c.JSON(http.StatusInternalServerError, "Failed to get user")
			}
		} else {
			logger.InfoW("Get ", "user", &user)
			c.JSON(http.StatusOK, user)
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	}
	c.Abort()
	return
}

func (u *userController) Create(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)
	userID, err := u.service.CreateUser(&user)
	if err != nil {
		logger.ErrorF("Failed to create user %s %v", user.Name, err)
		if errors.Is(err, BadRequestError) {
			c.JSON(http.StatusBadRequest, "Failed to create user")
		} else {
			c.JSON(http.StatusInternalServerError, "Failed to create user")
		}
		c.Abort()
		return
	}
	logger.Info("Created user %s %s", user.Name, userID.String())
	c.JSON(http.StatusOK, userID)
	c.Abort()
	return
}

func (u *userController) Delete(c *gin.Context) {
	if c.Param("id") != "" {
		_, err := u.service.DeleteUser(c.Param("id"), "system")
		if err != nil {
			if errors.Is(err, BadRequestError) {
				c.JSON(http.StatusBadRequest, "Failed to delete user")
			} else {
				c.JSON(http.StatusInternalServerError, "Failed to delete user")
			}
		} else {
			c.Writer.WriteHeader(http.StatusOK)
		}
	} else {
		c.JSON(http.StatusBadRequest, "bad request")
	}
	c.Abort()
	return
}

func (u *userController) Update(c *gin.Context) {
	var userInfo model.User
	userID, error := uuid.Parse(c.Param("id"))
	if error != nil {
		c.JSON(http.StatusBadRequest, "bad request")
	} else {
		userInfo.ID = userID
		c.BindJSON(&userInfo)
		_, err := u.service.UpdateUser(&userInfo, "")
		if err != nil {
			logger.Error(err)
			if errors.Is(err, BadRequestError) {
				c.JSON(http.StatusBadRequest, "Failed to update user")
			} else {
				c.JSON(http.StatusInternalServerError, "Failed to update user")
			}
		} else {
			c.Writer.WriteHeader(http.StatusOK)
		}
	}
	c.Abort()
	return

}
