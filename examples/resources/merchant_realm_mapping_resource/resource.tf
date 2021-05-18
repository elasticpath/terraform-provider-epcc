terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source = "elasticpath.com/elasticpath/epcc"
    }
  }
}

resource "epcc_merchant_realm_mapping" "hello_world_mapping" {
  type = "merchant-realm-mappings"
  prefix = "test"
}