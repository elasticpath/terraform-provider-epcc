terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source  = "elasticpath.com/elasticpath/epcc"
    }
  }
}

data "epcc_hierarchy" "example_hierarchy" {
  id = "8d5ba73f-f44e-429d-b427-5dddbcaa6430"
}

data "epcc_node" "example_node" {
  id           = "a609e1df-f64a-454d-b321-645d9d9a2834"
  hierarchy_id = data.epcc_hierarchy.example_hierarchy.id
}

output "node_name" {
  value = data.epcc_node.example_node.name
}

output "node_products" {
  value = data.epcc_node.example_node.products
}
