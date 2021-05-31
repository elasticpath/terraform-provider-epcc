terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source  = "elasticpath.com/elasticpath/epcc"
    }
  }
}

data "epcc_entry" "order_flow_entry" {
  id   = "8cd0a8ef-a5c4-49bb-8012-186fad3f7917"
  slug = "orders"
}

//noinspection HILUnresolvedReference
output "shipping" {
  value = data.epcc_entry.order_flow_entry.strings.shipping
}
