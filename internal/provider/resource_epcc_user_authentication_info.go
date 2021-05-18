package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEpccUserAuthenticationInfo() *schema.Resource {
	return &schema.Resource{
		Description:   "Represents the EPCC API User Authentication Info",
		CreateContext: addDiagToContext(resourceEpccUserAuthenticationInfoCreate),
		ReadContext:   addDiagToContext(resourceEpccUserAuthenticationInfoRead),
		UpdateContext: addDiagToContext(resourceEpccUserAuthenticationInfoUpdate),
		DeleteContext: addDiagToContext(resourceEpccUserAuthenticationInfoDelete),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id":       {Type: schema.TypeString, Computed: true},
			"name":     {Type: schema.TypeString, Required: true},
			"email":    {Type: schema.TypeString, Required: true},
			"realm_id": {Type: schema.TypeString, Required: true},
		},
	}

}

func resourceEpccUserAuthenticationInfoDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)
	userAuthenticationInfoID := d.Id()
	realmID := d.Get("realm_id").(string)

	err := epcc.UserAuthenticationInfos.Delete(&ctx, client, userAuthenticationInfoID, realmID)

	if err != nil {
		FromAPIError(err)
	}

	d.SetId("")

	return *ctx.Value("diags").(*diag.Diagnostics)
}

func resourceEpccUserAuthenticationInfoUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	userAuthenticationInfoId := d.Id()

	userAuthenticationInfo := &epcc.UserAuthenticationInfo{
		Id:      userAuthenticationInfoId,
		Type:    "user-authentication-info",
		Name:    d.Get("name").(string),
		Email:   d.Get("email").(string),
		RealmId: d.Get("realm_id").(string),
	}

	updatedUserAuthenticationInfoData, apiError := epcc.UserAuthenticationInfos.Update(&ctx, client, userAuthenticationInfo)

	if apiError != nil {
		return FromAPIError(apiError)
	}

	d.SetId(updatedUserAuthenticationInfoData.Data.Id)

	return resourceEpccUserAuthenticationInfoRead(ctx, d, m)
}

func resourceEpccUserAuthenticationInfoRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)
	userAuthenticationInfoId := d.Id()

	realmId := d.Get("realm_id").(string)

	userAuthenticationInfo, err := epcc.UserAuthenticationInfos.Get(&ctx, client, realmId, userAuthenticationInfoId)

	if err != nil {
		return FromAPIError(err)
	}

	//if err := d.Set("type", "user-authentication-info"); err != nil {
	//	return diag.FromErr(err)
	//}
	if err := d.Set("name", userAuthenticationInfo.Data.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("email", userAuthenticationInfo.Data.Email); err != nil {
		return diag.FromErr(err)
	}

	return *ctx.Value("diags").(*diag.Diagnostics)
}

func resourceEpccUserAuthenticationInfoCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)
	userAuthenticationInfo := &epcc.UserAuthenticationInfo{
		Type:    "user-authentication-info",
		Id:      d.Get("id").(string),
		Name:    d.Get("name").(string),
		Email:   d.Get("email").(string),
		RealmId: d.Get("realm_id").(string),
	}
	createdUserAuthenticationInfoData, apiError := epcc.UserAuthenticationInfos.Create(&ctx, client, userAuthenticationInfo)
	if apiError != nil {
		return FromAPIError(apiError)
	}

	d.SetId(createdUserAuthenticationInfoData.Data.Id)

	resourceEpccUserAuthenticationInfoRead(ctx, d, m)

	return *ctx.Value("diags").(*diag.Diagnostics)
}
