package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudNeedFixVpcNotifyRoutesResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccVpcNotifyRoutes,
				Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttrSet("cloud_vpc_notify_routes.notify_routes", "id")),
			},
			{
				ResourceName:      "cloud_vpc_notify_routes.notify_routes",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccVpcNotifyRoutes = `

resource "cloud_vpc_notify_routes" "notify_routes" {
  route_table_id = ""
  route_item_ids = []
}

`
