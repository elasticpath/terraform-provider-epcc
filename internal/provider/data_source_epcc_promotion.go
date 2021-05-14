package provider

import (
	"context"
	"github.com/elasticpath/epcc-terraform-provider/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEpccPromotion() *schema.Resource {
	return &schema.Resource{
		ReadContext: addDiagToContext(dataSourceEpccPromotionRead),
		Schema: map[string]*schema.Schema{
			"id":             {Type: schema.TypeString, Required: true},
			"type":           {Type: schema.TypeString, Computed: true},
			"name":           {Type: schema.TypeString, Computed: true},
			"description":    {Type: schema.TypeString, Computed: true},
			"enabled":        {Type: schema.TypeBool, Computed: true},
			"automatic":      {Type: schema.TypeBool, Optional: true},
			"promotion_type": {Type: schema.TypeString, Computed: true},
			"schema": {Type: schema.TypeList, Required: true,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"currencies":
					{Type: schema.TypeList, Optional: true,
						Elem: &schema.Resource{Schema: map[string]*schema.Schema{
							"currency": {Type: schema.TypeString, Optional: true},
							"amount":   {Type: schema.TypeInt, Optional: true},
						}}},
				}}},
			"start": {Type: schema.TypeString, Computed: true},
			"end":   {Type: schema.TypeString, Computed: true},
			"min_cart_value": {Type: schema.TypeList, Optional: true,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"promotion": {Type: schema.TypeString, Optional: true},
					"amount":    {Type: schema.TypeInt, Optional: true},
				}}},
			"max_discount_value": {Type: schema.TypeList, Optional: true,
				Elem: &schema.Resource{Schema: map[string]*schema.Schema{
					"promotion": {Type: schema.TypeString, Optional: true},
					"amount":    {Type: schema.TypeInt, Optional: true},
				}}},
		},
	}
}

func dataSourceEpccPromotionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	promotionId := d.Get("id").(string)
	promotion, err := epcc.Promotions.Get(&ctx, client, promotionId)
	if err != nil {
		return FromAPIError(err)
	}

	d.Set("type", promotion.Data.Type)
	d.Set("name", promotion.Data.Name)
	d.Set("description", promotion.Data.Description)
	d.Set("enabled", promotion.Data.Enabled)
	d.Set("automatic", promotion.Data.Automatic)
	d.Set("promotion_type", promotion.Data.PromotionType)
	d.Set("schema", promotion.Data.Schema)
	d.Set("start", promotion.Data.Start)
	d.Set("end", promotion.Data.End)
	d.Set("min_cart_value", promotion.Data.MinCartValue)
	d.Set("max_discount_value", promotion.Data.MaxDiscountValue)

	d.SetId(promotion.Data.Id)

	return *ctx.Value("diags").(*diag.Diagnostics)
}
