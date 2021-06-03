package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceEntry(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: // language=HCL
				`
                	resource "epcc_flow" "test_flow" {
                	  name = "Test flow"
                	  slug = "test"
                	  description = "This is a Terraform test"
                	  enabled = true
                	}
                	
                	resource "epcc_field" "test_flow_string_field" {
                	  name = "Test string field"
                	  slug = "string"
                	  field_type = "string"
                	  description = "string field"
                	  required = false
                	  default = "default"
                	  omit_null = false
                	  enabled = true
                	  flow_id = epcc_flow.test_flow.id
                	}
                	
                	resource "epcc_field" "test_flow_integer_field" {
                	  name = "Test integer field"
                	  slug = "integer"
                	  field_type = "integer"
                	  description = "integer field"
                	  required = false
                	  default = 1
                	  omit_null = false
                	  enabled = true
                	  flow_id = epcc_flow.test_flow.id
                	}

                	resource "epcc_field" "test_flow_float_field" {
                	  name = "Test float field"
                	  slug = "float"
                	  field_type = "float"
                	  description = "float field"
                	  required = false
                	  default = 1.0
                	  omit_null = false
                	  enabled = true
                	  flow_id = epcc_flow.test_flow.id
                	}

                	resource "epcc_field" "test_flow_boolean_field" {
                	  name = "Test boolean field"
                	  slug = "boolean"
                	  field_type = "boolean"
                	  description = "boolean field"
                	  required = false
                	  omit_null = false
                	  enabled = true
                	  flow_id = epcc_flow.test_flow.id
                	}
                	
                	resource "epcc_entry" "test_entry" {
                	  slug = epcc_flow.test_flow.slug
                	  strings = {
                	    (epcc_field.test_flow_string_field.slug) = "netherlands"
                	  }
                	  numbers = {
                	    (epcc_field.test_flow_integer_field.slug) = 2,
                	    (epcc_field.test_flow_float_field.slug) = 3.1,
					  }
                	  booleans = {
                	    (epcc_field.test_flow_boolean_field.slug) = true
                	  }
                	}
                `,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("epcc_entry.test_entry", "slug", "test"),
					resource.TestCheckResourceAttr("epcc_entry.test_entry", "strings.string", "netherlands"),
					resource.TestCheckResourceAttr("epcc_entry.test_entry", "numbers.integer", "2"),
					resource.TestCheckResourceAttr("epcc_entry.test_entry", "numbers.float", "3.1"),
					resource.TestCheckResourceAttr("epcc_entry.test_entry", "booleans.boolean", "true"),
				),
			},
		},
	})
}
