---
page_title: "epcc-terraform-provider Provider"
subcategory: ""
description: |-
  
---

# epcc-terraform-provider Provider



## Example Usage

```terraform
provider "epcc" {
  // Can set via `EPCC_CLIENT_ID` environment variable.
  client_id = "some_client_id"

  // Can set via `EPCC_CLIENT_SECRET` environment variable.
  client_secret = "some_client_secret"

  // Can set via `EPCC_API_BASE_URL` environment variable
  api_base_url = "https://api.moltin.com/"

  // Can set via `EPCC_BETA_API_FEATURES` environment variable.
  beta_features = "account-management"
}
```

## Schema

### Optional

- **additional_headers** (Map of String)
- **api_base_url** (String)
- **beta_features** (String)
- **client_id** (String)
- **client_secret** (String, Sensitive)
- **enable_authentication** (Boolean)
