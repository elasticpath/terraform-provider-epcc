package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceMerchantRealmMappingPrefix(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceMerchantRealmMappingPrefix,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("epcc_merchant-realm-mapping-prefix.test_merchant_realm_mapping_prefix", "prefix", "abcdefgh"),
				),
			},
		},
	})
}

// language=HCL
const testAccResourceMerchantRealmMappingPrefix = `
data "epcc_merchant_realm_mappings" "test_merchant_realm_mappings" {
}
resource "epcc_merchant_realm_mapping_prefix" "test_merchant_realm_mapping_prefix" {
  merchant_realm_mapping_id = data.epcc_merchant_realm_mappings.test_merchant_realm_mappings.merchant_realm_mapping_id
  prefix = "abcdefgh"
}
`
