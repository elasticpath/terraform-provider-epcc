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
  url  = "https://sqs.us-east-2.amazonaws.com/123456789012/MyQueue"
  observes = [
    "cart.updated",
    "cart.deleted",
  ]
  integration_type      = "aws_sqs"
  region                = "us-east-2"
  aws_access_key_id     = "foofoofoofoofoo"
  aws_secret_access_key = "barbarbarbarbar"
  enabled               = true
}
