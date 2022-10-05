module "firewall_okta" {
  source       = "../_modules/firewall_okta"
  project      = local.project.project_id
  network_name = local.project.network_name
}
