---
page_title: "epcc_settings Data Source - terraform-provider-epcc"
subcategory: ""
description: |-
  Represents the EPCC API Settings https://documentation.elasticpath.com/commerce-cloud/docs/api/advanced/settings/index.html.
  Note: The epcc_settings resource behaves different from normal resources, in that Terraform does not create this reosurce, but instead "adopts" it into management.
---

# Data Source `epcc_settings`

Represents the EPCC API [Settings](https://documentation.elasticpath.com/commerce-cloud/docs/api/advanced/settings/index.html).
Note: The `epcc_settings` resource behaves different from normal resources, in that Terraform does not *create* this reosurce, but instead "adopts" it into management.



## Schema

### Optional

- **id** (String) The ID of this resource.

### Read-only

- **additional_languages** (List of String)
- **calculation_method** (String)
- **list_child_products** (Boolean)
- **page_length** (Number)


