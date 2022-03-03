package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMerchantRealmMappings() *schema.Resource {
	return &schema.Resource{
		Description: "Represents the EPCC API [Merchant Realm Mapping](https://documentation.elasticpath.com/commerce-cloud/docs/api/advanced/settings/merchant-realm-mappings/index.html)",
		ReadContext: addDiagToContext(dataSourceEpccMerchantRealmMappingsRead),
		Schema: map[string]*schema.Schema{
			"id":       {Type: schema.TypeString, Computed: true, Description: "The ID for the requested merchant realm mapping."},
			"prefix":   {Type: schema.TypeString, Computed: true, Description: "The prefix name to associate with a store"},
			"realm_id": {Type: schema.TypeString, Computed: true, Description: "The ID of the authentication realm used to sign in as administrator."},
			"store_id": {Type: schema.TypeString, Computed: true, Description: "System-generated store ID."},
		},
	}
}

func dataSourceEpccMerchantRealmMappingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	MerchantRealmMappings, err := epcc.MerchantRealmMappings.Get(&ctx, client)

	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	d.SetId(MerchantRealmMappings.Data.ID)
	d.Set("prefix", MerchantRealmMappings.Data.Prefix)
	d.Set("realm_id", MerchantRealmMappings.Data.RealmID)
	d.Set("store_id", MerchantRealmMappings.Data.StoreID)
}
