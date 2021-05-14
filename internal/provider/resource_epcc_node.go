package provider

import (
	"context"
	"github.com/elasticpath/epcc-terraform-provider/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEpccNode() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceEpccNodeCreate,
		ReadContext:   resourceEpccNodeRead,
		UpdateContext: resourceEpccNodeUpdate,
		DeleteContext: resourceEpccNodeDelete,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"hierarchy_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"slug": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Computed: false,
				Optional: true,
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

func resourceEpccNodeDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	var diags diag.Diagnostics

	err := epcc.Nodes.Delete(client, d.Get("hierarchy_id").(string), d.Id())

	if err != nil {
		FromAPIError(err)
	}

	d.SetId("")

	return diags
}

func resourceEpccNodeUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	hierarchyId :=  d.Get("hierarchy_id").(string)

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

	updatedNodeData, apiError := epcc.Nodes.Update(client, hierarchyId, d.Id(), node)

	if apiError != nil {
		return FromAPIError(apiError)
	}

	newProducts := convertIdsToTypeIdRelationship("product", convertSetToStringSlice(d.Get("products").(*schema.Set)))

	// Update Node Products Updates All the Products on the Node
	apiError = epcc.Nodes.UpdateNodeProducts(client, hierarchyId, d.Id(), epcc.DataForTypeIdRelationshipList{Data: &newProducts})

	d.SetId(updatedNodeData.Data.Id)

	return resourceEpccNodeRead(ctx, d, m)
}

func resourceEpccNodeRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	var diags diag.Diagnostics

	hierarchyId := d.Get("hierarchy_id").(string)

	nodeId := d.Id()
	node, err := epcc.Nodes.Get(client, hierarchyId, nodeId)

	if err != nil {
		return FromAPIError(err)
	}

	nodeProducts, err := epcc.Nodes.GetNodeProducts(client, hierarchyId, nodeId)

	if err != nil {
		return FromAPIError(err)
	}

	if err := d.Set("name", node.Data.Attributes.Name); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("slug", node.Data.Attributes.Slug); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("description", node.Data.Attributes.Description); err != nil {
		return diag.FromErr(err)
	}

	if node.Data.Relationships != nil && node.Data.Relationships.Parent != nil && node.Data.Relationships.Parent.Data != nil {
		if err := d.Set("parent_id", node.Data.Relationships.Parent.Data.Id); err != nil {
			return diag.FromErr(err)
		}
	}

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

	return diags
}

func resourceEpccNodeCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	var diags diag.Diagnostics

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

	createdNodeData, apiError := epcc.Nodes.Create(client, hierarchyId, node)

	if apiError != nil {
		return FromAPIError(apiError)
	}

	d.SetId(createdNodeData.Data.Id)


	files := d.Get("products").(*schema.Set)

	relationships := convertIdsToTypeIdRelationship("product", convertSetToStringSlice(files))

	if len(relationships) > 0 {
		apiError = epcc.Nodes.CreateNodeProducts(client, hierarchyId, createdNodeData.Data.Id, epcc.DataForTypeIdRelationshipList{
			Data: &relationships,
		})

		if apiError != nil {
			return FromAPIError(apiError)
		}
	}

	resourceEpccNodeRead(ctx, d, m)

	return diags
}
