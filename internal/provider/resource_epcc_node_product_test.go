package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceNodeProduct(t *testing.T) {
	//TODO Determine how to validate that the hierarchy moved
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: nodeProductTestSteps[0],
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: nodeProductTestSteps[1],
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
	})
}

var nodeProductTestSteps = [...]string{
	// language=HCL
	fmt.Sprintf(`
resource "epcc_hierarchy" "my_first_hierarchy" {
  name        = "Node Hierarchy"
  description = "Foo"
  slug        = "test-node-hierarchy-%[1]d"
}

resource "epcc_node" "my_parent_node" {
  name         = "Node #2"
  description  = "Node"
  slug         = "node-%[1]d"
  hierarchy_id = epcc_hierarchy.my_first_hierarchy.id
}

resource "epcc_product" "simple_product" {
  sku            = "457-%[1]d"
  name           = "Product With File"
  commodity_type = "physical"
}

resource "epcc_product" "simple_product_2" {
  sku            = "458-%[1]d"
  name           = "Product With File"
  commodity_type = "physical"
}

resource "epcc_node_product" "simple_product_node_relationship" {
  node_id = epcc_node.my_parent_node.id
  hierarchy_id = epcc_hierarchy.my_first_hierarchy.id
  product_id = epcc_product.simple_product.id
}
`, timestamp+1024),
	// language=HCL
	fmt.Sprintf(`
resource "epcc_hierarchy" "my_first_hierarchy" {
  name        = "Node Hierarchy"
  description = "Foo"
  slug        = "test-node-hierarchy-%[1]d"
}

resource "epcc_node" "my_parent_node" {
  name         = "Node #2"
  description  = "Node"
  slug         = "node-%[1]d"
  hierarchy_id = epcc_hierarchy.my_first_hierarchy.id
}

resource "epcc_product" "simple_product" {
  sku            = "457-%[1]d"
  name           = "Product With File"
  commodity_type = "physical"
}

resource "epcc_product" "simple_product_2" {
  sku            = "458-%[1]d"
  name           = "Product With File"
  commodity_type = "physical"
}

resource "epcc_node_product" "simple_product_node_relationship" {
  node_id = epcc_node.my_parent_node.id
  hierarchy_id = epcc_hierarchy.my_first_hierarchy.id
  product_id = epcc_product.simple_product_2.id
}

data "epcc_node_product" "data_simple_product_node_relationship" {
  id = epcc_node_product.simple_product_node_relationship.id
  node_id = epcc_node.my_parent_node.id
  hierarchy_id = epcc_hierarchy.my_first_hierarchy.id
  product_id = epcc_product.simple_product_2.id
}
`, timestamp+1024),
}
