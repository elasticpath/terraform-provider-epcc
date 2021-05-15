package provider

import (
	"context"
	"github.com/elasticpath/epcc-terraform-provider/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEpccAccount() *schema.Resource {
	return &schema.Resource{
		Description: "Represents the EPCC API [Account resource](https://documentation.elasticpath.com/commerce-cloud/docs/api/account-management/accounts/index.html#the-account-object).",
		CreateContext: addDiagToContext(resourceEpccAccountCreate),
		ReadContext:   addDiagToContext(resourceEpccAccountRead),
		UpdateContext: addDiagToContext(resourceEpccAccountUpdate),
		DeleteContext: addDiagToContext(resourceEpccAccountDelete),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"legal_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"registration_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,

				// You can't change the parent id of an account, must be recreated.
				ForceNew: true,
			},
		},
	}

}

func resourceEpccAccountDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)
	accountID := d.Id()

	err := epcc.Accounts.Delete(&ctx, client, accountID)

	if err != nil {
		FromAPIError(err)
	}

	d.SetId("")

	return *ctx.Value("diags").(*diag.Diagnostics)
}

func resourceEpccAccountUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	accountId := d.Id()

	account := &epcc.Account{
		Type:           "account",
		Name:           d.Get("name").(string),
		LegalName:      d.Get("legal_name").(string),
		RegistrationId: d.Get("registration_id").(string),
		ParentId:       d.Get("parent_id").(string),
	}

	createdAccountData, apiError := epcc.Accounts.Update(&ctx, client, accountId, account)

	if apiError != nil {
		return FromAPIError(apiError)
	}

	d.SetId(createdAccountData.Data.Id)

	return resourceEpccAccountRead(ctx, d, m)
}

func resourceEpccAccountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	accountID := d.Id()

	account, err := epcc.Accounts.Get(&ctx, client, accountID)

	if err != nil {
		return FromAPIError(err)
	}

	if err := d.Set("name", account.Data.Name); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("legal_name", account.Data.LegalName); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("registration_id", account.Data.RegistrationId); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("parent_id", account.Data.ParentId); err != nil {
		return diag.FromErr(err)
	}

	return *ctx.Value("diags").(*diag.Diagnostics)
}

func resourceEpccAccountCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)
	account := &epcc.Account{
		Type:           "account",
		Name:           d.Get("name").(string),
		LegalName:      d.Get("legal_name").(string),
		RegistrationId: d.Get("registration_id").(string),
		ParentId:       d.Get("parent_id").(string),
	}

	createdAccountData, apiError := epcc.Accounts.Create(&ctx, client, account)

	if apiError != nil {
		return FromAPIError(apiError)
	}

	d.SetId(createdAccountData.Data.Id)

	resourceEpccAccountRead(ctx, d, m)

	return *ctx.Value("diags").(*diag.Diagnostics)
}
