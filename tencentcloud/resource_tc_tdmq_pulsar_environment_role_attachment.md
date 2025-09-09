Provide a resource to create a TDMQ environment role.

Example Usage

```hcl
resource "cloud_tdmq_pulsar_cluster" "example" {
  cluster_name = "tf_example"
  remark       = "remark."
  tags = {
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

resource "cloud_tdmq_pulsar_role" "example" {
  role_name  = "tf_example"
  cluster_id = cloud_tdmq_pulsar_cluster.example.id
  remark     = "remark."
}

resource "cloud_tdmq_pulsar_environment_role_attachment" "example" {
  environ_id  = cloud_tdmq_pulsar_environment.example.environ_name
  role_name   = cloud_tdmq_pulsar_role.example.role_name
  permissions = ["produce", "consume"]
  cluster_id  = cloud_tdmq_pulsar_cluster.example.id
}
```