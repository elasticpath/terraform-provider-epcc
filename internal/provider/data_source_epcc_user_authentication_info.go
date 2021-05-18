package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEpccUserAuthenticationInfo() *schema.Resource {
	return &schema.Resource{
		ReadContext: addDiagToContext(dataSourceEpccUserAuthenticationInfoRead),
		Schema: map[string]*schema.Schema{
			"id":    {Type: schema.TypeString, Required: true},
			"name":  {Type: schema.TypeString, Computed: true},
			"email": {Type: schema.TypeString, Computed: true},
		},
	}
}

func dataSourceEpccUserAuthenticationInfoRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	userAuthenticationInfoId := d.Get("id").(string)
	realmId := d.Get("realm_id").(string)
	userAuthenticationInfo, err := epcc.UserAuthenticationInfos.Get(&ctx, client, realmId, userAuthenticationInfoId)
	if err != nil {
		return FromAPIError(err)
	}

	d.Set("id", userAuthenticationInfo.Data.Id)
	d.Set("type", userAuthenticationInfo.Data.Type)
	d.Set("name", userAuthenticationInfo.Data.Name)
	d.Set("email", userAuthenticationInfo.Data.Email)
	d.SetId(userAuthenticationInfo.Data.Id)

	return *ctx.Value("diags").(*diag.Diagnostics)
}
