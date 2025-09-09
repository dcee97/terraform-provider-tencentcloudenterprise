---
subcategory: "Cloud Load Balancer(CLB)"
layout: "cloud"
page_title: "TencentCloud: cloud_clb_resources"
sidebar_current: "docs-cloud-datasource-clb_resources"
description: |-
  Use this data source to query detailed information of clb resources
---

# cloud_clb_resources

Use this data source to query detailed information of clb resources

## Example Usage

```hcl
data "cloud_clb_resources" "resources" {
  filters {
    name   = "isp"
    values = ["BGP"]
  }
}
```

## Argument Reference

The following arguments are supported:

* `filters` - (Optional, List) Filter to query the list of AZ resources as detailed below: zone - String - Optional - Filter by AZ, such as ap-guangzhou-1. isp -- String - Optional - Filter by the ISP. Values: BGP, CMCC, CUCC and CTCC.
* `result_output_file` - (Optional, String) Used to save results.

The `filters` object supports the following:

* `name` - (Required, String) Filter name.
* `values` - (Required, Set) Filter value array.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `resource_set` - List of resources supported by the AZ.
  * `domain` - domain of clb.
  * `load_balancer_id` - id of clb.
  * `load_balancer_name` - name of clb.
  * `load_balancer_type` - type of clb.
  * `load_balancer_vips` - vips of clb.
  * `master_zone` - master zone of clb.
    * `local_zone` - is the local zone.
    * `zone_id` - zone id of clb.
    * `zone_name` - zone name of clb.
    * `zone_region` - zone region of clb.
    * `zone` - zone of clb.


