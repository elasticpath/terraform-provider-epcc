package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gitlab.elasticpath.com/Steve.Ramage/epcc-terraform-provider/external/sdk/epcc"
)

func dataSourceEpccPricebook() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceEpccPricebookRead,
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

func dataSourceEpccPricebookRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*epcc.Client)

	var diags diag.Diagnostics

	pricebookId := d.Get("id").(string)

	pricebook, err := epcc.Pricebooks.Get(client, pricebookId)

	if err != nil {
		return FromAPIError(err)
	}

	d.Set("type", pricebook.Data.Type)
	d.Set("name", pricebook.Data.Attributes.Name)
	d.Set("description", pricebook.Data.Attributes.Description)

	d.SetId(pricebook.Data.Id)

	return diags
}
