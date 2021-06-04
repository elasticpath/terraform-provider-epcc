---
page_title: "epcc_integration Resource - terraform-provider-epcc"
subcategory: ""
description: |-
  Allows to configure webhooks, and corresponds to EPCC API Event (Webhooks) Object https://documentation.elasticpath.com/commerce-cloud/docs/api/advanced/events/index.html#event-object
---

# Resource `epcc_integration`

Allows to configure webhooks, and corresponds to EPCC API [Event (Webhooks) Object](https://documentation.elasticpath.com/commerce-cloud/docs/api/advanced/events/index.html#event-object)



## Schema

### Required

- **name** (String)
- **url** (String) Webhook endpoint

### Optional

- **description** (String)
- **enabled** (Boolean) Should the event trigger or not. Default: `false`
- **observes** (List of String) [observable event type](https://documentation.elasticpath.com/commerce-cloud/docs/api/advanced/events/create-an-event.html)
- **secret_key** (String) Value that is passed to webhook as `X-Moltin-Secret-Key` header

### Read-only

- **id** (String) The ID of this resource.


