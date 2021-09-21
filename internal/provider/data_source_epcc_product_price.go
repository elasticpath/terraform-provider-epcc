package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEpccProductPrice() *schema.Resource {
	return &schema.Resource{
		ReadContext: addDiagToContext(dataSourceEpccProductPriceRead),
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique identifier of the price."},
			"pricebook_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sku": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"currency": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"code": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"amount": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"includes_tax": {
							Type:     schema.TypeBool,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceEpccProductPriceRead(ctx context.Context, d *schema.ResourceData, m interface{}) {

	client := m.(*epcc.Client)
	priceId := d.Get("id").(string)

	productPrice, err := epcc.ProductPrices.Get(&ctx, client, d.Get("pricebook_id").(string), priceId)

	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	if err := d.Set("sku", productPrice.Data.Attributes.Sku); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	currencies := make([]map[string]interface{}, 0)

	for code, apiCurrency := range productPrice.Data.Attributes.Currencies {
		currency := make(map[string]interface{})
		currency["code"] = code
		currency["amount"] = apiCurrency.Amount
		currency["includes_tax"] = apiCurrency.IncludesTax
		currencies = append(currencies, currency)
	}

	if err := d.Set("currency", currencies); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	d.SetId(productPrice.Data.Id)
}
