package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEpccOidcProfile() *schema.Resource {
	return &schema.Resource{
		ReadContext: addDiagToContext(dataSourceEpccOidcProfileRead),
		Schema: map[string]*schema.Schema{
			"id":            {Type: schema.TypeString, Required: true, Description: "The unique identifier for this OpenID Connect profile."},
			"name":          {Type: schema.TypeString, Computed: true, Description: "The name of the OpenID Connect profile."},
			"discovery_url": {Type: schema.TypeString, Computed: true, Description: "The url of the OpenID Connect discovery document."},
			"client_id":     {Type: schema.TypeString, Computed: true, Description: "The client id to be used with the external authentication provider."},
			"client_secret": {Type: schema.TypeString, Computed: true, Description: "The client secret for the OpenID Provider."},
			"realm_id":      {Type: schema.TypeString, Computed: true, Description: "The ID of the authentication-realm containing the OpenID Connect profiles."},
		},
	}
}

func dataSourceEpccOidcProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	profileId := d.Get("id").(string)
	realmId := d.Get("realm_id").(string)
	oidcProfile, err := epcc.OidcProfiles.Get(&ctx, client, realmId, profileId)
	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	d.Set("id", oidcProfile.Data.Id)
	d.Set("type", oidcProfile.Data.Type)
	d.Set("name", oidcProfile.Data.Name)
	d.Set("discovery_url", oidcProfile.Data.DiscoveryUrl)
	d.Set("client_id", oidcProfile.Data.ClientID)
	d.Set("client_secret", oidcProfile.Data.ClientSecret)
	d.SetId(oidcProfile.Data.Id)
}
