package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEpccSettings() *schema.Resource {
	return &schema.Resource{
		Description:   "Represents the EPCC API [Settings](https://documentation.elasticpath.com/commerce-cloud/docs/api/advanced/settings/index.html).\nNote: The `epcc_settings` resource behaves different from normal resources, in that Terraform does not *create* this reosurce, but instead \"adopts\" it into management.",
		CreateContext: addDiagToContext(resourceEpccSettingsUpdate),
		ReadContext:   addDiagToContext(resourceEpccSettingsRead),
		UpdateContext: addDiagToContext(resourceEpccSettingsUpdate),
		DeleteContext: addDiagToContext(resourceEpccSettingsDelete),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id":                   {Type: schema.TypeString, Computed: true},
			"page_length":          {Type: schema.TypeInt, Optional: true, Description: "The number of results per page when paginating results."},
			"list_child_products":  {Type: schema.TypeBool, Optional: true, Description: "Whether to display child products in product listings."},
			"additional_languages": {Type: schema.TypeList, Optional: true, Elem: &schema.Schema{Type: schema.TypeString}, Description: "You can define additional language codes that are enabled for a project, this applies only to the legacy catalog and does not apply to PCM products, hierarchies, and catalogs."},
			"calculation_method":   {Type: schema.TypeString, Optional: true, Description: "This option defines the method used to calculate cart and order totals."},
		},
	}
}

func resourceEpccSettingsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) {
	//no-op
}

func resourceEpccSettingsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	settings := &epcc.Settings{
		Id:                  "0",
		Type:                "settings",
		PageLength:          d.Get("page_length").(int),
		ListChildProducts:   d.Get("list_child_products").(bool),
		AdditionalLanguages: d.Get("additional_languages").([]interface{}),
		CalculationMethod:   d.Get("calculation_method").(string),
	}

	_, apiError := epcc.SettingsVar.Update(&ctx, client, *settings)
	d.Set("id", "0")
	d.SetId("0")
	if apiError != nil {
		ReportAPIError(ctx, apiError)
		return
	}
}

func resourceEpccSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	settings, err := epcc.SettingsVar.Get(&ctx, client)

	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	d.SetId("0")

	if err := d.Set("id", settings.Data.Id); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("additional_languages", settings.Data.AdditionalLanguages); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("calculation_method", settings.Data.CalculationMethod); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("list_child_products", settings.Data.ListChildProducts); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("page_length", settings.Data.PageLength); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
}
