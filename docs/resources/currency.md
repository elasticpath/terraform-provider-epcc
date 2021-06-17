---
page_title: "epcc_currency Resource - terraform-provider-epcc"
subcategory: ""
description: |-
  Represents the EPCC API Currency Object https://documentation.elasticpath.com/commerce-cloud/docs/api/advanced/currencies/index.html#the-currency-object.
---

# Resource `epcc_currency`

Represents the EPCC API [Currency Object](https://documentation.elasticpath.com/commerce-cloud/docs/api/advanced/currencies/index.html#the-currency-object).

!> **WARNING:** If multiple currencies are defined, please ensure that the `default` tag is set to `true` on only one of them


## Schema

### Required

- **code** (String)
- **decimal_places** (Number)
- **decimal_point** (String)
- **default** (Boolean)
- **enabled** (Boolean)
- **exchange_rate** (Number)
- **format** (String)
- **thousand_separator** (String)

### Read-only

- **id** (String) The ID of this resource.


