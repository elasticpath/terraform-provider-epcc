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

data "epcc_settings" "example" {
}

output "page_length" {
  value = data.epcc_settings.example.page_length
}

