terraform {
  backend "gcs" {
    bucket = "global-sre-dev-terraform"
    prefix = "emissary/demo"
  }
}
