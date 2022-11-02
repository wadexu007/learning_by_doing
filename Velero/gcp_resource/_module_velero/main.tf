resource "google_storage_bucket" "velero-backup" {
  name     = format("lr-velero-%s", var.project_name)
  project  = var.project_name
  location = var.bucket_location
  lifecycle_rule {
    condition {
      age = var.bucket_lifecycle_age
    }
    action {
      type = "Delete"
    }
  }
}

resource "google_service_account" "velero" {
  project      = var.project_name
  account_id   = "velero"
  display_name = "velero"
}

resource "google_project_iam_custom_role" "velero-backup-server" {
  role_id     = "velero.backup.server"
  title       = "Velero Backup Server"
  description = "This role contains permissions required for Velero to backup Kubernetes"
  permissions = [
    "compute.disks.get",
    "compute.disks.create",
    "compute.disks.createSnapshot",
    "compute.snapshots.get",
    "compute.snapshots.create",
    "compute.snapshots.useReadOnly",
    "compute.snapshots.delete",
    "compute.zones.get",
    "storage.objects.create",
    "storage.objects.delete",
    "storage.objects.get",
    "storage.objects.list",
    "iam.serviceAccounts.signBlob",
  ]
}

resource "google_project_iam_member" "velero-backup-server" {
  project = var.project_name
  role    = google_project_iam_custom_role.velero-backup-server.id
  member  = format("serviceAccount:%s", google_service_account.velero.email)
}

resource "google_service_account_iam_member" "velero-workload-identity-user" {
  service_account_id = google_service_account.velero.id
  role               = "roles/iam.workloadIdentityUser"
  member             = format("serviceAccount:%s.svc.id.goog[velero/velero]", var.project_name)
}

