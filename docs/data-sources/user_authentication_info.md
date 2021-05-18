---
page_title: "epcc_user_authentication_info Data Source - terraform-provider-epcc"
subcategory: ""
description: |- Represents the EPCC
API OpenID Connect User Authentication Info
---

# Data Source `epcc_user_authentication_info`

Represents the EPCC
API OpenID Connect User Authentication Info

## Schema

### Required

- **id** (String) The ID of this resource.
- **realm_id** (String) The ID of the parent realm.

### Read-only

- **name** (String)
- **email** (String)