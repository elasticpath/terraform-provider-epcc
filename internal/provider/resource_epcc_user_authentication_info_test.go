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
					resource.TestMatchResourceAttr("epcc_user_authentication_info.acc_test_user_authentication_info", "name", regexp.MustCompile("test_user_authentication_info")),
					resource.TestMatchResourceAttr("epcc_user_authentication_info.acc_test_user_authentication_info", "discovery_url", regexp.MustCompile("https://shared-keycloak.env.am.pd.elasticpath.cloud/auth/realms/epcc-integrations-env/.well-known/openid-configuration")),
					resource.TestMatchResourceAttr("epcc_user_authentication_info.acc_test_user_authentication_info", "client_id", regexp.MustCompile("epcc-integrations")),
					resource.TestMatchResourceAttr("epcc_user_authentication_info.acc_test_user_authentication_info", "client_secret", regexp.MustCompile("")),
					resource.TestMatchResourceAttr("epcc_realm.acc_test_realm_for_user_authentication_info", "name", regexp.MustCompile("test_realm")),
				),
			},
		},
	})
}

// language=HCL
const testAccResourceUserAuthenticationInfo = `
resource "epcc_realm" "acc_test_realm_for_user_authentication_info" {
  name = "test_realm"
  redirect_uris = [
    "https://google.com/"
  ]
  duplicate_email_policy = "allowed"
  origin_id = "hello-world"
  origin_type = "customer-authentication-settings"
}

resource "epcc_user_authentication_info" "acc_test_user_authentication_info" {
  name = "John Doe"
  email = "john.doe@banks.com"
  realm_id = epcc_realm.acc_test_realm_for_user_authentication_info.id
}
`
