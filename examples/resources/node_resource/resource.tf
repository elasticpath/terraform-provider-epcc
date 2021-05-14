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
  slug        = "test-node-hierarchy"
}

resource "epcc_node" "my_parent_node" {
  name         = "Node #2"
  description  = "Node"
  slug         = "node-3"
  hierarchy_id = epcc_hierarchy.my_first_hierarchy.id
}



resource "epcc_product" "simple_product" {
  sku = "12345567"
  name = "Product With File"
  commodity_type = "physical"
}


resource "epcc_product" "simple_product_2" {
  sku = "123455678"
  name = "Product With File"
  commodity_type = "physical"
}


// TODO Docs say node name is unique amoung siblings
// But the API errors if the name is the same between this and parent node.
resource "epcc_node" "my_child_node" {
  name         = "Node #3"
  description  = "Node"
  slug         = "node-4"
  hierarchy_id = epcc_hierarchy.my_first_hierarchy.id
  parent_id    = epcc_node.my_parent_node.id
  products = [ epcc_product.simple_product.id, epcc_product.simple_product_2.id ]
}
