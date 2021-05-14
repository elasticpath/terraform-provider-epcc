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
					resource.TestMatchResourceAttr("epcc_field.sports_season_field", "name", regexp.MustCompile("Sport season")),
					resource.TestMatchResourceAttr("epcc_field.sports_season_field", "slug", regexp.MustCompile("season")),
					resource.TestMatchResourceAttr("epcc_field.sports_season_field", "field_type", regexp.MustCompile("string")),
					resource.TestMatchResourceAttr("epcc_field.sports_season_field", "description", regexp.MustCompile("Season the sport is played in")),
					resource.TestMatchResourceAttr("epcc_field.sports_season_field", "required", regexp.MustCompile("false")),
					resource.TestMatchResourceAttr("epcc_field.sports_season_field", "default", regexp.MustCompile("summer")),
					resource.TestMatchResourceAttr("epcc_field.sports_season_field", "omit_null", regexp.MustCompile("false")),
					resource.TestMatchResourceAttr("epcc_field.sports_season_field", "enabled", regexp.MustCompile("true")),
					resource.TestMatchResourceAttr("epcc_field.sports_season_field", "flow_id", regexp.MustCompile("^[{]?[0-9a-fA-F]{8}-([0-9a-fA-F]{4}-){3}[0-9a-fA-F]{12}[}]?$")),
				),
			},
		},
	})
}

const testAccResourceField =
// language=HCL
`
resource "epcc_flow" "sports_flow" {
  name        = "Flow for sports"
  slug        = "hockey"
  description = "This is a Terraform test"
  enabled     = true
}

resource "epcc_field" "sports_season_field" {
  name = "Sport season"
  slug = "season"
  field_type = "string"
  description = "Season the sport is played in"
  required = false
  default = "summer"
  omit_null = false
  enabled = true
  flow_id = epcc_flow.sports_flow.id
}
`
