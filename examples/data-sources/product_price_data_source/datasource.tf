terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source  = "elasticpath.com/elasticpath/epcc"
    }
  }
}

data "epcc_product_price" "price" {
  id           = "634c1afd-db6a-473c-ad39-7d06bd274659"
  pricebook_id = "667ea035-688e-4387-87b1-303cf8fd1ca0"
}


output "prices" {
  value = data.epcc_product_price.price.currency[*]
}
