package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEpccPromotion() *schema.Resource {
	return &schema.Resource{
		Description:   "Represents the EPCC API [Promotion Object](https://documentation.elasticpath.com/commerce-cloud/docs/api/carts-and-checkout/promotions/index.html#the-promotion-object).",
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

func resourceEpccPromotionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	promotionID := d.Id()

	err := epcc.Promotions.Delete(&ctx, client, promotionID)

	if err != nil {
		ReportAPIError(ctx, err)
	}

	d.SetId("")
}

func resourceEpccPromotionUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) {
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
		ReportAPIError(ctx, apiError)
		return
	}

	d.SetId(updatedPromotionData.Data.Id)

	resourceEpccPromotionRead(ctx, d, m)
}

func resourceEpccPromotionRead(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	promotionId := d.Id()

	promotion, err := epcc.Promotions.Get(&ctx, client, promotionId)

	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	if err := d.Set("type", promotion.Data.Type); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("name", promotion.Data.Name); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("description", promotion.Data.Description); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("enabled", promotion.Data.Enabled); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("automatic", promotion.Data.Automatic); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("promotion_type", promotion.Data.PromotionType); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("start", promotion.Data.Start); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("end", promotion.Data.End); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("schema", [1]interface{}{promotion.Data.Schema}); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("min_cart_value", promotion.Data.MinCartValue); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("max_discount_value", promotion.Data.MaxDiscountValue); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
}

func resourceEpccPromotionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) {
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
		ReportAPIError(ctx, apiError)
		return
	}

	d.SetId(createdPromotionData.Data.Id)

	resourceEpccPromotionRead(ctx, d, m)
}
