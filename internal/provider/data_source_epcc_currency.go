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
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier for this currency.",
			},
			"code": {Type: schema.TypeString,
				Required:    true,
				Description: "The currency code.",
			},
			"exchange_rate": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The exchange rate.",
			},
			"format": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "How to structure a currency; e.g., `${price}`.",
			},
			"decimal_point": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The decimal point character.",
			},
			"thousand_separator": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The thousand separator character.",
			},
			"decimal_places": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The amount of decimal places the currency is formatted to.",
			},
			"default": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether this is the default currency in the store.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Is this currency available for products? `true` or `false`",
			},
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
