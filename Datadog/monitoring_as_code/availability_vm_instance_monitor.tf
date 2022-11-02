resource "datadog_monitor" "instance_down_monitor" {
  name               = "The instance is down on {{host.name}}"
  type               = "metric alert"
  message            = "VM instance is down. Notify: @slack-sre-oncall"
  escalation_message = "Please investigate the vm instance issue, escalate to admin @wade.xu@demo.com"

  query = "min(last_1m):avg:system.uptime{env:sre-eng-dev,platform:gce} by {host} < 1"

  monitor_thresholds {
    critical = 1
  }

  notify_no_data = true

  tags = ["app:test", "env:demo", "team:sre"]
}