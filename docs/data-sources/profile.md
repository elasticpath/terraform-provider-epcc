---
page_title: "epcc_profile Data Source - terraform-provider-epcc"
subcategory: ""
description: |- Represents the EPCC
API [OpenID Connect Profiles] https://documentation.elasticpath.com/commerce-cloud/docs/api/single-sign-on/oidc-profiles/index.html.
---

# Data Source `epcc_profile`

Represents the EPCC
API [OpenID Connect Profiles](https://documentation.elasticpath.com/commerce-cloud/docs/api/single-sign-on/oidc-profiles/index.html).

## Schema

### Required

- **id** (String) The ID of this resource.
- **realm_id** (String) The ID of the parent realm.

### Read-only

- **discovery_url** (String)
- **client_id** (String)
- **client_secret** (String)


