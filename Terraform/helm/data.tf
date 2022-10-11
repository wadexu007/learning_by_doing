data "google_project" "this" {
  project_id = "global-sre-dev"
}

data "google_client_config" "this" {}

data "google_container_cluster" "this" {
  name     = "sre-mgmt"
  location = "us-west1"
  project  = data.google_project.this.project_id
}
