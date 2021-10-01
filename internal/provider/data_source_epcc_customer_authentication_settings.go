package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCustomerAuthenticationSettings() *schema.Resource {
	return &schema.Resource{
		Description: "Represents the EPCC API [Customer Authentication Settings] (https://documentation.elasticpath.com/commerce-cloud/docs/api/advanced/settings/customer-authentication-settings/index.html)",
		ReadContext: addDiagToContext(dataSourceEpccCustomerAuthenticationSettingsRead),
		Schema: map[string]*schema.Schema{
			"realm_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The authentication realm id used for authentication for this store.",
			},
			"client_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The client id to be used in Single Sign On authentication flows for customers.",
			}},
	}
}

func dataSourceEpccCustomerAuthenticationSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	customerAuthenticationSettings, err := epcc.CustomerAuthenticationSettings.Get(&ctx, client)

	if err != nil {
		ReportAPIError(ctx, err)
		return
	}
	d.SetId("singleton")
	d.Set("realm_id", customerAuthenticationSettings.Data.Relationships.AuthenticationRealm.Data.Id)
	d.Set("client_id", customerAuthenticationSettings.Data.Meta.ClientId)
}
