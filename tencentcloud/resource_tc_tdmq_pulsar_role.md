Provide a resource to create a TDMQ Pulsar role.

Example Usage

```hcl
resource "cloud_tdmq_pulsar_cluster" "example" {
  cluster_name = "tf_example"
  remark       = "remark."
  tags         = {
    "createdBy" = "terraform"
  }
}

resource "cloud_tdmq_pulsar_role" "example" {
  role_name  = "role_example"
  cluster_id = cloud_tdmq_pulsar_cluster.example.id
  remark     = "remark."
}
```
