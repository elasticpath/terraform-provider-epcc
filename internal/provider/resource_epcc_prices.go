package provider

import (
	"context"

	"github.com/elasticpath/epcc-terraform-provider/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEpccPrices() *schema.Resource {
	return &schema.Resource{
		Description:   "Allows the caller to create, update or delete Elastic Path Commerce Cloud PCM Pricebook [prices](https://documentation.elasticpath.com/commerce-cloud/docs/concepts/products-pcm.html).",
		CreateContext: resourceEpccPricesCreate,
		ReadContext:   resourceEpccPricesRead,
		UpdateContext: resourceEpccPricesUpdate,
		DeleteContext: resourceEpccPricesDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id":           {Type: schema.TypeString, Computed: true, Description: "The unique identifier of the price."},
			"pricebook_id": {Type: schema.TypeString, Required: true},
			"sku":          {Type: schema.TypeString, Required: true},
			"prices": {Type: schema.TypeList, Required: true,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"price": {Type: schema.TypeList, Optional: true,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"currency":     {Type: schema.TypeString, Optional: true},
							"amount":       {Type: schema.TypeInt, Optional: true},
							"includes_tax": {Type: schema.TypeInt, Optional: true},
						}}},
				}}},
		},
	}

}

func resourceEpccPricesDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	var diags diag.Diagnostics

	pricebookId := d

	err := epcc.Prices.Delete(client, pricebookId, priceId)

	if err != nil {
		FromAPIError(err)
	}

	d.SetId("")

	return diags
}

func resourceEpccPricesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	promotionId := d.Id()
	schemas := d.Get("schema").([]interface{})
	singleSchema := schemas[0]
	promotion := &epcc.Prices{
		Id:               promotionId,
		Type:             "promotion",
		Name:             d.Get("name").(string),
		Description:      d.Get("description").(string),
		Enabled:          d.Get("enabled").(bool),
		Automatic:        d.Get("automatic").(bool),
		PricesType:       d.Get("promotion_type").(string),
		Start:            d.Get("start").(string),
		End:              d.Get("end").(string),
		Schema:           singleSchema,
		MinCartValue:     d.Get("min_cart_value").(interface{}),
		MaxDiscountValue: d.Get("max_discount_value").(interface{}),
	}

	updatedPricesData, apiError := epcc.Pricess.Update(client, promotionId, promotion)

	if apiError != nil {
		return FromAPIError(apiError)
	}

	d.SetId(updatedPricesData.Data.Id)

	return resourceEpccPricesRead(ctx, d, m)
}

func resourceEpccPricesRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	var diags diag.Diagnostics

	promotionId := d.Id()

	promotion, err := epcc.Pricess.Get(client, promotionId)

	if err != nil {
		return FromAPIError(err)
	}

	if err := d.Set("type", promotion.Data.Type); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", promotion.Data.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", promotion.Data.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("enabled", promotion.Data.Enabled); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("automatic", promotion.Data.Automatic); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("promotion_type", promotion.Data.PricesType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("start", promotion.Data.Start); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("end", promotion.Data.End); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("schema", [1]interface{}{promotion.Data.Schema}); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("min_cart_value", promotion.Data.MinCartValue); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("max_discount_value", promotion.Data.MaxDiscountValue); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceEpccPricesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	var diags diag.Diagnostics

	pricebookId := d.Get("pricebookId").(string)

	allPrices := d.Get("prices").([]interface{})
	singlePrice := allPrices[0]

	productPrice := &epcc.Prices{
		Type: "product-price",
		Attributes: epcc.PriceAttributes{
			Currencies: singlePrice,
		},
	}
	createdPricesData, apiError := epcc.Prices.Create(client, pricebookId, productPrice)
	if apiError != nil {
		return FromAPIError(apiError)
	}

	d.SetId(createdPricesData.Data.Id)

	resourceEpccPricesRead(ctx, d, m)

	return diags
}
