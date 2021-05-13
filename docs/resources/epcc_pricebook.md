---
page_title: "epcc_pricebook Resource - epcc-terraform-provider"
subcategory: ""
description: |-
  
---

# Resource `epcc_pricebook`

Allows the caller to create, modify or delete an Elastic Path Commerce Cloud [price book](https://documentation.elasticpath.com/commerce-cloud/docs/concepts/price-books.html).

## Schema

### Required

- **name** (String) A unique name for the price book. This is required when creating a new Elastic Path Commerce Cloud price book.

### Optional

- **description** (String) The purpose for the price book, such as flash sale pricing or preferred customer pricing.

### Read-only

- **id** (String) The unique identifier of the price book. This will be set for you.
