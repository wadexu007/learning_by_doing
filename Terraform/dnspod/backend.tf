terraform {
  backend "gcs" {
    bucket = "global-sre-dev-terraform"
    prefix = "dns/state"
  }
}
