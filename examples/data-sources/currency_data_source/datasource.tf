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
  id = "d8339926-3ee0-4d69-acd5-af107f7da42f"
}

output "currency_code" {
  value = data.epcc_currency.example.code
}

