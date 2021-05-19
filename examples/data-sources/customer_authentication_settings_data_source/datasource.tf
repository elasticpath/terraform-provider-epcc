terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source  = "elasticpath.com/elasticpath/epcc"
    }
  }
}
provider "epcc" {
}
data "epcc_customer_authentication_settings" "example" {
}

output "epcc_customer_authentication_settings_realm_id" {
  value = data.epcc_customer_authentication_settings.example.realm_id
}

output "epcc_customer_authentication_settings_client_id" {
  value = data.epcc_customer_authentication_settings.example.client_id
}