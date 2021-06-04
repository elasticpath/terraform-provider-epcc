---
page_title: "epcc_payment_gateway Data Source - terraform-provider-epcc"
subcategory: ""
description: |-
  Payment gateway connectivity configuration
---

# Data Source `epcc_payment_gateway`

Payment gateway connectivity configuration



## Schema

### Required

- **slug** (String)

### Optional

- **id** (String) The ID of this resource.
- **test** (Boolean) Is this a sandbox environment. Default: `false`

### Read-only

- **enabled** (Boolean) Should the gateway process payments. Default: `false`


