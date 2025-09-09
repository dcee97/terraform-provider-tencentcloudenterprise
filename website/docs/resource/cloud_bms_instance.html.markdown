---
subcategory: "Bare Metal Server(BMS)"
layout: "cloud"
page_title: "TencentCloud: cloud_bms_instance"
sidebar_current: "docs-cloud-resource-bms_instance"
description: |-
  Provide a resource to create a bms instance
---

# cloud_bms_instance

Provide a resource to create a bms instance

## Example Usage

```hcl
resource "cloud_bms_instance" "cloud_bms_instance-wuua" {
  operating_system      = "your operating system name"
  operating_system_type = "Linux"
  password              = "your password"
  subnet_id             = "your subnet id"
  vpc_id                = "your vpc id"
  flavor_id             = "your flavorid"
  instance_name         = "your instance"
  ipv6_address          = true
  raid_type             = "NORAID"
  availability_zone     = "your zone name"
  security_service      = true
  whistle_service       = true
}
```

## Argument Reference

The following arguments are supported:

* `availability_zone` - (Required, String, ForceNew) The available zone for the bms instance.
* `flavor_id` - (Required, String) The flavor of the instance.
* `instance_name` - (Required, String) The name of the instance. The max length of instance_name is 60, and default value is `Terraform-bms-Instance`.
* `operating_system_type` - (Required, String) The operating system type of the instance.
* `operating_system` - (Required, String) The operate system of bms instance.
* `password` - (Required, String) Password for the instance. In order for the new password to take effect, the instance will be restarted after the password change. Modifying will cause the instance reset.
* `raid_type` - (Required, String) The type of raid.
* `subnet_id` - (Required, String) The ID of a VPC subnet.
* `vpc_id` - (Required, String) The ID of a VPC network.
* `allocate_public_ip` - (Optional, Bool, ForceNew) Associate a public IP address with an instance in a VPC or Classic. Boolean value, Default is false.
* `force_delete` - (Optional, Bool) Indicate whether to force delete the instance. Default is `false`. If set true, the instance will be permanently deleted instead of being moved into the recycle bin. Note: only works for `PREPAID` instance.
* `hostname` - (Optional, String) The hostname of the instance. Windows instance: The name should be a combination of 2 to 15 characters comprised of letters (case insensitive), numbers, and hyphens (-). Period (.) is not supported, and the name cannot be a string of pure numbers. Other types (such as Linux) of instances: The name should be a combination of 2 to 60 characters, supporting multiple periods (.). The piece between two periods is composed of letters (case insensitive), numbers, and hyphens (-). Modifying will cause the instance reset.
* `instance_count` - (Optional, Int) The number of instances to be purchased. Value range:[1,100]; default value: 1.
* `internet_max_bandwidth_out` - (Optional, Int) Maximum outgoing bandwidth to the public network, measured in Mbps (Mega bits per second). This value does not need to be set when `allocate_public_ip` is false. Size is related to eip.
* `placement_group_id` - (Optional, String, ForceNew) The ID of a placement group.
* `private_ip` - (Optional, String) The private IP to be assigned to this instance, must be in the provided subnet and available.
* `project_id` - (Optional, Int) The project the instance belongs to, default to 0.
* `security_service` - (Optional, Bool) enhance service for security, it is disable by default.
* `system_disk_size` - (Optional, Int) Size of the system disk. unit is GB, Default is 50GB. If modified, the instance may force stop.
* `tags` - (Optional, Map) A mapping of tags to assign to the resource. For tag limits, please refer to [Use Limits](https://intl.cloud.tencent.com/document/product/651/13354).

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `create_time` - Create time of the instance.
* `instance_status` - Current status of the instance.


## Import

Placement group can be imported using the id, e.g.

```
$ terraform import cloud_placement_group.foo ps-ilan8vjf
```

