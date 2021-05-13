---
page_title: "epcc_product Resource - epcc-terraform-provider"
subcategory: ""
description: |-
  Allows the caller to create, update or delete an Elastic Path Commerce Cloud PCM product https://documentation.elasticpath.com/commerce-cloud/docs/concepts/products-pcm.html.
---

# Resource `epcc_product`

Allows the caller to create, update or delete an Elastic Path Commerce Cloud PCM [product](https://documentation.elasticpath.com/commerce-cloud/docs/concepts/products-pcm.html).



## Schema

### Required

- **commodity_type** (String) Valid values: `physical` or `digital`.
- **name** (String) The product name to display to customers.
- **sku** (String) The unique _stock keeping unit_ of the product.

### Optional

- **description** (String) The product description to display to customers.
- **mpn** (String) The _manufacturer part number_ of the product.
- **slug** (String) The unique slug of the product.
- **status** (String) Valid values: `draft` or `live`. Default is `draft`.
- **upc_ean** (String) The _universal product code_ or _european article number_ of the product.

### Read-only

- **id** (String) The unique identifier of the product.


