# Demo how to use Terraform deploy resource on GCP

### Prerequisites
create a bucket for project_1/backend.tf to store Terraform state file

* Kubernetes 1.19+
* Terraform 1.2.9

### Deployment
```
cd project_1

terraform init

terraform plan

terraform apply
```

#### Cleanup
```
terraform destroy
```
