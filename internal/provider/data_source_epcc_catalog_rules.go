package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEpccCatalogRule() *schema.Resource {
	return &schema.Resource{
		ReadContext: addDiagToContext(dataSourceEpccCatalogRuleRead),
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
			"catalog": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"customers": &schema.Schema{
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceEpccCatalogRuleRead(ctx context.Context, d *schema.ResourceData, m interface{}) {

	client := m.(*epcc.Client)
	catalogRuleId := d.Get("id").(string)
	catalogRule, err := epcc.CatalogRules.Get(&ctx, client, catalogRuleId)

	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	if err := d.Set("name", catalogRule.Data.Attributes.Name); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	if err := d.Set("description", catalogRule.Data.Attributes.Description); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	if err := d.Set("catalog", catalogRule.Data.Attributes.Catalog); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	if err := d.Set("customers", catalogRule.Data.Attributes.Customers); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	d.SetId(catalogRule.Data.Id)
}
