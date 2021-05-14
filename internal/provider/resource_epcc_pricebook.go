package provider

import (
	"context"

	"github.com/elasticpath/epcc-terraform-provider/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEpccPricebook() *schema.Resource {
	return &schema.Resource{
		CreateContext: addDiagToContext(resourceEpccPricebookCreate),
		ReadContext:   addDiagToContext(resourceEpccPricebookRead),
		UpdateContext: addDiagToContext(resourceEpccPricebookUpdate),
		DeleteContext: addDiagToContext(resourceEpccPricebookDelete),
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}

}

func resourceEpccPricebookDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)
	pricebookId := d.Id()

	err := epcc.Pricebooks.Delete(&ctx, client, pricebookId)

	if err != nil {
		FromAPIError(err)
	}

	d.SetId("")

	return *ctx.Value("diags").(*diag.Diagnostics)
}

func resourceEpccPricebookUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	pricebookId := d.Id()

	pricebook := &epcc.Pricebook{
		Id:   pricebookId,
		Type: "pricebook",
		Attributes: epcc.PricebookAttributes{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
		},
	}

	updatedPricebookData, apiError := epcc.Pricebooks.Update(&ctx, client, pricebookId, pricebook)

	if apiError != nil {
		return FromAPIError(apiError)
	}

	d.SetId(updatedPricebookData.Data.Id)

	return resourceEpccPricebookRead(ctx, d, m)
}

func resourceEpccPricebookRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)
	pricebookId := d.Id()

	pricebook, err := epcc.Pricebooks.Get(&ctx, client, pricebookId)

	if err != nil {
		return FromAPIError(err)
	}

	if err := d.Set("name", pricebook.Data.Attributes.Name); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("description", pricebook.Data.Attributes.Description); err != nil {
		return diag.FromErr(err)
	}

	return *ctx.Value("diags").(*diag.Diagnostics)
}

func resourceEpccPricebookCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)
	pricebook := &epcc.Pricebook{
		Type: "pricebook",
		Attributes: epcc.PricebookAttributes{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
		},
	}

	createdPricebookData, apiError := epcc.Pricebooks.Create(&ctx, client, pricebook)

	if apiError != nil {
		return FromAPIError(apiError)
	}

	d.SetId(createdPricebookData.Data.Id)

	resourceEpccPricebookRead(ctx, d, m)

	return *ctx.Value("diags").(*diag.Diagnostics)
}
