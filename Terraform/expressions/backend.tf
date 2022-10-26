terraform {
  backend "gcs" {
    bucket = "global-sre-dev-terraform"
    prefix = "expressions/state"
  }
}