terraform {
  backend "gcs" {
    bucket = "wadexu007-terraform-dev"
    prefix = "state/gke"
  }
}