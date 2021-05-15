terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source  = "elasticpath.com/elasticpath/epcc"
    }
  }
}

data "epcc_catalog" "example_catalog" {
  id = "2f91defa-fe27-44e0-a5bc-a6706043be93"
}

output "catalog_name" {
  value = data.epcc_catalog.example_catalog.name
}