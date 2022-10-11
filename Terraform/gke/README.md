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

### Provision the GKE cluster
Update backend.tf -> bucket

```
# valid LOCATION values are `asia`, `eu` or `us`
gsutil mb -l $LOCATION gs://$BUCKET_NAME
gsutil versioning set on gs://$BUCKET_NAME
```

### Authentication
The account should have full permission to GCP project
```
gcloud auth login

gcloud auth list
```

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
