terraform {
  required_version = ">= 1.2.9"

  backend "gcs" {
    bucket = "global-sre-dev-terraform"
    prefix = "kustomize/state"
  }

  required_providers {
    kustomization = {
      source = "kbst/kustomization"
      version = "0.9.0"
    }
  }
}

# The easiest way is to supply a path to your kubeconfig file using the config_path attribute or using the KUBE_CONFIG_PATH environment variable. 
# A kubeconfig file may have multiple contexts. If config_context is not specified, the provider will use the default context.

provider "kustomization" {
  kubeconfig_path = "~/.kube/config"
}
