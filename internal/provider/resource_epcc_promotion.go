package provider

import (
	"context"
	"github.com/elasticpath/epcc-terraform-provider/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEpccPromotion() *schema.Resource {
	return &schema.Resource{
		CreateContext: addDiagToContext(resourceEpccPromotionCreate),
		ReadContext:   addDiagToContext(resourceEpccPromotionRead),
		UpdateContext: addDiagToContext(resourceEpccPromotionUpdate),
		DeleteContext: addDiagToContext(resourceEpccPromotionDelete),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id":             {Type: schema.TypeString, Computed: true},
			"type":           {Type: schema.TypeString, Required: true},
			"name":           {Type: schema.TypeString, Required: true},
			"description":    {Type: schema.TypeString, Required: true},
			"enabled":        {Type: schema.TypeBool, Required: true},
			"automatic":      {Type: schema.TypeBool, Optional: true},
			"promotion_type": {Type: schema.TypeString, Required: true},
			"schema": {Type: schema.TypeList, Required: true,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"currencies": {Type: schema.TypeList, Optional: true,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"currency": {Type: schema.TypeString, Optional: true},
							"amount":   {Type: schema.TypeInt, Optional: true},
						}}},
				}}},
			"start": {Type: schema.TypeString, Required: true},
			"end":   {Type: schema.TypeString, Required: true},
			"min_cart_value": {Type: schema.TypeList, Optional: true,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"currency": {Type: schema.TypeString, Optional: true},
					"amount":   {Type: schema.TypeInt, Optional: true},
				}}},
			"max_discount_value": {Type: schema.TypeList, Optional: true,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"currency": {Type: schema.TypeString, Optional: true},
					"amount":   {Type: schema.TypeInt, Optional: true},
				}}},
		},
	}

}

func resourceEpccPromotionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)
	promotionID := d.Id()

	err := epcc.Promotions.Delete(&ctx, client, promotionID)

	if err != nil {
		FromAPIError(err)
	}

	d.SetId("")

	return *ctx.Value("diags").(*diag.Diagnostics)
}

func resourceEpccPromotionUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	promotionId := d.Id()
	schemas := d.Get("schema").([]interface{})
	singleSchema := schemas[0]
	promotion := &epcc.Promotion{
		Id:               promotionId,
		Type:             "promotion",
		Name:             d.Get("name").(string),
		Description:      d.Get("description").(string),
		Enabled:          d.Get("enabled").(bool),
		Automatic:        d.Get("automatic").(bool),
		PromotionType:    d.Get("promotion_type").(string),
		Start:            d.Get("start").(string),
		End:              d.Get("end").(string),
		Schema:           singleSchema,
		MinCartValue:     d.Get("min_cart_value").(interface{}),
		MaxDiscountValue: d.Get("max_discount_value").(interface{}),
	}

	updatedPromotionData, apiError := epcc.Promotions.Update(&ctx, client, promotionId, promotion)

	if apiError != nil {
		return FromAPIError(apiError)
	}

	d.SetId(updatedPromotionData.Data.Id)

	return resourceEpccPromotionRead(ctx, d, m)
}

func resourceEpccPromotionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	promotionId := d.Id()

	promotion, err := epcc.Promotions.Get(&ctx, client, promotionId)

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
	if err := d.Set("promotion_type", promotion.Data.PromotionType); err != nil {
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

	return *ctx.Value("diags").(*diag.Diagnostics)
}

func resourceEpccPromotionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)
	schemas := d.Get("schema").([]interface{})
	singleSchema := schemas[0]
	promotion := &epcc.Promotion{
		Type:             "promotion",
		Name:             d.Get("name").(string),
		Description:      d.Get("description").(string),
		Enabled:          d.Get("enabled").(bool),
		Automatic:        d.Get("automatic").(bool),
		PromotionType:    d.Get("promotion_type").(string),
		Schema:           singleSchema,
		Start:            d.Get("start").(string),
		End:              d.Get("end").(string),
		MinCartValue:     d.Get("min_cart_value").(interface{}),
		MaxDiscountValue: d.Get("max_discount_value").(interface{}),
	}
	createdPromotionData, apiError := epcc.Promotions.Create(&ctx, client, promotion)
	if apiError != nil {
		return FromAPIError(apiError)
	}

	d.SetId(createdPromotionData.Data.Id)

	resourceEpccPromotionRead(ctx, d, m)

	return *ctx.Value("diags").(*diag.Diagnostics)
}
