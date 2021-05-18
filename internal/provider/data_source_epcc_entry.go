package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEpccEntry() *schema.Resource {
	return &schema.Resource{
		ReadContext: addDiagToContext(dataSourceEpccEntryRead),
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"payload": &schema.Schema{
				Type:     schema.TypeMap,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceEpccEntryRead(ctx context.Context, d *schema.ResourceData, m interface{}) {

	client := m.(*epcc.Client)
	entryId := d.Get("id").(string)
	flowSlug := d.Get("slug").(string)
	entry, err := epcc.Entries.Get(&ctx, client, flowSlug, entryId)
	if err != nil {
		ReportAPIError(ctx, err)
	} else {
		d.SetId(entry.Data.Id)
	}
}
