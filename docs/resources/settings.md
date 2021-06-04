---
page_title: "epcc_settings Resource - terraform-provider-epcc"
subcategory: ""
description: |-
  Represents the EPCC API Settings https://documentation.elasticpath.com/commerce-cloud/docs/api/advanced/settings/index.html.
  Note: The epcc_settings resource behaves different from normal resources, in that Terraform does not create this reosurce, but instead "adopts" it into management.
---

# Resource `epcc_settings`

Represents the EPCC API [Settings](https://documentation.elasticpath.com/commerce-cloud/docs/api/advanced/settings/index.html).
Note: The `epcc_settings` resource behaves different from normal resources, in that Terraform does not *create* this reosurce, but instead "adopts" it into management.



## Schema

### Optional

- **additional_languages** (List of String)
- **calculation_method** (String)
- **list_child_products** (Boolean)
- **page_length** (Number)

### Read-only

- **id** (String) The ID of this resource.


