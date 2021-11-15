package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccResourceIntegration(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				// language=HCL
				Config: `
                   resource "epcc_integration" "test" {
                     name = "Test Integration"
                     description = "Test Integration Description"
                     url = "https://webhook"
                     secret_key = "secret"
                     integration_type = "webhook"
                     enabled = true
                     observes = [
                       "cart.updated",
                       "cart.deleted",
                     ]
                   }
                `,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("epcc_integration.test", "id"),
					resource.TestCheckResourceAttr("epcc_integration.test", "name", "Test Integration"),
					resource.TestCheckResourceAttr("epcc_integration.test", "description", "Test Integration Description"),
					resource.TestCheckResourceAttr("epcc_integration.test", "url", "https://webhook"),
					resource.TestCheckResourceAttr("epcc_integration.test", "secret_key", "secret"),
					resource.TestCheckResourceAttr("epcc_integration.test", "integration_type", "webhook"),
					resource.TestCheckResourceAttr("epcc_integration.test", "enabled", "true"),
					resource.TestCheckResourceAttr("epcc_integration.test", "observes.0", "cart.updated"),
					resource.TestCheckResourceAttr("epcc_integration.test", "observes.1", "cart.deleted"),
					resource.TestCheckResourceAttr("epcc_integration.test", "observes.#", "2"),
				),
			},
			{
				// language=HCL
				Config: `
                   resource "epcc_integration" "test" {
                     name = "Test Integration Updated"
                     description = "Test Integration Description Updated"
                     url = "https://webhook-updated"
                     secret_key = "secret-updated"
                     integration_type = "webhook"
                     enabled = false
                      observes = [
                       "order.updated",
                       "order.deleted",
                     ]
                   }
                `,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("epcc_integration.test", "id"),
					resource.TestCheckResourceAttr("epcc_integration.test", "name", "Test Integration Updated"),
					resource.TestCheckResourceAttr("epcc_integration.test", "description", "Test Integration Description Updated"),
					resource.TestCheckResourceAttr("epcc_integration.test", "url", "https://webhook-updated"),
					resource.TestCheckResourceAttr("epcc_integration.test", "secret_key", "secret-updated"),
					resource.TestCheckResourceAttr("epcc_integration.test", "integration_type", "webhook"),
					resource.TestCheckResourceAttr("epcc_integration.test", "enabled", "false"),
					resource.TestCheckResourceAttr("epcc_integration.test", "observes.0", "order.updated"),
					resource.TestCheckResourceAttr("epcc_integration.test", "observes.1", "order.deleted"),
					resource.TestCheckResourceAttr("epcc_integration.test", "observes.#", "2"),
				),
			},
			{
				// language=HCL
				Config: `
                   resource "epcc_integration" "test_sqs" {
                     name = "Test Integration Updated"
                     description = "Test Integration Description Updated"
                     url = "https://webhook-updated"
                     integration_type = "aws_sqs"
                     region = "us-west-2"
                     aws_access_key_id = "FOOFOOFOOFOOFOO"
                     aws_secret_access_key = "BARBARBARBARBAR"
                     enabled = false
                     observes = [
                       "order.updated",
                       "order.deleted",
                     ]
                   }
                `,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("epcc_integration.test_sqs", "id"),
					resource.TestCheckResourceAttr("epcc_integration.test_sqs", "name", "Test Integration Updated"),
					resource.TestCheckResourceAttr("epcc_integration.test_sqs", "description", "Test Integration Description Updated"),
					resource.TestCheckResourceAttr("epcc_integration.test_sqs", "url", "https://webhook-updated"),
					resource.TestCheckResourceAttr("epcc_integration.test_sqs", "integration_type", "aws_sqs"),
					resource.TestCheckResourceAttr("epcc_integration.test_sqs", "enabled", "false"),
					resource.TestCheckResourceAttr("epcc_integration.test_sqs", "observes.0", "order.updated"),
					resource.TestCheckResourceAttr("epcc_integration.test_sqs", "observes.1", "order.deleted"),
					resource.TestCheckResourceAttr("epcc_integration.test_sqs", "observes.#", "2"),
					resource.TestCheckResourceAttr("epcc_integration.test_sqs", "region", "us-west-2"),
					resource.TestCheckResourceAttr("epcc_integration.test_sqs", "aws_access_key_id", "FOOFOOFOOFOOFOO"),
					resource.TestCheckResourceAttr("epcc_integration.test_sqs", "aws_secret_access_key", "BARBARBARBARBAR"),
				),
			},
		},
	})
}
