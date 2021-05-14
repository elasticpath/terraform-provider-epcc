package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceField(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceField,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("epcc_field.customer_hobby_field", "name", regexp.MustCompile("Customer Hobby")),
					resource.TestMatchResourceAttr("epcc_field.customer_hobby_field", "slug", regexp.MustCompile("hobby")),
					resource.TestMatchResourceAttr("epcc_field.customer_hobby_field", "field_type", regexp.MustCompile("string")),
					resource.TestMatchResourceAttr("epcc_field.customer_hobby_field", "description", regexp.MustCompile("Activity the customer is fond of")),
					resource.TestMatchResourceAttr("epcc_field.customer_hobby_field", "required", regexp.MustCompile("false")),
					resource.TestMatchResourceAttr("epcc_field.customer_hobby_field", "default", regexp.MustCompile("biking")),
					resource.TestMatchResourceAttr("epcc_field.customer_hobby_field", "omit_null", regexp.MustCompile("false")),
					resource.TestMatchResourceAttr("epcc_field.customer_hobby_field", "enabled", regexp.MustCompile("true")),
					resource.TestMatchResourceAttr("epcc_field.customer_hobby_field", "flow_id", regexp.MustCompile("cf47328b-e80e-42e4-9428-11cae225ce3d")),
				),
			},
		},
	})
}

const testAccResourceField =
// language=HCL
`
resource "epcc_field" "customer_hobby_field" {
  name = "Customer Hobby"
  slug = "hobby"
  field_type = "string"
  description = "Activity the customer is fond of"
  required = false
  default = "biking"
  omit_null = false
  enabled = true
  flow_id = "cf47328b-e80e-42e4-9428-11cae225ce3d"
}
`


