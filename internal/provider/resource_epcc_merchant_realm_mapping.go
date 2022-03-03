package provider

import (
	"context"
	"fmt"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"time"
)

func resourceEpccMerchantRealmMapping() *schema.Resource {
	return &schema.Resource{
		Description:   "Represents the EPCC API [Merchant Realm Mapping](https://documentation.elasticpath.com/commerce-cloud/docs/api/advanced/settings/merchant-realm-mappings/index.html).",
		CreateContext: addDiagToContext(resourceEpccMerchantRealmMappingUpdate),
		ReadContext:   addDiagToContext(resourceEpccMerchantRealmMappingRead),
		UpdateContext: addDiagToContext(resourceEpccMerchantRealmMappingUpdate),
		DeleteContext: addDiagToContext(resourceEpccMerchantRealmMappingDelete),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id":       {Type: schema.TypeString, Computed: true, Description: "The ID for the requested merchant realm mapping."},
			"prefix":   {Type: schema.TypeString, Required: true, Description: "The prefix name to associate with a store."},
			"store_id": {Type: schema.TypeString, Computed: true, Description: "System-generated store ID."},
			"realm_id": {Type: schema.TypeString, Computed: true, Description: "The ID of the authentication realm used to sign in as administrator."},
		},
	}
}

func resourceEpccMerchantRealmMappingDelete(ctx context.Context, d *schema.ResourceData, m interface{}) {
	// no-op
	client := m.(*epcc.Client)

	id := d.Id()
	prefix := fmt.Sprintf("deletedat%d", time.Now().UnixNano())
	_, apiError := epcc.MerchantRealmMappings.Update(&ctx, client, &id, &prefix)

	if apiError != nil {
		ReportAPIError(ctx, apiError)
		return
	}
	d.SetId("")
}

func resourceEpccMerchantRealmMappingUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	// We read before so we can get the ID.
	// This is only needed for create, but we use the same function for create an update
	merchantRealmMappingPrefix, err := epcc.MerchantRealmMappings.Get(&ctx, client)

	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	merchantRealmMappingId := merchantRealmMappingPrefix.Data.ID

	get := d.Get("prefix").(string)
	updatedMerchantRealmMappingPrefixData, apiError := epcc.MerchantRealmMappings.Update(&ctx, client, &merchantRealmMappingId, &get)
	if apiError != nil {
		ReportAPIError(ctx, apiError)
		return
	}
	d.SetId(updatedMerchantRealmMappingPrefixData.Data.ID)
	resourceEpccMerchantRealmMappingRead(ctx, d, m)
}

func resourceEpccMerchantRealmMappingRead(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	merchantRealmMappingPrefix, err := epcc.MerchantRealmMappings.Get(&ctx, client)

	if err != nil {
		ReportAPIError(ctx, err)
		return
	}
	if err := d.Set("prefix", merchantRealmMappingPrefix.Data.Prefix); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	if err := d.Set("store_id", merchantRealmMappingPrefix.Data.StoreID); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	if err := d.Set("realm_id", merchantRealmMappingPrefix.Data.RealmID); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	d.SetId(merchantRealmMappingPrefix.Data.ID)
}
