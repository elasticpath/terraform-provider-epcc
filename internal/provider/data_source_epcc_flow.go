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
			"id": {
				Type:     schema.TypeString,
				Required: true,
				Description: "The unique identifier for this flow.",
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: "The name of the flow.",
			},
			"slug": {
				Type:     schema.TypeString,
				Computed: true,
				Description: "A unique slug identifier for the flow.",
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
				Description: "Any description for this flow.",
			},
			"enabled": {
				Type:     schema.TypeBool,
				Computed: true,
				Description: "true if enabled, false if not.",
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
