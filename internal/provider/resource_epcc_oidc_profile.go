package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEpccOidcProfile() *schema.Resource {
	return &schema.Resource{
		Description:   "Represents the EPCC API [OpenID Connect Profiles](https://documentation.elasticpath.com/commerce-cloud/docs/api/single-sign-on/oidc-profiles/index.html).",
		CreateContext: addDiagToContext(resourceEpccOidcProfileCreate),
		ReadContext:   addDiagToContext(resourceEpccOidcProfileRead),
		UpdateContext: addDiagToContext(resourceEpccOidcProfileUpdate),
		DeleteContext: addDiagToContext(resourceEpccOidcProfileDelete),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique identifier for this OpenID Connect profile.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the OpenID Connect profile.",
			},
			"discovery_url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The url of the OpenID Connect discovery document.",
			},
			"client_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The client id to be used with the external authentication provider.",
			},
			"client_secret": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The client secret for the OpenID Provider.",
			},
			"realm_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the authentication-realm containing the OpenID Connect profiles.",
			},
			"callback_endpoint": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "This link is the endpoint that should be supplied as the callback url to the upstream authentication provider.",
			},
			"authorization_endpoint": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "This link represents the endpoint that front end applications should use to authenticate with this OpenID Connect profile. The front end application is responsible for appending all of the [required parameters](https://openid.net/specs/openid-connect-core-1_0.html#AuthRequest) to this request.",
			},
			"client_discovery_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A link to the [OpenID Connect Discovery](https://openid.net/specs/openid-connect-discovery-1_0.html) document for this provider.",
			},
		},
	}

}

func resourceEpccOidcProfileDelete(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	profileID := d.Id()
	realmID := d.Get("realm_id").(string)

	err := epcc.OidcProfiles.Delete(&ctx, client, profileID, realmID)

	if err != nil {
		ReportAPIError(ctx, err)
	}

	d.SetId("")
}

func resourceEpccOidcProfileUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	profileId := d.Id()

	oidcProfile := &epcc.OidcProfile{
		Id:           profileId,
		Type:         "oidc-profile",
		Name:         d.Get("name").(string),
		DiscoveryUrl: d.Get("discovery_url").(string),
		ClientID:     d.Get("client_id").(string),
		ClientSecret: d.Get("client_secret").(string),
		RealmId:      d.Get("realm_id").(string),
	}

	updatedOidcProfileData, apiError := epcc.OidcProfiles.Update(&ctx, client, oidcProfile)

	if apiError != nil {
		ReportAPIError(ctx, apiError)
		return
	}

	d.SetId(updatedOidcProfileData.Data.Id)

	resourceEpccOidcProfileRead(ctx, d, m)
}

func resourceEpccOidcProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	profileId := d.Id()

	realmId := d.Get("realm_id").(string)

	oidcProfile, err := epcc.OidcProfiles.Get(&ctx, client, realmId, profileId)

	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	//if err := d.Set("type", "oidc-oidcProfile"); err != nil {
	//	addToDiag(ctx, diag.FromErr(err)); return
	//}
	if err := d.Set("name", oidcProfile.Data.Name); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("discovery_url", oidcProfile.Data.DiscoveryUrl); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("client_id", oidcProfile.Data.ClientID); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	if oidcProfile.Links != nil {
		if err := d.Set("callback_endpoint", oidcProfile.Links.CallbackEndpoint); err != nil {
			addToDiag(ctx, diag.FromErr(err))
			return
		}

		if err := d.Set("authorization_endpoint", oidcProfile.Links.AuthorizationEndpoint); err != nil {
			addToDiag(ctx, diag.FromErr(err))
			return
		}

		if err := d.Set("client_discovery_url", oidcProfile.Links.ClientDiscoveryUrl); err != nil {
			addToDiag(ctx, diag.FromErr(err))
			return
		}
	} else {
		if err := d.Set("callback_endpoint", ""); err != nil {
			addToDiag(ctx, diag.FromErr(err))
			return
		}

		if err := d.Set("authorization_endpoint", ""); err != nil {
			addToDiag(ctx, diag.FromErr(err))
			return
		}

		if err := d.Set("client_discovery_url", ""); err != nil {
			addToDiag(ctx, diag.FromErr(err))
			return
		}
	}
}

func resourceEpccOidcProfileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	oidcProfile := &epcc.OidcProfile{
		Type:         "oidc-profile",
		Id:           d.Get("id").(string),
		Name:         d.Get("name").(string),
		DiscoveryUrl: d.Get("discovery_url").(string),
		ClientID:     d.Get("client_id").(string),
		ClientSecret: d.Get("client_secret").(string),
		RealmId:      d.Get("realm_id").(string),
	}
	createdOidcProfileData, apiError := epcc.OidcProfiles.Create(&ctx, client, oidcProfile)
	if apiError != nil {
		ReportAPIError(ctx, apiError)
		return
	}

	d.SetId(createdOidcProfileData.Data.Id)

	resourceEpccOidcProfileRead(ctx, d, m)
}
