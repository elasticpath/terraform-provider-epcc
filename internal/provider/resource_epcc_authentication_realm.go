package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEpccAuthenticationRealm() *schema.Resource {
	return &schema.Resource{
		Description:   "Represents the EPCC API [Authentication Realms](https://documentation.elasticpath.com/commerce-cloud/docs/api/single-sign-on/authentication-realms/index.html).",
		CreateContext: addDiagToContext(resourceEpccAuthenticationRealmCreate),
		ReadContext:   addDiagToContext(resourceEpccAuthenticationRealmRead),
		UpdateContext: addDiagToContext(resourceEpccAuthenticationRealmUpdate),
		DeleteContext: addDiagToContext(resourceEpccAuthenticationRealmDelete),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier for the authentication realm.",
			},
			"authentication_realm_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique identifier for the authentication realm. You must supply this value and retrieve it from another linked resource or data source.",
				ForceNew:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the authentication realm.",
			},
			"redirect_uris": {
				Type:        schema.TypeList,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "An array of Storefront URIs that can start Single Sign On authentication. These URIs must follow the rules for [redirection endpoints in OAuth 2.0](https://tools.ietf.org/html/rfc6749#section-3.1.2). All URIs must start with `https://` except for `http://localhost`.",
			},
			"duplicate_email_policy": {
				Type:        schema.TypeString,
				Required:    true,
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

func resourceEpccAuthenticationRealmDelete(ctx context.Context, d *schema.ResourceData, m interface{}) {

	// When we delete a realm we will reset it to the initial state
	client := m.(*epcc.Client)

	originType := d.Get("origin_type")

	authenticationRealmId := d.Id()

	authenticationRealm := &epcc.Realm{
		Id:                   authenticationRealmId,
		Type:                 "authentication-realm",
		Name:                 "Authentication Realm (Cleared by Terraform)",
		RedirectUris:         make([]interface{}, 0),
		DuplicateEmailPolicy: "allowed",
	}

	switch originType {
	// Try and reset the name to the right state
	case "account_authentication_settings":
		authenticationRealm.Name = "Account Management Realm"
	case "customer-authentication-settings":
		authenticationRealm.Name = "Buyer Organization"
	case "merchant-realm-mappings":
		authenticationRealm.Name = "Merchant Organization"
	}

	_, apiError := epcc.Realms.Update(&ctx, client, authenticationRealmId, authenticationRealm)

	if apiError != nil {
		ReportAPIError(ctx, apiError)
		return
	}

	d.SetId("")
}

func resourceEpccAuthenticationRealmUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	authenticationRealmId := d.Id()

	authenticationRealm := &epcc.Realm{
		Id:                   authenticationRealmId,
		Type:                 "authentication-realm",
		Name:                 d.Get("name").(string),
		RedirectUris:         d.Get("redirect_uris").([]interface{}),
		DuplicateEmailPolicy: d.Get("duplicate_email_policy").(string),
	}

	updatedAuthenticationRealmData, apiError := epcc.Realms.Update(&ctx, client, authenticationRealmId, authenticationRealm)

	if apiError != nil {
		ReportAPIError(ctx, apiError)
		return
	}

	d.SetId(updatedAuthenticationRealmData.Data.Id)

	resourceEpccAuthenticationRealmRead(ctx, d, m)
}

func resourceEpccAuthenticationRealmRead(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	realmId := d.Id()

	authenticationRealm, err := epcc.Realms.Get(&ctx, client, realmId)

	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	if err := d.Set("name", authenticationRealm.Data.Name); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("redirect_uris", authenticationRealm.Data.RedirectUris); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("duplicate_email_policy", authenticationRealm.Data.DuplicateEmailPolicy); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("origin_id", authenticationRealm.Data.Relationships.Origin.Data.Id); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("origin_type", authenticationRealm.Data.Relationships.Origin.Data.Type); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
}

func resourceEpccAuthenticationRealmCreate(ctx context.Context, d *schema.ResourceData, m interface{}) {

	client := m.(*epcc.Client)
	realmId := d.Get("authentication_realm_id").(string)

	_, err := epcc.Realms.Get(&ctx, client, realmId)

	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	// Save original values

	updatedAuthenticationRealm := &epcc.Realm{
		Id:                   realmId,
		Type:                 "authentication-realm",
		Name:                 d.Get("name").(string),
		RedirectUris:         d.Get("redirect_uris").([]interface{}),
		DuplicateEmailPolicy: d.Get("duplicate_email_policy").(string),
	}

	_, apiError := epcc.Realms.Update(&ctx, client, realmId, updatedAuthenticationRealm)

	if apiError != nil {
		ReportAPIError(ctx, apiError)
		return
	}

	d.SetId(realmId)

	resourceEpccAuthenticationRealmRead(ctx, d, m)
}
