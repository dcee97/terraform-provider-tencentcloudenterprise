resource "cloud_eip_instance" "foo" {
  name = "awesome_eip_example"

  tags = {
    "test" = "test"
  }
}

data "cloud_eips" "tags" {
  tags = cloud_eip_instance.foo.tags
}