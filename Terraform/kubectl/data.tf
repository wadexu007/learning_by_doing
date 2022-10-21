data "google_project" "this" {
  project_id = local.project_id
}

data "google_client_config" "this" {}

data "google_container_cluster" "this" {
  name     = local.cluster_name
  location = local.cluster_region
  project  = data.google_project.this.project_id
}
