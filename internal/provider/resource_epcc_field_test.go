package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceField_StringFormat(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				// language=HCL
				Config: `
					resource "epcc_flow" "sports_flow" {
					  name        = "Flow for sports"
					  slug        = "hockey"
					  description = "This is a Terraform test"
					  enabled     = true
					}
					
					resource "epcc_field" "season" {
					  name = "Sport season"
					  slug = "season"
					  field_type = "string"
					  description = "Season the sport is played in"
					  required = false
					  enabled = true
					  default = "summer"
  					  omit_null = false
					  flow_id = epcc_flow.sports_flow.id
					  valid_string_format = "uuid"
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("epcc_field.season", "name", "Sport season"),
					resource.TestCheckResourceAttr("epcc_field.season", "slug", "season"),
					resource.TestCheckResourceAttr("epcc_field.season", "field_type", "string"),
					resource.TestCheckResourceAttr("epcc_field.season", "description", "Season the sport is played in"),
					resource.TestCheckResourceAttr("epcc_field.season", "required", "false"),
					resource.TestCheckResourceAttr("epcc_field.season", "enabled", "true"),
					resource.TestCheckResourceAttr("epcc_field.season", "default", "summer"),
					resource.TestCheckResourceAttr("epcc_field.season", "omit_null", "false"),
					resource.TestMatchResourceAttr("epcc_field.season", "flow_id",
						regexp.MustCompile("^[{]?[0-9a-fA-F]{8}-([0-9a-fA-F]{4}-){3}[0-9a-fA-F]{12}[}]?$")),
					resource.TestCheckResourceAttr("epcc_field.season", "valid_string_format", "uuid"),
				),
			},
			{
				ResourceName:      "epcc_field.season",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceField_StringEnum(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				// language=HCL
				Config: `
					resource "epcc_flow" "string_flow" {
					  name        = "String enum flow"
					  slug        = "string_flow"
					  description = "This is a Terraform test"
					  enabled     = true
					}
					
					resource "epcc_field" "string_field" {
					  name = "String Field"
					  slug = "string_field"
					  field_type = "string"
					  description = "This is a Terraform test"
					  required = false
					  enabled = true
					  flow_id = epcc_flow.string_flow.id
  					  valid_string_format = "slug"
					  valid_string_enum = ["spring", "summer", "winter", "autumn"]
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("epcc_field.string_field", "valid_string_format", "slug"),
					resource.TestCheckResourceAttr("epcc_field.string_field", "valid_string_enum.0", "spring"),
					resource.TestCheckResourceAttr("epcc_field.string_field", "valid_string_enum.1", "summer"),
					resource.TestCheckResourceAttr("epcc_field.string_field", "valid_string_enum.2", "winter"),
					resource.TestCheckResourceAttr("epcc_field.string_field", "valid_string_enum.3", "autumn"),
				),
			},
			{
				ResourceName:      "epcc_field.string_field",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceField_Relationship(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				// language=HCL
				Config: `
					resource "epcc_flow" "shipping_provider" {
					  name        = "Shipping Provider"
					  slug        = "shipping_provider"
					  description = "This is a Terraform test"
					  enabled     = true
					}

					resource "epcc_flow" "shipping_cost" {
					  name        = "Shipping Cost"
					  slug        = "shipping_cost"
					  description = "This is a Terraform test"
					  enabled     = true
					}
					
					resource "epcc_field" "cost" {
					  name = "Associated Cost"
					  slug = "cost"
					  field_type = "relationship"
					  description = "This is a Terraform test"
					  required = false
					  enabled  = true
					  flow_id = epcc_flow.shipping_provider.id
					  relationship_to_one = "shipping_cost"
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("epcc_field.cost", "relationship_to_one", "shipping_cost"),
				),
			},
			{
				ResourceName:      "epcc_field.cost",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceField_FloatRange(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				// language=HCL
				Config: `
					resource "epcc_flow" "float_test" {
					  name        = "Float Flow"
					  slug        = "float_flow"
					  description = "This is a Terraform test"
					  enabled     = true
					}

					resource "epcc_field" "float_field" {
					  name = "Float Field"
					  slug = "float_field"
					  field_type = "float"
					  description = "This is a Terraform test"
					  required = false
					  enabled  = true
					  flow_id = epcc_flow.float_test.id
					  valid_float_enum = [1.5, 2.5, 3.33]
					  valid_float_range = [
					  	{
							from = 10.7
							to = 16.24
					  	},
					  	{
							from = 74.9
							to = 75.05
					  	}
					  ]
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("epcc_field.float_field", "valid_float_enum.0", "1.5"),
					resource.TestCheckResourceAttr("epcc_field.float_field", "valid_float_enum.1", "2.5"),
					resource.TestCheckResourceAttr("epcc_field.float_field", "valid_float_enum.2", "3.33"),
					resource.TestCheckResourceAttr("epcc_field.float_field", "valid_float_range.0.from", "10.7"),
					resource.TestCheckResourceAttr("epcc_field.float_field", "valid_float_range.0.to", "16.24"),
					resource.TestCheckResourceAttr("epcc_field.float_field", "valid_float_range.1.from", "74.9"),
					resource.TestCheckResourceAttr("epcc_field.float_field", "valid_float_range.1.to", "75.05"),
				),
			},
			{
				ResourceName:      "epcc_field.float_field",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceField_IntRange(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				// language=HCL
				Config: `
					resource "epcc_flow" "float_test" {
					  name        = "Int Flow"
					  slug        = "int_flow"
					  description = "This is a Terraform test"
					  enabled     = true
					}

					resource "epcc_field" "int_field" {
					  name = "Int Field"
					  slug = "int_field"
					  field_type = "integer"
					  description = "This is a Terraform test"
					  required = false
					  enabled  = true
					  flow_id = epcc_flow.float_test.id
					  valid_int_enum = [1, 2, 3]
					  valid_int_range = [
					  	{
							from = 10
							to = 16
					  	},
					  	{
							from = 74
							to = 75
					  	}
					  ]
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("epcc_field.int_field", "valid_int_enum.0", "1"),
					resource.TestCheckResourceAttr("epcc_field.int_field", "valid_int_enum.1", "2"),
					resource.TestCheckResourceAttr("epcc_field.int_field", "valid_int_enum.1", "3"),
					resource.TestCheckResourceAttr("epcc_field.int_field", "valid_int_range.0.from", "10"),
					resource.TestCheckResourceAttr("epcc_field.int_field", "valid_int_range.0.to", "16"),
					resource.TestCheckResourceAttr("epcc_field.int_field", "valid_int_range.1.from", "74"),
					resource.TestCheckResourceAttr("epcc_field.int_field", "valid_int_range.2.to", "75"),
				),
			},
			{
				ResourceName:      "epcc_field.int_field",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
