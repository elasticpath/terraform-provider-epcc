---
page_title: "epcc_product_price Data Source - epcc-terraform-provider"
subcategory: ""
description: |-
  Represents the EPCC API Price Object https://documentation.elasticpath.com/commerce-cloud/docs/api/pcm/pricebooks/prices/create-product-prices.html.
---

# Data Source `epcc_product_price`

Represents the EPCC API [Price Object](https://documentation.elasticpath.com/commerce-cloud/docs/api/pcm/pricebooks/prices/create-product-prices.html).



## Schema

### Required

- **id** (String) The unique identifier of the price.
- **pricebook_id** (String)

### Read-only

- **currency** (Set of Object) (see [below for nested schema](#nestedatt--currency))
- **sku** (String)

<a id="nestedatt--currency"></a>
### Nested Schema for `currency`

Read-only:

- **amount** (Number)
- **code** (String)
- **includes_tax** (Boolean)

