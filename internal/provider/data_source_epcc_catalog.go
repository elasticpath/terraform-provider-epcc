package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEpccCatalog() *schema.Resource {
	return &schema.Resource{
		ReadContext: addDiagToContext(dataSourceEpccCatalogRead),
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of the catalog.",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the catalog.",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: " A description of the catalog, such as the purpose for the catalog.",
			},
			"hierarchies": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The unique identifiers of the hierarchies to associate with this catalog.",
			},
			"pricebook": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier of the price book to associate with this catalog.",
			},
		},
	}
}

func dataSourceEpccCatalogRead(ctx context.Context, d *schema.ResourceData, m interface{}) {

	client := m.(*epcc.Client)
	catalogId := d.Get("id").(string)
	catalog, err := epcc.Catalogs.Get(&ctx, client, catalogId)

	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	if err := d.Set("name", catalog.Data.Attributes.Name); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	if err := d.Set("description", catalog.Data.Attributes.Description); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	if err := d.Set("pricebook", catalog.Data.Attributes.PriceBook); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	if err := d.Set("hierarchies", catalog.Data.Attributes.Hierarchies); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	d.SetId(catalog.Data.Id)
}
