---
page_title: "epcc_promotion Data Source - epcc-terraform-provider"
subcategory: ""
description: |-
  
---

# Data Source `epcc_promotion`





## Schema

### Required

- **id** (String) The ID of this resource.

### Optional

- **automatic** (Boolean)
- **max_discount_value** (Block List) (see [below for nested schema](#nestedblock--max_discount_value))
- **min_cart_value** (Block List) (see [below for nested schema](#nestedblock--min_cart_value))

### Read-only

- **description** (String)
- **enabled** (Boolean)
- **end** (String)
- **name** (String)
- **promotion_type** (String)
- **schema** (List of Object) (see [below for nested schema](#nestedatt--schema))
- **start** (String)
- **type** (String)

<a id="nestedblock--max_discount_value"></a>
### Nested Schema for `max_discount_value`

Optional:

- **amount** (Number)
- **promotion** (String)


<a id="nestedblock--min_cart_value"></a>
### Nested Schema for `min_cart_value`

Optional:

- **amount** (Number)
- **promotion** (String)


<a id="nestedatt--schema"></a>
### Nested Schema for `schema`

Read-only:


