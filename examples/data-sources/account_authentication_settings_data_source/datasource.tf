terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source = "elasticpath.com/elasticpath/epcc"
    }
  }
}
provider "epcc" {
  beta_features = "account-authentication-settings"
}
data "epcc_account_authentication_settings" "example" {
}

output "epcc_account_authentication_settings_realm_id" {
  value = data.epcc_account_authentication_settings.example.realm_id
}

output "epcc_account_authentication_settings_client_id" {
  value = data.epcc_account_authentication_settings.example.client_id
}