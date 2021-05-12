terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source  = "elasticpath.com/elasticpath/epcc"
    }
  }
}

resource "epcc_customer" "my_first_terraform_customer" {
  name  = "A customer not an account"
  email = "steve@ramage.com"
}


resource "epcc_customer" "my_second_terraform_customer" {
  name  = "A second customer not an account"
  email = "ste4e@ramage.com"
}