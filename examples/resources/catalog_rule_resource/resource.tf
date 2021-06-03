terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source  = "elasticpath.com/elasticpath/epcc"
    }
  }
}

resource "epcc_hierarchy" "sneakers_hierarchy" {
  name = "Hierarchy"
}

resource "epcc_pricebook" "sneakers_prices" {
  name = "Sneakers Prices"
}

resource "epcc_catalog" "sneakers_catalog" {
  name        = "Sneakers Catalog"
  hierarchies = [epcc_hierarchy.sneakers_hierarchy.id]
  pricebook   = epcc_pricebook.sneakers_prices.id
}

resource "epcc_customer" "customer_1" {
  name  = "Banana"
  email = "banana@food.com"
}

resource "epcc_customer" "customer_2" {
  name  = "Orange"
  email = "orange@food.com"
}

resource "epcc_catalog_rule" "sneaker_catalog_rule" {
  name      = "Special running shoes"
  catalog   = epcc_catalog.sneakers_catalog.id
  customers = [epcc_customer.customer_1.id, epcc_customer.customer_2.id]
}
