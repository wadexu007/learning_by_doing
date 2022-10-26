## What's in this folder?

The folder contains more features with Terraform

### Features
* lookup
* for_each
* count
* function
* dynamic blocks
* conditional expressions
* for expressions

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
