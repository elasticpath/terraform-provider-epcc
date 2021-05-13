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

	flowId := d.Get("id").(string)

	flow, err := epcc.Flows.Get(client, flowId)

	if err != nil {
		return FromAPIError(err)
	}

	d.Set("type", flow.Data.Type)
	d.Set("name", flow.Data.Name)
	d.Set("slug", flow.Data.Slug)
	d.Set("description", flow.Data.Description)
	d.Set("enabled", flow.Data.Enabled)




	d.SetId(flow.Data.Id)

	return diags
}
