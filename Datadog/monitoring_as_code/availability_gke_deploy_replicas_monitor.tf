## monitor all deployment replicas in the cluster
resource "datadog_monitor" "gke_deployment_monitor" {
  name               = "Kubernetes Deploy Replicas Availability"
  type               = "query alert"
  message            = "Kubernetes Deploy Replicas is zero. Service is down. Notify: @slack-sre-oncall"
  escalation_message = "Please investigate the Kubernetes Deployment issue, escalate to admin @wade.xu@demo.com"

  query = "avg(last_5m):avg:kubernetes_state.deployment.replicas_available{cluster_name:sre-gke} by {kube_deployment} < 1"

  monitor_thresholds {
    critical = 1
  }

  notify_no_data = true

  tags = ["app:bitnami/nginx", "env:demo", "team:sre"]
}

