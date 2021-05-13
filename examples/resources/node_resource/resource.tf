terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source  = "elasticpath.com/elasticpath/epcc"
    }
  }
}

resource "epcc_hierarchy" "my_first_hierarchy" {
  name  = "Node Hierarchy"
  description = "Foo"
  slug = "test-node-hierarchy"
}

resource "epcc_node" "my_parent_node" {
  name = "Node #2"
  description = "Node"
  slug = "node-3"
  hierarchy_id = epcc_hierarchy.my_first_hierarchy.id
}

// TODO Docs say node name is unique amoung siblings
// But the API errors if the name is the same between this and parent node.
resource "epcc_node" "my_child_node" {
  name = "Node #3"
  description = "Node"
  slug = "node-4"
  hierarchy_id = epcc_hierarchy.my_first_hierarchy.id
  parent_id = epcc_node.my_parent_node.id
}