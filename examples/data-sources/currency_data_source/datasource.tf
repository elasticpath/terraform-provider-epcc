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

data "epcc_currency" "USD" {
  code = "USD"
}

output "usd_currency_id" {
  value = data.epcc_currency.USD.id
}

