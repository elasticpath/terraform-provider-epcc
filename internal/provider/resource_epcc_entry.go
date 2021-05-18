package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEpccEntry() *schema.Resource {
	return &schema.Resource{
		Description:   "Represents the EPCC API [Entry Object](https://documentation.elasticpath.com/commerce-cloud/docs/api/advanced/custom-data/entries/index.html).",
		CreateContext: addDiagToContext(resourceEpccEntryCreate),
		ReadContext:   addDiagToContext(resourceEpccEntryRead),
		UpdateContext: addDiagToContext(resourceEpccEntryUpdate),
		DeleteContext: addDiagToContext(resourceEpccEntryDelete),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"slug": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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

func resourceEpccEntryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	flowSlug := d.Get("slug").(string)
	entryID := d.Id()

	err := epcc.Entries.Delete(&ctx, client, flowSlug, entryID)

	if err != nil {
		FromAPIError(err)
	}

	d.SetId("")

	return *ctx.Value("diags").(*diag.Diagnostics)
}

func resourceEpccEntryUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	flowSlug := d.Get("slug").(string)
	entryID := d.Id()

	createdEntryData, apiError := epcc.Entries.Update(&ctx, client, flowSlug, entryID, d.Get("payload").(map[string]interface{}))

	if apiError != nil {
		return FromAPIError(apiError)
	}

	d.SetId(createdEntryData.Data.Id)

	return resourceEpccEntryRead(ctx, d, m)
}

func resourceEpccEntryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)
	entryID := d.Id()

	flowSlug := d.Get("slug").(string)
	_, err := epcc.Entries.Get(&ctx, client, flowSlug, entryID)

	if err != nil {
		return FromAPIError(err)
	}

	return *ctx.Value("diags").(*diag.Diagnostics)
}

func resourceEpccEntryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	entry := &epcc.Entry{
		Type: "entry",
	}

	flowSlug := d.Get("slug").(string)

	createdEntryData, apiError := epcc.Entries.Create(&ctx, client, flowSlug, entry, d.Get("payload").(map[string]interface{}))

	if apiError != nil {
		return FromAPIError(apiError)
	}

	d.SetId(createdEntryData.Data.Id)

	resourceEpccEntryRead(ctx, d, m)

	return *ctx.Value("diags").(*diag.Diagnostics)
}
