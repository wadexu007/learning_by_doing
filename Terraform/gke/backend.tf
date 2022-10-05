terraform {
  backend "gcs" {
    bucket = "xperiences-eng-cn-dev-dmz-terraform-test"
    prefix = "demo/state"
  }
}