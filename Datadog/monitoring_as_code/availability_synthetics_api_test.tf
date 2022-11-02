# Example Usage (Synthetics API test)
resource "datadog_synthetics_test" "test_syn_api" {
  type    = "api"
  subtype = "http"

  request_definition {
    method = "GET"
    url    = "https://www.example.org"
  }

  assertion {
    type     = "statusCode"
    operator = "is"
    target   = "200"
  }

  locations = ["aws:us-west-2"]
  options_list {
    tick_every          = 900
    min_location_failed = 1
  }

  name    = "Synthetics test on example.org URL"
  message = <<EOT
Oh no! Light from the your app is no longer shining!
Notify: @wade.xu@demo.com
EOT
  tags    = ["app:demo", "env:demo", "team:sre"]

  status = "live"
}
