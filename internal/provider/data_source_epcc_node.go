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
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"hierarchy_id": &schema.Schema{
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
			"slug": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"products": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceEpccNodeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*epcc.Client)
	nodeId := d.Get("id").(string)
	hierarchyId := d.Get("hierarchy_id").(string)
	node, err := epcc.Nodes.Get(&ctx, client, hierarchyId, nodeId)

	if err != nil {
		return FromAPIError(err)
	}

	nodeProducts, err := epcc.Nodes.GetNodeProducts(&ctx, client, hierarchyId, nodeId)

	if err != nil {
		return FromAPIError(err)
	}

	d.Set("name", node.Data.Attributes.Name)
	d.Set("type", node.Data.Type)
	d.Set("slug", node.Data.Attributes.Slug)

	if node.Data.Relationships != nil && node.Data.Relationships.Parent != nil && node.Data.Relationships.Parent.Data != nil {
		if err := d.Set("parent_id", node.Data.Relationships.Parent.Data.Id); err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId(node.Data.Id)

	if nodeProducts != nil && nodeProducts.Data != nil {
		fileIds := convertJsonTypesToIds(nodeProducts.Data)

		if err := d.Set("products", fileIds); err != nil {
			return diag.FromErr(err)
		}
	} else {
		if err := d.Set("products", [0]string{}); err != nil {
			return diag.FromErr(err)
		}
	}

	return *ctx.Value("diags").(*diag.Diagnostics)
}
