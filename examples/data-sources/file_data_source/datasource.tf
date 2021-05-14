terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source  = "elasticpath.com/elasticpath/epcc"
    }
  }
}

data "epcc_file" "simple_file" {
  id = "62f6f9d4-4d14-4ff1-9af0-ff23da418e5f"
}

output "file_size" {
  value = data.epcc_file.simple_file.file_size
}

output "is_public" {
  value = data.epcc_file.simple_file.public
}

output "mime_type" {
  value = data.epcc_file.simple_file.mime_type
}

output "file_link" {
  value = data.epcc_file.simple_file.file_link
}