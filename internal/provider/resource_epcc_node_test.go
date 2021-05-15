package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceNode(t *testing.T) {
	//TODO Determine how to validate that the hierarchy moved
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			/*{
				Config: nodeTestSteps[0],
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("epcc_node.my_test_node", "name", regexp.MustCompile("First Node")),
					resource.TestMatchResourceAttr("epcc_node.my_test_node", "description", regexp.MustCompile("Node Description")),
					resource.TestMatchResourceAttr("epcc_node.my_test_node", "slug", regexp.MustCompile("first-node-slug")),
				),
			},
			{
				Config: nodeTestSteps[1],
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("epcc_node.my_test_node", "name", regexp.MustCompile("First Updated Node")),
					resource.TestMatchResourceAttr("epcc_node.my_test_node", "description", regexp.MustCompile("Node Updated Description")),
					resource.TestMatchResourceAttr("epcc_node.my_test_node", "slug", regexp.MustCompile("first-updated-node-slug")),
				),
			},
			{
				Config: nodeTestSteps[2],
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("epcc_node.my_test_node", "name", regexp.MustCompile("First Updated Node")),
					resource.TestMatchResourceAttr("epcc_node.my_test_node", "description", regexp.MustCompile("Node Updated Description")),
					resource.TestMatchResourceAttr("epcc_node.my_test_node", "slug", regexp.MustCompile("first-updated-node-slug")),
				),
			},*/
			{
				Config: nodeTestSteps[3],
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
	})
}

func getId(resourceName string, idValue **string) resource.TestCheckFunc {
	*idValue = new(string)
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("Could not determine ID")
		}

		*idValue = &rs.Primary.ID

		return nil
	}
}

var nodeTestSteps = [...]string{
	// language=HCL
	`
resource "epcc_hierarchy" "my_test_hierarchy" {
  	name = "Test Hierarchy"
  	description = "Hierarchy Description"
  	slug = "hierarchy-slug"
}

resource "epcc_node" "my_test_node" { 
	name = "First Node"
	description = "Node Description"
	slug = "first-node-slug"
	hierarchy_id = epcc_hierarchy.my_test_hierarchy.id
}
`,
	// language=HCL
	`
resource "epcc_hierarchy" "my_test_hierarchy" {
  	name = "Test Hierarchy"
  	description = "Hierarchy Description"
  	slug = "hierarchy-slug"
}

resource "epcc_node" "my_test_node" { 
	name = "First Updated Node"
	description = "Node Updated Description"
	slug = "first-updated-node-slug"
	hierarchy_id = epcc_hierarchy.my_test_hierarchy.id
}
`,
	// language=HCL
	`
resource "epcc_hierarchy" "my_test_hierarchy" {
  	name = "Test Hierarchy"
  	description = "Hierarchy Description"
  	slug = "hierarchy-slug"
}

resource "epcc_hierarchy" "my_second_hierarchy" {
  	name = "Test Second Hierarchy"
  	description = "Second Hierarchy Description"
  	slug = "second-hierarchy-slug"
}

resource "epcc_node" "my_test_node" { 
	name = "First Updated Node"
	description = "Node Updated Description"
	slug = "first-updated-node-slug"
	hierarchy_id = epcc_hierarchy.my_second_hierarchy.id
}
`,
	fmt.Sprintf(
		// language=HCL
		`resource "epcc_hierarchy" "my_test_hierarchy" {
		name = "Test Hierarchy"
		description = "Hierarchy Description"
		slug = "hierarchy-slug"
	}

	resource "epcc_product" "acc_test_epcc_product_1_physical_draft" {
		sku            = "test-%d"
		name           = "TestAccResourceProduct1"
		commodity_type = "physical"
		description    = "Draft description"
		mpn            = "mfg-acc_test_epcc_product_1"
		status         = "draft"
}

	resource "epcc_node" "my_test_node" { 
		name = "First Updated Node"
		description = "Node Updated Description"
		slug = "first-updated-node-slug"
		hierarchy_id = epcc_hierarchy.my_test_hierarchy.id
		products = [
		]
	}`, timestamp),
}
