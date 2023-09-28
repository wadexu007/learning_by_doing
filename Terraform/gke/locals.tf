locals {
  # project details
  project = {
    project_id       = "adept-presence-396401"
    region           = "us-central1"
    network_name     = "wade-test-network"
  }

  # cluster details
  wade_cluster = {
    cluster_name                = "wade-gke"
    cluster_version             = "1.22.12-gke.500"
    subnet_name                 = "wade-gke"
    subnet_range                = "10.254.71.0/24"
    secondary_ip_range_pods     = "172.20.72.0/21"
    secondary_ip_range_services = "10.127.8.0/24"
    region                      = "us-central1"

    node_pools = [
      {
        name               = "app-pool"
        machine_type       = "n1-standard-2"
        node_locations     = join(",", slice(data.google_compute_zones.available.names, 0, 3))
        initial_node_count = 1
        min_count          = 1
        max_count          = 10
        max_pods_per_node  = 64
        disk_size_gb       = 100
        disk_type          = "pd-standard"
        image_type         = "COS"
        auto_repair        = true
        auto_upgrade       = false
        preemptible        = false
        max_surge          = 1
        max_unavailable    = 0
      }
    ]

    node_pools_labels = {
      all = {}
    }

    node_pools_tags = {
      all = ["k8s-nodes"]
    }

    node_pools_metadata = {
      all = {
        disable-legacy-endpoints = "true"
      }
    }

    node_pools_taints = {
      all = []
    }

    oauth_scopes = {
      all = [
        "https://www.googleapis.com/auth/monitoring",
        "https://www.googleapis.com/auth/compute",
        "https://www.googleapis.com/auth/devstorage.full_control",
        "https://www.googleapis.com/auth/logging.write",
        "https://www.googleapis.com/auth/service.management",
        "https://www.googleapis.com/auth/servicecontrol",
      ]
    }

    master_authorized_networks = [
      {
        display_name = "Whitelist 1"
        cidr_block   = "0.0.0.0/0" # need to change to your whitelist
      },
    ]
  }
}