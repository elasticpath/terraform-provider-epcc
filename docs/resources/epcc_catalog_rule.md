---
page_title: "epcc_catalog_rule Resource - epcc-terraform-provider"
subcategory: ""
description: |-
  Represents the EPCC API PCM Catalog Rule Object https://documentation.elasticpath.com/commerce-cloud/docs/api/pcm/catalogs/rules/get-a-catalog-rule.html.
---

# Resource `epcc_catalog_rule`

Represents the EPCC API [*PCM* Catalog Rule Object](https://documentation.elasticpath.com/commerce-cloud/docs/api/pcm/catalogs/rules/get-a-catalog-rule.html).



## Schema

### Required

- **catalog** (String)
- **name** (String)

### Optional

- **customers** (Set of String)
- **description** (String)

### Read-only

- **id** (String) The ID of this resource.


