package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEpccAccount() *schema.Resource {
	return &schema.Resource{
		ReadContext: addDiagToContext(dataSourceEpccAccountRead),
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"legal_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"registration_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceEpccAccountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*epcc.Client)
	accountId := d.Get("id").(string)

	account, err := epcc.Accounts.Get(&ctx, client, accountId)

	if err != nil {
		return FromAPIError(err)
	}

	d.Set("name", account.Data.Name)
	d.Set("type", account.Data.Type)
	d.Set("legal_name", account.Data.Type)
	d.Set("registration_id", account.Data.RegistrationId)
	d.Set("parent_id", account.Data.ParentId)
	d.SetId(account.Data.Id)

	return *ctx.Value("diags").(*diag.Diagnostics)
}
