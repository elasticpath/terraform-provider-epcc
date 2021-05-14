terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source  = "elasticpath.com/elasticpath/epcc"
    }
  }
}

resource "epcc_customer" "my_first_terraform_customer" {
  name  = "Jon Smith"
  email = "jon@smith.com"
}

resource "epcc_customer" "my_second_terraform_customer" {
  name  = "Jane Smith"
  email = "jane@smith.com"
}