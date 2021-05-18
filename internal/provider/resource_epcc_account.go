package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEpccAccount() *schema.Resource {
	return &schema.Resource{
		Description:   "Represents the EPCC API [Account resource](https://documentation.elasticpath.com/commerce-cloud/docs/api/account-management/accounts/index.html#the-account-object).",
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

func resourceEpccAccountDelete(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	accountID := d.Id()

	err := epcc.Accounts.Delete(&ctx, client, accountID)

	if err != nil {
		ReportAPIError(ctx, err)
	}

	d.SetId("")
}

func resourceEpccAccountUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) {
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
		ReportAPIError(ctx, apiError)
		return
	}

	d.SetId(createdAccountData.Data.Id)

	resourceEpccAccountRead(ctx, d, m)
}

func resourceEpccAccountRead(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	accountID := d.Id()

	account, err := epcc.Accounts.Get(&ctx, client, accountID)

	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	if err := d.Set("name", account.Data.Name); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	if err := d.Set("legal_name", account.Data.LegalName); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	if err := d.Set("registration_id", account.Data.RegistrationId); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	if err := d.Set("parent_id", account.Data.ParentId); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
}

func resourceEpccAccountCreate(ctx context.Context, d *schema.ResourceData, m interface{}) {
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
		ReportAPIError(ctx, apiError)
		return
	}

	d.SetId(createdAccountData.Data.Id)

	resourceEpccAccountRead(ctx, d, m)
}
