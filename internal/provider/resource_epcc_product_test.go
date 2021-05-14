package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceProduct(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: productTestSteps[0],
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("epcc_product.acc_test_epcc_product_1_physical_draft", "sku", regexp.MustCompile("acc_test_epcc_product_1")),
					resource.TestMatchResourceAttr("epcc_product.acc_test_epcc_product_1_physical_draft", "name", regexp.MustCompile("TestAccResourceProduct1")),
					resource.TestMatchResourceAttr("epcc_product.acc_test_epcc_product_1_physical_draft", "commodity_type", regexp.MustCompile("physical")),
					resource.TestMatchResourceAttr("epcc_product.acc_test_epcc_product_1_physical_draft", "description", regexp.MustCompile("Draft description")),
					resource.TestMatchResourceAttr("epcc_product.acc_test_epcc_product_1_physical_draft", "mpn", regexp.MustCompile("mfg-acc_test_epcc_product_1")),
					resource.TestMatchResourceAttr("epcc_product.acc_test_epcc_product_1_physical_draft", "status", regexp.MustCompile("draft")),
				),
			},
			{
				Config: productTestSteps[1],
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("epcc_product.acc_test_epcc_product_1_physical_live", "sku", regexp.MustCompile("acc_test_epcc_product_1")),
					resource.TestMatchResourceAttr("epcc_product.acc_test_epcc_product_1_physical_live", "name", regexp.MustCompile("Test Acc Resource Product1 Updated")),
					resource.TestMatchResourceAttr("epcc_product.acc_test_epcc_product_1_physical_live", "commodity_type", regexp.MustCompile("physical")),
					resource.TestMatchResourceAttr("epcc_product.acc_test_epcc_product_1_physical_live", "description", regexp.MustCompile("Live description")),
					resource.TestMatchResourceAttr("epcc_product.acc_test_epcc_product_1_physical_live", "mpn", regexp.MustCompile("mfg-acc_test_epcc_product_1")),
					resource.TestMatchResourceAttr("epcc_product.acc_test_epcc_product_1_physical_live", "status", regexp.MustCompile("live")),
				),
			},
		},
	})
}

var productTestSteps = [...]string{
	`
resource "epcc_product" "acc_test_epcc_product_1_physical_draft" {
	sku            = "acc_test_epcc_product_1"
	name           = "TestAccResourceProduct1"
	commodity_type = "physical"
	description    = "Draft description"
	mpn            = "mfg-acc_test_epcc_product_1"
	status         = "draft"
  }`,
	`
 resource "epcc_product" "acc_test_epcc_product_1_physical_live" {
	sku            = "acc_test_epcc_product_1"
	name           = "Test Acc Resource Product1 Updated"
	commodity_type = "physical"
	description    = "Live description"
	mpn            = "mfg-acc_test_epcc_product_1"
	status         = "live"
  }`,
}
