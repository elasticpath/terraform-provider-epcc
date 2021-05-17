terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source  = "elasticpath.com/elasticpath/epcc"
    }
  }
}

resource "epcc_hierarchy" "my_first_hierarchy" {
  name        = "Node Hierarchy"
  description = "Foo"
  slug        = "test-node-hierarchy-7"
}

resource "epcc_node" "my_parent_node" {
  name         = "Node #2"
  description  = "Node"
  slug         = "node-6"
  hierarchy_id = epcc_hierarchy.my_first_hierarchy.id
}

resource "epcc_product" "simple_product" {
  sku            = "457"
  name           = "Product With File"
  commodity_type = "physical"
}

resource "epcc_node_product" "simple_product_node_relationship" {
  node_id      = epcc_node.my_parent_node.id
  hierarchy_id = epcc_hierarchy.my_first_hierarchy.id
  product_id   = epcc_product.simple_product.id
}

data "epcc_node_product" "data_simple_product_node_relationship" {
  id           = epcc_node_product.simple_product_node_relationship.id
  node_id      = epcc_node.my_parent_node.id
  hierarchy_id = epcc_hierarchy.my_first_hierarchy.id
  product_id   = epcc_product.simple_product.id
}