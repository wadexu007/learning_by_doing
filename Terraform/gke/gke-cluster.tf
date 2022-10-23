
data "google_compute_zones" "available" {
  region = "us-central1"
  status = "UP"
}

resource "google_compute_network" "default" {
  project                 = local.project.project_id
  name                    = local.project.network_name
  auto_create_subnetworks = false
  routing_mode            = "GLOBAL"
}

resource "google_compute_subnetwork" "wade-gke" {
  project       = local.project.project_id
  network       = google_compute_network.default.name
  name          = local.wade_cluster.subnet_name
  ip_cidr_range = local.wade_cluster.subnet_range
  region        = local.wade_cluster.region

  secondary_ip_range {
    range_name    = format("%s-secondary1", local.wade_cluster.cluster_name)
    ip_cidr_range = local.wade_cluster.secondary_ip_range_pods
  }

  secondary_ip_range {
    range_name    = format("%s-secondary2", local.wade_cluster.cluster_name)
    ip_cidr_range = local.wade_cluster.secondary_ip_range_services
  }

  private_ip_google_access = true

}

resource "google_service_account" "sa-wade-test" {
  account_id   = "sa-wade-test"
  display_name = "sa-wade-test"
}


module "wade-gke" {
  source = "terraform-google-modules/kubernetes-engine/google//modules/beta-private-cluster"
  version = "23.1.0"

  project_id = local.project.project_id
  name       = local.wade_cluster.cluster_name

  kubernetes_version     = local.wade_cluster.cluster_version
  region                 = local.wade_cluster.region
  network                = google_compute_network.default.name
  subnetwork             = google_compute_subnetwork.wade-gke.name
  master_ipv4_cidr_block = "10.1.0.0/28"
  ip_range_pods          = google_compute_subnetwork.wade-gke.secondary_ip_range.0.range_name
  ip_range_services      = google_compute_subnetwork.wade-gke.secondary_ip_range.1.range_name

  service_account                 = google_service_account.sa-wade-test.email
  master_authorized_networks      = local.wade_cluster.master_authorized_networks
  master_global_access_enabled    = false
  istio                           = false
  issue_client_certificate        = false
  enable_private_endpoint         = false
  enable_private_nodes            = true
  remove_default_node_pool        = true
  enable_shielded_nodes           = false
  identity_namespace              = "enabled"
  node_metadata                   = "GKE_METADATA"
  horizontal_pod_autoscaling      = true
  enable_vertical_pod_autoscaling = false

  node_pools              = local.wade_cluster.node_pools
  node_pools_oauth_scopes = local.wade_cluster.oauth_scopes
  node_pools_labels       = local.wade_cluster.node_pools_labels
  node_pools_metadata     = local.wade_cluster.node_pools_metadata
  node_pools_taints       = local.wade_cluster.node_pools_taints
  node_pools_tags         = local.wade_cluster.node_pools_tags

}