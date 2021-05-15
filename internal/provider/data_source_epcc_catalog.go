package provider

import (
	"context"
	"github.com/elasticpath/epcc-terraform-provider/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEpccCatalog() *schema.Resource {
	return &schema.Resource{
		ReadContext: addDiagToContext(dataSourceEpccCatalogRead),
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
			"hierarchies": &schema.Schema{
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"pricebook": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceEpccCatalogRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*epcc.Client)
	catalogId := d.Get("id").(string)
	catalog, err := epcc.Catalogs.Get(&ctx, client, catalogId)

	if err != nil {
		return FromAPIError(err)
	}

	if err := d.Set("name", catalog.Data.Attributes.Name); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("description", catalog.Data.Attributes.Description); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("pricebook", catalog.Data.Attributes.PriceBook); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("hierarchies", catalog.Data.Attributes.Hierarchies); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(catalog.Data.Id)

	return *ctx.Value("diags").(*diag.Diagnostics)
}
