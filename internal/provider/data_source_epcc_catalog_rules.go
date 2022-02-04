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
			"id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique identifier of the catalog rule.",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the rule without spaces.",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The purpose for this rule.",
			},
			"catalog": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier of the catalog for this rule. If you want to display a catalog that contains V2 Products, Brands, Categories, and Collections, specify `legacy`",
			},
			"customers": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The list of customers who are eligible to see this catalog. If empty, the rule matches all customers.",
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
