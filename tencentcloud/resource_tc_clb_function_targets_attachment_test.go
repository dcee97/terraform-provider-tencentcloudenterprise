package tencentcloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTencentCloudClbFunctionTargetsAttachmentResource_basic(t *testing.T) {
	t.Parallel()
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAccPreCheck(t)
		},
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccClbFunctionTargetsAttachment,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("cloud_clb_function_targets_attachment.function_targets", "id"),
					resource.TestCheckResourceAttr("cloud_clb_function_targets_attachment.function_targets", "function_targets.0.weight", "10"),
				),
			},
			{
				Config: testAccClbFunctionTargetsAttachmentUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("cloud_clb_function_targets_attachment.function_targets", "function_targets.0.weight", "20"),
				),
			},
			{
				ResourceName:      "cloud_clb_function_targets_attachment.function_targets",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

const testAccClbFunctionTargetsAttachment = `

resource "cloud_clb_function_targets_attachment" "function_targets" {
  domain           = "xxx.com"
  listener_id      = "lbl-nonkgvc2"
  load_balancer_id = "lb-5dnrkgry"
  url              = "/"

  function_targets {
    weight = 10

    function {
      function_name           = "keep-1676351130"
      function_namespace      = "default"
      function_qualifier      = "$LATEST"
      function_qualifier_type = "VERSION"
    }
  }
}

`

const testAccClbFunctionTargetsAttachmentUpdate = `

resource "cloud_clb_function_targets_attachment" "function_targets" {
  domain           = "xxx.com"
  listener_id      = "lbl-nonkgvc2"
  load_balancer_id = "lb-5dnrkgry"
  url              = "/"

  function_targets {
    weight = 20

    function {
      function_name           = "keep-1676351130"
      function_namespace      = "default"
      function_qualifier      = "$LATEST"
      function_qualifier_type = "VERSION"
    }
  }
}

`
