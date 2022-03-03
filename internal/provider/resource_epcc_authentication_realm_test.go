package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceAuthenticationRealmWithCustomerAuthenticationSettings(t *testing.T) {
	// language=HCL
	config := ` # Example usage to update the realm associated with Customer Authentication
				data "epcc_customer_authentication_settings" "customer_auth_settings" {
				}
				
				resource "epcc_authentication_realm" "acc_test_realm" {
				  name = "test_realm"
				  authentication_realm_id = data.epcc_customer_authentication_settings.customer_auth_settings.realm_id
				  redirect_uris = [
					"https://google.com/"]
				  duplicate_email_policy = "api_only"
				}`

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: config,

				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("epcc_authentication_realm.acc_test_realm", "name", regexp.MustCompile("test_realm")),
					resource.TestMatchResourceAttr("epcc_authentication_realm.acc_test_realm", "redirect_uris.#", regexp.MustCompile("1")),
					resource.TestMatchResourceAttr("epcc_authentication_realm.acc_test_realm", "redirect_uris.0", regexp.MustCompile("https://google.com/")),
					resource.TestMatchResourceAttr("epcc_authentication_realm.acc_test_realm", "duplicate_email_policy", regexp.MustCompile("api_only")),
					resource.TestMatchResourceAttr("epcc_authentication_realm.acc_test_realm", "origin_type", regexp.MustCompile("customer-authentication-settings")),
				),
			},
			{
				// language=HCL
				Config:  config,
				Destroy: true,
				Check:   resource.ComposeTestCheckFunc(),
			},
			{
				// language=HCL
				Config: `
				# Example usage to update the realm associated with Customer Authentication
				data "epcc_customer_authentication_settings" "customer_auth_settings" {
				}
				
				data "epcc_authentication_realm" "acc_test_realm" {
				  id = data.epcc_customer_authentication_settings.customer_auth_settings.realm_id
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("data.epcc_authentication_realm.acc_test_realm", "name", regexp.MustCompile("Buyer Organization")),
					resource.TestMatchResourceAttr("data.epcc_authentication_realm.acc_test_realm", "redirect_uris.#", regexp.MustCompile("0")),
					resource.TestMatchResourceAttr("data.epcc_authentication_realm.acc_test_realm", "duplicate_email_policy", regexp.MustCompile("allowed")),
					resource.TestMatchResourceAttr("data.epcc_authentication_realm.acc_test_realm", "origin_type", regexp.MustCompile("customer-authentication-settings")),
				),
			},
		},
	})
}

func TestAccResourceAuthenticationRealmWithAccountAuthenticationSettings(t *testing.T) {
	// language=HCL
	config := `# Example usage to update the realm associated with Customer Authentication
				data "epcc_account_authentication_settings" "account_authentication_settings" {
				}
				
				resource "epcc_authentication_realm" "acc_test_realm" {
				  name = "test_realm_2"
				  authentication_realm_id = data.epcc_account_authentication_settings.account_authentication_settings.realm_id
				  redirect_uris = [
					"https://yahoo.co.uk/"
					]
				  duplicate_email_policy = "api_only"
				}`
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: config,

				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("epcc_authentication_realm.acc_test_realm", "name", regexp.MustCompile("test_realm_2")),
					resource.TestMatchResourceAttr("epcc_authentication_realm.acc_test_realm", "redirect_uris.#", regexp.MustCompile("1")),
					resource.TestMatchResourceAttr("epcc_authentication_realm.acc_test_realm", "redirect_uris.0", regexp.MustCompile("https://yahoo.co.uk/")),
					resource.TestMatchResourceAttr("epcc_authentication_realm.acc_test_realm", "duplicate_email_policy", regexp.MustCompile("api_only")),
					resource.TestMatchResourceAttr("epcc_authentication_realm.acc_test_realm", "origin_type", regexp.MustCompile("account_authentication_settings")),
				),
			},
			{
				// language=HCL
				Config:  config,
				Destroy: true,
				Check:   resource.ComposeTestCheckFunc(),
			},
			{
				// language=HCL
				Config: `
				# Example usage to update the realm associated with Customer Authentication
				data "epcc_account_authentication_settings" "account_authentication_settings" {
				}
				
				data "epcc_authentication_realm" "acc_test_realm" {
				  id = data.epcc_account_authentication_settings.account_authentication_settings.realm_id
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("data.epcc_authentication_realm.acc_test_realm", "name", regexp.MustCompile("Account Management Realm")),
					resource.TestMatchResourceAttr("data.epcc_authentication_realm.acc_test_realm", "redirect_uris.#", regexp.MustCompile("0")),
					resource.TestMatchResourceAttr("data.epcc_authentication_realm.acc_test_realm", "duplicate_email_policy", regexp.MustCompile("allowed")),
					resource.TestMatchResourceAttr("data.epcc_authentication_realm.acc_test_realm", "origin_type", regexp.MustCompile("account_authentication_settings")),
				),
			},
		},
	})
}

func TestAccResourceAuthenticationRealmWithMerchantRealmMappings(t *testing.T) {
	// language=HCL
	config := `# Example usage to update the realm associated with Customer Authentication
				data "epcc_merchant_realm_mappings" "mrm" {
				}
				
				resource "epcc_authentication_realm" "acc_test_realm" {
				  name = "test_realm_2"
				  authentication_realm_id = data.epcc_merchant_realm_mappings.mrm.realm_id
				  redirect_uris = [
					"https://yahoo.co.uk/"
					]
				  duplicate_email_policy = "api_only"
				}`
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: config,

				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("epcc_authentication_realm.acc_test_realm", "name", regexp.MustCompile("test_realm_2")),
					resource.TestMatchResourceAttr("epcc_authentication_realm.acc_test_realm", "redirect_uris.#", regexp.MustCompile("1")),
					resource.TestMatchResourceAttr("epcc_authentication_realm.acc_test_realm", "redirect_uris.0", regexp.MustCompile("https://yahoo.co.uk/")),
					resource.TestMatchResourceAttr("epcc_authentication_realm.acc_test_realm", "duplicate_email_policy", regexp.MustCompile("api_only")),
					resource.TestMatchResourceAttr("epcc_authentication_realm.acc_test_realm", "origin_type", regexp.MustCompile("merchant-realm-mappings")),
				),
			},
			{
				// language=HCL
				Config:  config,
				Destroy: true,
				Check:   resource.ComposeTestCheckFunc(),
			},
			{
				// language=HCL
				Config: `
				# Example usage to update the realm associated with Customer Authentication
				data "epcc_account_authentication_settings" "account_authentication_settings" {
				}
				
				data "epcc_authentication_realm" "acc_test_realm" {
				  id = data.epcc_account_authentication_settings.account_authentication_settings.realm_id
				}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("data.epcc_authentication_realm.acc_test_realm", "name", regexp.MustCompile("Account Management Realm")),
					resource.TestMatchResourceAttr("data.epcc_authentication_realm.acc_test_realm", "redirect_uris.#", regexp.MustCompile("0")),
					resource.TestMatchResourceAttr("data.epcc_authentication_realm.acc_test_realm", "duplicate_email_policy", regexp.MustCompile("allowed")),
					resource.TestMatchResourceAttr("data.epcc_authentication_realm.acc_test_realm", "origin_type", regexp.MustCompile("account_authentication_settings")),
				),
			},
		},
	})
}
