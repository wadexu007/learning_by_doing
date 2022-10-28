package main

import (
	"helm-go-client/config"
	"helm-go-client/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	log.Println("Init Configuration")
	log.Println("kube config path: " + config.Conf.KUBE_CONFIG_PATH)

	r := mux.NewRouter()

	// Routes consist of a path and a handler function.
	r.HandleFunc("/healthz", controllers.HealthzHandler).Methods("GET")
	r.HandleFunc("/list", controllers.ListHelmReleaseHandler).Methods("GET")
	r.HandleFunc("/get", controllers.GetHelmReleaseHandler).Methods("GET")
	r.HandleFunc("/install", controllers.InstallHelmChartHandler).Methods("POST")
	r.HandleFunc("/delete", controllers.DeleteHelmReleaseHandler).Methods("DELETE")
	r.HandleFunc("/updateOnlyResource", controllers.UpdateHelmChartOnlyResourceHandler).Methods("PUT")
	r.HandleFunc("/updateHelmAnyValue", controllers.UpdateHelmChartAnyValueHandler).Methods("PUT")

	// Bind to a port and pass our router in
	log.Println("Start http server")
	log.Fatal(http.ListenAndServe(":8080", r))
}
