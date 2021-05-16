---
page_title: "epcc_node Data Source - terraform-provider-epcc"
subcategory: ""
description: |-
  Represents the EPCC API Node Object https://documentation.elasticpath.com/commerce-cloud/docs/api/pcm/hierarchies/index.html#the-node-object.
---

# Data Source `epcc_node`

Represents the EPCC API [Node Object](https://documentation.elasticpath.com/commerce-cloud/docs/api/pcm/hierarchies/index.html#the-node-object).



## Schema

### Required

- **hierarchy_id** (String)
- **id** (String) The ID of this resource.

### Optional

- **parent_id** (String)
- **products** (Set of String)

### Read-only

- **description** (String)
- **name** (String)
- **slug** (String)


