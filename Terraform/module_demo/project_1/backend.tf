terraform {
  backend "gcs" {
    bucket = "wadexu007-terraform-dev"
    prefix = "stage/tf_module_demo"
  }
}
