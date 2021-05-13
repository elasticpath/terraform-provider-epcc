terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source  = "elasticpath.com/elasticpath/epcc"
    }
  }
}

data "epcc_pricebook" "example" {
  id = "90c5733a-6db5-4b42-9c7d-48f689bca3b8"
}

output "pricebook_name" {
  value = data.epcc_pricebook.example.name
}

