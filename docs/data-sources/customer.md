---
page_title: "epcc_customer Data Source - terraform-provider-epcc"
subcategory: ""
description: |-
  Represents the EPCC API Customer Object https://documentation.elasticpath.com/commerce-cloud/docs/api/orders-and-customers/customers/index.html#the-customer-object.
---

# Data Source `epcc_customer`

Represents the EPCC API [Customer Object](https://documentation.elasticpath.com/commerce-cloud/docs/api/orders-and-customers/customers/index.html#the-customer-object).



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) The unique identifier for this customer.

### Read-Only

- **email** (String) The `email` of the customer.
- **name** (String) The `name` of the customer.

