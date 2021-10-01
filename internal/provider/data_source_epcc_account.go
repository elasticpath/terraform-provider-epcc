package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEpccAccount() *schema.Resource {
	return &schema.Resource{
		ReadContext: addDiagToContext(dataSourceEpccAccountRead),
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique identifier for an Account",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the account.",
			},
			"legal_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The legal name of the account.",
			},
			"registration_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The registration ID of the account.",
			},
			"parent_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Specifies the ID of the parent account.",
			},
		},
	}
}

func dataSourceEpccAccountRead(ctx context.Context, d *schema.ResourceData, m interface{}) {

	client := m.(*epcc.Client)
	accountId := d.Get("id").(string)

	account, err := epcc.Accounts.Get(&ctx, client, accountId)

	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	d.Set("name", account.Data.Name)
	d.Set("type", account.Data.Type)
	d.Set("legal_name", account.Data.Type)
	d.Set("registration_id", account.Data.RegistrationId)
	d.Set("parent_id", account.Data.ParentId)
	d.SetId(account.Data.Id)
}
