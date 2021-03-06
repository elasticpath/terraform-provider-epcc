---
page_title: "epcc_authentication_realm Data Source - terraform-provider-epcc"
subcategory: ""
description: |-
  Represents the EPCC API Authentication Realms https://documentation.elasticpath.com/commerce-cloud/docs/api/single-sign-on/authentication-realms/index.html.
---

# Data Source `epcc_authentication_realm`

Represents the EPCC API [Authentication Realms](https://documentation.elasticpath.com/commerce-cloud/docs/api/single-sign-on/authentication-realms/index.html).



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **id** (String) The unique identifier for the authentication realm.

### Read-Only

- **duplicate_email_policy** (String) The values permitted for this parameter are, `allowed` or `api_only`. When an unfamiliar user signs in for the first time, a value of `allowed` always creates a new user with the name and e-mail address supplied by the identity provider. With the `api_only` value, the system assigns the user to an existing user with a matching e-mail address, if one already exists. The `api_only` setting is recommended only when all configured identity providers treat e-mail address as a unique identifier for the user, otherwise a user might get access to another user’s account and data. Thus the `api_only` value can simplify administration of users.
- **name** (String) The name of the authentication realm.
- **origin_id** (String) The ID of the origin entity.
- **origin_type** (String) The type of the origin entity.
- **redirect_uris** (List of String) An array of Storefront URIs that can start Single Sign On authentication. These URIs must follow the rules for [redirection endpoints in OAuth 2.0](https://tools.ietf.org/html/rfc6749#section-3.1.2). All URIs must start with `https://` except for `http://localhost`.

