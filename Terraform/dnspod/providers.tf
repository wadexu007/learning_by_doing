terraform {
  required_version = ">= 1.2.9"

  required_providers {
    tencentcloud = {
      source = "tencentcloudstack/tencentcloud"
      version = "1.78.3"
    }
  }
}

provider "tencentcloud" {
  region     = "ap-guangzhou" # This is the TencentCloud region. It must be provided, but it can also be sourced from the `TENCENTCLOUD_REGION` environment variables. The default input value is ap-guangzhou.
}
