package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourcePricebook(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePricebook,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("epcc_pricebook.acc_test_epcc_pricebook", "name", regexp.MustCompile("Pricebook for acc tests name")),
					resource.TestMatchResourceAttr("epcc_pricebook.acc_test_epcc_pricebook", "description", regexp.MustCompile("Pricebook for acc tests description")),
				),
			},
		},
	})
}

const testAccResourcePricebook =
// language=HCL
`
resource "epcc_pricebook" "acc_test_epcc_pricebook" {
  name            = "Pricebook for acc tests name"
  description      = "Pricebook for acc tests description"
}
`
