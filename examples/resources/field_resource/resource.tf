terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source = "elasticpath.com/elasticpath/epcc"
    }
  }
}

resource "epcc_flow" "customer_flow" {
  name = "Flow for all customers"
  slug = "customers"
  description = "This is a Terraform test"
  enabled = true
}

resource "epcc_field" "customer_age_field" {
  name = "Customer Age"
  slug = "age"
  field_type = "integer"
  description = "Age is a customer resource extension"
  required = false
  default = 18
  omit_null = false
  flow_id = epcc_flow.customer_flow.id
}