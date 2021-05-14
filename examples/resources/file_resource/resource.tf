terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source  = "elasticpath.com/elasticpath/epcc"
    }
  }
}

resource "epcc_file" "my_image_file" {
  file      = filebase64("ep.png")
  file_name = "ep.png"
  public    = true
}

resource "epcc_file" "my_text_file" {
  file      = base64encode("Hello World!")
  file_name = "hello-world.txt"
  public    = false
}

resource "epcc_file" "my_binary_file" {
  file      = "iVBORw0KGgoAAAANSUhEUgAAACAAAAAgAQMAAABJtOi3AAAABlBMVEX///8AAABVwtN+AAAACXBIWXMAAA7EAAAOxAGVKw4bAAAAfklEQVQImR3OsQkDMQwF0G9cpPQIGkVrpTDWdSlvpYMrUmaEOGQBQxoXwspPVDwQEvoC/qUbsYNEB9KPHAO4fEh5E7kT3TkwIU3Pjmr7QA2ZqEsd1Y1co6Elb7A8G/QxHBovh8Q5UeI2eLl0pCUHsY1tMCgWg8LJcxKuIfOdL8H1PrJhZV++AAAAAElFTkSuQmCC"
  file_name = "file.bin"
  public    = true
}

resource "epcc_file" "my_image_link" {
  file_location = "https://my.example.com/images/abc.png"
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
output "my_image_link_link" {
  value = epcc_file.my_image_link.file_link
}