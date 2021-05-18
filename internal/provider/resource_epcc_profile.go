package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEpccProfile() *schema.Resource {
	return &schema.Resource{
		Description:   "Represents the EPCC API [OpenID Connect Profiles](https://documentation.elasticpath.com/commerce-cloud/docs/api/single-sign-on/oidc-profiles/index.html).",
		CreateContext: addDiagToContext(resourceEpccProfileCreate),
		ReadContext:   addDiagToContext(resourceEpccProfileRead),
		UpdateContext: addDiagToContext(resourceEpccProfileUpdate),
		DeleteContext: addDiagToContext(resourceEpccProfileDelete),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id":            {Type: schema.TypeString, Computed: true},
			"name":          {Type: schema.TypeString, Required: true},
			"discovery_url": {Type: schema.TypeString, Required: true},
			"client_id":     {Type: schema.TypeString, Required: true},
			"client_secret": {Type: schema.TypeString, Required: true},
			"realm_id":      {Type: schema.TypeString, Required: true},
		},
	}

}

func resourceEpccProfileDelete(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	profileID := d.Id()
	realmID := d.Get("realm_id").(string)

	err := epcc.Profiles.Delete(&ctx, client, profileID, realmID)

	if err != nil {
		ReportAPIError(ctx, err)
	}

	d.SetId("")
}

func resourceEpccProfileUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	profileId := d.Id()

	profile := &epcc.Profile{
		Id:           profileId,
		Type:         "oidc-profile",
		Name:         d.Get("name").(string),
		DiscoveryUrl: d.Get("discovery_url").(string),
		ClientID:     d.Get("client_id").(string),
		ClientSecret: d.Get("client_secret").(string),
		RealmId:      d.Get("realm_id").(string),
	}

	updatedProfileData, apiError := epcc.Profiles.Update(&ctx, client, profile)

	if apiError != nil {
		ReportAPIError(ctx, apiError)
		return
	}

	d.SetId(updatedProfileData.Data.Id)

	resourceEpccProfileRead(ctx, d, m)
}

func resourceEpccProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	profileId := d.Id()

	realmId := d.Get("realm_id").(string)

	profile, err := epcc.Profiles.Get(&ctx, client, realmId, profileId)

	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	//if err := d.Set("type", "oidc-profile"); err != nil {
	//	addToDiag(ctx, diag.FromErr(err)); return
	//}
	if err := d.Set("name", profile.Data.Name); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("discovery_url", profile.Data.DiscoveryUrl); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("client_id", profile.Data.ClientID); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
}

func resourceEpccProfileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	profile := &epcc.Profile{
		Type:         "oidc-profile",
		Id:           d.Get("id").(string),
		Name:         d.Get("name").(string),
		DiscoveryUrl: d.Get("discovery_url").(string),
		ClientID:     d.Get("client_id").(string),
		ClientSecret: d.Get("client_secret").(string),
		RealmId:      d.Get("realm_id").(string),
	}
	createdProfileData, apiError := epcc.Profiles.Create(&ctx, client, profile)
	if apiError != nil {
		ReportAPIError(ctx, apiError)
		return
	}

	d.SetId(createdProfileData.Data.Id)

	resourceEpccProfileRead(ctx, d, m)
}
