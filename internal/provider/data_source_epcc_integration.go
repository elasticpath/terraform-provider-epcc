package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type IntegrationDataSourceProvider struct {
}

func (ds IntegrationDataSourceProvider) DataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Allows to configure webhooks",
		ReadContext: addDiagToContext(ds.read),
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
			"integration_type": {
				Type:        schema.TypeString,
				Description: "Specifies how the event is delivered, either webhook or aws_sqs",
				Computed:    true,
			},
			"aws_access_key_id": {
				Type:        schema.TypeString,
				Description: "The required AWS access key ID. Note: The EPCC API only returns the 4 characters of this value",
				Computed:    true,
			},
			"aws_secret_access_key": {
				Type:        schema.TypeString,
				Description: "The required AWS secret key ID. Note: The EPCC API only returns the 4 characters of this value",
				Computed:    true,
				Sensitive:   true,
			},
			"region": {
				Type:        schema.TypeString,
				Description: "The required AWS region.",
				Computed:    true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func (ds IntegrationDataSourceProvider) read(ctx context.Context, data *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	result, err := epcc.Integrations.Get(&ctx, client, data.Get("id").(string))
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

	if err := data.Set("aws_access_key_id", result.Data.Configuration.AwsAccessKeyId); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	if err := data.Set("aws_secret_access_key", result.Data.Configuration.AwsSecretAccessKey); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	if err := data.Set("region", result.Data.Configuration.Region); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	data.SetId(result.Data.Id)
}
