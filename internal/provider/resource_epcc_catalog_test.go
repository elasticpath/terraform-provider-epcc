package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceCatalog(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceCatalog[0],
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("epcc_catalog.my_first_catalog", "name", "My Catalog Name"),
					resource.TestCheckResourceAttr("epcc_catalog.my_first_catalog", "description", "My Catalog Description"),
				),
			},
			{
				Config: testAccResourceCatalog[1],
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("epcc_catalog.my_first_catalog", "name", "My Updated Catalog Name"),
					resource.TestCheckResourceAttr("epcc_catalog.my_first_catalog", "description", "My Updated Catalog Description"),
				),
			},
		},
	})
}


var testAccResourceCatalog = [...]string{
	// language=HCL
fmt.Sprintf(`
resource "epcc_product" "my_first_terraform_physical_product" {
  sku            = "%[1]d"
  name           = "TFProduct1Physical"
  commodity_type = "physical"
  description    = "Terraform Physical Product 1"
  mpn            = "mfg-part-1"
  status         = "live"
}

resource "epcc_hierarchy" "my_first_hierarchy" {
  name        = "Hierarchy"
  description = "Foo"
  slug        = "test%[1]d"
}

resource "epcc_node" "my_first_node" {
  name         = "Node #2"
  description  = "Node"
  slug         = "node-3%[1]d"
  hierarchy_id = epcc_hierarchy.my_first_hierarchy.id
  products = [ epcc_product.my_first_terraform_physical_product.id ]
}

resource "epcc_pricebook" "my_first_terraform_pricebook" {
  name        = "TFPricebook1%[1]d"
  description = "Terraform 1"
}

resource "epcc_catalog" "my_first_catalog" {
  name = "My Catalog Name"
  description = "My Catalog Description"
  hierarchies = [ epcc_hierarchy.my_first_hierarchy.id]
  pricebook = epcc_pricebook.my_first_terraform_pricebook.id
}
`, timestamp+10),
	// language=HCL
	fmt.Sprintf(`
resource "epcc_product" "my_first_terraform_physical_product" {
  sku            = "%[1]d"
  name           = "TFProduct1Physical"
  commodity_type = "physical"
  description    = "Terraform Physical Product 1"
  mpn            = "mfg-part-1"
  status         = "live"
}

resource "epcc_hierarchy" "my_first_hierarchy" {
  name        = "Hierarchy"
  description = "Foo"
  slug        = "test%[1]d"
}

resource "epcc_node" "my_first_node" {
  name         = "Node #2"
  description  = "Node"
  slug         = "node-3%[1]d"
  hierarchy_id = epcc_hierarchy.my_first_hierarchy.id
  products = [ epcc_product.my_first_terraform_physical_product.id ]
}

resource "epcc_pricebook" "my_first_terraform_pricebook" {
  name        = "TFPricebook1%[1]d"
  description = "Terraform 1"
}

resource "epcc_catalog" "my_first_catalog" {
  name = "My Updated Catalog Name"
  description = "My Updated Catalog Description"
  hierarchies = [ epcc_hierarchy.my_first_hierarchy.id]
  pricebook = epcc_pricebook.my_first_terraform_pricebook.id
}
`, timestamp+10),
}
