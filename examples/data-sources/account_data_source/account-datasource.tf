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
data "epcc_account" "example" {
  id = "99915f2a-1b74-4860-b8df-325cb44a9f63"
}

output "account_name" {
  value = data.epcc_account.example.name
}

