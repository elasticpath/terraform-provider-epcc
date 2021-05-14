terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source  = "elasticpath.com/elasticpath/epcc"
    }
  }
}

resource "epcc_flow" "hello_world_flow" {
  name        = "test5"
  slug        = "test5"
  description = "This is a Terraform test"
  enabled     = true
}