package test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceFile(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: fileTestSteps[0],
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("epcc_file.my_image_file", "file_name", "ep.png"),
				),
			},
			{
				Config: fileTestSteps[1],
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("epcc_file.my_text_file", "file_name", "hello_world.txt"),
				),
			},
			{
				Config: fileTestSteps[1],
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("epcc_file.my_binary_file", "file_name", "binary_data.bin"),
				),
			},
		},
	})
}

var fileTestSteps = [...]string{
	// language=HCL
	`resource "epcc_file" "my_image_file" {
		file_name = "ep.png"
		file_hash = filemd5("ep.png")
		public = true
		}`, // language=HCL
	`resource "epcc_file" "my_text_file" {
		file_name = "hello_world.txt"
		file_hash = filemd5("hello_world.txt")
		public = false
		}`,
	// language=HCL
	`resource "epcc_file" "my_binary_file" {
		file_name = "binary_data.bin"
		file_hash = filemd5("binary_data.bin")
		public = true
		}`,
}
