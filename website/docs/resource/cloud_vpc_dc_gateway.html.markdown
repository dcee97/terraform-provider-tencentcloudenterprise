---
subcategory: "Virtual Private Cloud(VPC)"
layout: "cloud"
page_title: "TencentCloud: cloud_vpc_dc_gateway"
sidebar_current: "docs-cloud-resource-vpc_dc_gateway"
description: |-
  Provides a resource to creating direct connect gateway instance.
---

# cloud_vpc_dc_gateway

Provides a resource to creating direct connect gateway instance.

## Example Usage

```hcl
resource "cloud_vpc" "main" {
  name       = "ci-vpc-instance-test"
  cidr_block = "10.0.0.0/16"
}

resource "cloud_vpc_dc_gateway" "vpc_main" {
  name                = "ci-cdg-vpc-test"
  network_instance_id = cloud_vpc.main.id
  network_type        = "VPC"
  gateway_type        = "NAT"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, String) Name of the DCG.
* `network_instance_id` - (Required, String, ForceNew) If the `network_type` value is `VPC`, the available value is VPC ID. But when the `network_type` value is `CCN`, the available value is CCN instance ID.
* `network_type` - (Required, String, ForceNew) Type of associated network. Valid value: `VPC` and `CCN`.
* `band_with` - (Optional, Int) the bandwith speed limit of the gateway.
* `gateway_type` - (Optional, String, ForceNew) Type of the gateway. Valid value: `NORMAL` and `NAT`. Default is `NORMAL`. NOTES: CCN only supports `NORMAL` and a VPC can create two DCGs, the one is NAT type and the other is non-NAT type.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `cnn_route_type` - Type of CCN route. Valid value: `BGP` and `STATIC`. The property is available when the DCG type is CCN gateway and BGP enabled.
* `create_time` - Creation time of resource.
* `enable_bgp` - Indicates whether the BGP is enabled.


## Import

Direct connect gateway instance can be imported, e.g.

```
$ terraform import cloud_vpc_dc_gateway.instance dcg-id
```

