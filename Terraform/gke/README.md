## Deployment GKE Terraform

#### Prerequisites
* Terraform 1.2.9
* Google Cloud SDK

```
gcloud version

Google Cloud SDK 397.0.0
alpha 2022.08.05
beta 2022.08.05
bq 2.0.75
core 2022.08.05
gsutil 5.11
```

### Authentication
The account should have full permission to GCP project
```
gcloud auth application-default login
```

### Provision the GKE cluster
```
# valid LOCATION values are `asia`, `eu` or `us`
export BUCKET_NAME=wadexu007-terraform-dev
gsutil mb -l $LOCATION gs://$BUCKET_NAME
gsutil versioning set on gs://$BUCKET_NAME
```
Update backend.tf -> bucket to your name
Update locals.tf -> project_id and others if you want

### Deployment
```
terraform init

terraform plan

terraform apply
```


### Adding the cluster to your context
```shell
gcloud container clusters get-credentials $(terraform output -raw cluster_name) \
    --region $(terraform output -raw region) \
    --project $(terraform output -raw project_id)
```

From this point you can now use [kubectl](https://kubernetes.io/docs/reference/kubectl/) to manage your cluster and deploy Kubernetes configurations to it.
