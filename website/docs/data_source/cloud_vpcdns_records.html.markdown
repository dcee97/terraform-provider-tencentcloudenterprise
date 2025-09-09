---
subcategory: "Virtual Private Cloud DNS(VPCDNS)"
layout: "cloud"
page_title: "TencentCloud: cloud_vpcdns_records"
sidebar_current: "docs-cloud-datasource-vpcdns_records"
description: |-
  Provide a resource to query VPCDNS records.
---

# cloud_vpcdns_records

Provide a resource to query VPCDNS records.

## Example Usage

```hcl
data "cloud_vpcdns_domain" "foo" {
  domain             = "brucezylin.cc"
  dns_forward_status = "ENABLED"
  tags = {
    "createdBy" = "terraform3"
  }
}
```

## Argument Reference

The following arguments are supported:

* `domain_id` - (Required, Int) The domain id of the VpcDns.
* `result_output_file` - (Optional, String) The file path to output the result.
* `sub_domain` - (Optional, String) The domain of the VpcDns.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `list` - list of records.
  * `create_time` - create time of record.
  * `domain_id` - ID of the domain.
  * `enabled` - is enabled, 0: disabled, 1: enabled.
  * `mx` - MX.
  * `record_id` - ID of the record.
  * `record_type` - type of record.
  * `status` - status of record.
  * `sub_domain` - sub domain of the record.
  * `ttl` - TTL.
  * `update_time` - update time of record.


## Import

Vpc subnet instance can be imported, e.g.

```
$ terraform import cloud_vpcdns_domain.test domain_id
```

