---
page_title: "epcc_flow Data Source - terraform-provider-epcc"
subcategory: ""
description: |-
  Represents the EPCC API Flow Object https://documentation.elasticpath.com/commerce-cloud/docs/api/advanced/custom-data/flows/index.html#the-flow-object.
---

# Data Source `epcc_flow`

Represents the EPCC API [Flow Object](https://documentation.elasticpath.com/commerce-cloud/docs/api/advanced/custom-data/flows/index.html#the-flow-object).



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) The unique identifier for this flow.

### Read-Only

- **description** (String) Any description for this flow.
- **enabled** (Boolean) true if enabled, false if not.
- **name** (String) The name of the flow.
- **slug** (String) A unique slug identifier for the flow.

