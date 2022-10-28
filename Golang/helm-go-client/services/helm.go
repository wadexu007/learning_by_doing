package services

import (
	"errors"
	"log"
)

type SimpleHelmReleaseOutput struct {
	Name        string `json:"name"`
	NAMESPACE   string `json:"namespace"`
	REVISION    int64  `json:"revision"`
	STATUS      string `json:"status"`
	CHARTPATH   string `json:"chart_path"`
	APP_VERSION string `json:"app_version"`
}

var (
	ErrListAllHelmRelease = errors.New("failed to get all helm release")
	ErrGetOneHelmRelease  = errors.New("failed to get helm release")
	ErrInstallHelmChart   = errors.New("failed to install helm chart")
	ErrDeleteHelmRelease  = errors.New("failed to delete helm release")
	ErrUpdateHelmChart    = errors.New("failed to update helm chart")
)

type HelmService interface {
	ListHelmRelease(clusterContext string, namespace string) ([]SimpleHelmReleaseOutput, error)
	GetHelmRelease(clusterContext, namespace, releaseName string) (SimpleHelmReleaseOutput, error)
	InstalHelmChart(clusterContext, chartPath, namespace, releaseName string) (SimpleHelmReleaseOutput, error)
	DeleteHelmRelease(clusterContext, namespace, releaseName string) (SimpleHelmReleaseOutput, error)
	UpdateHelmChart(clusterContext, namespace, chartPath, releaseName string, vals map[string]string) (SimpleHelmReleaseOutput, error)
	UpdateHelmChartAnyValue(clusterContext, namespace, chartPath, releaseName string, vals map[string]interface{}) (SimpleHelmReleaseOutput, error)
}

var helmServiceClient HelmService

func NewHelmService() HelmService {
	if helmServiceClient == nil {
		helmServiceClient = &helmService{}
	}
	return helmServiceClient
}

type helmService struct{}

func (s *helmService) ListHelmRelease(clusterContext, namespace string) ([]SimpleHelmReleaseOutput, error) {
	var outputs []SimpleHelmReleaseOutput
	results, err := coreV1Client.HelmClient(clusterContext, namespace).List()
	if err != nil {
		log.Printf("Failed to get all helm release, cluster %v, error: %v", clusterContext, err)
		return nil, ErrListAllHelmRelease
	}
	for _, result := range results {
		output := SimpleHelmReleaseOutput{
			Name:        result.Name,
			NAMESPACE:   result.Namespace,
			REVISION:    int64(result.Version),
			STATUS:      string(result.Info.Status),
			CHARTPATH:   result.Chart.ChartFullPath(),
			APP_VERSION: result.Chart.AppVersion(),
		}
		outputs = append(outputs, output)
	}
	return outputs, nil
}

func (s *helmService) GetHelmRelease(clusterContext, namespace, releaseName string) (SimpleHelmReleaseOutput, error) {
	var output SimpleHelmReleaseOutput
	result, err := coreV1Client.HelmClient(clusterContext, namespace).Get(releaseName)
	if result == nil {
		log.Printf("Nothing to get helm release with name: %s", releaseName)
		return output, ErrGetOneHelmRelease
	}

	if err != nil {
		log.Printf("Failed to get helm release with name: %s", releaseName)
		return output, ErrGetOneHelmRelease
	}
	output = SimpleHelmReleaseOutput{
		Name:        result.Name,
		NAMESPACE:   result.Namespace,
		REVISION:    int64(result.Version),
		STATUS:      string(result.Info.Status),
		CHARTPATH:   result.Chart.ChartFullPath(),
		APP_VERSION: result.Chart.AppVersion(),
	}

	return output, nil
}

func (s *helmService) InstalHelmChart(clusterContext, chartPath, namespace, releaseName string) (SimpleHelmReleaseOutput, error) {
	var output SimpleHelmReleaseOutput
	result, err := coreV1Client.HelmClient(clusterContext, namespace).InstallChart(chartPath, releaseName)
	if result == nil {
		log.Printf("No release, failed to Install helm chart: %s", releaseName)
		return output, ErrInstallHelmChart
	}

	if err != nil {
		log.Printf("Failed to Install helm release with name: %s", releaseName)
		return output, ErrInstallHelmChart
	}
	output = SimpleHelmReleaseOutput{
		Name:        result.Name,
		NAMESPACE:   result.Namespace,
		REVISION:    int64(result.Version),
		STATUS:      string(result.Info.Status),
		CHARTPATH:   result.Chart.ChartFullPath(),
		APP_VERSION: result.Chart.AppVersion(),
	}

	return output, nil
}

func (s *helmService) DeleteHelmRelease(clusterContext, namespace, releaseName string) (SimpleHelmReleaseOutput, error) {
	var output SimpleHelmReleaseOutput
	result, err := coreV1Client.HelmClient(clusterContext, namespace).Delete(releaseName)
	if result == nil {
		log.Printf("No release after install helm chart: %s", releaseName)
		return output, ErrGetOneHelmRelease
	}

	if err != nil {
		log.Printf("Failed to Install helm release with name: %s", releaseName)
		return output, ErrGetOneHelmRelease
	}
	output = SimpleHelmReleaseOutput{
		Name:        result.Release.Name,
		NAMESPACE:   result.Release.Namespace,
		REVISION:    int64(result.Release.Version),
		STATUS:      string(result.Info),
		CHARTPATH:   result.Release.Chart.ChartFullPath(),
		APP_VERSION: result.Release.Chart.AppVersion(),
	}
	return output, nil
}

func (s *helmService) UpdateHelmChart(clusterContext, namespace, chartPath, releaseName string, vals map[string]string) (SimpleHelmReleaseOutput, error) {
	var output SimpleHelmReleaseOutput
	result, err := coreV1Client.HelmClient(clusterContext, namespace).UpdateOnlyResource(chartPath, releaseName, vals)
	if result == nil {
		log.Printf("No release after update helm chart: %s", releaseName)
		return output, ErrGetOneHelmRelease
	}

	if err != nil {
		log.Printf("Failed to update helm release with name: %s", releaseName)
		return output, ErrGetOneHelmRelease
	}
	output = SimpleHelmReleaseOutput{
		Name:        result.Name,
		NAMESPACE:   result.Namespace,
		REVISION:    int64(result.Version),
		STATUS:      string(result.Info.Status),
		CHARTPATH:   result.Chart.ChartFullPath(),
		APP_VERSION: result.Chart.AppVersion(),
	}
	return output, nil
}

func (s *helmService) UpdateHelmChartAnyValue(clusterContext, namespace, chartPath, releaseName string, vals map[string]interface{}) (SimpleHelmReleaseOutput, error) {
	var output SimpleHelmReleaseOutput
	result, err := coreV1Client.HelmClient(clusterContext, namespace).UpdateAnyValue(chartPath, releaseName, vals)
	if result == nil {
		log.Printf("No release after update helm chart: %s", releaseName)
		return output, ErrGetOneHelmRelease
	}

	if err != nil {
		log.Printf("Failed to update helm release with name: %s", releaseName)
		return output, ErrGetOneHelmRelease
	}
	output = SimpleHelmReleaseOutput{
		Name:        result.Name,
		NAMESPACE:   result.Namespace,
		REVISION:    int64(result.Version),
		STATUS:      string(result.Info.Status),
		CHARTPATH:   result.Chart.ChartFullPath(),
		APP_VERSION: result.Chart.AppVersion(),
	}
	return output, nil
}
