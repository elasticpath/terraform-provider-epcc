terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source  = "elasticpath.com/elasticpath/epcc"
    }
  }
}

resource "epcc_hierarchy" "my_first_hierarchy" {
  name  = "Hierarchy"
  description = "Foo"
  slug = "test"
}