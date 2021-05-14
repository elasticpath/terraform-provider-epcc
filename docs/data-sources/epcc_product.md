---
page_title: "epcc_product Data Source - epcc-terraform-provider"
subcategory: ""
description: |-
  Allows the caller to look up details of an Elastic Path Commerce Cloud PCM product https://documentation.elasticpath.com/commerce-cloud/docs/concepts/products-pcm.html.
---

# Data Source `epcc_product`

Allows the caller to look up details of an Elastic Path Commerce Cloud PCM [product](https://documentation.elasticpath.com/commerce-cloud/docs/concepts/products-pcm.html).



## Schema

### Required

- **id** (String) The unique identifier of the product.

### Read-only

- **commodity_type** (String) The type of the product; either `physical` or `digital`.
- **description** (String) The product description to display to customers.
- **mpn** (String) The _manufacturer part number_ of the product.
- **name** (String) The product name to display to customers.
- **sku** (String) The unique _stock keeping unit_ of the product.
- **slug** (String) The unique slug of the product.
- **status** (String) The status of the product; either `draft` or `live`. Default is `draft`.
- **upc_ean** (String) The _universal product code_ or _european article number_ of the product.


