terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source  = "elasticpath.com/elasticpath/epcc"
    }
  }
}

resource "epcc_flow" "tourism_flow" {
  name        = "Flow for tourism"
  slug        = "tourism"
  description = "This is a Terraform test"
  enabled     = true
}

resource "epcc_field" "tourism_season_field" {
  name        = "tourism season"
  slug        = "season"
  field_type  = "string"
  description = "Season for travelling"
  required    = false
  default     = "summer"
  omit_null   = false
  enabled     = true
  flow_id     = epcc_flow.tourism_flow.id
}

resource "epcc_field" "tourism_place_field" {
  name        = "tourism place"
  slug        = "place"
  field_type  = "string"
  description = "place for travelling"
  required    = false
  default     = "vancouver"
  omit_null   = false
  enabled     = true
  flow_id     = epcc_flow.tourism_flow.id
}

resource "epcc_entry" "tourism_netherlands" {
  slug = epcc_flow.tourism_flow.slug
  payload = {
    (epcc_field.tourism_season_field.slug) = "spring",
    (epcc_field.tourism_place_field.slug)  = "netherlands"
  }
}
