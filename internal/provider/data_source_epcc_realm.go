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
			"id":                     {Type: schema.TypeString, Required: true},
			"name":                   {Type: schema.TypeString, Computed: true},
			"redirect_uris":          {Type: schema.TypeList, Computed: true, Elem: &schema.Schema{Type: schema.TypeString}},
			"duplicate_email_policy": {Type: schema.TypeString, Computed: true},
			"origin_id":              {Type: schema.TypeString, Computed: true},
			"origin_type":            {Type: schema.TypeString, Computed: true},
		},
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
