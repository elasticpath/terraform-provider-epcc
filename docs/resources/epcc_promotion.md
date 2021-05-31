---
page_title: "epcc_promotion Resource - epcc-terraform-provider"
subcategory: ""
description: |-
  Represents the EPCC API Promotion Object https://documentation.elasticpath.com/commerce-cloud/docs/api/carts-and-checkout/promotions/index.html#the-promotion-object.
---

# Resource `epcc_promotion`

Represents the EPCC API [Promotion Object](https://documentation.elasticpath.com/commerce-cloud/docs/api/carts-and-checkout/promotions/index.html#the-promotion-object).



## Schema

### Required

- **description** (String)
- **enabled** (Boolean)
- **end** (String)
- **name** (String)
- **promotion_type** (String)
- **schema** (Block List, Min: 1) (see [below for nested schema](#nestedblock--schema))
- **start** (String)
- **type** (String)

### Optional

- **automatic** (Boolean)
- **max_discount_value** (Block List) (see [below for nested schema](#nestedblock--max_discount_value))
- **min_cart_value** (Block List) (see [below for nested schema](#nestedblock--min_cart_value))

### Read-only

- **id** (String) The ID of this resource.

<a id="nestedblock--schema"></a>
### Nested Schema for `schema`

Optional:

- **currencies** (Block List) (see [below for nested schema](#nestedblock--schema--currencies))

<a id="nestedblock--schema--currencies"></a>
### Nested Schema for `schema.currencies`

Optional:

- **amount** (Number)
- **currency** (String)



<a id="nestedblock--max_discount_value"></a>
### Nested Schema for `max_discount_value`

Optional:

- **amount** (Number)
- **currency** (String)


<a id="nestedblock--min_cart_value"></a>
### Nested Schema for `min_cart_value`

Optional:

- **amount** (Number)
- **currency** (String)


