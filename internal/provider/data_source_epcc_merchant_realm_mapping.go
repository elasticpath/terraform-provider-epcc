package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEpccMerchantRealmMapping() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceEpccMerchantRealmMappingRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"prefix": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
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
				Computed: true,
			},
		},
	}
}

func dataSourceEpccMerchantRealmMappingRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*epcc.Client)
	var diags diag.Diagnostics

	merchantRealmMappingId := d.Get("id").(string)

	merchantRealmMapping, err := epcc.MerchantRealmMappings.Get(&ctx, client, merchantRealmMappingId)

	if err != nil {
		return FromAPIError(err)
	}

	d.Set("prefix", merchantRealmMapping.Data.Prefix)
	d.Set("realm_id", merchantRealmMapping.Data.RealmId)
	d.Set("store_id", merchantRealmMapping.Data.StoreId)
	d.Set("type", merchantRealmMapping.Data.Type)

	d.SetId(merchantRealmMapping.Data.Id)

	return diags
}
