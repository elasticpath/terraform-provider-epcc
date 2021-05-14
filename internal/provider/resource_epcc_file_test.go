package provider

import (
	"regexp"
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
					resource.TestMatchResourceAttr("epcc_file.my_image_link", "file_location", regexp.MustCompile("https://my.example.com/images/abc.png")),
				),
			},
			{
				Config: fileTestSteps[1],
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("epcc_file.my_image_link", "file_location", regexp.MustCompile("https://other.example.com/image/abc.png")),
				),
			},
			{
				Config: fileTestSteps[2],
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: fileTestSteps[3],
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
	})
}

var fileTestSteps = [...]string{
	// language=HCL
	`resource "epcc_file" "my_image_link" {
  file_location = "https://my.example.com/images/abc.png"
}`, // language=HCL
	`resource "epcc_file" "my_image_link" {
  file_location = "https://other.example.com/image/abc.png"
}`, // language=HCL
	`
resource "epcc_file" "my_image_file" {
  file = base64encode("hello")
  file_name = "hello.txt"
  public = true
}`,
	// language=HCL
	`resource "epcc_file" "my_image_file" {
  file = base64encode("hello2")
  file_name = "hello2.txt"
}`,
}
