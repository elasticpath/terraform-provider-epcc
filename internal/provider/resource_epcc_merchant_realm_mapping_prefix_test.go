package provider

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceMerchantRealmMappingPrefix(t *testing.T) {
	myRandSeq := randSeq(10)
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccResourceMerchantRealmMappingPrefix, myRandSeq),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("epcc_merchant-realm-mapping-prefix.test_merchant_realm_mapping_prefix", "prefix", myRandSeq),
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
  prefix = "%s"
}
`

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

func randSeq(n int) string {

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
