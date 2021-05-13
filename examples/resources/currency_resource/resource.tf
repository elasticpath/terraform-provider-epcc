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

resource "epcc_currency" "british_pound_currency" {
  type = "currency"
  code = "GBP"
  exchange_rate = 2
  format = "£{price}"
  decimal_point = "."
  thousand_separator = ","
  decimal_places = 0
  default = false
  enabled = true
}