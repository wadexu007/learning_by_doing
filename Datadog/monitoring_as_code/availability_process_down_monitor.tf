resource "datadog_monitor" "process_down_monitor" {
  name               = "System {{proc.name}} process is down"
  message            = "This monitor triggers when a process running on a server/instance goes down. Notify: @slack-sre-oncall"
  escalation_message = "Please investigate the system process, escalate to admin @wade.xu@demo.com"

  type  = "process alert"
  query = "processes('mongo').over('env:sre-eng-dev').rollup('count').last('10m') < 2"

  # type  = "service check"
  # query = "\"process.up\".over(\"process:sre-eng-dev-process\").by(\"proc\").last(1).pct_by_status()"

  monitor_thresholds {
    critical = 2
  }

  notify_no_data = true

  tags = ["app:test", "env:demo", "team:sre"]
}