---
subcategory: "TDSQL for MySQL(DCDB)"
layout: "cloud"
page_title: "TencentCloud: cloud_dcdb_db_sync_mode_config"
sidebar_current: "docs-cloud-resource-dcdb_db_sync_mode_config"
description: |-
  Provides a resource to create a dcdb db_sync_mode_config
---

# cloud_dcdb_db_sync_mode_config

Provides a resource to create a dcdb db_sync_mode_config

## Example Usage

```hcl
resource "cloud_dcdb_db_sync_mode_config" "config" {
  instance_id = "%s"
  sync_mode   = 2
}
```

### # Import

dcdb db_sync_mode_config can be imported using the id, e.g.

```hcl
terraform import cloud_dcdb_db_sync_mode_config.db_sync_mode_config db_sync_mode_config_id
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required, String) ID of the instance for which to modify the sync mode. The ID is in the format of `tdsql-ow728lmc`.
* `sync_mode` - (Required, Int) Sync mode. Valid values: `0` (async), `1` (strong sync), `2` (downgradable strong sync).

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



