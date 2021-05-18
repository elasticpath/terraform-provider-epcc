package provider

import (
	"context"
	"fmt"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEpccProductPrice() *schema.Resource {
	return &schema.Resource{
		Description:   "Represents the EPCC API [Price Object](https://documentation.elasticpath.com/commerce-cloud/docs/api/pcm/pricebooks/prices/create-product-prices.html).",
		CreateContext: addDiagToContext(resourceEpccProductPriceCreate),
		ReadContext:   addDiagToContext(resourceEpccProductPriceRead),
		UpdateContext: addDiagToContext(resourceEpccProductPriceUpdate),
		DeleteContext: addDiagToContext(resourceEpccProductPriceDelete),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier of the price."},
			"pricebook_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"sku": {
				Type:     schema.TypeString,
				Required: true,
			},
			"currency": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"code": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"amount": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
						"includes_tax": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
						},
					},
				},
			},
		},
	}

}

func resourceEpccProductPriceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	priceID := d.Id()

	err := epcc.ProductPrices.Delete(&ctx, client, d.Get("pricebook_id").(string), priceID)

	if err != nil {
		ReportAPIError(ctx, err)
	}

	d.SetId("")
}

func resourceEpccProductPriceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	priceId := d.Id()

	productPrice := &epcc.ProductPrice{
		Id:   priceId,
		Type: "product-price",
		Attributes: epcc.ProductPriceAttributes{
			Sku:        d.Get("sku").(string),
			Currencies: map[string]epcc.ProductPriceInCurrency{},
		},
	}

	currencies, err := convertSchemaCurrencyToApi(d.Get("currency").(*schema.Set))

	if err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	productPrice.Attributes.Currencies = currencies

	updatedPriceData, apiError := epcc.ProductPrices.Update(&ctx, client, d.Get("pricebook_id").(string), priceId, productPrice)

	if apiError != nil {
		ReportAPIError(ctx, apiError)
		return
	}

	d.SetId(updatedPriceData.Data.Id)

	resourceEpccProductPriceRead(ctx, d, m)
}

func resourceEpccProductPriceRead(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	priceId := d.Id()

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
}

func resourceEpccProductPriceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	productPrice := &epcc.ProductPrice{
		Type: "product-price",
		Attributes: epcc.ProductPriceAttributes{
			Sku:        d.Get("sku").(string),
			Currencies: map[string]epcc.ProductPriceInCurrency{},
		},
	}

	currencies, err := convertSchemaCurrencyToApi(d.Get("currency").(*schema.Set))

	if err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	productPrice.Attributes.Currencies = currencies

	createdPriceData, apiError := epcc.ProductPrices.Create(&ctx, client, d.Get("pricebook_id").(string), productPrice)
	if apiError != nil {
		ReportAPIError(ctx, apiError)
		return
	}

	d.SetId(createdPriceData.Data.Id)

	resourceEpccProductPriceRead(ctx, d, m)
}

func convertSchemaCurrencyToApi(currencies *schema.Set) (ref map[string]epcc.ProductPriceInCurrency, err error) {

	ref = make(map[string]epcc.ProductPriceInCurrency)
	for _, currencyInt := range currencies.List() {
		currency := currencyInt.(map[string]interface{})

		code := currency["code"].(string)

		if _, alreadyExists := ref[code]; alreadyExists {
			return nil, fmt.Errorf("Currency %s is specified more than once", code)
		}

		priceCurrency := epcc.ProductPriceInCurrency{
			Amount:      currency["amount"].(int),
			IncludesTax: currency["includes_tax"].(bool),
		}

		ref[code] = priceCurrency
	}

	return ref, err
}
