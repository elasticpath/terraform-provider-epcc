package provider

import (
	"context"
	"github.com/elasticpath/epcc-terraform-provider/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type IntegrationDataSourceProvider struct {
}

func (ds IntegrationDataSourceProvider) DataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Allows to configure webhooks",
		ReadContext: ds.read,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enabled": {
				Type:        schema.TypeBool,
				Description: "Should the event trigger or not. Default: `false`",
				Computed:    true,
			},
			"url": {
				Type:        schema.TypeString,
				Description: "Webhook endpoint",
				Computed:    true,
			},
			"secret_key": {
				Type:        schema.TypeString,
				Description: "Value that is passed to webhook as `X-Moltin-Secret-Key` header",
				Computed:    true,
			},
			"observes": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "[observable event type](https://documentation.elasticpath.com/commerce-cloud/docs/api/advanced/events/create-an-event.html)",
				Computed:    true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func (ds IntegrationDataSourceProvider) read(_ context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	result, err := epcc.Integrations.Get(client, data.Get("id").(string))
	if err != nil {
		return FromAPIError(err)
	}

	if err := data.Set("name", result.Data.Name); err != nil {
		return diag.FromErr(err)
	}

	if err := data.Set("description", result.Data.Description); err != nil {
		return diag.FromErr(err)
	}

	if err := data.Set("enabled", result.Data.Enabled); err != nil {
		return diag.FromErr(err)
	}

	if err := data.Set("url", result.Data.Configuration.Url); err != nil {
		return diag.FromErr(err)
	}

	if err := data.Set("secret_key", result.Data.Configuration.SecretKey); err != nil {
		return diag.FromErr(err)
	}

	if err := data.Set("observes", result.Data.Observes); err != nil {
		return diag.FromErr(err)
	}

	data.SetId(result.Data.Id)

	return nil
}
