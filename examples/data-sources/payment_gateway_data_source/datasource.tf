terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source  = "elasticpath.com/elasticpath/epcc"
    }
  }
}

data "epcc_payment_gateway" "manual" {
  slug = "manual"
}

data "epcc_payment_gateway" "stripe_payment_intents" {
  slug = "stripe_payment_intents"
}

data "epcc_payment_gateway" "braintree" {
  slug = "braintree"
}
