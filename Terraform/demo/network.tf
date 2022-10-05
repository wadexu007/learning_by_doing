resource "google_compute_network" "default" {
  project                 = local.project.project_id
  name                    = local.project.network_name
  auto_create_subnetworks = true
  routing_mode            = "GLOBAL"
}
