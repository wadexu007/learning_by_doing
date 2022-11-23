package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Status(c *gin.Context) {
	c.String(http.StatusOK, "pong!")
}
