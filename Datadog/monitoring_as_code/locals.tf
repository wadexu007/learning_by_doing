locals {
  project_id             = "global-sre-dev"
  cluster_name           = "finops-gke"
  cluster_region         = "us-central1"
}

// read dd api key from google's secret manager
data "google_secret_manager_secret_version" "datadog_api_key" {
  project = local.project_id
  secret  = "datadog-api-key"
}

// read dd app key from google's secret manager
data "google_secret_manager_secret_version" "datadog_app_key" {
  project = local.project_id
  secret  = "datadog-app-key"
}