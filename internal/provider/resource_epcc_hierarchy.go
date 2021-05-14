package provider

import (
	"context"
	"github.com/elasticpath/epcc-terraform-provider/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEpccHierarchy() *schema.Resource {
	return &schema.Resource{
		CreateContext: addDiagToContext(resourceEpccHierarchyCreate),
		ReadContext:   addDiagToContext(resourceEpccHierarchyRead),
		UpdateContext: addDiagToContext(resourceEpccHierarchyUpdate),
		DeleteContext: addDiagToContext(resourceEpccHierarchyDelete),
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
				Computed: false,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Computed: false,
				Optional: true,
			},
			"slug": &schema.Schema{
				Type:     schema.TypeString,
				Required: false,
				Computed: false,
				Optional: true,
			},
		},
	}

}

func resourceEpccHierarchyDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)
	hierarchyID := d.Id()

	err := epcc.Hierarchies.Delete(&ctx, client, hierarchyID)

	if err != nil {
		FromAPIError(err)
	}

	d.SetId("")

	return *ctx.Value("diags").(*diag.Diagnostics)
}

func resourceEpccHierarchyUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	hierarchyId := d.Id()

	hierarchy := &epcc.Hierarchy{
		Type: "hierarchy",
		Id:   hierarchyId,
		Attributes: epcc.HierarchyAttributes{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
			Slug:        d.Get("slug").(string),
		},
	}

	updatedHierarchyData, apiError := epcc.Hierarchies.Update(&ctx, client, hierarchyId, hierarchy)

	if apiError != nil {
		return FromAPIError(apiError)
	}

	d.SetId(updatedHierarchyData.Data.Id)

	return resourceEpccHierarchyRead(ctx, d, m)
}

func resourceEpccHierarchyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)
	hierarchyId := d.Id()

	hierarchy, err := epcc.Hierarchies.Get(&ctx, client, hierarchyId)

	if err != nil {
		return FromAPIError(err)
	}

	if err := d.Set("name", hierarchy.Data.Attributes.Name); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("slug", hierarchy.Data.Attributes.Slug); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("description", hierarchy.Data.Attributes.Description); err != nil {
		return diag.FromErr(err)
	}

	return *ctx.Value("diags").(*diag.Diagnostics)
}

func resourceEpccHierarchyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)
	hierarchy := &epcc.Hierarchy{
		Type: "hierarchy",
		Attributes: epcc.HierarchyAttributes{
			Description: d.Get("description").(string),
			Name:        d.Get("name").(string),
			Slug:        d.Get("slug").(string),
		},
	}

	createdHierarchyData, apiError := epcc.Hierarchies.Create(&ctx, client, hierarchy)

	if apiError != nil {
		return FromAPIError(apiError)
	}

	d.SetId(createdHierarchyData.Data.Id)

	resourceEpccHierarchyRead(ctx, d, m)

	return *ctx.Value("diags").(*diag.Diagnostics)
}
