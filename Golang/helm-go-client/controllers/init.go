package controllers

import (
	"helm-go-client/services"
)

var helmService services.HelmService

func init() {
	helmService = services.NewHelmService()

}
