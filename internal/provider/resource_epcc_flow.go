package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEpccFlow() *schema.Resource {
	return &schema.Resource{
		Description:   "Represents the EPCC API [Flow Object](https://documentation.elasticpath.com/commerce-cloud/docs/api/advanced/custom-data/flows/index.html#the-flow-object).",
		CreateContext: addDiagToContext(resourceEpccFlowCreate),
		ReadContext:   addDiagToContext(resourceEpccFlowRead),
		UpdateContext: addDiagToContext(resourceEpccFlowUpdate),
		DeleteContext: addDiagToContext(resourceEpccFlowDelete),
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
			"slug": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
		},
	}

}

func resourceEpccFlowDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)
	flowID := d.Id()

	err := epcc.Flows.Delete(&ctx, client, flowID)

	if err != nil {
		FromAPIError(err)
	}

	d.SetId("")

	return *ctx.Value("diags").(*diag.Diagnostics)
}

func resourceEpccFlowUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	flowId := d.Id()

	flow := &epcc.Flow{
		Type:        "flow",
		Name:        d.Get("name").(string),
		Slug:        d.Get("slug").(string),
		Description: d.Get("description").(string),
		Enabled:     d.Get("enabled").(bool),
	}

	createdFlowData, apiError := epcc.Flows.Update(&ctx, client, flowId, flow)

	if apiError != nil {
		return FromAPIError(apiError)
	}

	d.SetId(createdFlowData.Data.Id)

	return resourceEpccFlowRead(ctx, d, m)
}

func resourceEpccFlowRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)
	flowID := d.Id()

	flow, err := epcc.Flows.Get(&ctx, client, flowID)

	if err != nil {
		return FromAPIError(err)
	}

	if err := d.Set("name", flow.Data.Name); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("slug", flow.Data.Slug); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("description", flow.Data.Description); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("enabled", flow.Data.Enabled); err != nil {
		return diag.FromErr(err)
	}

	return *ctx.Value("diags").(*diag.Diagnostics)
}

func resourceEpccFlowCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)
	flow := &epcc.Flow{
		Type:        "flow",
		Name:        d.Get("name").(string),
		Slug:        d.Get("slug").(string),
		Description: d.Get("description").(string),
		Enabled:     d.Get("enabled").(bool),
	}

	createdFlowData, apiError := epcc.Flows.Create(&ctx, client, flow)

	if apiError != nil {
		return FromAPIError(apiError)
	}

	d.SetId(createdFlowData.Data.Id)

	resourceEpccFlowRead(ctx, d, m)

	return *ctx.Value("diags").(*diag.Diagnostics)
}
