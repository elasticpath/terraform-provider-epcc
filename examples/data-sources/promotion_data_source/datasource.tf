terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source  = "elasticpath.com/elasticpath/epcc"
    }
  }
}
provider "epcc" {
  beta_features = "account-management"
}

resource "epcc_currency" "chf" {
  code               = "CHF"
  exchange_rate      = 1
  format             = "£{price}"
  decimal_point      = "."
  thousand_separator = ","
  decimal_places     = 0
  default            = false
  enabled            = true
}

resource "epcc_promotion" "acc_test_promotion" {
  type           = "promotion"
  name           = "Promo #1"
  description    = "Initial Promotion"
  enabled        = true
  promotion_type = "fixed_discount"
  schema {
    currencies {
      currency = epcc_currency.chf.code
      amount   = 900
    }
  }

  max_discount_value {
    currency = epcc_currency.chf.code
    amount   = 960
  }
  min_cart_value {
    currency = epcc_currency.chf.code
    amount   = 100
  }
  start = "2019-05-12T00:00:00Z"
  end   = "2019-10-12T00:00:00Z"
}