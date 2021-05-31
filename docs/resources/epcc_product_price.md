---
page_title: "epcc_product_price Resource - epcc-terraform-provider"
subcategory: ""
description: |-
  Represents the EPCC API Price Object https://documentation.elasticpath.com/commerce-cloud/docs/api/pcm/pricebooks/prices/create-product-prices.html.
---

# Resource `epcc_product_price`

Represents the EPCC API [Price Object](https://documentation.elasticpath.com/commerce-cloud/docs/api/pcm/pricebooks/prices/create-product-prices.html).



## Schema

### Required

- **currency** (Block Set, Min: 1) (see [below for nested schema](#nestedblock--currency))
- **pricebook_id** (String)
- **sku** (String)

### Read-only

- **id** (String) The unique identifier of the price.

<a id="nestedblock--currency"></a>
### Nested Schema for `currency`

Required:

- **amount** (Number)
- **code** (String)

Optional:

- **includes_tax** (Boolean)


