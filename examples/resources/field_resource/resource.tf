terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source  = "elasticpath.com/elasticpath/epcc"
    }
  }
}

resource "epcc_flow" "sports_flow" {
  name        = "Flow for sports"
  slug        = "hockey"
  description = "This is a Terraform test"
  enabled     = true
}

resource "epcc_field" "sports_season_field" {
  name        = "Sport season"
  slug        = "season"
  field_type  = "string"
  description = "Season the sport is played in"
  required    = false
  default     = "summer"
  omit_null   = false
  enabled     = true
  flow_id     = epcc_flow.sports_flow.id
  valid_string_format = "slug"
  valid_string_enum = ["spring", "summer", "winter", "autumn"]
}
