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
			"integration_type": {
				Type:        schema.TypeString,
				Description: "Specifies how the event is delivered, either webhook or aws_sqs",
				Required:    true,
			},
			"aws_access_key_id": {
				Type:        schema.TypeString,
				Description: "The required AWS access key ID. Note: The EPCC API only returns the 4 characters of this value",
				Optional:    true,
			},
			"aws_secret_access_key": {
				Type:        schema.TypeString,
				Description: "The required AWS secret key ID. Note: The EPCC API only returns the 4 characters of this value",
				Optional:    true,
				Sensitive:   true,
			},
			"region": {
				Type:        schema.TypeString,
				Description: "The required AWS region.",
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
		IntegrationType: data.Get("integration_type").(string),
		Name:            data.Get("name").(string),
		Description:     data.Get("description").(string),
		Enabled:         data.Get("enabled").(bool),
		Configuration: epcc.IntegrationConfiguration{
			Url:                data.Get("url").(string),
			SecretKey:          data.Get("secret_key").(string),
			AwsAccessKeyId:     data.Get("aws_access_key_id").(string),
			AwsSecretAccessKey: data.Get("aws_secret_access_key").(string),
			Region:             data.Get("region").(string),
		},
		Observes: convertArrayToStringSlice(observes),
	}

	result, err := epcc.Integrations.Create(&ctx, client, integrationObject)
	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	data.SetId(result.Data.Id)

	// EPCC API only returns the last 4 characters for these two keys, so we will store in state the correct value from the request
	if err := data.Set("aws_access_key_id", integrationObject.Configuration.AwsAccessKeyId); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	if err := data.Set("aws_secret_access_key", integrationObject.Configuration.AwsSecretAccessKey); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	r.read(ctx, data, m)
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
		Type:            epcc.IntegrationType,
		IntegrationType: data.Get("integration_type").(string),
		Name:            data.Get("name").(string),
		Description:     data.Get("description").(string),
		Enabled:         data.Get("enabled").(bool),
		Configuration: epcc.IntegrationConfiguration{
			Url:                data.Get("url").(string),
			SecretKey:          data.Get("secret_key").(string),
			AwsAccessKeyId:     data.Get("aws_access_key_id").(string),
			AwsSecretAccessKey: data.Get("aws_secret_access_key").(string),
			Region:             data.Get("region").(string),
		},
		Observes: convertArrayToStringSlice(observes),
	}

	result, apiError := epcc.Integrations.Update(&ctx, client, data.Id(), integrationObject)

	if apiError != nil {
		ReportAPIError(ctx, apiError)
		return
	}

	data.SetId(result.Data.Id)

	// EPCC API only returns the last 4 characters for these two keys, so we will store in state the correct value from the request
	if err := data.Set("aws_access_key_id", integrationObject.Configuration.AwsAccessKeyId); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	if err := data.Set("aws_secret_access_key", integrationObject.Configuration.AwsSecretAccessKey); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	r.read(ctx, data, m)
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

	if err := data.Set("integration_type", result.Data.IntegrationType); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	if err := data.Set("region", result.Data.Configuration.Region); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	// Update the state to the API value only if the last 4 characters don't match.
	if accessKeyIdFromState, ok := data.GetOk("aws_access_key_id"); ok {
		if lastFourChars(accessKeyIdFromState.(string)) != lastFourChars(result.Data.Configuration.AwsAccessKeyId) {
			if err := data.Set("aws_access_key_id", result.Data.Configuration.AwsAccessKeyId); err != nil {
				addToDiag(ctx, diag.FromErr(err))
				return
			}
		}
	}

	// Update the state to the API value only if the last 4 characters don't match.
	if secretAccessKeyFromState, ok := data.GetOk("aws_secret_access_key"); ok {
		if lastFourChars(secretAccessKeyFromState.(string)) != lastFourChars(result.Data.Configuration.AwsSecretAccessKey) {
			if err := data.Set("aws_secret_access_key", result.Data.Configuration.AwsSecretAccessKey); err != nil {
				addToDiag(ctx, diag.FromErr(err))
				return
			}
		}
	}
}

func lastFourChars(s string) string {
	if len(s) > 4 {
		// not UTF-8 safe, but who knows if AWS access keys will ever be UTF-8
		return s[len(s)-4:]
	}

	return s
}
