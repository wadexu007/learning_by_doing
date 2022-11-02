resource "datadog_monitor" "log_monitor" {
  name               = "Log monitor - Catch an error from log"
  type               = "log alert"
  message            = <<-EOT
  An error log xxx has been caught by Datadog Log Alert.
  Notify: @slack-sre-oncall
  EOT
  escalation_message = "Please investigate the issue from this log, escalate to admin @wade.xu@demo.com"

  query = "logs(\"env:sre-eng-dev \"*Error: xxx is currently suspended*\"\").index(\"*\").rollup(\"count\").last(\"5m\") >= 3"

  monitor_thresholds {
    critical = 3
    warning  = 1
  }

  notify_no_data = true

  tags = ["app:test", "env:demo", "team:sre"]
}