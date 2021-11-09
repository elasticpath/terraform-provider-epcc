package provider

import (
	"context"
	"fmt"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/elasticpath/terraform-provider-epcc/internal/provider/util"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var mutex = util.NewMutexKV()

func resourceEpccNodeProduct() *schema.Resource {
	return &schema.Resource{
		Description:   "Represents the EPCC API [Node and Product Relationship](https://documentation.elasticpath.com/commerce-cloud/docs/api/pcm/hierarchies/relationships/create-node-product-relationships.html).",
		CreateContext: addDiagToContext(resourceEpccNodeProductCreate),
		ReadContext:   addDiagToContext(resourceEpccNodeProductRead),
		DeleteContext: addDiagToContext(resourceEpccNodeProductDelete),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hierarchy_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"node_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"product_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}

}

func resourceEpccNodeProductDelete(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	hierarchyId := d.Get("hierarchy_id").(string)
	nodeId := d.Get("node_id").(string)
	productId := d.Get("product_id").(string)

	d.Get("hierarchy")

	list := convertIdsToTypeIdRelationship("product", []string{productId})

	wrappedList := epcc.DataForTypeIdRelationshipList{
		Data: &list,
	}

	// PCM seems to have issues if we concurrently call endpoint
	// Even if we try to lock based on node id
	mutex.Lock("nodes")
	err := epcc.Nodes.DeleteNodeProduct(&ctx, client, hierarchyId, nodeId, wrappedList)
	mutex.Unlock("nodes")

	if err != nil {
		ReportAPIError(ctx, err)
	}

	d.SetId("")
}

func resourceEpccNodeProductRead(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	hierarchyId := d.Get("hierarchy_id").(string)
	nodeId := d.Get("node_id").(string)
	productId := d.Get("product_id").(string)

	mutex.Lock("nodes")
	nodeProduct, err := epcc.Nodes.GetNodeProducts(&ctx, client, hierarchyId, nodeId)
	mutex.Unlock("nodes")

	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	var foundMatch = false

	var allIds = []string{}
	if nodeProduct.Data != nil {
		allIds = convertJsonTypesToIds(nodeProduct.Data)

		for _, relationship := range *nodeProduct.Data {

			if relationship.Id == productId {
				foundMatch = true
				break
			}

		}

	}

	if !foundMatch {
		addToDiag(ctx, diag.FromErr(fmt.Errorf("Could not find node product relationship for hierarchy %s node %s product %s\nAll ids: %s", hierarchyId, nodeId, productId, allIds)))
	} else {
		if err := d.Set("hierarchy_id", hierarchyId); err != nil {
			addToDiag(ctx, diag.FromErr(err))
			return
		}

		if err := d.Set("node_id", nodeId); err != nil {
			addToDiag(ctx, diag.FromErr(err))
			return
		}

		if err := d.Set("product_id", productId); err != nil {
			addToDiag(ctx, diag.FromErr(err))
			return
		}

		d.SetId(productId)

	}
}

func resourceEpccNodeProductCreate(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	hierarchyId := d.Get("hierarchy_id").(string)
	nodeId := d.Get("node_id").(string)
	productId := d.Get("product_id").(string)

	relationships := convertIdsToTypeIdRelationship("product", []string{productId})

	mutex.Lock("nodes")
	apiError := epcc.Nodes.CreateNodeProducts(&ctx, client, hierarchyId, nodeId, epcc.DataForTypeIdRelationshipList{
		Data: &relationships,
	})
	mutex.Unlock("nodes")

	if apiError != nil {
		ReportAPIError(ctx, apiError)
		return
	}

	d.SetId(productId)

	resourceEpccNodeProductRead(ctx, d, m)
}
