package test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceFlow(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFlow,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("epcc_flow.acc_test_flow", "name", regexp.MustCompile("flow_test")),
					resource.TestMatchResourceAttr("epcc_flow.acc_test_flow", "slug", regexp.MustCompile("flow_test")),
					resource.TestMatchResourceAttr("epcc_flow.acc_test_flow", "description", regexp.MustCompile("This is a Terraform test")),
					resource.TestMatchResourceAttr("epcc_flow.acc_test_flow", "enabled", regexp.MustCompile("true")),
				),
			},
		},
	})
}

const testAccResourceFlow = `
resource "epcc_flow" "acc_test_flow" {
  name			= "flow_test"
  slug			= "flow_test"
  description	= "This is a Terraform test"
  enabled		= true
}
`
