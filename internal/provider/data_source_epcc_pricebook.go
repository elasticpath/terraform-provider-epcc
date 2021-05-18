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
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
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
