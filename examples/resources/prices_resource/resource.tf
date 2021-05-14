terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source = "elasticpath.com/elasticpath/epcc"
    }
  }
}
provider "epcc" {
  beta_features = "account-management"
}

resource "epcc_currency" "GBP" {
  type = "currency"
  code = "GBP"
  exchange_rate = 1
  format = "£{price}"
  decimal_point = "."
  thousand_separator = ","
  decimal_places = 0
  default = false
  enabled = true
}

resource "epcc_currency" "KRW" {
  type = "currency"
  code = "KRW"
  exchange_rate = 1
  format = "₩{price}"
  decimal_point = "."
  thousand_separator = ","
  decimal_places = 0
  default = false
  enabled = true
}

data "epcc_product" "acc_test_example_product" {
  id = "5fc6b278-5a5f-4ed9-8f8e-0a60e3341310"
}

output "product_name" {
  value = data.epcc_product.acc_test_example_product.name
}

data "epcc_pricebook" "acc_test_example_pricebook" {
  id = "90c5733a-6db5-4b42-9c7d-48f689bca3b8"
}

output "pricebook_name" {
  value = data.epcc_pricebook.acc_test_example_pricebook.name
}

resource "epcc_prices" "acc_test_prices" {
  sku = data.epcc_product.acc_test_example_product.sku
  pricebook_id = data.epcc_pricebook.acc_test_example_pricebook.id
  
  prices = tomap({
     epcc_currency.GBP.code = {
      amount = 100
      includes_tax = false
    }
    epcc_currency.KRW.code = {
      amount = 900
      includes_tax = true
    }
  })

}

