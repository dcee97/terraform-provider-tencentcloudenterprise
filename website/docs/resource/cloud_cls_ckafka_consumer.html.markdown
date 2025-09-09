---
subcategory: "Cloud Log Service(CLS)"
layout: "cloud"
page_title: "TencentCloud: cloud_cls_ckafka_consumer"
sidebar_current: "docs-cloud-resource-cls_ckafka_consumer"
description: |-
  Provides a resource to create a cls ckafka_consumer
---

# cloud_cls_ckafka_consumer

Provides a resource to create a cls ckafka_consumer

## Example Usage

```hcl
resource "cloud_cls_ckafka_consumer" "ckafka_consumer" {
  compression  = 1
  need_content = true
  topic_id     = "7e34a3a7-635e-4da8-9005-88106c1fde69"

  ckafka {
    instance_id   = "ckafka-qzoeaqx8"
    instance_name = "ckafka-instance"
    topic_id      = "topic-c6tm4kpm"
    topic_name    = "name"
    vip           = "172.16.112.23"
    vport         = "9092"
  }

  content {
    enable_tag = true
    meta_fields = [
      "__FILENAME__",
      "__HOSTNAME__",
      "__PKGID__",
      "__SOURCE__",
      "__TIMESTAMP__",
    ]
    tag_json_not_tiled = true
    timestamp_accuracy = 2
  }
}
```

## Argument Reference

The following arguments are supported:

* `topic_id` - (Required, String, ForceNew) Topic id.
* `ckafka` - (Optional, List) Ckafka info.
* `compression` - (Optional, Int) Compression method. 0 for NONE, 2 for SNAPPY, 3 for LZ4.
* `content` - (Optional, List) Metadata information.
* `need_content` - (Optional, Bool) Whether to deliver the metadata information of the log.

The `ckafka` object supports the following:

* `instance_id` - (Required, String) Instance id.
* `instance_name` - (Required, String) Instance name.
* `topic_id` - (Required, String) Topic id of ckafka.
* `topic_name` - (Required, String) Topic name of ckafka.
* `vip` - (Required, String) Vip.
* `vport` - (Required, String) Vport.

The `content` object supports the following:

* `enable_tag` - (Required, Bool) Whether to deliver the TAG info.
* `meta_fields` - (Required, Set) Metadata info list.
* `tag_json_not_tiled` - (Optional, Bool) Whether to tiling tag json.
* `timestamp_accuracy` - (Optional, Int) Delivery timestamp precision,1 for second, 2 for millisecond.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.



## Import

cls ckafka_consumer can be imported using the id, e.g.

```
terraform import cloud_cls_ckafka_consumer.ckafka_consumer topic_id
```

