resource "datadog_monitor" "host_low_disk_monitor" {
  name               = "Low disk space on {{host.name}}"
  type               = "metric alert"
  message            = <<-EOT
  This monitor triggers when a server or instance is running low on disk space.
  Notify: @slack-sre-oncall
  EOT
  escalation_message = "Please investigate the instance low disk space issue, escalate to admin @wade.xu@demo.com"

  query = "max(last_30m):max:system.disk.in_use{env:sre-eng-dev,region:cn} by {device,host} > 0.9"

  monitor_thresholds {
    warning           = 0.7
    warning_recovery  = 0.65
    critical          = 0.9
    critical_recovery = 0.85
  }

  notify_no_data = true

  tags = ["app:test", "env:demo", "team:sre"]
}