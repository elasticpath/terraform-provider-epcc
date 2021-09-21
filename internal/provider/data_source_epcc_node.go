package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEpccNode() *schema.Resource {
	return &schema.Resource{
		ReadContext: addDiagToContext(dataSourceEpccNodeRead),
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"hierarchy_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"slug": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceEpccNodeRead(ctx context.Context, d *schema.ResourceData, m interface{}) {

	client := m.(*epcc.Client)
	nodeId := d.Get("id").(string)
	hierarchyId := d.Get("hierarchy_id").(string)
	node, err := epcc.Nodes.Get(&ctx, client, hierarchyId, nodeId)

	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	d.Set("name", node.Data.Attributes.Name)
	d.Set("type", node.Data.Type)
	d.Set("slug", node.Data.Attributes.Slug)

	if node.Data.Relationships != nil && node.Data.Relationships.Parent != nil && node.Data.Relationships.Parent.Data != nil {
		if err := d.Set("parent_id", node.Data.Relationships.Parent.Data.Id); err != nil {
			addToDiag(ctx, diag.FromErr(err))
			return
		}
	}

	d.SetId(node.Data.Id)
}
