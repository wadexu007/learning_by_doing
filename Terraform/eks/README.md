# Deployment EKS Terraform

#### Provision the EKS cluster
Update providers.tf -> bucket

1. Option 1: Export AWS access and security to environment variables

```shell
export AWS_ACCESS_KEY_ID=xxx
export AWS_SECRET_ACCESS_KEY=xxx
```

2. Option 2: Add a profile to your AWS credentials file
```
aws configure
# or 
vim ~/.aws/credentials

[default]
aws_access_key_id=xxx
aws_secret_access_key=xxx

```

3. Verify account
```
aws sts get-caller-identity 
```

4. Execute TF commands

```
terraform init

terraform plan

terraform apply
```

This process should take approximately 10 minutes. Upon completion, Terraform will print your configuration's outputs.
```
Apply complete! Resources: 51 added, 0 changed, 0 destroyed.

Outputs:

cluster_endpoint = "https://xxxxxxxxx.yl4.cn-north-1.eks.amazonaws.com.cn"
cluster_id = "test-eks-2022"
cluster_name = "test-eks-2022"
region = "cn-north-1"

```

#### Adding the cluster to your context
```shell
aws eks --region $(terraform output -raw region) update-kubeconfig \
    --name $(terraform output -raw cluster_name)
```
From this point you can now use [kubectl](https://kubernetes.io/docs/reference/kubectl/) to manage your cluster and deploy Kubernetes configurations to it.

Example:
```shell
kubectl cluster-info

kubectl get nodes
```