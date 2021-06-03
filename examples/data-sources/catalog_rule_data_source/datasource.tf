terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source  = "elasticpath.com/elasticpath/epcc"
    }
  }
}

data "epcc_catalog_rule" "example_catalog_rule" {
  id = "f74b8c23-8bdd-4e67-98b8-8c8e14c3bfdc"
}

output "catalog_rule_name" {
  value = data.epcc_catalog_rule.example_catalog_rule.name
}