package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"regexp"
	"strconv"
	"testing"
	"time"
)

var timestamp = time.Now().UnixNano()

func TestAccResourceProduct(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: productTestSteps[0],
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("epcc_product.acc_test_epcc_product_1_physical_draft", "sku", regexp.MustCompile(strconv.FormatInt(timestamp, 10))),
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
					resource.TestMatchResourceAttr("epcc_product.acc_test_epcc_product_1_physical_live", "sku", regexp.MustCompile(strconv.FormatInt(timestamp, 10))),
					resource.TestMatchResourceAttr("epcc_product.acc_test_epcc_product_1_physical_live", "name", regexp.MustCompile("Test Acc Resource Product1 Updated")),
					resource.TestMatchResourceAttr("epcc_product.acc_test_epcc_product_1_physical_live", "commodity_type", regexp.MustCompile("physical")),
					resource.TestMatchResourceAttr("epcc_product.acc_test_epcc_product_1_physical_live", "description", regexp.MustCompile("Live description")),
					resource.TestMatchResourceAttr("epcc_product.acc_test_epcc_product_1_physical_live", "mpn", regexp.MustCompile("mfg-acc_test_epcc_product_1")),
					resource.TestMatchResourceAttr("epcc_product.acc_test_epcc_product_1_physical_live", "status", regexp.MustCompile("live")),
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
var productTestSteps = [...]string{
	// language=HCL
	fmt.Sprintf(`
resource "epcc_product" "acc_test_epcc_product_1_physical_draft" {
	sku            = "%d"
	name           = "TestAccResourceProduct1"
	commodity_type = "physical"
	description    = "Draft description"
	mpn            = "mfg-acc_test_epcc_product_1"
	status         = "draft"
  }`, timestamp),
	// language=HCL
	fmt.Sprintf(`
 resource "epcc_product" "acc_test_epcc_product_1_physical_live" {
	sku            = "%d"
	name           = "Test Acc Resource Product1 Updated"
	commodity_type = "physical"
	description    = "Live description"
	mpn            = "mfg-acc_test_epcc_product_1"
	status         = "live"
  }`, timestamp),
	fmt.Sprintf(
		// language=HCL
		`resource "epcc_file" "product_logo" {
	  file = "iVBORw0KGgoAAAANSUhEUgAAACAAAAAgAQMAAABJtOi3AAAABlBMVEX///8AAABVwtN+AAAACXBIWXMAAA7EAAAOxAGVKw4bAAAAfklEQVQImR3OsQkDMQwF0G9cpPQIGkVrpTDWdSlvpYMrUmaEOGQBQxoXwspPVDwQEvoC/qUbsYNEB9KPHAO4fEh5E7kT3TkwIU3Pjmr7QA2ZqEsd1Y1co6Elb7A8G/QxHBovh8Q5UeI2eLl0pCUHsY1tMCgWg8LJcxKuIfOdL8H1PrJhZV++AAAAAElFTkSuQmCC"
	  file_name = "file.png"
	  public = true
	}

	resource "epcc_product" "acc_test_product_with_file" { 
		sku = "%d"
		name = "Test Product"
		commodity_type = "physical"
		status = "live"
		files = [ epcc_file.product_logo.id ]
	}
	`, timestamp),
	fmt.Sprintf(
		// language=HCL
		`resource "epcc_file" "product_logo_1" {
	  file = "iVBORw0KGgoAAAANSUhEUgAAACAAAAAgAQMAAABJtOi3AAAABlBMVEX///8AAABVwtN+AAAACXBIWXMAAA7EAAAOxAGVKw4bAAAAfklEQVQImR3OsQkDMQwF0G9cpPQIGkVrpTDWdSlvpYMrUmaEOGQBQxoXwspPVDwQEvoC/qUbsYNEB9KPHAO4fEh5E7kT3TkwIU3Pjmr7QA2ZqEsd1Y1co6Elb7A8G/QxHBovh8Q5UeI2eLl0pCUHsY1tMCgWg8LJcxKuIfOdL8H1PrJhZV++AAAAAElFTkSuQmCC"
	  file_name = "file.png"
	  public = false
	}

	resource "epcc_product" "acc_test_product_with_file" { 
		sku = "%d"
		name = "Test Product"
		commodity_type = "physical"
		status = "live"
		files = [ epcc_file.product_logo_1.id ]
	}
	`, timestamp),
	fmt.Sprintf(
		// language=HCL
		`resource "epcc_file" "product_logo_1" {
	  file = "iVBORw0KGgoAAAANSUhEUgAAACAAAAAgAQMAAABJtOi3AAAABlBMVEX///8AAABVwtN+AAAACXBIWXMAAA7EAAAOxAGVKw4bAAAAfklEQVQImR3OsQkDMQwF0G9cpPQIGkVrpTDWdSlvpYMrUmaEOGQBQxoXwspPVDwQEvoC/qUbsYNEB9KPHAO4fEh5E7kT3TkwIU3Pjmr7QA2ZqEsd1Y1co6Elb7A8G/QxHBovh8Q5UeI2eLl0pCUHsY1tMCgWg8LJcxKuIfOdL8H1PrJhZV++AAAAAElFTkSuQmCC"
	  file_name = "file.png"
	  public = false
	}

	resource "epcc_file" "product_logo_2" {
		file = "iVBORw0KGgoAAAANSUhEUgAAACAAAAAgAQMAAABJtOi3AAAABlBMVEX///8AAABVwtN+AAAACXBIWXMAAA7EAAAOxAGVKw4bAAAAfklEQVQImR3OsQkDMQwF0G9cpPQIGkVrpTDWdSlvpYMrUmaEOGQBQxoXwspPVDwQEvoC/qUbsYNEB9KPHAO4fEh5E7kT3TkwIU3Pjmr7QA2ZqEsd1Y1co6Elb7A8G/QxHBovh8Q5UeI2eLl0pCUHsY1tMCgWg8LJcxKuIfOdL8H1PrJhZV++AAAAAElFTkSuQmCC"
		file_name = "file.png"
		public = false
	}

	resource "epcc_product" "acc_test_product_with_file" { 
		sku = "%d"
		name = "Test Product"
		commodity_type = "physical"
		status = "live"
		files = [ epcc_file.product_logo_2.id ]
	}
	`, timestamp),
	fmt.Sprintf(
		// language=HCL
		`resource "epcc_file" "product_logo_1" {
	  file = "iVBORw0KGgoAAAANSUhEUgAAACAAAAAgAQMAAABJtOi3AAAABlBMVEX///8AAABVwtN+AAAACXBIWXMAAA7EAAAOxAGVKw4bAAAAfklEQVQImR3OsQkDMQwF0G9cpPQIGkVrpTDWdSlvpYMrUmaEOGQBQxoXwspPVDwQEvoC/qUbsYNEB9KPHAO4fEh5E7kT3TkwIU3Pjmr7QA2ZqEsd1Y1co6Elb7A8G/QxHBovh8Q5UeI2eLl0pCUHsY1tMCgWg8LJcxKuIfOdL8H1PrJhZV++AAAAAElFTkSuQmCC"
	  file_name = "file.png"
	  public = false
	}

	resource "epcc_file" "product_logo_2" {
		file = "iVBORw0KGgoAAAANSUhEUgAAACAAAAAgAQMAAABJtOi3AAAABlBMVEX///8AAABVwtN+AAAACXBIWXMAAA7EAAAOxAGVKw4bAAAAfklEQVQImR3OsQkDMQwF0G9cpPQIGkVrpTDWdSlvpYMrUmaEOGQBQxoXwspPVDwQEvoC/qUbsYNEB9KPHAO4fEh5E7kT3TkwIU3Pjmr7QA2ZqEsd1Y1co6Elb7A8G/QxHBovh8Q5UeI2eLl0pCUHsY1tMCgWg8LJcxKuIfOdL8H1PrJhZV++AAAAAElFTkSuQmCC"
		file_name = "file.png"
		public = false
	}

	resource "epcc_product" "acc_test_product_with_file" { 
		sku = "%d"
		name = "Test Product"
		commodity_type = "physical"
		status = "live"
		files = [ epcc_file.product_logo_2.id, epcc_file.product_logo_1.id]
	}
	`, timestamp),
	fmt.Sprintf(
		// language=HCL
		`resource "epcc_file" "product_logo_1" {
	  file = "iVBORw0KGgoAAAANSUhEUgAAACAAAAAgAQMAAABJtOi3AAAABlBMVEX///8AAABVwtN+AAAACXBIWXMAAA7EAAAOxAGVKw4bAAAAfklEQVQImR3OsQkDMQwF0G9cpPQIGkVrpTDWdSlvpYMrUmaEOGQBQxoXwspPVDwQEvoC/qUbsYNEB9KPHAO4fEh5E7kT3TkwIU3Pjmr7QA2ZqEsd1Y1co6Elb7A8G/QxHBovh8Q5UeI2eLl0pCUHsY1tMCgWg8LJcxKuIfOdL8H1PrJhZV++AAAAAElFTkSuQmCC"
	  file_name = "file.png"
	  public = false
	}

	resource "epcc_file" "product_logo_2" {
		file = "iVBORw0KGgoAAAANSUhEUgAAACAAAAAgAQMAAABJtOi3AAAABlBMVEX///8AAABVwtN+AAAACXBIWXMAAA7EAAAOxAGVKw4bAAAAfklEQVQImR3OsQkDMQwF0G9cpPQIGkVrpTDWdSlvpYMrUmaEOGQBQxoXwspPVDwQEvoC/qUbsYNEB9KPHAO4fEh5E7kT3TkwIU3Pjmr7QA2ZqEsd1Y1co6Elb7A8G/QxHBovh8Q5UeI2eLl0pCUHsY1tMCgWg8LJcxKuIfOdL8H1PrJhZV++AAAAAElFTkSuQmCC"
		file_name = "file.png"
		public = false
	}

	resource "epcc_product" "acc_test_product_with_file" { 
		sku = "%d"
		name = "Test Product"
		commodity_type = "physical"
		status = "live"
		files = [ ]
	}
	`, timestamp),
}
