terraform {
  required_version = ">= 1.2.9"

  backend "gcs" {
    bucket = "global-sre-dev-terraform"
    prefix = "state/monitor"
  }

  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 4.0"
    }
    kubectl = {
      source  = "gavinbunney/kubectl"
      version = ">= 1.7.0"
    }
    helm = {
      source = "hashicorp/helm"
      version = "= 2.6.0"
    }
    datadog = {
      source  = "datadog/datadog"
      version = "~> 3.0"
    }
  }
}

data "google_client_config" "this" {}

data "google_container_cluster" "this" {
  name     = local.cluster_name
  location = local.cluster_region
  project  = local.project_id
}

provider "google" {
  project = local.project_id
  region  = local.cluster_region
}

provider "kubectl" {
  host                   = "https://${data.google_container_cluster.this.endpoint}"
  token                  = data.google_client_config.this.access_token
  cluster_ca_certificate = base64decode(data.google_container_cluster.this.master_auth[0].cluster_ca_certificate)
  load_config_file       = false
}

provider "helm" {
  kubernetes {
    host                   = "https://${data.google_container_cluster.this.endpoint}"
    token                  = data.google_client_config.this.access_token
    cluster_ca_certificate = base64decode(data.google_container_cluster.this.master_auth[0].cluster_ca_certificate)
  }

  experiments {
    manifest = true
  }
}

provider "datadog" {
  api_key = data.google_secret_manager_secret_version.datadog_api_key.secret_data
  app_key = data.google_secret_manager_secret_version.datadog_app_key.secret_data
}