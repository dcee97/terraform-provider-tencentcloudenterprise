---
subcategory: "Cloud Object Storage(COS)"
layout: "cloud"
page_title: "TencentCloud: cloud_cos_buckets"
sidebar_current: "docs-cloud-datasource-cos_buckets"
description: |-
  Use this data source to query the COS buckets of the current Tencent Cloud user.
---

# cloud_cos_buckets

Use this data source to query the COS buckets of the current Tencent Cloud user.

## Example Usage

```hcl
data "cloud_cos_buckets" "cos_buckets" {
  bucket_prefix      = "tf-bucket-"
  result_output_file = "mytestpath"
}
```

## Argument Reference

The following arguments are supported:

* `result_output_file` - (Optional, String) Used to save results.


