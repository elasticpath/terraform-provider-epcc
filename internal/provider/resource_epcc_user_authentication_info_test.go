package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceUserAuthenticationInfo(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceUserAuthenticationInfo,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("epcc_user_authentication_info.acc_test_user_authentication_info", "name", regexp.MustCompile("John Doe")),
					resource.TestMatchResourceAttr("epcc_user_authentication_info.acc_test_user_authentication_info", "email", regexp.MustCompile("john.doe@banks.com")),
					resource.TestMatchResourceAttr("epcc_authentication_realm.acc_test_realm_for_user_authentication_info", "name", regexp.MustCompile("test_realm")),
				),
			},
		},
	})
}

// language=HCL
const testAccResourceUserAuthenticationInfo = `

data "epcc_customer_authentication_settings" "customer_auth_settings" {
}

resource "epcc_authentication_realm" "acc_test_realm_for_user_authentication_info" {
  name = "test_realm"
  authentication_realm_id = data.epcc_customer_authentication_settings.customer_auth_settings.realm_id
  redirect_uris = [
    "https://google.com/"
  ]
  duplicate_email_policy = "allowed"
}

resource "epcc_user_authentication_info" "acc_test_user_authentication_info" {
  name = "John Doe"
  email = "john.doe@banks.com"
  realm_id = epcc_authentication_realm.acc_test_realm_for_user_authentication_info.id
}
`
