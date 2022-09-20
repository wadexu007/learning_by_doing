terraform {
  backend "gcs" {
    bucket = "rogertest"
    prefix = "demo/state"
  }
}