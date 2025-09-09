package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudCfsAutoSnapshotPolicyAttachmentResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCfsAutoSnapshotPolicyAttachment,
				Check:  resource.ComposeTestCheckFunc(resource.TestCheckResourceAttrSet("cloud_cfs_auto_snapshot_policy_attachment.auto_snapshot_policy_attachment", "id")),
			},
			{
				ResourceName:      "cloud_cfs_auto_snapshot_policy_attachment.auto_snapshot_policy_attachment",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccCfsAutoSnapshotPolicyAttachment = `

resource "cloud_cfs_auto_snapshot_policy_attachment" "auto_snapshot_policy_attachment" {
  auto_snapshot_policy_id = "asp-basic"
  file_system_ids         = "cfs-iobiaxtj"
}

`
