package test

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
					resource.TestMatchResourceAttr("epcc_realm.acc_test_realm", "name", regexp.MustCompile("test_realm")),
					resource.TestMatchResourceAttr("epcc_realm.acc_test_realm", "redirect_uris.#", regexp.MustCompile("1")),
					resource.TestMatchResourceAttr("epcc_realm.acc_test_realm", "redirect_uris.0", regexp.MustCompile("https://google.com/")),
					resource.TestMatchResourceAttr("epcc_realm.acc_test_realm", "duplicate_email_policy", regexp.MustCompile("allowed")),
					resource.TestMatchResourceAttr("epcc_realm.acc_test_realm", "origin_id", regexp.MustCompile("hello-world")),
					resource.TestMatchResourceAttr("epcc_realm.acc_test_realm", "origin_type", regexp.MustCompile("customer-authentication-settings")),
				),
			},
		},
	})
}

// language=HCL
const testAccResourceRealm = `
resource "epcc_realm" "acc_test_realm" {
  name = "test_realm"
  redirect_uris = [
    "https://google.com/"]
  duplicate_email_policy = "allowed"
  origin_id = "hello-world"
  origin_type = "customer-authentication-settings"
}
`
