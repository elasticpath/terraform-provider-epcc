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

- **code** (String) The currency code.

### Read-only

- **decimal_places** (Number) The amount of decimal places the currency is formatted to.
- **decimal_point** (String) The decimal point character.
- **default** (Boolean) Whether this is the default currency in the store.
- **enabled** (Boolean) Is this currency available for products? `true` or `false`
- **exchange_rate** (Number) The exchange rate.
- **format** (String) How to structure a currency; e.g., `${price}`.
- **id** (String) The unique identifier for this currency.
- **thousand_separator** (String) The thousand separator character.


