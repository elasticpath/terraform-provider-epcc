package test

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccResourcePaymentGateway_Stripe(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				// language=HCL
				Config: `
                   resource "epcc_payment_gateway" "stripe" {
                     slug = "stripe"
                     enabled = true
					 options = {
					 	login = "test_login_creds"	
					 }
                   }
                `,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("epcc_payment_gateway.stripe", "slug", "stripe"),
					resource.TestCheckResourceAttr("epcc_payment_gateway.stripe", "enabled", "true"),
					resource.TestCheckResourceAttr("epcc_payment_gateway.stripe", "options.login", "test_login_creds"),
				),
			},
			{
				ResourceName:            "epcc_payment_gateway.stripe",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"options"}, // not importing secrets
			},
		},
	})
}

func TestAccResourcePaymentGateway_Manual(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				// language=HCL
				Config: `
                   resource "epcc_payment_gateway" "manual" {
                     slug = "manual"
                     enabled = true
                   }
                `,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("epcc_payment_gateway.manual", "slug", "manual"),
					resource.TestCheckResourceAttr("epcc_payment_gateway.manual", "enabled", "true"),
				),
			},
			{
				ResourceName:      "epcc_payment_gateway.manual",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourcePaymentGateway_CyberSource(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				// language=HCL
				Config: `
                   resource "epcc_payment_gateway" "cyber_source" {
                     slug = "cyber_source"
                     enabled = true
				 	 test = true
                     options = {
						login = "test_login"
						password = "test_password"
                     }
                   }
                `,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("epcc_payment_gateway.cyber_source", "slug", "cyber_source"),
					resource.TestCheckResourceAttr("epcc_payment_gateway.cyber_source", "enabled", "true"),
					resource.TestCheckResourceAttr("epcc_payment_gateway.cyber_source", "test", "true"),
					resource.TestCheckResourceAttr("epcc_payment_gateway.cyber_source", "options.login", "test_login"),
					resource.TestCheckResourceAttr("epcc_payment_gateway.cyber_source", "options.password", "test_password"),
				),
			},
			{
				ResourceName:            "epcc_payment_gateway.cyber_source",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"options"}, // not importing secrets
			},
		},
	})
}

func TestAccResourcePaymentGateway_Braintree(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				// language=HCL
				Config: `
                   resource "epcc_payment_gateway" "braintree" {
                     slug = "braintree"
                     enabled = true
                     options = {
						merchant_id = "test_merchant_id"
						private_key = "test_private_key"
						public_key = "test_public_key"
						environment = "sandbox"
                     }
                   }
                `,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("epcc_payment_gateway.braintree", "slug", "braintree"),
					resource.TestCheckResourceAttr("epcc_payment_gateway.braintree", "enabled", "true"),
					resource.TestCheckResourceAttr("epcc_payment_gateway.braintree", "options.merchant_id", "test_merchant_id"),
					resource.TestCheckResourceAttr("epcc_payment_gateway.braintree", "options.private_key", "test_private_key"),
					resource.TestCheckResourceAttr("epcc_payment_gateway.braintree", "options.public_key", "test_public_key"),
					resource.TestCheckResourceAttr("epcc_payment_gateway.braintree", "options.environment", "sandbox"),
				),
			},
			{
				ResourceName:            "epcc_payment_gateway.braintree",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"options"}, // not importing secrets
			},
		},
	})
}
