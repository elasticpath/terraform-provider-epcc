terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source  = "elasticpath.com/elasticpath/epcc"
    }
  }
}

data "epcc_hierarchy" "example_hierarchy" {
  id = "d9bde721-e30b-489d-bb35-1fca6aa0810e"
}

data "epcc_node" "example_node" {
  id           = "d056d40e-86d0-4081-a23d-5d190531d9a1"
  hierarchy_id = data.epcc_hierarchy.example_hierarchy.id
}

output "node_name" {
  value = data.epcc_node.example_node.name
}
