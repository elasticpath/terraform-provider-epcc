terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source  = "elasticpath.com/elasticpath/epcc"
    }
  }
}

resource "epcc_product" "my_first_terraform_physical_product" {
  sku            = "tf-product-1"
  name           = "TFProduct1Physical"
  commodity_type = "physical"
  description    = "Terraform Physical Product 1"
  mpn            = "mfg-part-1"
  status         = "live"
}

resource "epcc_product" "my_second_physical_terraform_product" {
  sku            = "tf-product-2"
  name           = "TFProduct2Physical"
  commodity_type = "physical"
}

resource "epcc_product" "my_digital_terraform_product" {
  sku            = "tf-product-3"
  name           = "TFProduct3Digital"
  commodity_type = "digital"
}

resource "epcc_file" "product_logo" {
  file      = "iVBORw0KGgoAAAANSUhEUgAAACAAAAAgAQMAAABJtOi3AAAABlBMVEX///8AAABVwtN+AAAACXBIWXMAAA7EAAAOxAGVKw4bAAAAfklEQVQImR3OsQkDMQwF0G9cpPQIGkVrpTDWdSlvpYMrUmaEOGQBQxoXwspPVDwQEvoC/qUbsYNEB9KPHAO4fEh5E7kT3TkwIU3Pjmr7QA2ZqEsd1Y1co6Elb7A8G/QxHBovh8Q5UeI2eLl0pCUHsY1tMCgWg8LJcxKuIfOdL8H1PrJhZV++AAAAAElFTkSuQmCC"
  file_name = "file.png"
  public    = true
}

resource "epcc_product" "product_with_file" {
  sku            = "1234556"
  name           = "Product With File"
  commodity_type = "physical"
  files          = [epcc_file.product_logo.id]
}

/* Intentionally invalid - status is set to an invalid value to test the validator
resource "epcc_product" "my_badstatus_terraform_prooduct" {
  name            = "TFBadProduct1"
  commodity_type  = "digital"
  sku             = "tf-product-bad1"
  slug            = "slimy"
  status          = "inprogress"
}
*/

/* Intentionally invalid - commodity_type is set to an invalid value to test the validator
resource "epcc_product" "my_bad_terraform_prooduct" {
  name            = "TFBadProduct2"`
  commodity_type  = "blahblah"
  sku             = "tf-product-bad2"
}
*/