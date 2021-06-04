---
page_title: "epcc_entry Resource - terraform-provider-epcc"
subcategory: ""
description: |-
  Represents the EPCC API Entry Object https://documentation.elasticpath.com/commerce-cloud/docs/api/advanced/custom-data/entries/index.html.
---

# Resource `epcc_entry`

Represents the EPCC API [Entry Object](https://documentation.elasticpath.com/commerce-cloud/docs/api/advanced/custom-data/entries/index.html).



## Schema

### Required

- **slug** (String)

### Optional

- **booleans** (Map of Boolean)
- **numbers** (Map of Number)
- **strings** (Map of String)
- **target_id** (String) Target core object identifier (can only be used for core flows)

### Read-only

- **id** (String) The ID of this resource.


