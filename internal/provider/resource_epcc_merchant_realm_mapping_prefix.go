package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEpccMerchantRealmMappingPrefix() *schema.Resource {
	return &schema.Resource{
		Description:   "Represents the EPCC API [OpenID Connect MerchantRealmMappingPrefixes](https://documentation.elasticpath.com/commerce-cloud/docs/api/single-sign-on/oidc-merchantRealmMappingPrefixes/index.html).",
		CreateContext: addDiagToContext(resourceEpccMerchantRealmMappingPrefixUpdate),
		ReadContext:   addDiagToContext(resourceEpccMerchantRealmMappingPrefixRead),
		UpdateContext: addDiagToContext(resourceEpccMerchantRealmMappingPrefixUpdate),
		DeleteContext: addDiagToContext(resourceEpccMerchantRealmMappingPrefixDelete),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id":                        {Type: schema.TypeString, Computed: true},
			"merchant_realm_mapping_id": {Type: schema.TypeString, Required: true},
			"prefix":                    {Type: schema.TypeString, Required: true},
		},
	}
}

func resourceEpccMerchantRealmMappingPrefixDelete(ctx context.Context, d *schema.ResourceData, m interface{}) {
	// no-op
}

func resourceEpccMerchantRealmMappingPrefixUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	merchantRealmMappingId := d.Get("merchant_realm_mapping_id").(string)
	get := d.Get("prefix").(string)
	updatedMerchantRealmMappingPrefixData, apiError := epcc.MerchantRealmMappingPrefixes.Update(&ctx, client, &merchantRealmMappingId, &get)
	if apiError != nil {
		ReportAPIError(ctx, apiError)
		return
	}
	d.SetId(updatedMerchantRealmMappingPrefixData.Data.Id)
	resourceEpccMerchantRealmMappingPrefixRead(ctx, d, m)
}

func resourceEpccMerchantRealmMappingPrefixRead(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	merchantRealmMappingPrefix, err := epcc.MerchantRealmMappingPrefixes.Get(&ctx, client)

	if err != nil {
		ReportAPIError(ctx, err)
		return
	}
	if err := d.Set("prefix", merchantRealmMappingPrefix.Data.Prefix); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("merchant_realm_mapping_id", merchantRealmMappingPrefix.Data.Id); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	d.SetId(merchantRealmMappingPrefix.Data.Id)
}
