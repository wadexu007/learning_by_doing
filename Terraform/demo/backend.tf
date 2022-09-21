terraform {
  backend "gcs" {
    bucket = "wadexu007"
    prefix = "demo/state"
  }
}