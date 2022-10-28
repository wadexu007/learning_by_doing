package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type helmSimpleValues struct {
	Request_cpu    string
	Request_memory string
}

type helmReleaseInput struct {
	Cluster   string                 `json:"cluster"`
	ChartPath string                 `json:"chartPath"`
	Namespace string                 `json:"namespace"`
	Name      string                 `json:"name"`
	Vaules    map[string]interface{} `json:"values"`
}

type helmSpecficUpdateInput struct {
	HelmReleaseInput helmReleaseInput `json:"input"`
	HelmSimpleValues helmSimpleValues `json:"values"`
}

func ListHelmReleaseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	clusterContext := strings.TrimSpace(r.URL.Query().Get("clusterContext"))
	//if not provide namespace, then get all namespace
	namespace := strings.TrimSpace(r.URL.Query().Get("namespace"))

	result, err := helmService.ListHelmRelease(clusterContext, namespace)
	if err != nil {
		http.Error(w, "Not found any Helm release under deployed status", http.StatusNotFound)
		return
	}
	log.Println("Get all Helm release")
	json.NewEncoder(w).Encode(result)
}

func GetHelmReleaseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	clusterContext := strings.TrimSpace(r.URL.Query().Get("clusterContext"))
	namespace := strings.TrimSpace(r.URL.Query().Get("namespace"))
	if namespace == "" {
		http.Error(w, "Namespace can't be an empty string", http.StatusBadRequest)
		return
	}
	name := strings.TrimSpace(r.URL.Query().Get("name"))
	if name == "" {
		http.Error(w, "Please provide name of Helm release to get", http.StatusBadRequest)
		return
	}
	result, err := helmService.GetHelmRelease(clusterContext, namespace, name)
	if err != nil {
		http.Error(w, "Not found this Helm release under deployed status", http.StatusNotFound)
		return
	}
	log.Printf("Get Helm release: %s", name)
	json.NewEncoder(w).Encode(result)
}

func InstallHelmChartHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	req := helmReleaseInput{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Failed to decode request body with error: %v", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	if req.Namespace == "" {
		http.Error(w, "Namespace can't be an empty string", http.StatusBadRequest)
		return
	}
	if req.Name == "" {
		http.Error(w, "Please provide name of Helm release to get", http.StatusBadRequest)
		return
	}
	result, err := helmService.InstalHelmChart(req.Cluster, req.ChartPath, req.Namespace, req.Name)
	if err != nil {
		http.Error(w, "Failed to install Helm Chart", http.StatusBadRequest)
		return
	}
	log.Printf("Install Helm release: %s", result.Name)
	json.NewEncoder(w).Encode(result)
}

func DeleteHelmReleaseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	req := helmReleaseInput{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Failed to decode request body with error: %v", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	if req.Namespace == "" {
		http.Error(w, "Namespace can't be an empty string", http.StatusBadRequest)
		return
	}
	if req.Name == "" {
		http.Error(w, "Please provide name of Helm release to delete", http.StatusBadRequest)
		return
	}
	result, err := helmService.DeleteHelmRelease(req.Cluster, req.Namespace, req.Name)
	if err != nil {
		http.Error(w, "Failed to delete Helm release", http.StatusNotFound)
		return
	}
	log.Printf("Delete Helm release: %s", result.Name)
	json.NewEncoder(w).Encode(result)
}

func UpdateHelmChartOnlyResourceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	req := helmSpecficUpdateInput{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Failed to decode request body with error: %v", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	if req.HelmReleaseInput.Namespace == "" {
		http.Error(w, "Namespace can't be an empty string", http.StatusBadRequest)
		return
	}
	if req.HelmReleaseInput.Name == "" {
		http.Error(w, "Please provide name of Helm release to delete", http.StatusBadRequest)
		return
	}
	vals := map[string]string{
		"request_cpu":    req.HelmSimpleValues.Request_cpu,
		"request_memory": req.HelmSimpleValues.Request_memory,
	}
	log.Printf("Update resource request cpu and mem are: %v", vals)
	result, err := helmService.UpdateHelmChart(req.HelmReleaseInput.Cluster, req.HelmReleaseInput.Namespace, req.HelmReleaseInput.ChartPath, req.HelmReleaseInput.Name, vals)
	if err != nil {
		http.Error(w, "Failed to Update Helm Chart", http.StatusNotFound)
		return
	}
	log.Printf("Update Helm release: %s", result.Name)
	json.NewEncoder(w).Encode(result)
}

func UpdateHelmChartAnyValueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	req := helmReleaseInput{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Failed to decode request body with error: %v", err)
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	if req.Namespace == "" {
		http.Error(w, "Namespace can't be an empty string", http.StatusBadRequest)
		return
	}
	if req.Name == "" {
		http.Error(w, "Please provide name of Helm release to delete", http.StatusBadRequest)
		return
	}

	result, err := helmService.UpdateHelmChartAnyValue(req.Cluster, req.Namespace, req.ChartPath, req.Name, req.Vaules)
	if err != nil {
		http.Error(w, "Failed to Update Helm Chart", http.StatusNotFound)
		return
	}
	log.Printf("Update Helm release: %s", result.Name)
	json.NewEncoder(w).Encode(result)
}
