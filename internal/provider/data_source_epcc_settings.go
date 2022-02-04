package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEpccSettings() *schema.Resource {
	return &schema.Resource{
		ReadContext: addDiagToContext(dataSourceEpccSettingsRead),
		Schema: map[string]*schema.Schema{
			"page_length":          {Type: schema.TypeInt, Computed: true, Description: "The number of results per page when paginating results"},
			"list_child_products":  {Type: schema.TypeBool, Computed: true, Description: "Whether to display child products in product listings."},
			"additional_languages": {Type: schema.TypeList, Computed: true, Elem: &schema.Schema{Type: schema.TypeString}, Description: "You can define additional language codes that are enabled for a project, this applies only to the legacy catalog and does not apply to PCM products, hierarchies, and catalogs."},
			"calculation_method":   {Type: schema.TypeString, Computed: true, Description: "This option defines the method used to calculate cart and order totals."},
		},
	}
}

func dataSourceEpccSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	settings, err := epcc.SettingsVar.Get(&ctx, client)
	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	d.SetId("0")
	d.Set("page_length", settings.Data.PageLength)
	d.Set("list_child_products", settings.Data.ListChildProducts)
	d.Set("additional_languages", settings.Data.AdditionalLanguages)
	d.Set("calculation_method", settings.Data.CalculationMethod)
}
