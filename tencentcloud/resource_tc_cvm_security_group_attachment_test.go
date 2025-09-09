package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudCvmSecurityGroupAttachmentResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCvmSecurityGroupAttachment,
				Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttrSet("cloud_cvm_security_group_attachment.security_group_attachment", "id")),
			},
			{
				ResourceName:      "cloud_cvm_security_group_attachment.security_group_attachment",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccCvmSecurityGroupAttachment = `
resource "cloud_cvm_security_group_attachment" "security_group_attachment" {
	security_group_id = "sg-a0212ii1"
	instance_id = "ins-cr2rfq78"
}
`
