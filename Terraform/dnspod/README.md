# Demo how to use Terraform create DNS records in Tecent Cloud DNSPod

### Prerequisites
* Create a bucket for backend.tf to store Terraform state file
* Generate API secrets in TecentCloud DNSPod [Console](https://console.dnspod.cn/account/token/apikey)
* Provide credentials via TENCENTCLOUD_SECRET_ID and TENCENTCLOUD_SECRET_KEY environment variables

```
$ export TENCENTCLOUD_SECRET_ID="my-secret-id"
$ export TENCENTCLOUD_SECRET_KEY="my-secret-key"
```

### Execution
```
terraform init

terraform plan

terraform apply
```
