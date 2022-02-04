package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceFile(t *testing.T) {

	resource.UnitTest(t, resource.TestCase{

		PreCheck: func() {
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				PreventDiskCleanup: true,
				Config: `resource "epcc_file" "my_logo"{
					file_name = "ep.png"
					file_hash = filemd5("ep.png")
					public = true
				}
				resource "epcc_file" "my_image_link" {
					file_location = "https://my.example.com/images/abc.png"
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("epcc_file.my_logo", "file_name", "ep.png"),
					resource.TestCheckResourceAttr("epcc_file.my_image_link", "file_location", "https://my.example.com/images/abc.png"),
				),
			},
		},
	})
}
