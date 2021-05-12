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

resource "epcc_account" "hello_world_account" {
  name            = "Steve's First Terraform Account IV"
  legal_name      = "Steve's First Terraform Account LTD"
  registration_id = "3"
}