terraform {
  required_providers {
    epcc = {
      version = "0.0.1"
      source  = "elasticpath.com/elasticpath/epcc"
    }
  }
}

resource "epcc_payment_gateway" "manual" {
  slug    = "manual"
  enabled = true
}

resource "epcc_payment_gateway" "braintree" {
  slug    = "braintree"
  enabled = true
  options = {
    merchant_id = "test_merchant_id"
    private_key = "test_private_key"
    public_key  = "test_public_key"
    environment = "sandbox"
  }
}
