Provides a resource to create a tdmq rabbitmq_user

Example Usage

```hcl
data "cloud_availability_zones" "zones" {
  name = "ap-guangzhou-6"
}

# create vpc
resource "cloud_vpc" "vpc" {
  name       = "vpc"
  cidr_block = "10.0.0.0/16"
}

# create vpc subnet
resource "cloud_vpc_subnet" "subnet" {
  name              = "subnet"
  vpc_id            = cloud_vpc.vpc.id
  availability_zone = "ap-guangzhou-6"
  cidr_block        = "10.0.20.0/28"
  is_multicast      = false
}

# create rabbitmq instance
resource "cloud_tdmq_rabbitmq_vip_instance" "example" {
  zone_ids                              = [data.cloud_availability_zones.zones.zones.0.id]
  vpc_id                                = cloud_vpc.vpc.id
  subnet_id                             = cloud_vpc_subnet.subnet.id
  cluster_name                          = "tf-example-rabbitmq-vip-instance"
  node_spec                             = "rabbit-vip-basic-1"
  node_num                              = 1
  storage_size                          = 200
  enable_create_default_ha_mirror_queue = false
  auto_renew_flag                       = true
  time_span                             = 1
}

# create rabbitmq user
resource "cloud_tdmq_rabbitmq_user" "example" {
  instance_id     = cloud_tdmq_rabbitmq_vip_instance.example.id
  user            = "tf-example-user"
  password        = "$Password"
  description     = "desc."
  tags            = ["management", "monitoring", "example"]
  max_connections = 3
  max_channels    = 3
}
```

Import

tdmq rabbitmq_user can be imported using the id, e.g.

```
terraform import cloud_tdmq_rabbitmq_user.example amqp-8xzx822q#tf-example-user
```
