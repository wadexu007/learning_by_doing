## only monitor the demo nginx pod
resource "datadog_monitor" "gke_pod_nginx_monitor" {
  name               = "Kubernetes Nginx Pod Health"
  type               = "metric alert"
  message            = "Kubernetes Pods are not in an optimal health state. Notify: @slack-sre-oncall"
  escalation_message = "Please investigate the Kubernetes Pods, escalate to admin @wade.xu@demo.com"

  query = "max(last_1m):sum:docker.containers.running{cluster-name:finops-gke,image_name:docker.io/bitnami/nginx} <= 1"

  monitor_thresholds {
    ok       = 3
    warning  = 2
    critical = 1
  }

  notify_no_data = true

  tags = ["app:bitnami/nginx", "env:demo", "team:sre"]
}

