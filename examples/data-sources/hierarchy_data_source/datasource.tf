terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source  = "elasticpath.com/elasticpath/epcc"
    }
  }
}

data "epcc_hierarchy" "example" {
  id = "c996cd23-bca1-48e0-9e1c-1d9186bd8a24"
}

output "hierarchy_name" {
  value = data.epcc_hierarchy.example.name
}

