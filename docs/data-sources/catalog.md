---
page_title: "epcc_catalog Data Source - terraform-provider-epcc"
subcategory: ""
description: |-
  Represents the EPCC API PCM Catalog Object https://documentation.elasticpath.com/commerce-cloud/docs/api/pcm/catalogs/index.html#the-catalog-object.
---

# Data Source `epcc_catalog`

Represents the EPCC API [*PCM* Catalog Object](https://documentation.elasticpath.com/commerce-cloud/docs/api/pcm/catalogs/index.html#the-catalog-object).



## Schema

### Required

- **id** (String) The ID of this resource.

### Read-only

- **description** (String) A description of the catalog, such as the purpose for the catalog.
- **hierarchies** (Set of String) The unique identifiers of the hierarchies to associate with this catalog.
- **name** (String) The name of the catalog.
- **pricebook** (String) The unique identifier of the price book to associate with this catalog.


