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

func resourceEpccUserAuthenticationInfoDelete(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	userAuthenticationInfoID := d.Id()
	realmID := d.Get("realm_id").(string)

	err := epcc.UserAuthenticationInfos.Delete(&ctx, client, userAuthenticationInfoID, realmID)

	if err != nil {
		ReportAPIError(ctx, err)
	}

	d.SetId("")
}

func resourceEpccUserAuthenticationInfoUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	userAuthenticationInfoId := d.Id()

	userAuthenticationInfo := &epcc.UserAuthenticationInfo{
		Id:      userAuthenticationInfoId,
		Type:    "user_authentication_info",
		Name:    d.Get("name").(string),
		Email:   d.Get("email").(string),
		RealmId: d.Get("realm_id").(string),
	}

	updatedUserAuthenticationInfoData, apiError := epcc.UserAuthenticationInfos.Update(&ctx, client, userAuthenticationInfo)

	if apiError != nil {
		ReportAPIError(ctx, apiError)
		return
	}

	d.SetId(updatedUserAuthenticationInfoData.Data.Id)

	resourceEpccUserAuthenticationInfoRead(ctx, d, m)
}

func resourceEpccUserAuthenticationInfoRead(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	userAuthenticationInfoId := d.Id()

	realmId := d.Get("realm_id").(string)

	userAuthenticationInfo, err := epcc.UserAuthenticationInfos.Get(&ctx, client, realmId, userAuthenticationInfoId)

	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	//if err := d.Set("type", "user-authentication-info"); err != nil {
	//	addToDiag(ctx, diag.FromErr(err)); return
	//}
	if err := d.Set("name", userAuthenticationInfo.Data.Name); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("email", userAuthenticationInfo.Data.Email); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
}

func resourceEpccUserAuthenticationInfoCreate(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	userAuthenticationInfo := &epcc.UserAuthenticationInfo{
		Type:    "user_authentication_info",
		Id:      d.Get("id").(string),
		Name:    d.Get("name").(string),
		Email:   d.Get("email").(string),
		RealmId: d.Get("realm_id").(string),
	}
	createdUserAuthenticationInfoData, apiError := epcc.UserAuthenticationInfos.Create(&ctx, client, userAuthenticationInfo)
	if apiError != nil {
		ReportAPIError(ctx, apiError)
		return
	}

	d.SetId(createdUserAuthenticationInfoData.Data.Id)

	resourceEpccUserAuthenticationInfoRead(ctx, d, m)
}
