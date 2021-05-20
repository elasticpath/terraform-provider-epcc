package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceCatalogRule(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceCatalogRule,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("epcc_catalog_rule.sneaker_catalog_rule", "name", regexp.MustCompile("Special running shoes")),
					resource.TestMatchResourceAttr("epcc_catalog_rule.sneaker_catalog_rule", "catalog",
						regexp.MustCompile("^[{]?[0-9a-fA-F]{8}-([0-9a-fA-F]{4}-){3}[0-9a-fA-F]{12}[}]?$")),
				),
			},
		},
	})
}

const testAccResourceCatalogRule =
// language=HCL
`
resource "epcc_hierarchy" "sneakers_hierarchy" {
  name        = "Hierarchy"
}

resource "epcc_pricebook" "sneakers_prices" {
  name        = "Sneakers Prices"
}

resource "epcc_catalog" "sneakers_catalog" {
  name        = "Sneakers Catalog"
  hierarchies = [epcc_hierarchy.sneakers_hierarchy.id]
  pricebook   = epcc_pricebook.sneakers_prices.id
}

resource "epcc_customer" "customer_1" {
  name  = "Banana"
  email = "banana@food.com"
}

resource "epcc_catalog_rule" "sneaker_catalog_rule" {
  name = "Special running shoes"
  catalog = epcc_catalog.sneakers_catalog.id
  customers = [epcc_customer.customer_1.id]
}
`
