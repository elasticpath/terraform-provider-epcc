package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEpccMerchantRealmMapping() *schema.Resource {
	return &schema.Resource{
		Description:   "Represents the EPCC API [MerchantRealmMapping Object](https://documentation.elasticpath.com/commerce-cloud/docs/api/advanced/settings/merchant-authentication-settings/index.html).",
		CreateContext: addDiagToContext(resourceEpccMerchantRealmMappingCreate),
		ReadContext:   addDiagToContext(resourceEpccMerchantRealmMappingRead),
		UpdateContext: addDiagToContext(resourceEpccMerchantRealmMappingUpdate),
		DeleteContext: addDiagToContext(resourceEpccMerchantRealmMappingDelete),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"prefix": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"realm_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"store_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}

}

func resourceEpccMerchantRealmMappingDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)
	merchantRealmMappingID := d.Id()

	err := epcc.MerchantRealmMappings.Delete(&ctx, client, merchantRealmMappingID)

	if err != nil {
		FromAPIError(err)
	}

	d.SetId("")

	return *ctx.Value("diags").(*diag.Diagnostics)
}

func resourceEpccMerchantRealmMappingUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	merchantRealmMappingId := d.Id()

	merchantRealmMapping := &epcc.MerchantRealmMapping{
		Prefix:  d.Get("prefix").(string),
		RealmId: d.Get("realm_id").(string),
		StoreId: d.Get("store_id").(string),
		Type:    "merchant-realm-mappings",
	}

	createdMerchantRealmMappingData, apiError := epcc.MerchantRealmMappings.Update(&ctx, client, merchantRealmMappingId, merchantRealmMapping)

	if apiError != nil {
		return FromAPIError(apiError)
	}

	d.SetId(createdMerchantRealmMappingData.Data.Id)

	return resourceEpccMerchantRealmMappingRead(ctx, d, m)
}

func resourceEpccMerchantRealmMappingRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)
	merchantRealmMappingID := d.Id()

	merchantRealmMapping, err := epcc.MerchantRealmMappings.Get(&ctx, client, merchantRealmMappingID)

	if err != nil {
		return FromAPIError(err)
	}

	if err := d.Set("prefix", merchantRealmMapping.Data.Prefix); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("realm_id", merchantRealmMapping.Data.RealmId); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("store_id", merchantRealmMapping.Data.StoreId); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("type", merchantRealmMapping.Data.Type); err != nil {
		return diag.FromErr(err)
	}

	return *ctx.Value("diags").(*diag.Diagnostics)
}

func resourceEpccMerchantRealmMappingCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)
	merchantRealmMapping := &epcc.MerchantRealmMapping{
		Prefix:  d.Get("prefix").(string),
		RealmId: d.Get("realm_id").(string),
		StoreId: d.Get("store_id").(string),
		Type:    "merchant-realm-mappings",
	}

	createdMerchantRealmMappingData, apiError := epcc.MerchantRealmMappings.Create(&ctx, client, merchantRealmMapping)

	if apiError != nil {
		return FromAPIError(apiError)
	}

	d.SetId(createdMerchantRealmMappingData.Data.Id)

	resourceEpccMerchantRealmMappingRead(ctx, d, m)

	return *ctx.Value("diags").(*diag.Diagnostics)
}
