package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEpccRealm() *schema.Resource {
	return &schema.Resource{
		ReadContext: addDiagToContext(dataSourceEpccRealmRead),
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
		DeprecationMessage: "This resource is deprecated please use epcc_authentication_realm",
	}
}

func dataSourceEpccRealmRead(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	realmId := d.Get("id").(string)
	realm, err := epcc.Realms.Get(&ctx, client, realmId)
	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	d.Set("id", realm.Data.Id)
	d.Set("type", realm.Data.Type)
	d.Set("name", realm.Data.Name)
	d.Set("redirect_uris", realm.Data.RedirectUris)
	d.Set("duplicate_email_policy", realm.Data.DuplicateEmailPolicy)
	d.Set("relationships", realm.Data.Relationships)
	d.Set("origin_id", realm.Data.Relationships.Origin.Data.Id)
	d.Set("origin_type", realm.Data.Relationships.Origin.Data.Type)
	d.SetId(realm.Data.Id)
}
