---
subcategory: "Virtual Private Cloud DNS(VPCDNS)"
layout: "cloud"
page_title: "TencentCloud: cloud_vpcdns_domains"
sidebar_current: "docs-cloud-datasource-vpcdns_domains"
description: |-
  Provide a resource to query VPCDNS domain.
---

# cloud_vpcdns_domains

Provide a resource to query VPCDNS domain.

## Example Usage

```hcl
data "cloud_vpcdns_domain" "foo" {
  domain = "brucezylin.cc"
}
```

## Argument Reference

The following arguments are supported:

* `domain` - (Required, String) The domain of the VpcDns.
* `result_output_file` - (Optional, String) The file path to output the result.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `list` - list of domain.
  * `create_time` - create time of domain.
  * `dns_forward_status` - dns forward status of domain.
  * `domain_id` - ID of the domain.
  * `domain` - content of domain.
  * `record_count` - record count of domain.
  * `remark` - remark of domain.
  * `update_time` - update time of domain.


## Import

Vpc subnet instance can be imported, e.g.

```
$ terraform import cloud_vpcdns_domain.test domain_id
```

