data "http" "okta_cloud_ip_ranges" {
  url = "https://s3.amazonaws.com/okta-ip-ranges/ip_ranges.json"

  request_headers = {
    Accept = "application/json"
  }
}

locals {
  # assign JSON output to a local variable
  okta_cloud_ip_ranges = jsondecode(data.http.okta_cloud_ip_ranges.response_body)
}

# output "okta_cloud_us1_ip_ranges_1" {
#   value       = local.okta_cloud_ip_ranges.us_cell_1.ip_ranges
# }

# firewall to okta in us_cell_1 region
resource "google_compute_firewall" "allow-cluster-okta" {
  name    = "fw-allow-egress-cluster-to-okta"
  network = var.network_name
  project = var.project

  allow {
    protocol = "tcp"
    ports    = ["443"]
  }

  priority    = 1000
  direction   = "EGRESS"
  target_tags = var.target_tags

  destination_ranges = slice(local.okta_cloud_ip_ranges.us_cell_1.ip_ranges, 0, length(local.okta_cloud_ip_ranges.us_cell_1.ip_ranges))
}
