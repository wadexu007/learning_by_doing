terraform {
  backend "gcs" {
    bucket = "global-sre-dev-terraform"
    prefix = "helm/state"
  }
}
