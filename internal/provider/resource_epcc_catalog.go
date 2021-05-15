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
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hierarchies": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"pricebook": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}

}

func resourceEpccCatalogDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)
	catalogID := d.Id()

	err := epcc.Catalogs.Delete(&ctx, client, catalogID)

	if err != nil {
		FromAPIError(err)
	}

	d.SetId("")

	return *ctx.Value("diags").(*diag.Diagnostics)
}

func resourceEpccCatalogUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		return FromAPIError(apiError)
	}

	d.SetId(updatedCatalogData.Data.Id)

	return resourceEpccCatalogRead(ctx, d, m)
}

func resourceEpccCatalogRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)
	catalogId := d.Id()

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

	return *ctx.Value("diags").(*diag.Diagnostics)
}

func resourceEpccCatalogCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		return FromAPIError(apiError)
	}

	d.SetId(createdCatalogData.Data.Id)

	resourceEpccCatalogRead(ctx, d, m)

	return *ctx.Value("diags").(*diag.Diagnostics)
}
