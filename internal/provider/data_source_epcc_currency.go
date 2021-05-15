package provider

import (
	"context"
	"github.com/elasticpath/epcc-terraform-provider/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEpccCurrency() *schema.Resource {
	return &schema.Resource{
		ReadContext: addDiagToContext(dataSourceEpccCurrencyRead),
		Schema: map[string]*schema.Schema{
			"id":                 {Type: schema.TypeString, Required: true},
			"code":               {Type: schema.TypeString, Computed: true},
			"exchange_rate":      {Type: schema.TypeInt, Computed: true},
			"format":             {Type: schema.TypeString, Computed: true},
			"decimal_point":      {Type: schema.TypeString, Computed: true},
			"thousand_separator": {Type: schema.TypeString, Computed: true},
			"decimal_places":     {Type: schema.TypeInt, Computed: true},
			"default":            {Type: schema.TypeBool, Computed: true},
			"enabled":            {Type: schema.TypeBool, Computed: true},
		},
	}
}

func dataSourceEpccCurrencyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	currencyId := d.Get("id").(string)
	currency, err := epcc.Currencies.Get(&ctx, client, currencyId)
	if err != nil {
		return FromAPIError(err)
	}

	d.Set("code", currency.Data.Code)
	d.Set("exchange_rate", currency.Data.ExchangeRate)
	d.Set("format", currency.Data.Format)
	d.Set("decimal_point", currency.Data.DecimalPoint)
	d.Set("thousand_separator", currency.Data.ThousandSeparator)
	d.Set("decimal_places", currency.Data.DecimalPlaces)
	d.Set("default", currency.Data.Default)
	d.Set("enabled", currency.Data.Enabled)

	d.SetId(currency.Data.Id)

	return *ctx.Value("diags").(*diag.Diagnostics)
}
