terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source  = "elasticpath.com/elasticpath/epcc"
    }
  }
}

data "epcc_integration" "example" {
  id = "37cecd4a-dbe1-4eef-bc9d-ad6c392ef9bb"
}

output "integration_name" {
  value = data.epcc_integration.example.name
}

