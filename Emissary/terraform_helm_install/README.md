## What's in this folder?
This folder contains TF code on how to deploy Emissary Helm charts with Terraform.

### Prerequisites
* Helm Provider 2.6.0
* Emissary-Ingress 3.2.0
* Helm Chart emissary-ingress 8.2.0

### Understanding my Helm Chart structure
There are 3 charts as below

 * emissary_crds, custom charts for install CRDs, default CRDs namespace scope is emissary-system, do not suggest to change.
```
resource "helm_release" "emissary_crds" {
     name             = "emissary-crds"
     create_namespace = true # create emissary default namespace `emissary-system`
     namespace        = local.emissary_ns
     chart            = "../common/helm/repos/emissary-crds-8.2.0.tgz"
}
```
 * emissary_ingress, official charts for emissary-ingress service and deployments, you can change namespace if you want.
 * emissary_config, custom charts for config listeners, host, mapping and TLSContext. Keep same namespace with emissary_ingress.
```
resource "helm_release" "emissary_config" {
     name      = "emissary-config"
     namespace = local.emissary_ns
     chart     = "../common/helm/repos/emissary-config-8.2.0.tgz"
    
     values = [
    templatefile("${local.common_yaml_d}/emissary-advance-template.yaml", local.emissary_advance_map),
    local.emissary_config_yaml
     ]
    
     depends_on = [
       helm_release.emissary_svc
     ]
}
```

### Deployment
* [Deployment from Kustomize](../kustomize_install/)
* Deploy by Helm Chart with Terraform

To start, you need update locals variable to your's env in dev/emissary.tf. e.g.
```
locals {
  project_id     = "global-sre-dev"
  cluster_name   = "sre-gke"
  cluster_region = "us-central1"
}
```

From emissary-ingress 2.1, it removed CRDs from its charts, refer to official installation docs, apply CRDs is first step now. `kubectl apply -f https://app.getambassador.io/yaml/emissary/3.2.0/emissary-crds.yaml`, that's why I make a custom helm charts for install Emissary CRDs with Terraform Helm provider.


if you have multiple Emissary to be installed in one cluster, CRDs only need to be installed once. Default CRDs namespace scope is `emissary-system`. Emissary is installed in `emissary` namespace in my example.

Test in local
```
cd terraform_helm_install/dev

terraform init
terraform plan 
terraform apply
```

Notes: If helm provider > 2.7.0, plan will prompt the below error. Workaround is apply CRDs first. `terraform apply -target helm_release.emissary_crds`
```
│ Error: unable to build kubernetes objects from release manifest: resource mapping not found for name: "ambassador" namespace: "emissary-system" from "": no matches for kind "Module" in version "getambassador.io/v2"
│ ensure CRDs are installed first
```

We also need create a tls secret manually for demo https listeners (Optional: it's better to use [External Secrets](https://external-secrets.io/v0.6.0/) and custom `emissary_config` helm chart to add this kind or use this way [Terraform kubectl provider](../../Terraform/kubectl/emissary.tf)).
```
kubectl create secret -n secret tls tls-secret \
    --save-config --dry-run=client \
    --key ./xxx.key \
    --cert ./xxx.pem \
    -o yaml | 
  kubectl apply -f -
```

Install result as below:
```
% helm list -n emissary-system
NAME         	NAMESPACE      	REVISION	UPDATED                            	STATUS  	CHART              	APP VERSION
emissary-crds	emissary-system	1       	2022-10-20 10:09:30.72553 +0800 CST	deployed	emissary-crds-8.2.0	3.2.0         
```
```
% helm list -n emissary                             
NAME            	NAMESPACE	REVISION	UPDATED                             	STATUS  	CHART                 	APP VERSION
emissary-config 	emissary 	1       	2022-10-20 10:31:24.819555 +0800 CST	deployed	emissary-config-8.2.0 	3.2.0      
emissary-ingress	emissary 	1       	2022-10-20 10:29:33.705888 +0800 CST	deployed	emissary-ingress-8.2.0	3.2.0  
```


### Verify

Looking at config.tf related to Mapping/Host/TLSContext config, create a Nginx service in nginx namespace for curl testing.
```
helm install my-nginx bitnami/nginx --set service.type="ClusterIP" -n nginx --create-namespace

```

```
% curl https://dev.wadexu.cloud
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
html { color-scheme: light dark; }
body { width: 35em; margin: 0 auto;
font-family: Tahoma, Verdana, Arial, sans-serif; }
</style>
</head>
<body>
<h1>Welcome to nginx!</h1>
<p>If you see this page, the nginx web server is successfully installed and
working. Further configuration is required.</p>

<p>For online documentation and support please refer to
<a href="http://nginx.org/">nginx.org</a>.<br/>
Commercial support is available at
<a href="http://nginx.com/">nginx.com</a>.</p>

<p><em>Thank you for using nginx.</em></p>
</body>
</html>

```

### FAQ:

* tls secret related errors, make sure it has been created.
```
error:1404B42E:SSL routines:ST_CONNECT:tlsv1 alert protocol version
```

* Connection refused, almost like Listeners problem, double check config of it.
```
curl: (7) Failed to connect to dev.wadexu.cloud port 443 after 255 ms: Connection refused
```

* CRDs related errors 
```
│ Error: unable to build kubernetes objects from release manifest: [resource mapping not found for name: "my-resolver" namespace: "emissary-system" from "": no matches for kind "KubernetesEndpointResolver" in version "getambassador.io/v2"
│ ensure CRDs are installed first, resource mapping not found for name: "ambassador" namespace: "emissary-system" from "": no matches for kind "Module" in version "getambassador.io/v2"
│ ensure CRDs are installed first]
```
```
│ Error: unable to build kubernetes objects from release manifest: [resource mapping not found for name: "my-host-dev" namespace: "emissary-system" from "": no matches for kind "Host" in version "getambassador.io/v3alpha1"
│ ensure CRDs are installed first, resource mapping not found for name: "http-listener" namespace: "emissary-system" from "": no matches for kind "Listener" in version "getambassador.io/v3alpha1"
│ ensure CRDs are installed first, resource mapping not found for name: "https-listener" namespace: "emissary-system" from "": no matches for kind "Listener" in version "getambassador.io/v3alpha1"
│ ensure CRDs are installed first, resource mapping not found for name: "my-nginx-mapping" namespace: "emissary-system" from "": no matches for kind "Mapping" in version "getambassador.io/v3alpha1"
│ ensure CRDs are installed first]
```


### Known issue
https://github.com/emissary-ingress/emissary/issues/4637

<br>
