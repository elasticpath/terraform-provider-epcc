package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEpccHierarchy() *schema.Resource {
	return &schema.Resource{
		Description:   "Represents the EPCC API [Hierarchy Object](https://documentation.elasticpath.com/commerce-cloud/docs/api/pcm/hierarchies/index.html#the-hierarchy-object).",
		CreateContext: addDiagToContext(resourceEpccHierarchyCreate),
		ReadContext:   addDiagToContext(resourceEpccHierarchyRead),
		UpdateContext: addDiagToContext(resourceEpccHierarchyUpdate),
		DeleteContext: addDiagToContext(resourceEpccHierarchyDelete),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				Computed: false,
			},
			"description": {
				Type:     schema.TypeString,
				Required: false,
				Computed: false,
				Optional: true,
			},
			"slug": {
				Type:     schema.TypeString,
				Required: false,
				Computed: false,
				Optional: true,
			},
		},
	}

}

func resourceEpccHierarchyDelete(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	hierarchyID := d.Id()

	err := epcc.Hierarchies.Delete(&ctx, client, hierarchyID)

	if err != nil {
		ReportAPIError(ctx, err)
	}

	d.SetId("")
}

func resourceEpccHierarchyUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) {
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
		ReportAPIError(ctx, apiError)
		return
	}

	d.SetId(updatedHierarchyData.Data.Id)

	resourceEpccHierarchyRead(ctx, d, m)
}

func resourceEpccHierarchyRead(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	hierarchyId := d.Id()

	hierarchy, err := epcc.Hierarchies.Get(&ctx, client, hierarchyId)

	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	if err := d.Set("name", hierarchy.Data.Attributes.Name); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	if err := d.Set("slug", hierarchy.Data.Attributes.Slug); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	if err := d.Set("description", hierarchy.Data.Attributes.Description); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
}

func resourceEpccHierarchyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) {
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
		ReportAPIError(ctx, apiError)
		return
	}

	d.SetId(createdHierarchyData.Data.Id)

	resourceEpccHierarchyRead(ctx, d, m)
}
