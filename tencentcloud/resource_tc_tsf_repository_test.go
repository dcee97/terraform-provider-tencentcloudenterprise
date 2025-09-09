package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudNeedFixTsfRepositoryResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheckCommon(t, ACCOUNT_TYPE_TSF) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccTsfRepository,
				Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttrSet("cloud_tsf_repository.repository", "id")),
			},
			{
				ResourceName:      "cloud_tsf_repository.repository",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccTsfRepository = `

resource "cloud_tsf_repository" "repository" {
    repository_name = ""
  repository_type = ""
  bucket_name = ""
  bucket_region = ""
  directory = ""
  repository_desc = ""
    }

`
