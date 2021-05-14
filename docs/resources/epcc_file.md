---
page_title: "epcc_file Resource - epcc-terraform-provider"
subcategory: ""
description: |-
  
---

# Resource `epcc_file`





## Schema

### Optional

- **file** (String)
- **file_location** (String)
- **file_name** (String)
- **public** (Boolean) TBD - But remember the behaviour of this differs from the API, in that terraform ignores this setting if you specify it for file_location.

### Read-only

- **file_link** (String)
- **file_size** (Number)
- **id** (String) The ID of this resource.
- **mime_type** (String)


