package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceOidcProfile(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceProfile,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("epcc_oidc_profile.acc_test_profile", "name", regexp.MustCompile("test_profile")),
					resource.TestMatchResourceAttr("epcc_oidc_profile.acc_test_profile", "discovery_url", regexp.MustCompile("https://elasticpath-customer.okta.com/.well-known/openid-configuration")),
					resource.TestMatchResourceAttr("epcc_oidc_profile.acc_test_profile", "client_id", regexp.MustCompile("epcc-integrations")),
					resource.TestMatchResourceAttr("epcc_oidc_profile.acc_test_profile", "client_secret", regexp.MustCompile("")),
					resource.TestMatchResourceAttr("epcc_authentication_realm.acc_test_realm_for_profile", "name", regexp.MustCompile("test_realm")),
				),
			},
		},
	})
}

// language=HCL
const testAccResourceProfile = `
data "epcc_customer_authentication_settings" "customer_auth_settings" {
}

resource "epcc_authentication_realm" "acc_test_realm_for_profile" {
  name = "test_realm"
  authentication_realm_id = data.epcc_customer_authentication_settings.customer_auth_settings.realm_id
  redirect_uris = [
    "https://google.com/"
  ]
  duplicate_email_policy = "allowed"
}

resource "epcc_oidc_profile" "acc_test_profile" {
  name = "test_profile"
  discovery_url = "https://elasticpath-customer.okta.com/.well-known/openid-configuration"
  client_id = "epcc-integrations"
  client_secret = "86c8986d-e1b2-4ce4-a24c-8430ec1ab383"
  realm_id = epcc_authentication_realm.acc_test_realm_for_profile.id
}
`
