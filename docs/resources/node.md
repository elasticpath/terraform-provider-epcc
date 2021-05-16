---
page_title: "epcc_node Resource - terraform-provider-epcc"
subcategory: ""
description: |-
  Represents the EPCC API Node Object https://documentation.elasticpath.com/commerce-cloud/docs/api/pcm/hierarchies/index.html#the-node-object.
---

# Resource `epcc_node`

Represents the EPCC API [Node Object](https://documentation.elasticpath.com/commerce-cloud/docs/api/pcm/hierarchies/index.html#the-node-object).



## Schema

### Required

- **hierarchy_id** (String)
- **name** (String)

### Optional

- **description** (String)
- **parent_id** (String)
- **products** (Set of String)
- **slug** (String)

### Read-only

- **id** (String) The ID of this resource.


