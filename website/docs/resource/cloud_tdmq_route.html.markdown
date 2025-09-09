---
subcategory: "TDMQ for Pulsar(tpulsar)"
layout: "cloud"
page_title: "TencentCloud: cloud_tdmq_route"
sidebar_current: "docs-cloud-resource-tdmq_route"
description: |-
  Provide a resource to create a TDMQ Route.
---

# cloud_tdmq_route

Provide a resource to create a TDMQ Route.

## Example Usage

```hcl
resource "cloud_tdmq_Route" "foo" {
  cluster_name = "example111"
  remark       = "this is description111."
  tags = {
    "createdBy" = "terraform"
    "test"      = "111"
  }
  bind_cluster_id   = 0
  bind_cluster_name = "default"
}
```

## Argument Reference

The following arguments are supported:

* `net_type` - (Required, Int) net type.
* `cluster_id` - (Required, String) The Dedicated Cluster Id.
* `remark` - (Optional, String) remark of the route.
* `unique_subnet_id` - (Required, String) uniq subnet id, the value contains a maximum of 128 characters.
* `unique_vpc_id` - (Required, String) uniq vpc id.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `router_id` - The id of the route.
* `vip` - The vip of the route.
* `vport` - The vport of the route.


## Import

Tdmq Route can be imported, e.g.

```
$ terraform import cloud_tdmq_Route.test tdmq_id
```

