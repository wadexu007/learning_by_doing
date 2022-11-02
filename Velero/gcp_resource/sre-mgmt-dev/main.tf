locals {
  project_id = "global-sre-dev"
}

module "velero" {
  source       = "../_module_velero"
  project_name = local.project_id
}