terraform {
  required_version = ">= 1.2.9"

  backend "gcs" {
    bucket = "global-sre-dev-terraform"
    prefix = "state/velero"
  }

  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 4.0"
    }
  }
}

provider "google" {
  project = local.project_id
}

