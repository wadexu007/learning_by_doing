
# resource "tencentcloud_dnspod_domain_instance" "foo" {
#   domain = "hello.com"
#   remark = "this is demo"
# }

resource "tencentcloud_dnspod_record" "demo" {
  domain      = "hello.com"
  record_type = "A"
  record_line = "默认"
  value       = "1.2.3.9"
  sub_domain  = "demo"
}