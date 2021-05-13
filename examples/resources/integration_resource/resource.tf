terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source  = "elasticpath.com/elasticpath/epcc"
    }
  }
}

resource "epcc_integration" "cart_observer" {
  name = "Cart Observer"
  url  = "http://localhost"
  observes = [
    "cart.updated",
    "cart.deleted",
  ]
}
