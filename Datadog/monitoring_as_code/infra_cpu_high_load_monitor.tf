resource "datadog_monitor" "host_high_cpu_monitor" {
  name               = "High CPU load on {{host.name}}"
  type               = "metric alert"
  message            = "This monitor triggers when a server or instance is running high CPU. Notify: @slack-sre-oncall"
  escalation_message = "Please investigate the issue, escalate to admin @wade.xu@demo.com"

  query = "avg(last_30m):avg:system.cpu.user{env:sre-eng-dev} by {host} > 90"

  monitor_thresholds {
    critical          = 90
    critical_recovery = 88
  }

  notify_no_data = true

  tags = ["app:test", "env:demo", "team:sre"]
}