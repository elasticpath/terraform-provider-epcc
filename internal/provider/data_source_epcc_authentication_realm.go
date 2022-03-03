package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEpccAuthenticationRealm() *schema.Resource {
	return &schema.Resource{
		ReadContext: addDiagToContext(dataSourceEpccAuthenticationRealmRead),
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique identifier for the authentication realm.",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the authentication realm.",
			},
			"redirect_uris": {
				Type:        schema.TypeList,
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "An array of Storefront URIs that can start Single Sign On authentication. These URIs must follow the rules for [redirection endpoints in OAuth 2.0](https://tools.ietf.org/html/rfc6749#section-3.1.2). All URIs must start with `https://` except for `http://localhost`.",
			},
			"duplicate_email_policy": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The values permitted for this parameter are, `allowed` or `api_only`. When an unfamiliar user signs in for the first time, a value of `allowed` always creates a new user with the name and e-mail address supplied by the identity provider. With the `api_only` value, the system assigns the user to an existing user with a matching e-mail address, if one already exists. The `api_only` setting is recommended only when all configured identity providers treat e-mail address as a unique identifier for the user, otherwise a user might get access to another user’s account and data. Thus the `api_only` value can simplify administration of users.",
			},
			"origin_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the origin entity.",
			},
			"origin_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of the origin entity.",
			},
		},
	}
}

func dataSourceEpccAuthenticationRealmRead(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	authenticationRealmId := d.Get("id").(string)
	authenticationRealm, err := epcc.Realms.Get(&ctx, client, authenticationRealmId)
	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	d.Set("id", authenticationRealm.Data.Id)
	d.Set("name", authenticationRealm.Data.Name)
	d.Set("redirect_uris", authenticationRealm.Data.RedirectUris)
	d.Set("duplicate_email_policy", authenticationRealm.Data.DuplicateEmailPolicy)
	d.Set("origin_id", authenticationRealm.Data.Relationships.Origin.Data.Id)
	d.Set("origin_type", authenticationRealm.Data.Relationships.Origin.Data.Type)
	d.SetId(authenticationRealm.Data.Id)
}
