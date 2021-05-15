package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEpccHierarchy() *schema.Resource {
	return &schema.Resource{
		ReadContext: addDiagToContext(dataSourceEpccHierarchyRead),
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
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
		},
	}
}

func dataSourceEpccHierarchyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*epcc.Client)
	hierarchyId := d.Get("id").(string)

	hierarchy, err := epcc.Hierarchies.Get(&ctx, client, hierarchyId)

	if err != nil {
		return FromAPIError(err)
	}

	d.Set("name", hierarchy.Data.Attributes.Name)
	d.Set("type", hierarchy.Data.Type)
	d.Set("slug", hierarchy.Data.Attributes.Slug)

	d.SetId(hierarchy.Data.Id)

	return *ctx.Value("diags").(*diag.Diagnostics)
}
