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


resource "epcc_realm" "test_realm_for_profile" {
  name = "test_realm"
  redirect_uris = [
    "https://google.com/"
  ]
  duplicate_email_policy = "allowed"
  origin_id              = "hello-world"
  origin_type            = "customer-authentication-settings"
}

resource "epcc_profile" "test_profile" {
  name          = "test_profile"
  discovery_url = "https://elasticpath-customer.okta.com/.well-known/openid-configuration"
  client_id     = "epcc-integrations"
  client_secret = "86c8986d-e1b2-4ce4-a24c-8430ec1ab383"
  realm_id      = epcc_realm.test_realm_for_profile.id
}