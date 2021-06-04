package provider

import (
	"io/ioutil"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceFile(t *testing.T) {
	tempDir := t.TempDir()

	resource.UnitTest(t, resource.TestCase{

		PreCheck: func() {
			testAccPreCheck(t)
			err := ioutil.WriteFile(tempDir+"/hello_world.txt", []byte("hello world"), 0644)
			if err != nil {
				t.Fatal(err)
			}
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				PreventDiskCleanup: true,
				Config: `resource "epcc_file" "my_text_file"{
					file_name = "` + tempDir + `/hello_world.txt"
					file_hash = filemd5("` + tempDir + `/hello_world.txt")
					public = true
				}
				resource "epcc_file" "my_image_link" {
					file_location = "https://my.example.com/images/abc.png"
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("epcc_file.my_text_file", "file_name", tempDir+"/hello_world.txt"),
					resource.TestCheckResourceAttr("epcc_file.my_image_link", "file_location", "https://my.example.com/images/abc.png"),
				),
			},
		},
	})
}
