terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source  = "elasticpath.com/elasticpath/epcc"
    }
  }
}
provider "epcc" {
}
data "epcc_merchant_realm_mappings" "example" {
}

output "epcc_merchant_realm_mappings_id" {
  value = data.epcc_merchant_realm_mappings.example.merchant_realm_mapping_id
}

output "epcc_merchant_realm_mappings_store_id" {
  value = data.epcc_merchant_realm_mappings.example.store_id
}

output "epcc_merchant_realm_mappings_store_prefix" {
  value = data.epcc_merchant_realm_mappings.example.prefix
}