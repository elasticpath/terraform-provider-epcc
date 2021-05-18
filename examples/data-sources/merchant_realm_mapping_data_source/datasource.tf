terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source  = "elasticpath.com/elasticpath/epcc"
    }
  }
}

data "epcc_merchant_realm_mapping" "example" {
  id = "987aab7d-b00d-433a-914a-bbf7a3705c73"
}

output "mrm_example" {
  value = data.epcc_merchant_realm_mapping.example
}

