terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source  = "elasticpath.com/elasticpath/epcc"
    }
  }
}

resource "epcc_file" "my_image_file" {
  file_name = "ep.png"
  file_hash = filemd5("ep.png")
  public    = true
}

resource "epcc_file" "my_text_file" {
  file_name = "hello_world.txt"
  file_hash = filemd5("hello_world.txt")
  public    = false
}

resource "epcc_file" "my_binary_file" {
  file_name = "binary_data.bin"
  file_hash = filemd5("binary_data.bin")
  public    = true
}

output "my_image_file_link" {
  value = epcc_file.my_image_file.file_link
}
output "my_text_file_link" {
  value = epcc_file.my_text_file.file_link
}
output "my_binary_file_link" {
  value = epcc_file.my_binary_file.file_link
}