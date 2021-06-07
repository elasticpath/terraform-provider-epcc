terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source = "elasticpath.com/elasticpath/epcc"
    }
  }
}

provider "epcc" {
}

data "epcc_merchant_realm_mappings" "test_merchant_realm_mappings" {
}

resource "epcc_merchant_realm_mapping_prefix" "test_merchant_realm_mapping_prefix" {
  merchant_realm_mapping_id = data.epcc_merchant_realm_mappings.test_merchant_realm_mappings.merchant_realm_mapping_id
  prefix = "abcdefgh"
}