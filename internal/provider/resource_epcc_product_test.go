package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"io/ioutil"
	"regexp"
	"strconv"
	"testing"
	"time"
)

var timestamp = time.Now().UnixNano()

func TestAccResourceProduct(t *testing.T) {
	tempDir := t.TempDir()
	productTestSteps := productTestSteps(tempDir)
	resource.UnitTest(t, resource.TestCase{
		PreCheck: func() {
			err := ioutil.WriteFile(tempDir+"/hello_world.txt", []byte("hello world"), 0644)
			if err != nil {
				t.Fatal(err)
			}
			testAccPreCheck(t)
		},
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: productTestSteps[0],
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("epcc_product.acc_test_epcc_product_1", "sku", regexp.MustCompile(strconv.FormatInt(timestamp, 10))),
					resource.TestMatchResourceAttr("epcc_product.acc_test_epcc_product_1", "name", regexp.MustCompile("TestAccResourceProduct1")),
					resource.TestMatchResourceAttr("epcc_product.acc_test_epcc_product_1", "commodity_type", regexp.MustCompile("physical")),
					resource.TestMatchResourceAttr("epcc_product.acc_test_epcc_product_1", "description", regexp.MustCompile("Draft description")),
					resource.TestMatchResourceAttr("epcc_product.acc_test_epcc_product_1", "mpn", regexp.MustCompile("mfg-acc_test_epcc_product_1")),
					resource.TestMatchResourceAttr("epcc_product.acc_test_epcc_product_1", "status", regexp.MustCompile("draft")),
				),
			},
			{
				Config: productTestSteps[1],
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("epcc_product.acc_test_epcc_product_1", "sku", regexp.MustCompile(strconv.FormatInt(timestamp, 10))),
					resource.TestMatchResourceAttr("epcc_product.acc_test_epcc_product_1", "name", regexp.MustCompile("Test Acc Resource Product1 Updated")),
					resource.TestMatchResourceAttr("epcc_product.acc_test_epcc_product_1", "commodity_type", regexp.MustCompile("physical")),
					resource.TestMatchResourceAttr("epcc_product.acc_test_epcc_product_1", "description", regexp.MustCompile("Live description")),
					resource.TestMatchResourceAttr("epcc_product.acc_test_epcc_product_1", "mpn", regexp.MustCompile("mfg-acc_test_epcc_product_1")),
					resource.TestMatchResourceAttr("epcc_product.acc_test_epcc_product_1", "status", regexp.MustCompile("live")),
				),
			},
			{
				Config: productTestSteps[2],
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: productTestSteps[3],
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: productTestSteps[4],
				Check:  resource.ComposeTestCheckFunc(),
			},
			{
				Config: productTestSteps[5],
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
	})
}

/*
 * Step 1: create the product
 * Step 2: Update the same product: change the name and status (at least that's the idea)
 */
func productTestSteps(tempDir string) []string {
	return []string{
		// language=HCL
		fmt.Sprintf(`
resource "epcc_product" "acc_test_epcc_product_1" {
	sku            = "%d-pr"
	name           = "TestAccResourceProduct1"
	commodity_type = "physical"
	description    = "Draft description"
	mpn            = "mfg-acc_test_epcc_product_1"
	status         = "draft"
  }`, timestamp),

		// language=HCL
		fmt.Sprintf(`
 resource "epcc_product" "acc_test_epcc_product_1" {
	sku            = "%d-pr"
	name           = "Test Acc Resource Product1 Updated"
	commodity_type = "physical"
	description    = "Live description"
	mpn            = "mfg-acc_test_epcc_product_1"
	status         = "live"
  }`, timestamp),
		fmt.Sprintf(
			// language=HCL
			`
	resource "epcc_file" "product_logo"{
		file_name = "%s/hello_world.txt"
		file_hash = filemd5("%s/hello_world.txt")
		public = true
	}

	resource "epcc_product" "acc_test_product_with_file" { 
		sku = "%d-pr-2"
		name = "Test Product"
		commodity_type = "physical"
		status = "live"
		files = [ epcc_file.product_logo.id ]
	}
	`, tempDir, tempDir, timestamp),
		fmt.Sprintf(
			// language=HCL
			`resource "epcc_file" "product_logo_1"{
		file_name = "%s/hello_world.txt"
		file_hash = filemd5("%s/hello_world.txt")
		public = true
	}

	resource "epcc_product" "acc_test_product_with_file" { 
		sku = "%d-pr-2"
		name = "Test Product"
		commodity_type = "physical"
		status = "live"
		files = [ epcc_file.product_logo_1.id ]
	}
	`, tempDir, tempDir, timestamp),
		fmt.Sprintf(
			// language=HCL
			`resource "epcc_file" "product_logo_1"{
		file_name = "%s/hello_world.txt"
		file_hash = filemd5("%s/hello_world.txt")
		public = true
	}

	resource "epcc_file" "product_logo_2"{
		file_name = "%s/hello_world.txt"
		file_hash = filemd5("%s/hello_world.txt")
		public = true
	}

	resource "epcc_product" "acc_test_product_with_file" { 
		sku = "%d-pr-2"
		name = "Test Product"
		commodity_type = "physical"
		status = "live"
		files = [ epcc_file.product_logo_2.id ]
	}
	`, tempDir, tempDir, tempDir, tempDir, timestamp),
		fmt.Sprintf(
			// language=HCL
			`resource "epcc_file" "product_logo_1"{
		file_name = "%s/hello_world.txt"
		file_hash = filemd5("%s/hello_world.txt")
		public = true
	}

	resource "epcc_file" "product_logo_2"{
		file_name = "%s/hello_world.txt"
		file_hash = filemd5("%s/hello_world.txt")
		public = true
	}

	resource "epcc_product" "acc_test_product_with_file" { 
		sku = "%d-pr-2"
		name = "Test Product"
		commodity_type = "physical"
		status = "live"
		files = [ epcc_file.product_logo_2.id, epcc_file.product_logo_1.id]
	}
	`, tempDir, tempDir, tempDir, tempDir, timestamp),
		fmt.Sprintf(
			// language=HCL
			`resource "epcc_file" "product_logo_1"{
		file_name = "%s/hello_world.txt"
		file_hash = filemd5("%s/hello_world.txt")
		public = true
	}

	resource "epcc_file" "product_logo_2"{
		file_name = "%s/hello_world.txt"
		file_hash = filemd5("%s/hello_world.txt")
		public = true
	}

	resource "epcc_product" "acc_test_product_with_file" { 
		sku = "%d-pr-2"
		name = "Test Product"
		commodity_type = "physical"
		status = "live"
		files = [ ]
	}
	`, tempDir, tempDir, tempDir, tempDir, timestamp),
	}
}
