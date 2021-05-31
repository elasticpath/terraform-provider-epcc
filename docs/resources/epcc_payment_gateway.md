---
page_title: "epcc_payment_gateway Resource - epcc-terraform-provider"
subcategory: ""
description: |-
  Payment gateway connectivity configuration
---

# Resource `epcc_payment_gateway`

Payment gateway connectivity configuration



## Schema

### Required

- **slug** (String)

### Optional

- **enabled** (Boolean) Should the gateway process payments. Default: `false`
- **id** (String) The ID of this resource.
- **options** (Map of String) Parameters specific to concrete payment provider
- **test** (Boolean) Is this a sandbox environment. Default: `false`


