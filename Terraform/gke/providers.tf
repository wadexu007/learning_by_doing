terraform {
  required_version = ">= 1.2.9"

  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 4.0"
    }
    google-beta = {
      source  = "hashicorp/google-beta"
      version = "~> 4.0"
    }
  }
}

provider "google" {
  project = local.project.project_id
  region  = local.project.region
}

provider "google-beta" {
  project = local.project.project_id
  region  = local.project.region
}
