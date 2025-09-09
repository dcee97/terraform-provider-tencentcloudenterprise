---
subcategory: "Cloud Object Storage(CSP)"
layout: "cloud"
page_title: "TencentCloud: cloud_csp_buckets"
sidebar_current: "docs-cloud-datasource-csp_buckets"
description: |-
  Use this data source to query the COS buckets of the current Tencent Cloud user.
---

# cloud_csp_buckets

Use this data source to query the COS buckets of the current Tencent Cloud user.

## Example Usage

```hcl
data "cloud_csp_buckets" "csp_buckets" {
  result_output_file = "mytestpath"
}
```

## Argument Reference

The following arguments are supported:

* `result_output_file` - (Optional, String) Used to save results.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `bucket_list` - csp bucket list.
  * `create_date` - bucket create date.
  * `location` - bucket region.
  * `name` - bucket name.
* `owner` - The owner of the bucket. Each element contains the following.


