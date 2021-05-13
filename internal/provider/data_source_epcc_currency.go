package provider

import (
	"context"
	"github.com/elasticpath/epcc-terraform-provider/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEpccCurrency() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceEpccCurrencyRead,
	}
}

func dataSourceEpccCurrencyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)
	var diags diag.Diagnostics
	currencyId := d.Get("id").(string)
	currency, err := epcc.Currencies.Get(client, currencyId)
	if err != nil {
		return FromAPIError(err)
	}

	d.SetId(currency.Data.Id)

	return diags
}
