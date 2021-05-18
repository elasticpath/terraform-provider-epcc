package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEpccRealm() *schema.Resource {
	return &schema.Resource{
		Description:   "Represents the EPCC API [Authentication Realms](https://documentation.elasticpath.com/commerce-cloud/docs/api/single-sign-on/authentication-realms/index.html).",
		CreateContext: addDiagToContext(resourceEpccRealmCreate),
		ReadContext:   addDiagToContext(resourceEpccRealmRead),
		UpdateContext: addDiagToContext(resourceEpccRealmUpdate),
		DeleteContext: addDiagToContext(resourceEpccRealmDelete),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id":                     {Type: schema.TypeString, Computed: true},
			"name":                   {Type: schema.TypeString, Required: true},
			"redirect_uris":          {Type: schema.TypeList, Required: true, Elem: &schema.Schema{Type: schema.TypeString}},
			"duplicate_email_policy": {Type: schema.TypeString, Required: true},
			"origin_id":              {Type: schema.TypeString, Required: true},
			"origin_type":            {Type: schema.TypeString, Required: true},
		},
	}

}

func resourceEpccRealmDelete(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	realmID := d.Id()

	err := epcc.Realms.Delete(&ctx, client, realmID)

	if err != nil {
		ReportAPIError(ctx, err)
	}

	d.SetId("")
}

func resourceEpccRealmUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	realmId := d.Id()

	realm := &epcc.Realm{
		Id:                   realmId,
		Type:                 "authentication-realm",
		Name:                 d.Get("name").(string),
		RedirectUris:         d.Get("redirect_uris").([]interface{}),
		DuplicateEmailPolicy: d.Get("duplicate_email_policy").(string),
		Relationships: &epcc.RealmRelationships{
			Origin: &epcc.RealmRelationshipsOrigin{
				Data: &epcc.RealmRelationshipsOriginData{
					Id:   d.Get("origin_id").(string),
					Type: d.Get("origin_type").(string),
				},
			},
		},
	}

	updatedRealmData, apiError := epcc.Realms.Update(&ctx, client, realmId, realm)

	if apiError != nil {
		ReportAPIError(ctx, apiError)
		return
	}

	d.SetId(updatedRealmData.Data.Id)

	resourceEpccRealmRead(ctx, d, m)
}

func resourceEpccRealmRead(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	realmId := d.Id()

	realm, err := epcc.Realms.Get(&ctx, client, realmId)

	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	//if err := d.Set("type", "authentication-realm"); err != nil {
	//	addToDiag(ctx, diag.FromErr(err)); return
	//}
	if err := d.Set("name", realm.Data.Name); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("redirect_uris", realm.Data.RedirectUris); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("duplicate_email_policy", realm.Data.DuplicateEmailPolicy); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("origin_id", realm.Data.Relationships.Origin.Data.Id); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("origin_type", realm.Data.Relationships.Origin.Data.Type); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
}

func resourceEpccRealmCreate(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	realm := &epcc.Realm{
		Type:                 "authentication-realm",
		Id:                   d.Get("id").(string),
		Name:                 d.Get("name").(string),
		RedirectUris:         d.Get("redirect_uris").([]interface{}),
		DuplicateEmailPolicy: d.Get("duplicate_email_policy").(string),
		Relationships: &epcc.RealmRelationships{
			Origin: &epcc.RealmRelationshipsOrigin{
				Data: &epcc.RealmRelationshipsOriginData{
					Id:   d.Get("origin_id").(string),
					Type: d.Get("origin_type").(string),
				},
			},
		},
	}
	createdRealmData, apiError := epcc.Realms.Create(&ctx, client, realm)
	if apiError != nil {
		ReportAPIError(ctx, apiError)
		return
	}

	d.SetId(createdRealmData.Data.Id)

	resourceEpccRealmRead(ctx, d, m)
}
