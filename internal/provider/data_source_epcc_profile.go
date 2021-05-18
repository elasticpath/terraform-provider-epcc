package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEpccProfile() *schema.Resource {
	return &schema.Resource{
		ReadContext: addDiagToContext(dataSourceEpccProfileRead),
		Schema: map[string]*schema.Schema{
			"id":            {Type: schema.TypeString, Required: true},
			"name":          {Type: schema.TypeString, Computed: true},
			"discovery_url": {Type: schema.TypeString, Computed: true},
			"client_id":     {Type: schema.TypeString, Computed: true},
			"client_secret": {Type: schema.TypeString, Computed: true},
			"realm_id":      {Type: schema.TypeString, Computed: true},
		},
	}
}

func dataSourceEpccProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	profileId := d.Get("id").(string)
	realmId := d.Get("realm_id").(string)
	profile, err := epcc.Profiles.Get(&ctx, client, realmId, profileId)
	if err != nil {
		return FromAPIError(err)
	}

	d.Set("id", profile.Data.Id)
	d.Set("type", profile.Data.Type)
	d.Set("name", profile.Data.Name)
	d.Set("discovery_url", profile.Data.DiscoveryUrl)
	d.Set("client_id", profile.Data.ClientID)
	d.Set("client_secret", profile.Data.ClientSecret)
	d.SetId(profile.Data.Id)

	return *ctx.Value("diags").(*diag.Diagnostics)
}
