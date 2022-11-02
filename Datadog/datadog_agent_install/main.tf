locals {
  project_id             = "global-sre-dev"
  cluster_name           = "sre-gke"
  cluster_region         = "us-central1"
  datadog_ns             = "datadog"
  dd_api_key_secret_name = "datadog-api-key"
  dd_app_key_secret_name = "datadog-app-key"

  datadogAttributes = {
    apiKey      = data.google_secret_manager_secret_version.datadog_api_key.secret_data
    appKey      = data.google_secret_manager_secret_version.datadog_app_key.secret_data
    clusterName = data.google_container_cluster.this.name
    projectName = data.google_container_cluster.this.project
  }
}

// read dd api key from google's secret manager
data "google_secret_manager_secret_version" "datadog_api_key" {
  project = local.project_id
  secret  = local.dd_api_key_secret_name
}

// read dd app key from google's secret manager
data "google_secret_manager_secret_version" "datadog_app_key" {
  project = local.project_id
  secret  = local.dd_app_key_secret_name
}

# create a namespace with Terraform kubectl provider
resource "kubectl_manifest" "datadog_ns" {
    yaml_body = <<YAML
apiVersion: v1
kind: Namespace
metadata:
  name: ${local.datadog_ns}
YAML
}

# install datadog agent with Terraform Helm provider
resource helm_release "datadog_agent" {
  name       = "datadog-agent"
  namespace  = local.datadog_ns
  repository = "https://helm.datadoghq.com"
  chart      = "datadog"
  version    = "3.1.11"

  values = [
    templatefile("yamls/datadog-agent-config.yaml.tftpl", local.datadogAttributes)
  ]
  depends_on = [
    kubectl_manifest.datadog_ns
  ]
}
