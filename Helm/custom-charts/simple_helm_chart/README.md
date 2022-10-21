## Create Your First Helm Chart

#### Prerequisites
* Helm 3.2.0+
* kubeconfig: Configured


#### Create a Helm chart
Replace with a chart name then start requires below simple command
```
helm create mychart

ls mychart/
Chart.yaml	charts		templates	values.yaml
```

#### Helm chart's structure
* Chart.yaml
* values.yaml
* templates
* charts

##### Chart.yaml
Chart.yaml define what the chart is.


##### templates
The most important part of the chart is the template directory. 
It holds all the configurations for your application that will be deployed into the cluster.

As you can see below, this application has a basic deployment, ingress, service account, and service. This directory also includes a test directory, which includes a test for a connection into the app. Each of these application features has its own template files under templates/:

* ingress.yaml
* service.yaml
* serviceaccount.yaml
* deployment.yaml
* hpa.yaml
* tests

##### charts
There is another directory, called charts, which is empty. It allows you to add dependent charts that are needed to deploy your application.

This is a far more advanced configuration, so leave the charts/ folder empty for now.


##### values.yaml
values.yaml define what values will be in it at deployment.

Template files are set up with formatting that collects deployment information from the values.yaml file. Therefore, to customize your Helm chart, you need to edit the values file. 

Update values.yaml these fields
```
nameOverride: "my-awesome-app"

fullnameOverride: "my-first-chart"
```

#### Deploy
Deploy from local
```
helm install my-first-chart mychart/ --values mychart/values.yaml 
```

#### Check Result
```
helm list

NAME          	NAMESPACE	REVISION	UPDATED                            	STATUS  	CHART        	APP VERSION
my-first-chart	default  	1       	2022-10-17 16:20:33.07764 +0800 CST	deployed	mychart-0.1.0	1.16.0  

```

## Another simple Helm Chart for create namespace
* user-namespaces
Template folder only contains a namespace.yaml

### Use via Terraform
```
helm package user-namespaces

# move output file user-namespaces-1.1.0.tgz to some local folder or upload to a bucket or publish to github
```

```
locals {
  project_id     = "global-sre-dev"
  cluster_name   = "sre-gke"
  cluster_region = "us-central1"
  emissary_ns    = "emissary-system"
  chart_version  = "8.2.0"

  emissary_ns_yaml = <<-EOT
    namespaces:
    - name: ${local.emissary_ns}
      owner: wadexu
  EOT
}

# from local chart
resource "helm_release" "emissary_ns" {
  name   = "emissary-ns"
  chart  = "some_folder/helm/repos/user-namespaces-1.1.0.tgz"
  values = [local.emissary_ns_yaml]
}
```

You can refer to [more example](../../../Emissary/terraform_helm_install/dev/emissary.tf) about how to Deploy Helm Charts With Terraform
