package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gitlab.elasticpath.com/Steve.Ramage/epcc-terraform-provider/external/sdk/epcc"
)

func dataSourceEpccFlow() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceEpccFlowRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"slug": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceEpccFlowRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*epcc.Client)

	var diags diag.Diagnostics

	accountId := d.Get("id").(string)

	account, err := epcc.Flows.Get(client, accountId)

	if err != nil {
		return FromAPIError(err)
	}

	d.Set("type", account.Data.Type)
	d.Set("name", account.Data.Name)
	d.Set("slug", account.Data.Slug)
	d.Set("description", account.Data.Description)
	d.Set("enabled", account.Data.Enabled)




	d.SetId(account.Data.Id)

	return diags
}
