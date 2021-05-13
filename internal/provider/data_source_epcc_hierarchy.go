package provider

import (
	"context"
	"github.com/elasticpath/epcc-terraform-provider/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEpccHierarchy() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceEpccHierarchyRead,
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

	var diags diag.Diagnostics

	hierarchyId := d.Get("id").(string)

	hierarchy, err := epcc.Hierarchies.Get(client, hierarchyId)

	if err != nil {
		return FromAPIError(err)
	}

	d.Set("name", hierarchy.Data.Attributes.Name)
	d.Set("type", hierarchy.Data.Type)
	d.Set("slug", hierarchy.Data.Attributes.Slug)

	d.SetId(hierarchy.Data.Id)

	return diags
}
