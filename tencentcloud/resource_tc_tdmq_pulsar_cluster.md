Provide a resource to create a TDMQ Pulsar cluster.

Example Usage

```hcl
resource "cloud_tdmq_pulsar_cluster" "example" {
  cluster_name = "tf_example"
  remark       = "remark."
  tags = {
    "createdBy" = "terraform"
  }
}
```

Import

Tdmq cluster can be imported, e.g.

```
$ terraform import cloud_tdmq_pulsar_cluster.example pulsar-78bwjaj8epxv
```