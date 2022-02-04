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
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique identifier of the node.",
			},
			"hierarchy_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique identifier of the hierarchy.",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A name for the node. Names must be unique among sibling nodes in the hierarchy, but otherwise a name can be non-unique. Cannot be null.",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A description of the node.",
			},
			"slug": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A slug for the node. Slugs must be unique among sibling nodes in the hierarchy, but otherwise a slug can be non-unique.",
			},
			"parent_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The node ID for the parent node. The new node is created as a child of this parent node.",
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
