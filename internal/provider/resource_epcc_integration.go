package provider

import (
	"context"
	"github.com/elasticpath/epcc-terraform-provider/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type IntegrationResourceProvider struct {
}

func (r IntegrationResourceProvider) Resource() *schema.Resource {
	return &schema.Resource{
		Description:   "Allows to configure webhooks",
		CreateContext: r.create,
		ReadContext:   r.read,
		UpdateContext: r.update,
		DeleteContext: r.delete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enabled": {
				Type:        schema.TypeBool,
				Description: "Should the event trigger or not. Default: `false`",
				Optional:    true,
			},
			"url": {
				Type:        schema.TypeString,
				Description: "Webhook endpoint",
				Required:    true,
			},
			"secret_key": {
				Type:        schema.TypeString,
				Description: "Value that is passed to webhook as `X-Moltin-Secret-Key` header",
				Optional:    true,
			},
			"observes": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "[observable event type](https://documentation.elasticpath.com/commerce-cloud/docs/api/advanced/events/create-an-event.html)",
				Optional:    true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func (r IntegrationResourceProvider) create(_ context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	observes := data.Get("observes").([]interface{})
	integrationObject := &epcc.Integration{
		Type:            epcc.IntegrationType,
		IntegrationType: epcc.Webhook,
		Name:            data.Get("name").(string),
		Description:     data.Get("description").(string),
		Enabled:         data.Get("enabled").(bool),
		Configuration: epcc.IntegrationConfiguration{
			Url:       data.Get("url").(string),
			SecretKey: data.Get("secret_key").(string),
		},
		Observes: convertToStringSlice(observes),
	}

	result, err := epcc.Integrations.Create(client, integrationObject)
	if err != nil {
		return FromAPIError(err)
	}

	data.SetId(result.Data.Id)

	return nil
}

func (r IntegrationResourceProvider) delete(_ context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	err := epcc.Integrations.Delete(client, data.Id())
	if err != nil {
		return FromAPIError(err)
	}

	data.SetId("")

	return nil
}

func (r IntegrationResourceProvider) update(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	integrationObject := &epcc.Integration{
		Type:        epcc.IntegrationType,
		Name:        data.Get("name").(string),
		Description: data.Get("description").(string),
		Enabled:     data.Get("enabled").(bool),
		Configuration: epcc.IntegrationConfiguration{
			Url:       data.Get("url").(string),
			SecretKey: data.Get("secret_key").(string),
		},
		Observes: data.Get("observes").([]string),
	}

	result, apiError := epcc.Integrations.Update(client, data.Id(), integrationObject)

	if apiError != nil {
		return FromAPIError(apiError)
	}

	data.SetId(result.Data.Id)

	return nil
}

func (r IntegrationResourceProvider) read(_ context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	result, err := epcc.Integrations.Get(client, data.Id())
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

	return nil
}

func convertToStringSlice(arr []interface{}) []string {
	var result []string
	for _, param := range arr {
		result = append(result, param.(string))
	}
	return result
}
