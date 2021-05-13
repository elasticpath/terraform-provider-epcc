package provider

import (
	"context"
	"github.com/elasticpath/epcc-terraform-provider/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEpccCurrency() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceEpccCurrencyCreate,
		ReadContext:   resourceEpccCurrencyRead,
		UpdateContext: resourceEpccCurrencyUpdate,
		DeleteContext: resourceEpccCurrencyDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id":                 {Type: schema.TypeString, Computed: true},
			"type":               {Type: schema.TypeString, Required: true},
			"code":               {Type: schema.TypeString, Required: true},
			"exchange_rate":      {Type: schema.TypeInt, Required: true},
			"format":             {Type: schema.TypeString, Required: true},
			"decimal_point":      {Type: schema.TypeString, Required: true},
			"thousand_separator": {Type: schema.TypeString, Required: true},
			"decimal_places":     {Type: schema.TypeInt, Required: true},
			"default":            {Type: schema.TypeBool, Required: true},
			"enabled":            {Type: schema.TypeBool, Required: true},
		},
	}

}

func resourceEpccCurrencyDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	var diags diag.Diagnostics

	currencyID := d.Id()

	err := epcc.Currencies.Delete(client, currencyID)

	if err != nil {
		FromAPIError(err)
	}

	d.SetId("")

	return diags
}

func resourceEpccCurrencyUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

	updatedCurrencyData, apiError := epcc.Currencies.Update(client, currencyId, currency)

	if apiError != nil {
		return FromAPIError(apiError)
	}

	d.SetId(updatedCurrencyData.Data.Id)

	return resourceEpccCurrencyRead(ctx, d, m)
}

func resourceEpccCurrencyRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	var diags diag.Diagnostics

	currencyId := d.Id()

	currency, err := epcc.Currencies.Get(client, currencyId)

	if err != nil {
		return FromAPIError(err)
	}

	if err := d.Set("type", currency.Data.Type); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("code", currency.Data.Code); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("exchange_rate", currency.Data.ExchangeRate); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("format", currency.Data.Format); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("decimal_point", currency.Data.DecimalPoint); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("thousand_separator", currency.Data.ThousandSeparator); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("decimal_places", currency.Data.DecimalPlaces); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("default", currency.Data.Default); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("enabled", currency.Data.Enabled); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceEpccCurrencyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	var diags diag.Diagnostics

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
	createdCurrencyData, apiError := epcc.Currencies.Create(client, currency)
	if apiError != nil {
		return FromAPIError(apiError)
	}

	d.SetId(createdCurrencyData.Data.Id)

	resourceEpccCurrencyRead(ctx, d, m)

	return diags
}