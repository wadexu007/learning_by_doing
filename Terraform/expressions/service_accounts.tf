locals {
  # service account details
  # sa_name is GCP SA name
  # kube_sa_name is kubernetes SA name which want to use Workload Identiy to impersonate GCP SA.
  sa_details = [
    {
        sa_name           = "uc-api"
        sa_display_name   = "User center api SA"
        kube_sa_name      = "my-kube-uc-api"
        kube_sa_namespace = "app"
    },
    {
        sa_name           = "mgs-api"
        sa_display_name   = "Msg center api SA"
        kube_sa_name      = "my-kube-msg-api"
        kube_sa_namespace = "app"
    },
    {
        sa_name           = "data-api"
        sa_display_name   = "Data process api SA"
        kube_sa_name      = "my-kube-data-api"
        kube_sa_namespace = "data"
    },
    {
        # no need workload identity binding
        sa_name           = "portal-fe"
        sa_display_name   = "Portal Web site"
    }
  ]
}

# for each create SA
resource "google_service_account" "sa-create" {
  for_each     = {for sa in local.sa_details: sa.sa_name => sa}
  account_id   = each.value.sa_name
  display_name = each.value.sa_display_name
  project      = local.project.project_id
}

# add conditions, not all SA need workload identity iam binding 
resource "google_service_account_iam_binding" "sa-workload-identity-user" {
  for_each           = {for sa in local.sa_details: sa.sa_name => sa if lookup(sa, "kube_sa_name", "") != ""}
  service_account_id = format("projects/%s/serviceAccounts/%s@%s.iam.gserviceaccount.com", local.project.project_id, each.value.sa_name, local.project.project_id)
  role               = "roles/iam.workloadIdentityUser"
  members            = [format("serviceAccount:%s.svc.id.goog[%s/%s]", local.project.project_id, each.value.kube_sa_namespace, each.value.kube_sa_name)]
}