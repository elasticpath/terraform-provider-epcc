---
page_title: "epcc_file Resource - terraform-provider-epcc"
subcategory: ""
description: |-
  Represents the EPCC API File Object https://documentation.elasticpath.com/commerce-cloud/docs/api/advanced/files/index.html#the-file-object.
---

# Resource `epcc_file`

Represents the EPCC API [File Object](https://documentation.elasticpath.com/commerce-cloud/docs/api/advanced/files/index.html#the-file-object).



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **file_hash** (String) A hash of the file contents
- **file_location** (String) The URL that points to an image file.
- **file_name** (String) The name of the file
- **public** (Boolean) Whether the file is public.

### Read-Only

- **file_link** (String) A link to the file
- **file_size** (Number) The size of the file
- **id** (String) The unique identifier for the file.
- **mime_type** (String) The MIME type of the file

