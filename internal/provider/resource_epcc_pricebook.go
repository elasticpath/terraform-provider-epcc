package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gitlab.elasticpath.com/elasticpathepcc-terraform-provider/external/sdk/epcc"
)

func resourceEpccPricebook() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceEpccPricebookCreate,
		ReadContext:   resourceEpccPricebookRead,
		UpdateContext: resourceEpccPricebookUpdate,
		DeleteContext: resourceEpccPricebookDelete,
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

func resourceEpccPricebookDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	var diags diag.Diagnostics

	pricebookId := d.Id()

	err := epcc.Pricebooks.Delete(client, pricebookId)

	if err != nil {
		FromAPIError(err)
	}

	d.SetId("")

	return diags
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

	updatedPricebookData, apiError := epcc.Pricebooks.Update(client, pricebookId, pricebook)

	if apiError != nil {
		return FromAPIError(apiError)
	}

	d.SetId(updatedPricebookData.Data.Id)

	return resourceEpccPricebookRead(ctx, d, m)
}

func resourceEpccPricebookRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	var diags diag.Diagnostics

	pricebookId := d.Id()

	pricebook, err := epcc.Pricebooks.Get(client, pricebookId)

	if err != nil {
		return FromAPIError(err)
	}

	if err := d.Set("name", pricebook.Data.Attributes.Name); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("description", pricebook.Data.Attributes.Description); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceEpccPricebookCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	var diags diag.Diagnostics

	pricebook := &epcc.Pricebook{
		Type: "pricebook",
		Attributes: epcc.PricebookAttributes{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
		},
	}

	createdPricebookData, apiError := epcc.Pricebooks.Create(client, pricebook)

	if apiError != nil {
		return FromAPIError(apiError)
	}

	d.SetId(createdPricebookData.Data.Id)

	resourceEpccPricebookRead(ctx, d, m)

	return diags
}
