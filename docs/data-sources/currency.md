---
page_title: "epcc_currency Data Source - terraform-provider-epcc"
subcategory: ""
description: |-
  Represents the EPCC API Currency Object https://documentation.elasticpath.com/commerce-cloud/docs/api/advanced/currencies/index.html#the-currency-object.
---

# Data Source `epcc_currency`

Represents the EPCC API [Currency Object](https://documentation.elasticpath.com/commerce-cloud/docs/api/advanced/currencies/index.html#the-currency-object).



## Schema

### Required

- **code** (String) Currency 3-letter unique code

### Read-only

- **id** (String) The ID of this resource.
- **decimal_places** (Number)
- **decimal_point** (String)
- **default** (Boolean)
- **enabled** (Boolean)
- **exchange_rate** (Number)
- **format** (String)
- **thousand_separator** (String)


