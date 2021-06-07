package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMerchantRealmMappings() *schema.Resource {
	return &schema.Resource{
		Description: "Represents the EPCC API [Merchant Realm Mappings] (https://documentation.elasticpath.com/commerce-cloud/docs/api/advanced/settings/merchant-authentication-settings/get-merchant-realm-mapping.html)",
		ReadContext: addDiagToContext(dataSourceEpccMerchantRealmMappingsRead),
		Schema: map[string]*schema.Schema{
			"merchant_realm_mapping_id": {Type: schema.TypeString, Computed: true},
			"prefix":                    {Type: schema.TypeString, Computed: true},
			"realm_id":                  {Type: schema.TypeString, Computed: true},
			"store_id":                  {Type: schema.TypeString, Computed: true},
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

	d.SetId("singleton")
	d.Set("merchant_realm_mapping_id", MerchantRealmMappings.Data.ID)
	d.Set("prefix", MerchantRealmMappings.Data.Prefix)
	d.Set("realm_id", MerchantRealmMappings.Data.RealmID)
	d.Set("store_id", MerchantRealmMappings.Data.StoreID)
}
