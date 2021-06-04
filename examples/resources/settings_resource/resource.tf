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


resource "epcc_settings" "test_settings" {
  type                = "settings"
  page_length         = 20
  list_child_products = false
  additional_languages = [
    "fr",
  "de"]
  calculation_method = "simple"
}