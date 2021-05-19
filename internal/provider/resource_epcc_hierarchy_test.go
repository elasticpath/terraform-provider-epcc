package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceHierarchy(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: hierarchyTestSteps[0],
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("epcc_hierarchy.my_test_hierarchy", "name", regexp.MustCompile("Test Hierarchy")),
					resource.TestMatchResourceAttr("epcc_hierarchy.my_test_hierarchy", "description", regexp.MustCompile("My Description")),
					resource.TestMatchResourceAttr("epcc_hierarchy.my_test_hierarchy", "slug", regexp.MustCompile("test-slug")),
				),
			},
			{
				Config: hierarchyTestSteps[1],
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("epcc_hierarchy.my_test_hierarchy", "name", regexp.MustCompile("Test Updated Hierarchy")),
					resource.TestMatchResourceAttr("epcc_hierarchy.my_test_hierarchy", "description", regexp.MustCompile("My Updated Description")),
					resource.TestMatchResourceAttr("epcc_hierarchy.my_test_hierarchy", "slug", regexp.MustCompile("updated-test-slug")),
				),
			},
		},
	})
}

var hierarchyTestSteps = [...]string{
	// language=HCL
	`
resource "epcc_hierarchy" "my_test_hierarchy" {
  name = "Test Hierarchy"
  description = "My Description"
  slug = "test-slug"
}`,
	// language=HCL
	`
resource "epcc_hierarchy" "my_test_hierarchy" {
  name = "Test Updated Hierarchy"
  description = "My Updated Description"
  slug = "updated-test-slug"
}`,
}
