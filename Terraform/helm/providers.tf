terraform {
  required_version = ">= 1.2.9"

  required_providers {
    helm = {
      source = "hashicorp/helm"
      version = "2.7.0"
    }
  }
}

# The easiest way is to supply a path to your kubeconfig file using the config_path attribute or using the KUBE_CONFIG_PATH environment variable. 
# A kubeconfig file may have multiple contexts. If config_context is not specified, the provider will use the default context.

# provider "helm" {
#   kubernetes {
#     config_path = "~/.kube/config"
#   }
# }


# OAuth2 access token 
provider "helm" {
  kubernetes {
    host                   = "https://${data.google_container_cluster.this.private_cluster_config[0].public_endpoint}"
    token                  = data.google_client_config.this.access_token
    cluster_ca_certificate = base64decode(data.google_container_cluster.this.master_auth[0].cluster_ca_certificate)
  }

  experiments {
    manifest = true
  }
}
