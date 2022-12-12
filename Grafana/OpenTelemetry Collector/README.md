## What is this folder?
The folder contains all resources to show how to automate instrument Golang app (OpenTelemetry SDK) to send `Trace` data --> `OpenTelemetry Collector` --> Grafana Tempo in Grafana Cloud.

Similar case to [this one](../Grafana%20Tempo/), but replace `Grafanan agent` to `OpenTelemetry Collector`

## OpenTelemetry Collector
Vendor-agnostic way to receive, process and export telemetry data.

## Configure
Review config in `otel-collector-trace.yaml` and do necessary update, e.g.
Replace grafana tempo endpoint and `<base64 data>` with below command output value:
```
echo -n "<Your user id>:<Your Grafana API key>" | base64
```

## Install otel-collector
```
kubectl apply -f otel-collector-trace.yaml
```

## Deploy a demo App
Here is k8s manifests [deployment yaml](./demo-app.yaml) and [source code](../../Golang/demo_app_with_instrumentation/) for reference, replace `OTEL_EXPORTER_OTLP_ENDPOINT` if need, my otel-collector endpoint is `otel-collector.otel.svc.cluster.local:4317`. 

```
kubectl apply -f demo-app.yaml
```

Kubectl port-forwarding can be used to connect to demo app API without exposing the service.
```
kubectl port-forward svc/bookstore -n exercise 8080:8080
```

Call demo app API
```
curl -X POST 'http://localhost:8080/books' -d '{"title":"Cloud Native","author":"wadexu"}' | jq
curl -X POST 'http://localhost:8080/books' -d '{"title":"Linux","author":"wadexu"}' | jq
curl -X GET http://localhost:8080/books | jq
```

## View Trace
Search for the trace in Grafana Cloud by navigating to Explore and choosing your traces data source.

## What's More
otel-collector also can collect logs and metrics, e.g. Add k8s cluster metrics to otel-collector.


Review `otel-collector-trace-metrics.yaml` and replace necessary value, e.g. prometheus `remotewrite` endpoint, user id and api key.
```
kubectl apply -f otel-rbac.yaml

kubectl apply -f otel-collector-trace-metrics.yaml
```

* auto instrumented and custom metrics, refer to [demo app](../../Golang/demo_app_with_instrumentation/) and [deployment yaml](./demo-app.yaml)
* k8s_cluster metrics
```
K8s_container_cpu_limit
K8s_container_cpu_request
K8s_container_memory_limit
K8s_container_memory_request
K8s_container_ready
K8s_container_restarts
K8s_daemonset_current_scheduled_nodes
K8s_daemonset_desired_scheduled_nodes
K8s_daemonset_misscheduled_nodes
K8s_daemonset_ready_nodes
K8s_deployment_available
K8s_deployment_desired
K8s_namespace_phase
K8s_node_condition_ready
k8s_pod_phase
k8s_replicaset_available
k8s_replicaset_desired
k8s_resource_quota_hard_limit
k8s_resource_quota_used
```

To collect Host level metrics `hostmetrics` and `kubeletstats`, we need deploy otel collector as daemonset, you can review `otel-collector-daemonset.yaml` for reference.

* Host and Kubelets
```
container_cpu_time
container_cpu_utilization
container_filesystem_available
container_filesystem_capacity
container_filesystem_usage
container_memory_available
container_memory_major_page_faults
container_memory_page_faults
container_memory_rss
container_memory_usage
container_memory_working_set
k8s_node_cpu_time
k8s_node_cpu_utilization
k8s_node_filesystem_available
k8s_node_filesystem_capacity
k8s_node_filesystem_usage
k8s_node_memory_available
k8s_node_memory_major_page_faults
k8s_node_memory_page_faults
k8s_node_memory_rss
k8s_node_memory_usage
k8s_node_memory_working_set
k8s_node_network_errors
k8s_node_network_io
k8s_pod_cpu_time
k8s_pod_cpu_utilization
k8s_pod_filesystem_available
k8s_pod_filesystem_capacity
k8s_pod_filesystem_usage
k8s_pod_memory_available
k8s_pod_memory_major_page_faults
k8s_pod_memory_page_faults
k8s_pod_memory_rss
k8s_pod_memory_usage
k8s_pod_memory_working_set
k8s_pod_network_errors
k8s_pod_network_io
system_cpu_load_average_15m
system_cpu_load_average_1m
system_cpu_load_average_5m
system_cpu_time
system_disk_io
system_disk_io_time
system_disk_merged
system_disk_operation_time
system_disk_operations
system_disk_pending_operations
system_disk_weighted_io_time
system_memory_usage
system_network_connections
system_network_dropped
system_network_errors
system_network_io
system_network_packets
```

## Useful links
* https://grafana.com/blog/2021/04/13/how-to-send-traces-to-grafana-clouds-tempo-service-with-opentelemetry-collector/
* https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/receiver/k8sclusterreceiver
* https://github.com/open-telemetry/opentelemetry-collector/blob/main/examples/k8s/otel-config.yaml

<br>