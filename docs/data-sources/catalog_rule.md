---
page_title: "epcc_catalog_rule Data Source - terraform-provider-epcc"
subcategory: ""
description: |-
  Represents the EPCC API PCM Catalog Rule Object https://documentation.elasticpath.com/commerce-cloud/docs/api/pcm/catalogs/rules/get-a-catalog-rule.html.
---

# Data Source `epcc_catalog_rule`

Represents the EPCC API [*PCM* Catalog Rule Object](https://documentation.elasticpath.com/commerce-cloud/docs/api/pcm/catalogs/rules/get-a-catalog-rule.html).



## Schema

### Required

- **id** (String) The ID of this resource.

### Read-only

- **catalog** (String) The unique identifier of the catalog for this rule. If you want to display a catalog that contains V2 Products, Brands, Categories, and Collections, specify `legacy`
- **customers** (Set of String) The list of customers who are eligible to see this catalog. If empty, the rule matches all customers.
- **description** (String) The purpose for this rule.
- **name** (String) The name of the rule without spaces.


