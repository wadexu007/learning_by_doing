package services

import (
	apiv1 "helm-go-client/api/v1"
)

var coreV1Client apiv1.CoreV1Interface

func init() {
	coreV1Client = apiv1.ReturnV1APIClient()
}
