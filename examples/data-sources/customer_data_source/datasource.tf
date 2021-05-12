terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source  = "elasticpath.com/elasticpath/epcc"
    }
  }
}

data "epcc_customer" "example" {
  id = "49d1443e-b7d7-42de-b624-ca85522f4bfd"
}

output "customer_name" {
  value = data.epcc_customer.example.name
}

