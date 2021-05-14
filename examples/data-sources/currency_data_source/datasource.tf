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
data "epcc_currency" "example" {
  id                 = "99915f2a-1b74-4860-b8df-325cb44a9f63"
  type               = "currency"
  code               = "CHF"
  exchange_rate      = 2
  format             = "£{price}"
  decimal_point      = "."
  thousand_separator = ","
  decimal_places     = 0
  default            = false
  enabled            = true
}

output "currency_code" {
  value = data.epcc_currency.example.code
}

