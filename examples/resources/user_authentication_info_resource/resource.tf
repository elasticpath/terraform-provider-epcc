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

resource "epcc_realm" "test_realm_for_user_authentication_info" {
  name = "test_realm"
  redirect_uris = [
    "https://google.com/"
  ]
  duplicate_email_policy = "allowed"
  origin_id              = "hello-world"
  origin_type            = "customer-authentication-settings"
}

resource "epcc_user_authentication_info" "test_user_authentication_info" {
  name     = "John Doe 2"
  email    = "john.doe@banks.com"
  realm_id = epcc_realm.test_realm_for_user_authentication_info.id
}