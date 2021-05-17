package provider

import (
	"context"
	"fmt"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEpccNodeProduct() *schema.Resource {
	return &schema.Resource{
		ReadContext: addDiagToContext(dataSourceEpccNodeProductRead),
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"hierarchy_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"node_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"product_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func dataSourceEpccNodeProductRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*epcc.Client)

	hierarchyId := d.Get("hierarchy_id").(string)
	nodeId := d.Get("node_id").(string)
	productId := d.Get("product_id").(string)

	nodeProduct, err := epcc.Nodes.GetNodeProducts(&ctx, client, hierarchyId, nodeId)

	if err != nil {
		return FromAPIError(err)
	}

	var foundMatch = false
	if nodeProduct.Data != nil {

		for _, relationship := range *nodeProduct.Data {

			if relationship.Id == productId {
				foundMatch = true
				break
			}
		}
	}

	if !foundMatch {
		return diag.FromErr(fmt.Errorf("Could not find node product relationship for hierarchy %s node %s product %s", hierarchyId, nodeId, productId))
	} else {
		if err := d.Set("hierarchy_id", hierarchyId); err != nil {
			return diag.FromErr(err)
		}

		if err := d.Set("node_id", nodeId); err != nil {
			return diag.FromErr(err)
		}

		if err := d.Set("product_id", productId); err != nil {
			return diag.FromErr(err)
		}

		d.SetId(productId)

		return *ctx.Value("diags").(*diag.Diagnostics)
	}

	return *ctx.Value("diags").(*diag.Diagnostics)
}
