package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEpccCurrency() *schema.Resource {
	return &schema.Resource{
		Description:   "Represents the EPCC API [Currency Object](https://documentation.elasticpath.com/commerce-cloud/docs/api/advanced/currencies/index.html#the-currency-object).",
		CreateContext: addDiagToContext(resourceEpccCurrencyCreate),
		ReadContext:   addDiagToContext(resourceEpccCurrencyRead),
		UpdateContext: addDiagToContext(resourceEpccCurrencyUpdate),
		DeleteContext: addDiagToContext(resourceEpccCurrencyDelete),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier for this currency.",
			},
			"code": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The currency code.",
			},
			"exchange_rate": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The exchange rate.",
			},
			"format": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "How to structure a currency; e.g., `${price}`.",
			},
			"decimal_point": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The decimal point character.",
			},
			"thousand_separator": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The thousand separator character.",
			},
			"decimal_places": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The amount of decimal places the currency is formatted to.",
			},
			"default": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Whether this is the default currency in the store.",

				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					if val.(bool) {
						warns = append(warns, "If multiple currencies are defined, please ensure that the `default` tag is set to `true` on only one of them")
					}
					return
				}},
			"enabled": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Is this currency available for products? `true` or `false`",
			},
		},
	}

}

func resourceEpccCurrencyDelete(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	currencyID := d.Id()

	err := epcc.Currencies.Delete(&ctx, client, currencyID)

	if err != nil {
		ReportAPIError(ctx, err)
	}

	d.SetId("")
}

func resourceEpccCurrencyUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	currencyId := d.Id()

	currency := &epcc.Currency{
		Id:                currencyId,
		Type:              "currency",
		Code:              d.Get("code").(string),
		ExchangeRate:      d.Get("exchange_rate").(int),
		Format:            d.Get("format").(string),
		DecimalPoint:      d.Get("decimal_point").(string),
		ThousandSeparator: d.Get("thousand_separator").(string),
		DecimalPlaces:     d.Get("decimal_places").(int),
		Default:           d.Get("default").(bool),
		Enabled:           d.Get("enabled").(bool),
	}

	updatedCurrencyData, apiError := epcc.Currencies.Update(&ctx, client, currencyId, currency)

	if apiError != nil {
		ReportAPIError(ctx, apiError)
		return
	}

	d.SetId(updatedCurrencyData.Data.Id)

	resourceEpccCurrencyRead(ctx, d, m)
}

func resourceEpccCurrencyRead(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	currencyId := d.Id()

	currency, err := epcc.Currencies.Get(&ctx, client, currencyId)

	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	if err := d.Set("code", currency.Data.Code); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("exchange_rate", currency.Data.ExchangeRate); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("format", currency.Data.Format); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("decimal_point", currency.Data.DecimalPoint); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("thousand_separator", currency.Data.ThousandSeparator); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("decimal_places", currency.Data.DecimalPlaces); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("default", currency.Data.Default); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("enabled", currency.Data.Enabled); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
}

func resourceEpccCurrencyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	currency := &epcc.Currency{
		Type:              "currency",
		Code:              d.Get("code").(string),
		ExchangeRate:      d.Get("exchange_rate").(int),
		Format:            d.Get("format").(string),
		DecimalPoint:      d.Get("decimal_point").(string),
		ThousandSeparator: d.Get("thousand_separator").(string),
		DecimalPlaces:     d.Get("decimal_places").(int),
		Default:           d.Get("default").(bool),
		Enabled:           d.Get("enabled").(bool),
	}
	createdCurrencyData, apiError := epcc.Currencies.Create(&ctx, client, currency)
	if apiError != nil {
		ReportAPIError(ctx, apiError)
		return
	}

	d.SetId(createdCurrencyData.Data.Id)

	resourceEpccCurrencyRead(ctx, d, m)
}
