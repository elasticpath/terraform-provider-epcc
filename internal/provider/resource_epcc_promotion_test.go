package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourcePromotion(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourcePromotion,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("epcc_promotion.acc_test_promotion", "name", "Promo #1"),
					resource.TestCheckResourceAttr("epcc_promotion.acc_test_promotion", "description", "Initial Promotion"),
					resource.TestCheckResourceAttr("epcc_promotion.acc_test_promotion", "enabled", "true"),
					resource.TestCheckResourceAttr("epcc_promotion.acc_test_promotion", "promotion_type", "fixed_discount"),
					resource.TestCheckResourceAttr("epcc_promotion.acc_test_promotion", "schema.#", "1"),
					resource.TestCheckResourceAttr("epcc_promotion.acc_test_promotion", "schema.0.currencies.#", "1"),
					resource.TestCheckResourceAttr("epcc_promotion.acc_test_promotion", "schema.0.currencies.0.currency", "CHF"),
					resource.TestCheckResourceAttr("epcc_promotion.acc_test_promotion", "schema.0.currencies.0.amount", "900"),
					resource.TestCheckResourceAttr("epcc_promotion.acc_test_promotion", "max_discount_value.#", "1"),
					resource.TestCheckResourceAttr("epcc_promotion.acc_test_promotion", "max_discount_value.0.currency", "CHF"),
					resource.TestCheckResourceAttr("epcc_promotion.acc_test_promotion", "max_discount_value.0.amount", "960"),
					resource.TestCheckResourceAttr("epcc_promotion.acc_test_promotion", "min_cart_value.#", "1"),
					resource.TestCheckResourceAttr("epcc_promotion.acc_test_promotion", "min_cart_value.0.currency", "CHF"),
					resource.TestCheckResourceAttr("epcc_promotion.acc_test_promotion", "min_cart_value.0.amount", "100"),
					resource.TestCheckResourceAttr("epcc_promotion.acc_test_promotion", "start", "2019-05-12T00:00:00Z"),
					resource.TestCheckResourceAttr("epcc_promotion.acc_test_promotion", "end", "2019-10-12T00:00:00Z"),
				),
			},
		},
	})
}

// language=HCL
const testAccResourcePromotion = `

resource "epcc_currency" "chf" {
  code = "CHF"
  exchange_rate = 1
  format = "£{price}"
  decimal_point = "."
  thousand_separator = ","
  decimal_places = 0
  default = false
  enabled = true
}

resource "epcc_promotion" "acc_test_promotion" {
  name = "Promo #1"
  description = "Initial Promotion"
  enabled = true
  promotion_type = "fixed_discount"
  schema {
    currencies {
      currency = epcc_currency.chf.code
      amount = 900
    }
  }

  max_discount_value {
    currency = epcc_currency.chf.code
    amount = 960
  }
  min_cart_value {
    currency = epcc_currency.chf.code
    amount = 100
  }
  start = "2019-05-12T00:00:00Z"
  end = "2019-10-12T00:00:00Z"
}

`
