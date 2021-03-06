package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc/payment_gateway"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type PaymentGatewayDataSourceProvider struct {
}

func (r PaymentGatewayDataSourceProvider) DataSource() *schema.Resource {
	return &schema.Resource{
		Description: "Payment gateway connectivity configuration",
		ReadContext: addDiagToContext(r.read),
		Schema: map[string]*schema.Schema{
			"slug": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The slug of the payment gateway.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Description: "Should the gateway process payments. Default: `false`",
				Computed:    true,
			},
			"test": {
				Type:        schema.TypeBool,
				Description: "Is this a sandbox environment. Default: `false`",
				Optional:    true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, data *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := data.Set("slug", data.Id())
				return []*schema.ResourceData{data}, err
			},
		},
	}
}

func (r PaymentGatewayDataSourceProvider) read(ctx context.Context, data *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	slug := payment_gateway.Slug(data.Get("slug").(string))
	result, err := epcc.PaymentGateways.Get(&ctx, client, slug)
	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	base := result.Data.Base()
	if err := data.Set("enabled", base.Enabled); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := data.Set("test", base.Test); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	data.SetId(result.Data.Type().AsString())
}
