package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAccountAuthenticationSettings() *schema.Resource {
	return &schema.Resource{
		Description: "Represents the EPCC API Account Authentication Settings",
		ReadContext: addDiagToContext(dataSourceEpccAccountAuthenticationSettingsRead),
		Schema: map[string]*schema.Schema{
			"realm_id":  {Type: schema.TypeString, Computed: true},
			"client_id": {Type: schema.TypeString, Computed: true}},
	}
}

func dataSourceEpccAccountAuthenticationSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	accountAuthenticationSettings, err := epcc.AccountAuthenticationSettings.Get(&ctx, client)

	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	d.SetId("singleton")
	d.Set("realm_id", accountAuthenticationSettings.Data.Relationships.AuthenticationRealm.Data.Id)
	d.Set("client_id", accountAuthenticationSettings.Data.Meta.ClientId)
}
