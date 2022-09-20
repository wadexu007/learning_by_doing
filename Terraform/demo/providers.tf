terraform {
  required_version = ">= 1.2.9"

  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 4"
    }
  }
}

provider "google" {
  project = local.project.project_id
  region  = local.project.region
}
