resource "datadog_dashboard" "demo" {
  title        = "Nginx Service"
  description  = "A Datadog Dashboard for the Nginx deployment"
  layout_type  = "ordered"

  widget {
    hostmap_definition {
      no_group_hosts  = true
      no_metric_hosts = true
      node_type       = "container"
      title           = "Kubernetes Pods"

      request {
        fill {
          q = "avg:process.stat.container.cpu.total_pct{image_name:docker.io/bitnami/nginx} by {host}"
        }
      }

      style {
        palette      = "hostmap_blues"
        palette_flip = false
      }
    }
  }

  widget {
    timeseries_definition {
      show_legend = false
      title       = "CPU Utilization"

      request {
        display_type = "line"
        q            = "top(avg:docker.cpu.usage{image_name:docker.io/bitnami/nginx} by {docker_image,container_id}, 10, 'mean', 'desc')"

        style {
          line_type  = "solid"
          line_width = "normal"
          palette    = "dog_classic"
        }
      }

      yaxis {
        include_zero = true
        max          = "auto"
        min          = "auto"
        scale        = "linear"
      }
    }
  }

  widget {
    alert_graph_definition {
      alert_id = datadog_monitor.gke_pod_nginx_monitor.id
      title    = "Kubernetes Node CPU"
      viz_type = "timeseries"
    }
  }

  widget {
    hostmap_definition {
      no_group_hosts  = true
      no_metric_hosts = true
      node_type       = "host"
      title           = "Kubernetes Nodes"

      request {
        fill {
          q = "avg:system.cpu.user{*} by {host}"
        }
      }

      style {
        palette      = "hostmap_blues"
        palette_flip = false
      }
    }
  }

  widget {
    timeseries_definition {
      show_legend = false
      title       = "Memory Utilization"
      request {
        display_type = "line"
        q            = "top(avg:docker.mem.in_use{image_name:docker.io/bitnami/nginx} by {container_name}, 10, 'mean', 'desc')"

        style {
          line_type  = "solid"
          line_width = "normal"
          palette    = "dog_classic"
        }
      }
      yaxis {
        include_zero = true
        max          = "auto"
        min          = "auto"
        scale        = "linear"
      }
    }
  }
}
