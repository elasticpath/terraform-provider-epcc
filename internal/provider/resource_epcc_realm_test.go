package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceRealm(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceRealm,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("epcc_realm.acc_test_realm", "code", regexp.MustCompile("CHF")),
					resource.TestMatchResourceAttr("epcc_realm.acc_test_realm", "exchange_rate", regexp.MustCompile("1")),
					resource.TestMatchResourceAttr("epcc_realm.acc_test_realm", "format", regexp.MustCompile("£{price}")),
					resource.TestMatchResourceAttr("epcc_realm.acc_test_realm", "decimal_point", regexp.MustCompile(".")),
					resource.TestMatchResourceAttr("epcc_realm.acc_test_realm", "thousand_separator", regexp.MustCompile(",")),
					resource.TestMatchResourceAttr("epcc_realm.acc_test_realm", "decimal_places", regexp.MustCompile("0")),
					resource.TestMatchResourceAttr("epcc_realm.acc_test_realm", "default", regexp.MustCompile("false")),
					resource.TestMatchResourceAttr("epcc_realm.acc_test_realm", "enabled", regexp.MustCompile("true")),
				),
			},
		},
	})
}

// language=HCL
const testAccResourceRealm = `
resource "epcc_realm" "test_realm" {
  name = "test_realm"
  redirect_uris = [
    "https://google.com/"]
  duplicate_email_policy = "ALLOWED"
  origin_id = "hello-world"
  origin_type = "customer-authentication-settings"
}
`
