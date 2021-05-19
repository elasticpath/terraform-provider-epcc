package test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceProductPrice(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceProductPrice,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
	})
}

var testAccResourceProductPrice =
// language=HCL
fmt.Sprintf(`
resource "epcc_product" "my_first_product" {
  sku            = "tf-product-1-3%[1]d"
  name           = "TFProduct1Physical%[1]d"
  commodity_type = "physical"
  description    = "Terraform Physical Product 1"
  mpn            = "mfg-part-1"
  status         = "live"
}


resource "epcc_pricebook" "my_first_terraform_pricebook" {
  name        = "TFPricebook1%[1]d"
  description = "Terraform 1"
}

resource "epcc_currency" "CAD" {
  code = "CYZ"
  exchange_rate = 1
  format = "$${price}"
  decimal_point = "."
  thousand_separator = ","
  decimal_places = 0
  default = false
  enabled = true
}

resource "epcc_currency" "NZD" {
  code = "NYZ"
  exchange_rate = 1
  format = "$${price}"
  decimal_point = "."
  thousand_separator = ","
  decimal_places = 0
  default = false
  enabled = true
}

resource "epcc_product_price" "price" {
  sku = epcc_product.my_first_product.sku
  pricebook_id = epcc_pricebook.my_first_terraform_pricebook.id

  currency {
    code = epcc_currency.CAD.code
    amount = 420
    includes_tax = true
  }

  currency {
    code = epcc_currency.NZD.code
    amount = 424
    includes_tax = false
  }
}

data "epcc_product_price" "price" {
  id = epcc_product_price.price.id
  pricebook_id = epcc_pricebook.my_first_terraform_pricebook.id
}
`, timestamp+151)
