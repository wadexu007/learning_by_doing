terraform {
  backend "gcs" {
    bucket = "global-sre-dev-terraform"
    prefix = "demo/module"
  }
}
