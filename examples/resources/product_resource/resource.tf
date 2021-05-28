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
  file_name = "ep.png"
  file_hash = filemd5("ep.png")
  public    = true
}

resource "epcc_product" "product_with_file" {
  sku            = "1234556"
  name           = "Product With File"
  commodity_type = "physical"
  files          = [epcc_file.product_logo.id]
  prices = {
    yo   = "bar"
    five = 3
    USD = {
      attributes = "foo"

    }
  }
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