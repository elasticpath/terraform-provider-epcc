package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceCustomer(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceCustomer,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("epcc_customer.acc_test_customer", "name", regexp.MustCompile("Customer for acc tests")),
					resource.TestMatchResourceAttr("epcc_customer.acc_test_customer", "email", regexp.MustCompile("customer@acc.tests")),
				),
			},
		},
	})
}

// language=HCL
const testAccResourceCustomer = `
resource "epcc_customer" "acc_test_customer" {
  name  = "Customer for acc tests"
  email = "customer@acc.tests"
}
`
