---
page_title: "epcc_field Resource - terraform-provider-epcc"
subcategory: ""
description: |-
  Represents the EPCC API Fields Object https://documentation.elasticpath.com/commerce-cloud/docs/api/advanced/custom-data/fields/index.html.
---

# Resource `epcc_field`

Represents the EPCC API [Fields Object](https://documentation.elasticpath.com/commerce-cloud/docs/api/advanced/custom-data/fields/index.html).



## Schema

### Required

- **description** (String)
- **enabled** (Boolean)
- **field_type** (String)
- **flow_id** (String)
- **name** (String)
- **required** (Boolean)
- **slug** (String)

### Optional

- **default** (String)
- **omit_null** (Boolean)
- **order** (Number)
- **relationship_to_many** (String)
- **relationship_to_one** (String)
- **valid_float_enum** (List of Number)
- **valid_float_range** (List of Map of Number)
- **valid_int_enum** (List of Number)
- **valid_int_range** (List of Map of Number)
- **valid_string_enum** (List of String)
- **valid_string_format** (String)

### Read-only

- **id** (String) The ID of this resource.


