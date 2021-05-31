---
page_title: "epcc_file Data Source - epcc-terraform-provider"
subcategory: ""
description: |-
  Represents the EPCC API File Object https://documentation.elasticpath.com/commerce-cloud/docs/api/advanced/files/index.html#the-file-object.
---

# Data Source `epcc_file`

Represents the EPCC API [File Object](https://documentation.elasticpath.com/commerce-cloud/docs/api/advanced/files/index.html#the-file-object).



## Schema

### Required

- **id** (String) The ID of this resource.

### Optional

- **file_name** (String)
- **public** (Boolean)

### Read-only

- **file_link** (String)
- **file_size** (Number)
- **mime_type** (String)

