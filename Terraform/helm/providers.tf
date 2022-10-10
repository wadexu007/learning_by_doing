terraform {
  required_version = ">= 1.2.9"

  required_providers {
    helm = {
      source = "hashicorp/helm"
      version = "2.7.0"
    }
  }
}

provider "helm" {
  kubernetes {
    config_path = "~/.kube/config"
  }
}
