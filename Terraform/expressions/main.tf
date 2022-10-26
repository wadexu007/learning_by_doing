locals {
  # project details
  project = {
    project_id    = "global-sre-dev"
    network_name  = "wade-test-network"
    region        = "asia-east2"
  }

  # subnets details, one dmz subnet, another gke subnet with secondary subnets
  subnets = [
    {
      subnet_name           = "dmz-zone-asia-east2"
      subnet_ip             = "10.127.1.0/25"
      subnet_region         = "asia-east2"
      subnet_private_access = true
      subnet_flow_logs      = false
    },
    {
      subnet_name           = "gke-zone-asia-east2"
      subnet_ip             = "10.127.1.128/25"
      subnet_region         = "asia-east2"
      subnet_private_access = true
      subnet_flow_logs      = false
      secondary_subnets     = {
        gke-pods = "10.117.0.0/20"
        gke-svcs = "10.117.136.0/24"
      }
    },
  ]
}

resource "google_compute_network" "default" {
  project                 = local.project.project_id
  name                    = local.project.network_name
  auto_create_subnetworks = false
  routing_mode            = "GLOBAL"
}

# create subnets with `for each` and `dynamic` expressions
resource "google_compute_subnetwork" "main" {
  for_each = { for s in local.subnets : s["subnet_name"] => s }

  name          = lookup(each.value, "subnet_name")
  project       = local.project.project_id
  ip_cidr_range = lookup(each.value, "subnet_ip")
  region        = lookup(each.value, "subnet_region")
  network       = google_compute_network.default.self_link

  dynamic "secondary_ip_range" {
    for_each = lookup(each.value, "secondary_subnets", {})

    content {
      range_name    = secondary_ip_range.key
      ip_cidr_range = secondary_ip_range.value
    }
  }

  private_ip_google_access = lookup(each.value, "subnet_private_access", false)
}

resource "google_compute_address" "nat_addr" {
  count   = 2
  name    = "nat-external-address-${count.index}"
  region  = local.project.region
  project = local.project.project_id
}

# manage cloud routers for dynamic routing
resource "google_compute_router" "router" {
  name    = "my-router"
  region  = local.project.region
  network = google_compute_network.default.self_link

  bgp {
    keepalive_interval = 20
    asn                = "64528"
    advertise_mode     = "CUSTOM"
    advertised_groups  = []
  }
}

# just a example for dynamic and functions feature
# set source_subnetwork_ip_ranges_to_nat=ALL_SUBNETWORKS_ALL_PRIMARY_IP_RANGES is an easiest way
resource "google_compute_router_nat" "my_nat" {
  name                               = "my-nat"
  project                            = local.project.project_id
  router                             = google_compute_router.router.name
  region                             = google_compute_router.router.region
  nat_ip_allocate_option             = "MANUAL_ONLY"
  source_subnetwork_ip_ranges_to_nat = "LIST_OF_SUBNETWORKS"
  nat_ips                            = google_compute_address.nat_addr.*.self_link

  icmp_idle_timeout_sec            = "30"
  tcp_established_idle_timeout_sec = "1200"
  tcp_transitory_idle_timeout_sec  = "30"
  udp_idle_timeout_sec             = "30"

  dynamic "subnetwork" {
    for_each =  { for s in local.subnets : s["subnet_name"] => s }

    content {
      name                    = lookup(subnetwork.value, "subnet_name")
      source_ip_ranges_to_nat = length(keys(lookup(subnetwork.value, "secondary_subnets", {}))) == 0 ? ["PRIMARY_IP_RANGE"] : ["LIST_OF_SECONDARY_IP_RANGES"]
      secondary_ip_range_names = keys(lookup(subnetwork.value, "secondary_subnets", {}))
    }
  }
  # Monitor errors because of manual NAT IP allocation.
  log_config {
    enable = true
    filter = "ERRORS_ONLY"
  }
}