package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceMerchantRealmMapping(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceMerchantRealmMapping,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("epcc_merchant_realm_mapping.acc_test_merchantRealmMapping", "type", regexp.MustCompile("merchant-realm-mappings")),
					resource.TestMatchResourceAttr("epcc_merchant_realm_mapping.acc_test_merchantRealmMapping", "slug", regexp.MustCompile("test")),
				),
			},
		},
	})
}

const testAccResourceMerchantRealmMapping = `
resource "epcc_merchant_realm_mapping" "hello_world_mapping" {
  type = "merchant-realm-mappings"
  prefix = "test"
}
`
