terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source  = "elasticpath.com/elasticpath/epcc"
    }
  }
}

data "epcc_product" "example" {
  id = "5fc6b278-5a5f-4ed9-8f8e-0a60e3341310"
}

output "product_name" {
  value = data.epcc_product.example.name
}

