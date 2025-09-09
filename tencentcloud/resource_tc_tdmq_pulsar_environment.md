Provide a resource to create a TDMQ namespace.

Example Usage

```hcl
resource "cloud_tdmq_pulsar_cluster" "example" {
  cluster_name = "tf_example"
  remark       = "remark."
  tags         = {
    "createdBy" = "terraform"
  }
}

resource "cloud_tdmq_pulsar_environment" "example" {
  environ_name = "tf_example"
  msg_ttl      = 300
  cluster_id   = cloud_tdmq_pulsar_cluster.example.id
  retention_policy {
    time_in_minutes = 60
    size_in_mb      = 10
  }
  remark = "remark."
}
```

Import

Tdmq namespace can be imported, e.g.

```
$ terraform import cloud_tdmq_pulsar_cluster.example tf_example#pulsar-78bwjaj8epxv
```
