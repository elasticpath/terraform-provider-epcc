package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEpccFlow() *schema.Resource {
	return &schema.Resource{
		ReadContext: addDiagToContext(dataSourceEpccFlowRead),
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
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSourceEpccFlowRead(ctx context.Context, d *schema.ResourceData, m interface{}) {

	client := m.(*epcc.Client)
	flowId := d.Get("id").(string)

	flow, err := epcc.Flows.Get(&ctx, client, flowId)

	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	d.Set("type", flow.Data.Type)
	d.Set("name", flow.Data.Name)
	d.Set("slug", flow.Data.Slug)
	d.Set("description", flow.Data.Description)
	d.Set("enabled", flow.Data.Enabled)

	d.SetId(flow.Data.Id)
}
