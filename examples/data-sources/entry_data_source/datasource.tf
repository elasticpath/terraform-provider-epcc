terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source  = "elasticpath.com/elasticpath/epcc"
    }
  }
}

data "epcc_entry" "example" {
  id   = "4f82aa4f-5ee8-478a-99ca-4a27f888d8ca"
  slug = "place"
}

output "entry_name" {
  value = data.epcc_entry.example
}

