package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEpccHierarchy() *schema.Resource {
	return &schema.Resource{
		ReadContext: addDiagToContext(dataSourceEpccHierarchyRead),
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique identifier of the hierarchy.",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the hierarchy. The maximum length is 1000 characters.",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A description of the hierarchy.",
			},
			"slug": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A unique slug for the hierarchy.",
			},
		},
	}
}

func dataSourceEpccHierarchyRead(ctx context.Context, d *schema.ResourceData, m interface{}) {

	client := m.(*epcc.Client)
	hierarchyId := d.Get("id").(string)

	hierarchy, err := epcc.Hierarchies.Get(&ctx, client, hierarchyId)

	if err != nil {
		ReportAPIError(ctx, err)
	}

	d.Set("name", hierarchy.Data.Attributes.Name)
	d.Set("type", hierarchy.Data.Type)
	d.Set("slug", hierarchy.Data.Attributes.Slug)

	d.SetId(hierarchy.Data.Id)
}
