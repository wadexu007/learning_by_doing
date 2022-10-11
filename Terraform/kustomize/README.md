## What's in this folder?

Demo how to use Terraform Kustomize provider to install resource on Kubernetes.

#### Prerequisites
* Kubernetes 1.19+
* Terraform 1.2.9
* Kustomization Provider 0.9.0


#### Requirements
* Create a bucket for backend.tf to store Terraform state file

```
# valid LOCATION values are `asia`, `eu` or `us`
gsutil mb -l $LOCATION gs://$BUCKET_NAME
gsutil versioning set on gs://$BUCKET_NAME
```

#### Authentication
Same as this [part](https://github.com/wadexu007/learning_by_doing/tree/main/Terraform/helm#authentication)

#### Installation
```
terraform init

terraform plan

terraform apply
```

Below two resources will be installed.
* [demo-app](../../Kustomize/demo-manifests/services/demo-app/)
* sealed-secrets

```
NAME                                         READY   STATUS    RESTARTS      AGE
demo-app-88f5c9f9c-qp7ns                     1/1     Running   0             4m9s
sealed-secrets-controller-5d67647cd8-9kw8p   1/1     Running   0             40s
```


#### Cleanup
```
terraform destroy
```

#### Reference
https://registry.terraform.io/providers/kbst/kustomization/latest/docs
https://www.kubestack.com/framework/documentation/cluster-service-modules/#custom-manifests
https://www.kubestack.com/catalog/sealed-secrets/