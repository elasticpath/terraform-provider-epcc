terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source  = "elasticpath.com/elasticpath/epcc"
    }
  }
}

resource "epcc_product" "my_first_terraform_physical_product" {
  sku            = "tf-product-1"
  name           = "TFProduct1Physical"
  commodity_type = "physical"
  description    = "Terraform Physical Product 1"
  mpn            = "mfg-part-1"
  status         = "live"
}

resource "epcc_hierarchy" "my_first_hierarchy" {
  name        = "Hierarchy"
  description = "Foo"
  slug        = "test"
}

resource "epcc_node" "my_first_node" {
  name         = "Node #2"
  description  = "Node"
  slug         = "node-3"
  hierarchy_id = epcc_hierarchy.my_first_hierarchy.id
  products     = [epcc_product.my_first_terraform_physical_product.id]
}

resource "epcc_pricebook" "my_first_terraform_pricebook" {
  name        = "TFPricebook1"
  description = "Terraform 1"
}

resource "epcc_catalog" "my_first_catalog" {
  name        = "My Second Catalog"
  description = "Catalog created  II by Terraform"
  hierarchies = [epcc_hierarchy.my_first_hierarchy.id]
  pricebook   = epcc_pricebook.my_first_terraform_pricebook.id
}