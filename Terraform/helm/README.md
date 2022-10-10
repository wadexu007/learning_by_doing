## What's in this folder?

Demo how to use Terraform Helm provider to install Helm Chart

Install below two Charts with a specific release name.
* Nginx
* Redis

```
NAME            	NAMESPACE	REVISION	UPDATED                             	STATUS  	CHART        	APP VERSION
my-nginx-release	default  	1       	2022-10-10 20:26:18.274088 +0800 CST	deployed	nginx-13.2.10	1.23.1     
my-redis-release	default  	1       	2022-10-10 20:32:31.556624 +0800 CST	deployed	redis-17.3.4 	7.0.5 
```

#### Prerequisites
* Kubernetes 1.19+
* Terraform 1.2.9
* Hashicorp Helm Provider 2.7.0


#### Requirements
* Create a bucket for backend.tf to store Terraform state file


#### Authentication
* config_path = "~/.kube/config" in providers.tf
<br>

The easiest way is to supply a path to your kubeconfig file using the config_path attribute or using the KUBE_CONFIG_PATH environment variable. A kubeconfig file may have multiple contexts. If config_context is not specified, the provider will use the default context.

#### Cleanup
```
terraform destroy
```

#### Reference
https://registry.terraform.io/providers/hashicorp/helm/latest/docs
