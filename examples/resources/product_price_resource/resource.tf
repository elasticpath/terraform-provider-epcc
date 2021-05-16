terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source  = "elasticpath.com/elasticpath/epcc"
    }
  }
}

resource "epcc_product" "my_first_product" {
  sku            = "tf-product-1-3"
  name           = "TFProduct1Physical"
  commodity_type = "physical"
  description    = "Terraform Physical Product 1"
  mpn            = "mfg-part-1"
  status         = "live"
}


resource "epcc_pricebook" "my_first_terraform_pricebook" {
  name        = "TFPricebook1"
  description = "Terraform 1"
}

resource "epcc_currency" "CAD" {
  code               = "CAD"
  exchange_rate      = 1
  format             = "$${price}"
  decimal_point      = "."
  thousand_separator = ","
  decimal_places     = 0
  default            = false
  enabled            = true
}

resource "epcc_currency" "NZD" {
  code               = "NZD"
  exchange_rate      = 1
  format             = "$${price}"
  decimal_point      = "."
  thousand_separator = ","
  decimal_places     = 0
  default            = false
  enabled            = true
}

resource "epcc_product_price" "price" {
  sku          = epcc_product.my_first_product.sku
  pricebook_id = epcc_pricebook.my_first_terraform_pricebook.id

  currency {
    code         = epcc_currency.CAD.code
    amount       = 420
    includes_tax = true
  }

  currency {
    code         = epcc_currency.NZD.code
    amount       = 424
    includes_tax = false
  }
}