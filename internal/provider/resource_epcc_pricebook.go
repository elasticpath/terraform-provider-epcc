package provider

import (
	"context"

	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEpccPricebook() *schema.Resource {
	return &schema.Resource{
		Description:   "Represents the EPCC API [PriceBook Object](https://documentation.elasticpath.com/commerce-cloud/docs/api/pcm/pricebooks/index.html#the-pricebook-object).",
		CreateContext: addDiagToContext(resourceEpccPricebookCreate),
		ReadContext:   addDiagToContext(resourceEpccPricebookRead),
		UpdateContext: addDiagToContext(resourceEpccPricebookUpdate),
		DeleteContext: addDiagToContext(resourceEpccPricebookDelete),
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier of the price book.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A unique name for the price book.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The purpose for the price book, such as flash sale pricing or preferred customer pricing.",
			},
		},
	}

}

func resourceEpccPricebookDelete(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	pricebookId := d.Id()

	err := epcc.Pricebooks.Delete(&ctx, client, pricebookId)

	if err != nil {
		ReportAPIError(ctx, err)
	}

	d.SetId("")
}

func resourceEpccPricebookUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) {
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
		ReportAPIError(ctx, apiError)
		return
	}

	d.SetId(updatedPricebookData.Data.Id)

	resourceEpccPricebookRead(ctx, d, m)
}

func resourceEpccPricebookRead(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	pricebookId := d.Id()

	pricebook, err := epcc.Pricebooks.Get(&ctx, client, pricebookId)

	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	if err := d.Set("name", pricebook.Data.Attributes.Name); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	if err := d.Set("description", pricebook.Data.Attributes.Description); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
}

func resourceEpccPricebookCreate(ctx context.Context, d *schema.ResourceData, m interface{}) {
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
		ReportAPIError(ctx, apiError)
		return
	}

	d.SetId(createdPricebookData.Data.Id)

	resourceEpccPricebookRead(ctx, d, m)
}
