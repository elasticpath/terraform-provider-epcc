package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type IntegrationResourceProvider struct {
}

func (r IntegrationResourceProvider) Resource() *schema.Resource {
	return &schema.Resource{
		Description:   "Allows to configure webhooks, and corresponds to EPCC API [Event (Webhooks) Object](https://documentation.elasticpath.com/commerce-cloud/docs/api/advanced/events/index.html#event-object)",
		CreateContext: addDiagToContext(r.create),
		ReadContext:   addDiagToContext(r.read),
		UpdateContext: addDiagToContext(r.update),
		DeleteContext: addDiagToContext(r.delete),
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

func (r IntegrationResourceProvider) create(ctx context.Context, data *schema.ResourceData, m interface{}) {
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
		Observes: convertArrayToStringSlice(observes),
	}

	result, err := epcc.Integrations.Create(&ctx, client, integrationObject)
	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	data.SetId(result.Data.Id)
}

func (r IntegrationResourceProvider) delete(ctx context.Context, data *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	err := epcc.Integrations.Delete(&ctx, client, data.Id())
	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	data.SetId("")
}

func (r IntegrationResourceProvider) update(ctx context.Context, data *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	observes := data.Get("observes").([]interface{})
	integrationObject := &epcc.Integration{
		Type:        epcc.IntegrationType,
		Name:        data.Get("name").(string),
		Description: data.Get("description").(string),
		Enabled:     data.Get("enabled").(bool),
		Configuration: epcc.IntegrationConfiguration{
			Url:       data.Get("url").(string),
			SecretKey: data.Get("secret_key").(string),
		},
		Observes: convertArrayToStringSlice(observes),
	}

	result, apiError := epcc.Integrations.Update(&ctx, client, data.Id(), integrationObject)

	if apiError != nil {
		ReportAPIError(ctx, apiError)
		return
	}

	data.SetId(result.Data.Id)
}

func (r IntegrationResourceProvider) read(ctx context.Context, data *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	result, err := epcc.Integrations.Get(&ctx, client, data.Id())
	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	if err := data.Set("name", result.Data.Name); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	if err := data.Set("description", result.Data.Description); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	if err := data.Set("enabled", result.Data.Enabled); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	if err := data.Set("url", result.Data.Configuration.Url); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	if err := data.Set("secret_key", result.Data.Configuration.SecretKey); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	if err := data.Set("observes", result.Data.Observes); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
}
