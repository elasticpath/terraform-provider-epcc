terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source  = "elasticpath.com/elasticpath/epcc"
    }
  }
}

data "epcc_flow" "orders" {
  id = "79b9cf49-dba6-4575-91ff-d793b01f126e"
}

resource "epcc_field" "shipping" {
  name        = "Shipping"
  slug        = "shipping"
  field_type  = "string"
  description = "Shipping option"
  required    = false
  enabled     = true
  flow_id     = data.epcc_flow.orders.id
}

resource "epcc_entry" "shipping_record" {
  slug      = "orders"
  target_id = "8cd0a8ef-a5c4-49bb-8012-186fad3f7917"
  strings = {
    (epcc_field.shipping.slug) = "shipping-id",
  }
}
