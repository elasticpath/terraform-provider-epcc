---
page_title: "epcc_integration Data Source - epcc-terraform-provider"
subcategory: ""
description: |-
  Allows to configure webhooks
---

# Data Source `epcc_integration`

Allows to configure webhooks



## Schema

### Required

- **id** (String) The ID of this resource.

### Read-only

- **description** (String)
- **enabled** (Boolean) Should the event trigger or not. Default: `false`
- **name** (String)
- **observes** (List of String) [observable event type](https://documentation.elasticpath.com/commerce-cloud/docs/api/advanced/events/create-an-event.html)
- **secret_key** (String) Value that is passed to webhook as `X-Moltin-Secret-Key` header
- **url** (String) Webhook endpoint


