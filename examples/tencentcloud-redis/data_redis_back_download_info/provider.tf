terraform {
  required_providers {
    tencentcloud = {
      source = "tencentcloudstack/tencentcloud"
      # 通过version指定版本
      # version = ">=1.60.18"
    }
  }
}


provider "tencentcloud" {
  secret_id = ""
  secret_key = ""
  region="chongqing"
}