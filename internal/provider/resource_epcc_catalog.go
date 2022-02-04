package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEpccCatalog() *schema.Resource {
	return &schema.Resource{
		Description:   "Represents the EPCC API [*PCM* Catalog Object](https://documentation.elasticpath.com/commerce-cloud/docs/api/pcm/catalogs/index.html#the-catalog-object).",
		CreateContext: addDiagToContext(resourceEpccCatalogCreate),
		ReadContext:   addDiagToContext(resourceEpccCatalogRead),
		UpdateContext: addDiagToContext(resourceEpccCatalogUpdate),
		DeleteContext: addDiagToContext(resourceEpccCatalogDelete),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of the catalog.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the catalog.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: " A description of the catalog, such as the purpose for the catalog.",
			},
			"hierarchies": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The unique identifiers of the hierarchies to associate with this catalog.",
			},
			"pricebook": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The unique identifier of the price book to associate with this catalog.",
			},
		},
	}

}

func resourceEpccCatalogDelete(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	catalogID := d.Id()

	err := epcc.Catalogs.Delete(&ctx, client, catalogID)

	if err != nil {
		ReportAPIError(ctx, err)
	}

	d.SetId("")
}

func resourceEpccCatalogUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	catalogId := d.Id()

	catalog := &epcc.Catalog{
		Id:   catalogId,
		Type: "catalog",
		Attributes: epcc.CatalogAttributes{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
			Hierarchies: convertSetToStringSlice(d.Get("hierarchies").(*schema.Set)),
			PriceBook:   d.Get("pricebook").(string),
		},
	}

	updatedCatalogData, apiError := epcc.Catalogs.Update(&ctx, client, catalogId, catalog)

	if apiError != nil {
		ReportAPIError(ctx, apiError)
		return
	}

	d.SetId(updatedCatalogData.Data.Id)

	resourceEpccCatalogRead(ctx, d, m)
}

func resourceEpccCatalogRead(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	catalogId := d.Id()

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
}

func resourceEpccCatalogCreate(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	catalog := &epcc.Catalog{
		Type: "catalog",
		Attributes: epcc.CatalogAttributes{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
			PriceBook:   d.Get("pricebook").(string),
			Hierarchies: convertSetToStringSlice(d.Get("hierarchies").(*schema.Set)),
		},
	}

	createdCatalogData, apiError := epcc.Catalogs.Create(&ctx, client, catalog)

	if apiError != nil {
		ReportAPIError(ctx, apiError)
		return
	}

	d.SetId(createdCatalogData.Data.Id)

	resourceEpccCatalogRead(ctx, d, m)
}
