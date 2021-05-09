package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gitlab.elasticpath.com/Steve.Ramage/epcc-terraform-provider/external/sdk/epcc"
)

func resourceEpccAccount() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceEpccAccountCreate,
		ReadContext:   resourceEpccAccountRead,
		UpdateContext: resourceEpccAccountUpdate,
		DeleteContext: resourceEpccAccountDelete,
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

func resourceEpccAccountDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	var diags diag.Diagnostics

	accountID := d.Id()

	err := epcc.Accounts.Delete(client, accountID)

	if err != nil {
		FromAPIError(err)
	}

	d.SetId("")

	return diags
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

	createdAccountData, apiError := epcc.Accounts.Update(client, accountId, account)

	if apiError != nil {
		return FromAPIError(apiError)
	}

	d.SetId(createdAccountData.Data.Id)

	return resourceEpccAccountRead(ctx, d, m)
}

func resourceEpccAccountRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	var diags diag.Diagnostics

	accountID := d.Id()

	account, err := epcc.Accounts.Get(client, accountID)

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

	return diags
}

func resourceEpccAccountCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	var diags diag.Diagnostics

	account := &epcc.Account{
		Type:           "account",
		Name:           d.Get("name").(string),
		LegalName:      d.Get("legal_name").(string),
		RegistrationId: d.Get("registration_id").(string),
		ParentId:       d.Get("parent_id").(string),
	}

	createdAccountData, apiError := epcc.Accounts.Create(client, account)

	if apiError != nil {
		return FromAPIError(apiError)
	}

	d.SetId(createdAccountData.Data.Id)

	resourceEpccAccountRead(ctx, d, m)

	return diags
}
