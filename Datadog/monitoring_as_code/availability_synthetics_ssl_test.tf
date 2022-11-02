# Example Usage (Synthetics SSL test)
resource "datadog_synthetics_test" "test_ssl" {
  type    = "api"
  subtype = "ssl"
  request_definition {
    host = "example.org"
    port = 443
  }
  assertion {
    type     = "certificate"
    operator = "isInMoreThan"
    target   = 30
  }
  locations = ["aws:eu-central-1"]
  options_list {
    tick_every         = 900
    accept_self_signed = true
  }
  name    = "SSL Expire test on example.org"
  message = "SSL Certificate will be expired in 30 days, notify @wade.xu@demo.com"
  tags    = ["app:demo", "env:demo", "team:sre"]

  status = "live"
}
