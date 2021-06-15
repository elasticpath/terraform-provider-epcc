package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEpccCurrency() *schema.Resource {
	return &schema.Resource{
		ReadContext: addDiagToContext(dataSourceEpccCurrencyRead),
		Schema: map[string]*schema.Schema{
			"id":                 {Type: schema.TypeString, Computed: true},
			"code":               {Type: schema.TypeString, Required: true},
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

func dataSourceEpccCurrencyRead(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	currencyCode := d.Get("code").(string)
	currencies, err := epcc.Currencies.GetAll(&ctx, client)
	if err != nil {
		ReportAPIError(ctx, err)
	} else {
		var currency *epcc.Currency
		for i, next := range currencies.Data {
			if next.Code == currencyCode {
				currency = &currencies.Data[i]
				break
			}
		}
		if currency == nil {
			addToDiag(ctx, diag.Errorf("currency with code %v not found", currencyCode))
			return
		}

		if err := d.Set("exchange_rate", currency.ExchangeRate); err != nil {
			addToDiag(ctx, diag.FromErr(err))
			return
		}
		if err := d.Set("format", currency.Format); err != nil {
			addToDiag(ctx, diag.FromErr(err))
			return
		}
		if err := d.Set("decimal_point", currency.DecimalPoint); err != nil {
			addToDiag(ctx, diag.FromErr(err))
			return
		}
		if err := d.Set("thousand_separator", currency.ThousandSeparator); err != nil {
			addToDiag(ctx, diag.FromErr(err))
			return
		}
		if err := d.Set("decimal_places", currency.DecimalPlaces); err != nil {
			addToDiag(ctx, diag.FromErr(err))
			return
		}
		if err := d.Set("default", currency.Default); err != nil {
			addToDiag(ctx, diag.FromErr(err))
			return
		}
		if err := d.Set("enabled", currency.Enabled); err != nil {
			addToDiag(ctx, diag.FromErr(err))
			return
		}

		d.SetId(currency.Id)
	}
}
