package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudClbSecurityGroupAttachmentResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccClbSecurityGroupAttachment,
				Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttrSet("cloud_clb_security_group_attachment.security_group_attachment", "id")),
			},
			{
				ResourceName:      "cloud_clb_security_group_attachment.security_group_attachment",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccClbSecurityGroupAttachment = `

resource "cloud_clb_security_group_attachment" "security_group_attachment" {
  security_group = "sg-ijato2x1"
  load_balancer_ids = ["lb-5dnrkgry"]
}

`
