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
			sourceFile, err := ioutil.ReadFile("../../ep.png")
			if err != nil {
				t.Fatal(err)
			}
			err = ioutil.WriteFile(tempDir+"/ep.png", sourceFile, 0644)
			if err != nil {
				t.Fatal(err)
			}
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				PreventDiskCleanup: true,
				Config: `resource "epcc_file" "my_logo"{
					file_name = "` + tempDir + `/ep.png"
					file_hash = filemd5("` + tempDir + `/ep.png")
					public = true
				}
				resource "epcc_file" "my_image_link" {
					file_location = "https://my.example.com/images/abc.png"
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("epcc_file.my_logo", "file_name", tempDir+`/ep.png`),
					resource.TestCheckResourceAttr("epcc_file.my_image_link", "file_location", "https://my.example.com/images/abc.png"),
				),
			},
		},
	})
}
