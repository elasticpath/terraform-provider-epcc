package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEpccNode() *schema.Resource {
	return &schema.Resource{
		Description:   "Represents the EPCC API [Node Object](https://documentation.elasticpath.com/commerce-cloud/docs/api/pcm/hierarchies/index.html#the-node-object).",
		CreateContext: addDiagToContext(resourceEpccNodeCreate),
		ReadContext:   addDiagToContext(resourceEpccNodeRead),
		UpdateContext: addDiagToContext(resourceEpccNodeUpdate),
		DeleteContext: addDiagToContext(resourceEpccNodeDelete),
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier of the node.",
			},
			"hierarchy_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The unique identifier of the hierarchy.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Computed:    false,
				Description: "A name for the node. Names must be unique among sibling nodes in the hierarchy, but otherwise a name can be non-unique. Cannot be null.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A description of the node.",
			},
			"slug": {
				Type:        schema.TypeString,
				Required:    false,
				Computed:    false,
				Optional:    true,
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

func resourceEpccNodeDelete(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	err := epcc.Nodes.Delete(&ctx, client, d.Get("hierarchy_id").(string), d.Id())

	if err != nil {
		ReportAPIError(ctx, err)
	}

	d.SetId("")
}

func resourceEpccNodeUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	hierarchyId := d.Get("hierarchy_id").(string)

	node := &epcc.Node{
		Type: "node",
		Id:   d.Id(),
		Attributes: epcc.NodeAttributes{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
			Slug:        d.Get("slug").(string),
		},
		Relationships: nil,
	}

	if len(d.Get("parent_id").(string)) > 0 {
		node.Relationships = &epcc.NodesRelationships{
			Parent: &epcc.DataForTypeIdRelationship{
				Data: &epcc.TypeIdRelationship{
					Id:   d.Get("parent_id").(string),
					Type: "node",
				},
			},
		}
	}

	updatedNodeData, apiError := epcc.Nodes.Update(&ctx, client, hierarchyId, d.Id(), node)

	if apiError != nil {
		ReportAPIError(ctx, apiError)
		return
	}

	d.SetId(updatedNodeData.Data.Id)

	resourceEpccNodeRead(ctx, d, m)
}

func resourceEpccNodeRead(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	hierarchyId := d.Get("hierarchy_id").(string)

	nodeId := d.Id()
	node, err := epcc.Nodes.Get(&ctx, client, hierarchyId, nodeId)

	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	if err := d.Set("name", node.Data.Attributes.Name); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	if err := d.Set("slug", node.Data.Attributes.Slug); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	if err := d.Set("description", node.Data.Attributes.Description); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	if node.Data.Relationships != nil && node.Data.Relationships.Parent != nil && node.Data.Relationships.Parent.Data != nil {
		if err := d.Set("parent_id", node.Data.Relationships.Parent.Data.Id); err != nil {
			addToDiag(ctx, diag.FromErr(err))
			return
		}
	}
}

func resourceEpccNodeCreate(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	node := &epcc.Node{
		Type: "node",
		Attributes: epcc.NodeAttributes{
			Description: d.Get("description").(string),
			Name:        d.Get("name").(string),
			Slug:        d.Get("slug").(string),
		},
		Relationships: nil,
	}

	hierarchyId := d.Get("hierarchy_id").(string)

	createdNodeData, apiError := epcc.Nodes.Create(&ctx, client, hierarchyId, node)

	if apiError != nil {
		ReportAPIError(ctx, apiError)
		return
	}

	d.SetId(createdNodeData.Data.Id)

	resourceEpccNodeRead(ctx, d, m)
}
