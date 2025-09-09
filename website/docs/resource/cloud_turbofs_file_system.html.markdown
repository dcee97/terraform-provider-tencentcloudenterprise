---
subcategory: "CFS Turbo(Turbofs)"
layout: "cloud"
page_title: "TencentCloud: cloud_turbofs_file_system"
sidebar_current: "docs-cloud-resource-turbofs_file_system"
description: |-
  Provides a resource to create a turbo file system(TurboFS).
---

# cloud_turbofs_file_system

Provides a resource to create a turbo file system(TurboFS).

## Example Usage

```hcl
resource "cloud_turbofs_file_system" "foo" {
  name              = "test_file_system"
  availability_zone = "ap-guangzhou-3"
  access_group_id   = "pgroup-7nx89k7l"
  protocol          = "NFS"
  vpc_id            = "vpc-ah9fbkap"
  subnet_id         = "subnet-9mu2t9iw"
}
```

## Argument Reference

The following arguments are supported:

* `availability_zone` - (Required, String, ForceNew) The available zone that the file system locates at.
* `capacity` - (Required, Float64) Capacity of the TURBO file system in TiB. The value should be an integer multiple of the spec policy.
* `name` - (Required, String) Name of a file system.
* `subnet_id` - (Required, String, ForceNew) ID of a subnet.
* `vpc_id` - (Required, String, ForceNew) ID of a VPC network.
* `access_group_id` - (Optional, String) ID of a access group.
* `bandwidth_resource_pkg_id` - (Optional, String) ID of the bandwidth resource package bound to the file system. Each file system can only be bound to one bandwidth resource package.
* `cfs_version` - (Optional, String) CFS file system version. Valid value, the default is `4.0`.
* `encrypted` - (Optional, Bool) Whether the file system is encrypted. If left empty, it defaults to not encrypted.
* `kms_key_id` - (Optional, String) ID of the KMS key used for encryption. This field is required when `encrypted` is set to true.
* `mount_ip` - (Optional, String, ForceNew) IP of mount point.
* `net_interface` - (Optional, String, ForceNew) Network type. Valid value is `vpc`.
* `pool_id` - (Optional, String) ID of the resource pool. If not specified, the system will automatically select the default resource pool.
* `project_id` - (Optional, String) ID of the project.
* `protocol` - (Optional, String, ForceNew) File service protocol. Valid value, the default is `TURBO`.
* `snapshot_id` - (Optional, String) Snapshot ID.
* `storage_resource_pkg_id` - (Optional, String) ID of the storage resource package bound to the file system. Each file system can only be bound to one storage resource package.
* `storage_type` - (Optional, String, ForceNew) File service StorageType. Valid values are `TP` and `TB`. and the default is `TB`.
* `tags` - (Optional, Map) Instance tags.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `create_time` - Create time of the file system.
* `fs_id` - Id of the file system.


## Import

Cloud file system can be imported using the id, e.g.

```
$ terraform import cloud_turbofs_file_system.foo turbofs-6hgquxmj
```

