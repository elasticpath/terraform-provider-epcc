package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceAccount(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceAccount,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("epcc_account.acc_test_account", "name", regexp.MustCompile("Account for acc tests name")),
					resource.TestMatchResourceAttr("epcc_account.acc_test_account", "legal_name", regexp.MustCompile("Account for acc tests legal name")),
					resource.TestMatchResourceAttr("epcc_account.acc_test_account", "registration_id", regexp.MustCompile("1")),
				),
			},
		},
	})
}

// language=HCL
const testAccResourceAccount = `
resource "epcc_account" "acc_test_account" {
  name            = "Account for acc tests name"
  legal_name      = "Account for acc tests legal name"
  registration_id = "1"
}
`
