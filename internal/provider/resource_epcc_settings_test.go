package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceSettings(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceSettings,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("epcc_settings.test_settings", "page_length", regexp.MustCompile("20")),
					resource.TestMatchResourceAttr("epcc_settings.test_settings", "list_child_products", regexp.MustCompile("false")),
					resource.TestMatchResourceAttr("epcc_settings.test_settings", "calculation_method", regexp.MustCompile("simple")),
				),
			},
		},
	})
}

// language=HCL
const testAccResourceSettings = `
resource "epcc_settings" "test_settings" {
  page_length = 20
  list_child_products = false
  additional_languages = [
    "fr",
    "de"]
  calculation_method = "simple"
}

data "epcc_settings" "example" {
}
`
