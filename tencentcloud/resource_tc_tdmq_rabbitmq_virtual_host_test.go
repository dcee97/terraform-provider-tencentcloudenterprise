package tencentcloud

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// go test -i; go test -test.run TestAccTencentCloudTdmqRabbitmqVirtualHostResource_basic -v
func TestAccTencentCloudTdmqRabbitmqVirtualHostResource_basic(t *testing.T) {
	//t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		CheckDestroy: testAccCheckTdmqRabbitmqVirtualHostDestroy,
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTdmqRabbitmqVirtualHost,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTdmqRabbitmqVirtualHostExists("cloud_tdmq_rabbitmq_virtual_host.example"),
					resource.TestCheckResourceAttrSet("cloud_tdmq_rabbitmq_virtual_host.example", "id"),
					resource.TestCheckResourceAttrSet("cloud_tdmq_rabbitmq_virtual_host.example", "instance_id"),
					resource.TestCheckResourceAttrSet("cloud_tdmq_rabbitmq_virtual_host.example", "virtual_host"),
					resource.TestCheckResourceAttrSet("cloud_tdmq_rabbitmq_virtual_host.example", "description"),
					resource.TestCheckResourceAttrSet("cloud_tdmq_rabbitmq_virtual_host.example", "trace_flag"),
				),
			},
			{
				ResourceName:      "cloud_tdmq_rabbitmq_virtual_host.example",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccTdmqRabbitmqVirtualHostUpdate,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTdmqRabbitmqVirtualHostExists("cloud_tdmq_rabbitmq_virtual_host.example"),
					resource.TestCheckResourceAttrSet("cloud_tdmq_rabbitmq_virtual_host.example", "id"),
					resource.TestCheckResourceAttrSet("cloud_tdmq_rabbitmq_virtual_host.example", "instance_id"),
					resource.TestCheckResourceAttrSet("cloud_tdmq_rabbitmq_virtual_host.example", "virtual_host"),
					resource.TestCheckResourceAttrSet("cloud_tdmq_rabbitmq_virtual_host.example", "description"),
					resource.TestCheckResourceAttrSet("cloud_tdmq_rabbitmq_virtual_host.example", "trace_flag"),
				),
			},
		},
	})
}

func testAccCheckTdmqRabbitmqVirtualHostDestroy(s *terraform.State) error {
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), logIdKey, logId)

	service := TdmqService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloud_tdmq_rabbitmq_virtual_host" {
			continue
		}

		ids := strings.Split(rs.Primary.ID, FILED_SP)
		if len(ids) != 2 {
			return fmt.Errorf("invalid ID %s", rs.Primary.ID)
		}
		instanceId := ids[0]
		virtualHost := ids[1]

		ret, err := service.DescribeTdmqRabbitmqVirtualHostById(ctx, instanceId, virtualHost)
		if err != nil {
			return err
		}

		if ret != nil {
			return fmt.Errorf("tdmq rabbitmq virtual_host still exist, id: %v", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckTdmqRabbitmqVirtualHostExists(re string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		logId := getLogId(contextNil)
		ctx := context.WithValue(context.TODO(), logIdKey, logId)

		rs, ok := s.RootModule().Resources[re]
		if !ok {
			return fmt.Errorf("tdcpg instance  %s is not found", re)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("tdcpg instance id is not set")
		}

		service := TdmqService{client: testAccProvider.Meta().(*TencentCloudClient).apiV3Conn}
		ids := strings.Split(rs.Primary.ID, FILED_SP)
		if len(ids) != 2 {
			return fmt.Errorf("invalid ID %s", rs.Primary.ID)
		}
		instanceId := ids[0]
		virtualHost := ids[1]

		ret, err := service.DescribeTdmqRabbitmqVirtualHostById(ctx, instanceId, virtualHost)
		if err != nil {
			return err
		}

		if ret == nil {
			return fmt.Errorf("tdmq rabbitmq virtual_host not found, id: %v", rs.Primary.ID)
		}

		return nil
	}
}

const testAccTdmqRabbitmqVirtualHost = `
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

# create virtual host
resource "cloud_tdmq_rabbitmq_virtual_host" "example" {
  instance_id  = cloud_tdmq_rabbitmq_vip_instance.example.id
  virtual_host = "tf-example-vhost"
  description  = "desc."
  trace_flag   = true
}
`

const testAccTdmqRabbitmqVirtualHostUpdate = `
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

# create virtual host
resource "cloud_tdmq_rabbitmq_virtual_host" "example" {
  instance_id  = cloud_tdmq_rabbitmq_vip_instance.example.id
  virtual_host = "tf-example-vhost"
  description  = "desc update."
  trace_flag   = false
}
`
