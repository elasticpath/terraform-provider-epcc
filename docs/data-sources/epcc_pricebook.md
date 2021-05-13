---
page_title: "epcc_pricebook Data Source - epcc-terraform-provider"
subcategory: ""
description: |-
  
---

# Data Source `epcc_pricebook`

Allows the caller to look up details of an Elastic Path Commerce Cloud [price book](https://documentation.elasticpath.com/commerce-cloud/docs/concepts/price-books.html).

## Schema

### Required

- **id** (String) The unique identifier of the price book.

### Read-only

- **name** (String) The name of the price book. This value will always be available as it is required in Elastic Path Commerce Cloud price books.
- **description** (String) The purpose for the price book, such as flash sale pricing or preferred customer pricing. This value may not be set as it is optional in Elastic Path Commerce Cloud price books.
