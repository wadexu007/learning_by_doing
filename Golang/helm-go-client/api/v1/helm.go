package apiv1

import (
	"helm-go-client/config"
	"log"
	"os"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/kube"
	"helm.sh/helm/v3/pkg/release"
)

type HelmAPIGetter interface {
	HelmClient(clusterContext, releaseNamespace string) HelmAPI
}

type HelmAPI interface {
	Get(releaseName string) (*release.Release, error)
	List() ([]*release.Release, error)
	InstallChart(chartPath, releaseName string) (*release.Release, error)
	Delete(releaseName string) (*release.UninstallReleaseResponse, error)
	UpdateOnlyResource(chartPath, releaseName string, vals map[string]string) (*release.Release, error)
	UpdateAnyValue(chartPath, releaseName string, vals map[string]interface{}) (*release.Release, error)
}

type helmAPI struct {
	actionConfig *action.Configuration
	context      string
	namespace    string
}

func NewHelmAPI(clusterContext, releaseNamespace string) HelmAPI {
	log.Println("Return new Helm API")

	actionConfig := new(action.Configuration)

	if err := actionConfig.Init(kube.GetConfig(config.Conf.KUBE_CONFIG_PATH, clusterContext, releaseNamespace), releaseNamespace, os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
		log.Printf("%+v", err)
		os.Exit(1)
	}

	var newHelmAPI HelmAPI
	if newHelmAPI == nil {
		newHelmAPI = &helmAPI{
			actionConfig: actionConfig,
			context:      clusterContext,
			namespace:    releaseNamespace,
		}
	}

	return newHelmAPI
}

func (h *helmAPI) List() ([]*release.Release, error) {
	log.Println("Call helmAPI API List method")
	client := action.NewList(h.actionConfig)
	// Only list deployed
	client.Deployed = true
	results, err := client.Run()
	if err != nil {
		log.Printf("%+v", err)
		// os.Exit(1)
	}

	return results, err
}

func (h *helmAPI) Get(releaseName string) (*release.Release, error) {
	log.Println("Call Helm API Get method")

	client := action.NewList(h.actionConfig)
	// Only list deployed
	client.Deployed = true
	results, err := client.Run()
	if err != nil {
		log.Printf("%+v", err)
		// os.Exit(1)
	}
	// var rel *release.Release
	for _, rel := range results {
		log.Printf("%+v: %+v", rel.Name, rel.Namespace)
		if rel.Name == releaseName {
			return rel, nil
		}
	}
	return nil, err
}

func (h *helmAPI) InstallChart(chartPath, releaseName string) (*release.Release, error) {
	log.Println("Call Helm Install Chart method")
	chart, err := loader.Load(chartPath)
	if err != nil {
		log.Printf("%+v", err)
	}

	client := action.NewInstall(h.actionConfig)
	client.Namespace = h.namespace
	client.ReleaseName = releaseName
	rel, err := client.Run(chart, nil)
	if err != nil {
		log.Printf("%+v", err)
	}
	return rel, err
}

func (h *helmAPI) Delete(releaseName string) (*release.UninstallReleaseResponse, error) {
	log.Println("Call Helm Delete release method")

	client := action.NewUninstall(h.actionConfig)
	res, err := client.Run(releaseName)
	if err != nil {
		log.Printf("%+v", err)
	}
	return res, err
}

func (h *helmAPI) UpdateOnlyResource(chartPath, releaseName string, vals map[string]string) (*release.Release, error) {
	log.Println("Call Helm Update Chart method")
	log.Printf("Resource will update to %v", vals)
	chart, err := loader.Load(chartPath)
	if err != nil {
		log.Printf("%+v", err)
	}

	resourceValues := map[string]interface{}{
		"resources": map[string]interface{}{
			"requests": map[string]string{
				"cpu":    vals["request_cpu"],
				"memory": vals["request_memory"],
			},
		},
	}

	client := action.NewUpgrade(h.actionConfig)
	res, err := client.Run(releaseName, chart, resourceValues)
	if err != nil {
		log.Printf("%+v", err)
	}
	return res, err
}

func (h *helmAPI) UpdateAnyValue(chartPath, releaseName string, vals map[string]interface{}) (*release.Release, error) {
	log.Println("Call Helm Update Chart method")
	log.Printf("Update char value: %v", vals)
	chart, err := loader.Load(chartPath)
	if err != nil {
		log.Printf("%+v", err)
	}

	client := action.NewUpgrade(h.actionConfig)
	res, err := client.Run(releaseName, chart, vals)
	if err != nil {
		log.Printf("%+v", err)
	}
	return res, err
}
