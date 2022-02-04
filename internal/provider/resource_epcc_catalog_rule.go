package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEpccCatalogRule() *schema.Resource {
	return &schema.Resource{
		Description:   "Represents the EPCC API [*PCM* Catalog Rule Object](https://documentation.elasticpath.com/commerce-cloud/docs/api/pcm/catalogs/rules/get-a-catalog-rule.html).",
		CreateContext: addDiagToContext(resourceEpccCatalogRuleCreate),
		ReadContext:   addDiagToContext(resourceEpccCatalogRuleRead),
		UpdateContext: addDiagToContext(resourceEpccCatalogRuleUpdate),
		DeleteContext: addDiagToContext(resourceEpccCatalogRuleDelete),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier of the catalog rule.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the rule without spaces.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The purpose for this rule.",
			},
			"catalog": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique identifier of the catalog for this rule. If you want to display a catalog that contains V2 Products, Brands, Categories, and Collections, specify `legacy`",
			},
			"customers": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The list of customers who are eligible to see this catalog. If empty, the rule matches all customers.",
			},
		},
	}

}

func resourceEpccCatalogRuleDelete(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	catalogRuleID := d.Id()

	err := epcc.CatalogRules.Delete(&ctx, client, catalogRuleID)

	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	d.SetId("")
}

func resourceEpccCatalogRuleUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	catalogRuleId := d.Id()

	catalogRule := &epcc.CatalogRule{
		Id:   catalogRuleId,
		Type: "catalog_rule",
		Attributes: epcc.CatalogRulesAttributes{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
			Catalog:     d.Get("catalog").(string),
			Customers:   convertSetToStringSlice(d.Get("customers").(*schema.Set)),
		},
	}

	updatedCatalogRuleData, apiError := epcc.CatalogRules.Update(&ctx, client, catalogRuleId, catalogRule)

	if apiError != nil {
		ReportAPIError(ctx, apiError)
		return
	}

	d.SetId(updatedCatalogRuleData.Data.Id)

	resourceEpccCatalogRuleRead(ctx, d, m)
}

func resourceEpccCatalogRuleRead(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	catalogRuleId := d.Id()

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
}

func resourceEpccCatalogRuleCreate(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	catalogRule := &epcc.CatalogRule{
		Type: "catalog_rule",
		Attributes: epcc.CatalogRulesAttributes{
			Name:        d.Get("name").(string),
			Description: d.Get("description").(string),
			Catalog:     d.Get("catalog").(string),
			Customers:   convertSetToStringSlice(d.Get("customers").(*schema.Set)),
		},
	}

	createdCatalogRuleData, apiError := epcc.CatalogRules.Create(&ctx, client, catalogRule)

	if apiError != nil {
		ReportAPIError(ctx, apiError)
		return
	}

	d.SetId(createdCatalogRuleData.Data.Id)

	resourceEpccCatalogRuleRead(ctx, d, m)
}
