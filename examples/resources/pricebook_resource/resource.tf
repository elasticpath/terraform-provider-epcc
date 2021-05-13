terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source  = "elasticpath.com/elasticpath/epcc"
    }
  }
}

resource "epcc_pricebook" "my_first_terraform_pricebook" {
  name        = "TFPricebook1"
  description = "Terraform 1"
}

resource "epcc_pricebook" "my_second_terraform_pricebook" {
  name = "TFPricebook2"
}