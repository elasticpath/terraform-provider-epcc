terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source  = "elasticpath.com/elasticpath/epcc"
    }
  }
}

data "epcc_flow" "example" {
  id = "6ef6abb7-0dcc-401e-9067-e7d052be7a63"
}

output "flow_name" {
  value = data.epcc_flow.example.name
}

