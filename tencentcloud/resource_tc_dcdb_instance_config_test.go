package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudDcdbInstanceConfigResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDcdbInstanceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("cloud_dcdb_instance_config.instance_config", "id"),
					resource.TestCheckResourceAttr("cloud_dcdb_instance_config.instance_config", "rs_access_strategy", "0"),
				),
			},
			{
				Config: testAccDcdbInstanceConfig_update,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("cloud_dcdb_instance_config.instance_config", "id"),
					resource.TestCheckResourceAttr("cloud_dcdb_instance_config.instance_config", "rs_access_strategy", "1"),
				),
			},
			{
				ResourceName:      "cloud_dcdb_instance_config.instance_config",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccDcdbInstanceConfig = CommonPresetDcdb + `

resource "cloud_dcdb_instance_config" "instance_config" {
  instance_id = local.dcdb_id
  rs_access_strategy = 0
}

`

const testAccDcdbInstanceConfig_update = CommonPresetDcdb + `

resource "cloud_dcdb_instance_config" "instance_config" {
  instance_id = local.dcdb_id
  rs_access_strategy = 1
}

`
