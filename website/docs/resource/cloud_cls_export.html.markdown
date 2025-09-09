---
subcategory: "Cloud Log Service(CLS)"
layout: "cloud"
page_title: "TencentCloud: cloud_cls_export"
sidebar_current: "docs-cloud-resource-cls_export"
description: |-
  Provides a resource to create a cls export
---

# cloud_cls_export

Provides a resource to create a cls export

## Example Usage

```hcl
resource "cloud_cls_export" "export" {
  topic_id  = "7e34a3a7-635e-4da8-9005-88106c1fde69"
  log_count = 2
  query     = "select count(*) as count"
  from      = 1607499107000
  to        = 1607499108000
  order     = "desc"
  format    = "json"
}
```

## Argument Reference

The following arguments are supported:

* `from` - (Required, Int, ForceNew) Export start time.
* `log_count` - (Required, Int, ForceNew) Export amount of log.
* `query` - (Required, String, ForceNew) Export query rules.
* `to` - (Required, Int, ForceNew) Export end time.
* `topic_id` - (Required, String, ForceNew) Topic id.
* `format` - (Optional, String, ForceNew) Log export format.
* `order` - (Optional, String, ForceNew) Log export time sorting. desc or asc.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

cls export can be imported using the id, e.g.

```
terraform import cloud_cls_export.export topic_id#export_id
```

