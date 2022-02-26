package provider

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestMain(m *testing.M) {
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	fmt.Printf("Seed is %d\n", seed)
	os.Exit(m.Run())
}

func TestAccResourceMerchantRealmMappingPrefix(t *testing.T) {
	myRandSeq := randSeq(20)
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccResourceMerchantRealmMappingPrefix, myRandSeq),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("epcc_merchant_realm_mapping.mrm", "prefix", myRandSeq),
				),
			},
		},
	})
}

// language=HCL
const testAccResourceMerchantRealmMappingPrefix = `
resource "epcc_merchant_realm_mapping" "mrm" {
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
