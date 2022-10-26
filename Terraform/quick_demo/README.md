# Demo how to use Terraform deploy resource on GCP

### Prerequisites
create a bucket for backend.tf to store Terraform state file

here use GCS as example
```
# valid LOCATION values are `asia`, `eu` or `us`
gsutil mb -l $LOCATION gs://$BUCKET_NAME
gsutil versioning set on gs://$BUCKET_NAME
```

* Kubernetes 1.19+
* Terraform 1.2.9

### Deployment
```
terraform init

terraform plan

terraform apply
```

#### Cleanup
```
terraform destroy
```
