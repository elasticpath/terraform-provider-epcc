---
page_title: "epcc_profile Resource - terraform-provider-epcc"
subcategory: ""
description: |- Represents the EPCC
API [OpenID Connect Profiles] https://documentation.elasticpath.com/commerce-cloud/docs/api/single-sign-on/oidc-profiles/index.html.
---

# Resource `epcc_profile`

Represents the EPCC
API [OpenID Connect Profiles](https://documentation.elasticpath.com/commerce-cloud/docs/api/single-sign-on/oidc-profiles/index.html).

## Schema

### Required

- **name** (String)
- **discovery_url** (String)
- **client_id** (String)
- **client_secret** (String)

### Read-only

- **id** (String) The ID of this resource.
- **realm_id** (String) The ID of the parent realm.


