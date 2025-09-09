---
subcategory: "Cloud Kafka(ckafka)"
layout: "cloud"
page_title: "TencentCloud: cloud_ckafka_instance"
sidebar_current: "docs-cloud-resource-ckafka_instance"
description: |-
  Use this resource to create ckafka instance.
---

# cloud_ckafka_instance

Use this resource to create ckafka instance.

~> **NOTE:** It only support create prepaid ckafka instance.

## Example Usage

### Basic Instance

```hcl
variable "vpc_id" {
  default = "vpc-68vi2d3h"
}

variable "subnet_id" {
  default = "subnet-ob6clqwk"
}

data "cloud_availability_zones" "gz" {
  name    = "ap-guangzhou-3"
  product = "ckafka"
}

resource "cloud_ckafka_instance" "kafka_instance" {
  instance_name      = "ckafka-instance-type-tf-test"
  zone_id            = data.cloud_availability_zones.gz.zones.0.id
  period             = 1
  vpc_id             = var.vpc_id
  subnet_id          = var.subnet_id
  msg_retention_time = 1300
  renew_flag         = 0
  kafka_version      = "2.4.1"
  disk_size          = 1000
  disk_type          = "CLOUD_BASIC"

  specifications_type = "standard"
  instance_type       = 2

  config {
    auto_create_topic_enable   = true
    default_num_partitions     = 3
    default_replication_factor = 3
  }

  dynamic_retention_config {
    enable = 1
  }
}
```

### Multi zone Instance

```hcl
variable "vpc_id" {
  default = "vpc-68vi2d3h"
}

variable "subnet_id" {
  default = "subnet-ob6clqwk"
}

data "cloud_availability_zones" "gz3" {
  name    = "ap-guangzhou-3"
  product = "ckafka"
}

data "cloud_availability_zones" "gz6" {
  name    = "ap-guangzhou-6"
  product = "ckafka"
}

resource "cloud_ckafka_instance" "kafka_instance" {
  instance_name   = "ckafka-instance-maz-tf-test"
  zone_id         = data.cloud_availability_zones.gz3.zones.0.id
  multi_zone_flag = true
  zone_ids = [
    data.cloud_availability_zones.gz3.zones.0.id,
    data.cloud_availability_zones.gz6.zones.0.id
  ]
  period             = 1
  vpc_id             = var.vpc_id
  subnet_id          = var.subnet_id
  msg_retention_time = 1300
  renew_flag         = 0
  kafka_version      = "1.1.1"
  disk_size          = 500
  disk_type          = "CLOUD_BASIC"

  config {
    auto_create_topic_enable   = true
    default_num_partitions     = 3
    default_replication_factor = 3
  }

  dynamic_retention_config {
    enable = 1
  }
}
```

## Argument Reference

The following arguments are supported:

* `band_width` - (Required, Int) Instance bandwidth in MBps.
* `disk_size` - (Required, Int) Disk Size. Its interval varies with bandwidth, and the input must be within the interval, which can be viewed through the control. If it is not within the interval, the plan will cause a change when first created.
* `instance_name` - (Required, String, ForceNew) Instance name.
* `subnet_id` - (Required, String) Subnet id.
* `vpc_id` - (Required, String) Vpc id.
* `zone_id` - (Required, Int) Available zone id.
* `cluster_id` - (Optional, Int, ForceNew) Cluster-ID represents  the cluster to which Ckafa belongs, associated with cluster_name only needs to select one of them.
* `instance_id` - (Optional, String) Ckafka instance ID.
* `msg_retention_time` - (Optional, Int) The maximum retention time of instance logs, in minutes. the default is 10080 (7 days), the maximum is 30 days, and the default 0 is not filled, which means that the log retention time recovery policy is not enabled.
* `renew_flag` - (Optional, Int) Prepaid automatic renewal mark, 0 means the default state, the initial state, 1 means automatic renewal, 2 means clear no automatic renewal (user setting).
* `tag_set` - (Optional, Map) Tag set of instance.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - ID of the resource.
* `vip` - Vip of instance.
* `vport` - Type of instance.


## Import

ckafka instance can be imported using the instance_id, e.g.

```
$ terraform import cloud_ckafka_instance.foo ckafka-f9ife4zz
```

