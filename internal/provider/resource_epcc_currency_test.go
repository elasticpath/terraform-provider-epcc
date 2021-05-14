package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceCurrency(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceCurrency,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("epcc_currency.acc_test_currency", "type", regexp.MustCompile("currency")),
					resource.TestMatchResourceAttr("epcc_currency.acc_test_currency", "code", regexp.MustCompile("CHF")),
					resource.TestMatchResourceAttr("epcc_currency.acc_test_currency", "exchange_rate", regexp.MustCompile("1")),
					resource.TestMatchResourceAttr("epcc_currency.acc_test_currency", "format", regexp.MustCompile("£{price}")),
					resource.TestMatchResourceAttr("epcc_currency.acc_test_currency", "decimal_point", regexp.MustCompile(".")),
					resource.TestMatchResourceAttr("epcc_currency.acc_test_currency", "thousand_separator", regexp.MustCompile(",")),
					resource.TestMatchResourceAttr("epcc_currency.acc_test_currency", "decimal_places", regexp.MustCompile("0")),
					resource.TestMatchResourceAttr("epcc_currency.acc_test_currency", "default", regexp.MustCompile("false")),
					resource.TestMatchResourceAttr("epcc_currency.acc_test_currency", "enabled", regexp.MustCompile("true")),
				),
			},
		},
	})
}

// language=HCL
const testAccResourceCurrency = `
resource "epcc_currency" "acc_test_currency" {
  type = "currency"
  code = "CHF"
  exchange_rate = 1
  format = "£{price}"
  decimal_point = "."
  thousand_separator = ","
  decimal_places = 0
  default = false
  enabled = true
}
`
