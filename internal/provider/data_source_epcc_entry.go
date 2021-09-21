package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type EntryDataSourceProvider struct {
}

func (p EntryDataSourceProvider) DataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: addDiagToContext(p.read),
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique identifier for this entry.",
			},
			"slug": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The slug for the Flow you are requesting an Entry for.",
			},
			"strings": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"numbers": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeFloat,
				},
			},
			"booleans": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeBool,
				},
			},
		},
	}
}

func (p EntryDataSourceProvider) read(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	flowSlug := d.Get("slug").(string)
	entryID := d.Get("id").(string)

	entry, err := epcc.Entries.Get(&ctx, client, flowSlug, entryID)
	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	if err := d.Set("strings", entry.Data.Strings); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("numbers", entry.Data.Numbers); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("booleans", entry.Data.Booleans); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	d.SetId(entryID)
}
