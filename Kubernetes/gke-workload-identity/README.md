## What is GKE Workload Identity

[Workload Identity](https://cloud.google.com/kubernetes-engine/docs/concepts/workload-identity) is the recommended way for your workloads running on Google Kubernetes Engine (GKE) to access Google Cloud services in a secure and manageable way.

### [Use Workload Identity](https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity)
This folder shows you how to enable and configure Workload Identity on your Google Kubernetes Engine (GKE) clusters. Workload Identity allows workloads in your GKE clusters to impersonate Identity and Access Management (IAM) service accounts to access Google Cloud services.

#### Enable Workload Identity
To enable workload identity on an existing cluster, first enable it on the cluster like below command.

Replace me first: 
* $CLUSTER_NAME = your cluster name
* $POOL = your cluster pool name
* $PROJECT_ID = your gcp project id
* $SECRETS_NAMESPACE = here use secrets as example


**GCloud Command**
```
gcloud container clusters update $CLUSTER_NAME --workload-pool=$PROJECT_ID.svc.id.goog
```

Next enable workload metadata config on the node pool in which the pod will run:
```
gcloud beta container node-pools update $POOL --cluster $CLUSTER_NAME --workload-metadata-from-node=GKE_METADATA_SERVER
```

**Terraform Code snippet**

```
resource "google_container_cluster" "primary" {
    ......

    workload_metadata_config {
      mode = "GKE_METADATA"
    }
}
```

if you use [terraform gke module](https://github.com/terraform-google-modules/terraform-google-kubernetes-engine), refer to [my example](../../Terraform/gke/gke-cluster.tf)
```
  node_metadata                   = "GKE_METADATA"
```

#### GCP IAM
**GCloud Command**

Create the policy binding:
```
gcloud iam service-accounts add-iam-policy-binding --role roles/iam.workloadIdentityUser --member "serviceAccount:$PROJECT_ID.svc.id.goog[$SECRETS_NAMESPACE/kubernetes-external-secrets]" my-secrets-sa@$PROJECT_ID.iam.gserviceaccount.com
```

Grant GCP service account access to secrets:
```
gcloud projects add-iam-policy-binding $PROJECT_ID --member=serviceAccount:my-secrets-sa@$PROJECT_ID.iam.gserviceaccount.com --role=roles/secretmanager.secretAccessor
```

**Terraform Code snippet**
```
#service account for external secrets
resource "google_service_account" "secrets-manager" {
  account_id   = "my-secrets-sa"
  display_name = "My secrets manager service account"
}

resource "google_service_account_iam_binding" "secrets-manager-binding" {
  service_account_id = google_service_account.secrets-manager.name
  role               = "roles/iam.workloadIdentityUser"
  members = [
    format("serviceAccount:%s.svc.id.goog[$SECRETS_NAMESPACE/kubernetes-external-secrets]", "$PROJECT_ID")
  ]
}

resource "google_project_iam_member" "secrets-manager-binding-accessor" {
  role    = "roles/secretmanager.secretAccessor"
  member  = format("serviceAccount:%s", google_service_account.secrets-manager.email)
  project = "$PROJECT_ID"
}
```

**Check binding result**
```
% gcloud iam service-accounts get-iam-policy my-secrets-sa@$PROJECT_ID.iam.gserviceaccount.com
bindings:
- members:
  - serviceAccount:$PROJECT_ID.svc.id.goog[secrets/kubernetes-external-secrets]
  role: roles/iam.workloadIdentityUser
etag: BwXaTccZWbo=
version: 1
```


#### Configure applications to use Workload Identity
Example about [external-secrets](../external-secrets/) to use Workload Identity

```
apiVersion: v1
kind: ServiceAccount
metadata:
  name: kubernetes-external-secrets
  annotations:
    iam.gke.io/gcp-service-account: my-secrets-sa@$PROJECT_ID.iam.gserviceaccount.com
```

[Deploy external secrets](../external-secrets/)

Double check SA annotations
```
kubectl describe serviceaccount kubernetes-external-secrets -n secrets 
Name:                kubernetes-external-secrets
Namespace:           secrets
Labels:              app.kubernetes.io/name=kubernetes-external-secrets
Annotations:         iam.gke.io/gcp-service-account: my-secrets-sa@$PROJECT_ID.iam.gserviceaccount.com
Mountable secrets:   kubernetes-external-secrets-token-pgj8t
Tokens:              kubernetes-external-secrets-token-pgj8t
Events:              <none>
```

#### Verify the Workload Identity setup
Create a Pod that uses the annotated Kubernetes service account and curl the service-accounts endpoint.
```
kubectl apply -f wi-test.yaml
```

```
kubectl exec -it workload-identity-test \
  --namespace secrets \
  -- /bin/bash
```

If the service accounts are correctly configured, the IAM service account email address is listed as the active (and only) identity. This demonstrates that by default, the Pod acts as the IAM service account's authority when calling Google Cloud APIs.
```
root@workload-identity-test:/# curl -H "Metadata-Flavor: Google" http://169.254.169.254/computeMetadata/v1/instance/service-accounts/default/email
my-secrets-sa@$PROJECT_ID.iam.gserviceaccount.comroot@workload-identity-test:/# 
```
<br>
