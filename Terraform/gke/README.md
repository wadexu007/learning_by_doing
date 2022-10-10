# Deployment GKE Terraform

### Provision the GKE cluster
Update backend.tf -> bucket

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
