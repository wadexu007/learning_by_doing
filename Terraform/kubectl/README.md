## What's in this folder?

Demo how to use Terraform kubectl provider to apply kubernetes resources yaml file directly.

#### Prerequisites
* Kubernetes 1.22
* Terraform 1.2.9
* Kubectl Provider 1.14.0

#### Requirements
* Create a bucket for backend.tf to store Terraform state file

```
# valid LOCATION values are `asia`, `eu` or `us`
gsutil mb -l $LOCATION gs://$BUCKET_NAME
gsutil versioning set on gs://$BUCKET_NAME
```

#### Installation
```
terraform init

terraform plan

terraform apply
```

#### Reference
https://registry.terraform.io/providers/gavinbunney/kubectl/latest/docs