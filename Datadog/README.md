## What Is Datadog

[Datadog](https://www.datadoghq.com/) is an observability service for cloud-scale applications, providing monitoring of servers, databases, tools, and services, through a SaaS-based data analytics platform. It provides cloud-scale monitoring and security for metrics, traces and logs in one unified platform.


### DataDog Agent Installation

There are 3 ways in [offical docs](https://docs.datadoghq.com/containers/kubernetes/installation/?tab=operator) for installing the Datadog Agent on Kubernetes, but here I tried another way with Terraform Helm Provider.

#### Prerequisites
* Kubernetes 1.22
* Terraform 1.2.9
* Terraform Helm Provider 2.6.0

My example setup DD agent and test in GCP. Similar way if you use AWS/Azure.
1. Get [Datadog API and application keys](https://app.datadoghq.com/organization-settings/api-keys?_gl=1*1b7ys11*_ga*OTUwMDQ4OTguMTY1OTUzMTIyMQ..*_ga_KN80RDFSQK*MTY2NzM3MjU5OC4xMS4xLjE2NjczNzMyODQuMjIuMC4w) first.
2. Upload keys to [GCP Secret Manager](https://cloud.google.com/secret-manager)
3. Execution

```
cd datadog_agent_install

terraform init
terraform plan 
terraform apply
```

### Monitoring as Code
MaC [examples](./monitoring_as_code/) as below:

* Availability level
  * availability_gke_demo_pods_monitor
  * availability_gke_deploy_replicas_monitor
  * availability_process_down_monitor
  * availability_synthetics_api_test
  * availability_synthetics_ssl_test
  * availability_vm_instance_monitor
* Infrastructure level
  * infra_cpu_high_load_monitor
  * infra_disk_low_space_monitor
* System level
  * system_log_monitor
  * system_process_high_cpu_monitor
* Dashboard
  * datadog_dashboard_example

#### Install from Terraform
```
cd monitoring_as_code

terraform init
terraform plan 
terraform apply
```

<br>
