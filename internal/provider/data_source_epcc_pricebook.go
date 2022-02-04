package provider

import (
	"context"

	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEpccPricebook() *schema.Resource {
	return &schema.Resource{
		ReadContext: addDiagToContext(dataSourceEpccPricebookRead),
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique identifier of the price book.",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A unique name for the price book.",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The purpose for the price book, such as flash sale pricing or preferred customer pricing.",
			},
		},
	}
}

func dataSourceEpccPricebookRead(ctx context.Context, d *schema.ResourceData, m interface{}) {

	client := m.(*epcc.Client)
	pricebookId := d.Get("id").(string)

	pricebook, err := epcc.Pricebooks.Get(&ctx, client, pricebookId)

	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	d.Set("name", pricebook.Data.Attributes.Name)
	d.Set("description", pricebook.Data.Attributes.Description)

	d.SetId(pricebook.Data.Id)
}
