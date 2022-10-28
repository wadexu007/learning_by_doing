package apiv1

import (
	"log"
)

// Core V1 Interface
type CoreV1Interface interface {
	HelmAPIGetter
}

// CoreV1Client实现了CoreV1Interface的接口
type coreV1Client struct {
}

var newCoreV1Client CoreV1Interface

func ReturnV1APIClient() CoreV1Interface {
	log.Println("return an new core v1 API Client")
	if newCoreV1Client == nil {
		newCoreV1Client = &coreV1Client{}
	}
	return newCoreV1Client

}

func (c *coreV1Client) HelmClient(clusterContext, releaseNamespace string) HelmAPI {

	return NewHelmAPI(clusterContext, releaseNamespace)
}
