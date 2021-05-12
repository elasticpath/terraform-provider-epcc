package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gitlab.elasticpath.com/Steve.Ramage/epcc-terraform-provider/external/sdk/epcc"
)

func resourceEpccCustomer() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceEpccCustomerCreate,
		ReadContext:   resourceEpccCustomerRead,
		UpdateContext: resourceEpccCustomerUpdate,
		DeleteContext: resourceEpccCustomerDelete,
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
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}

}

func resourceEpccCustomerDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	var diags diag.Diagnostics

	customerID := d.Id()

	err := epcc.Customers.Delete(client, customerID)

	if err != nil {
		FromAPIError(err)
	}

	d.SetId("")

	return diags
}

func resourceEpccCustomerUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	customerId := d.Id()

	customer := &epcc.Customer{
		Type:  "customer",
		Name:  d.Get("name").(string),
		Email: d.Get("email").(string),
	}

	updatedCustomerData, apiError := epcc.Customers.Update(client, customerId, customer)

	if apiError != nil {
		return FromAPIError(apiError)
	}

	d.SetId(updatedCustomerData.Data.Id)

	return resourceEpccCustomerRead(ctx, d, m)
}

func resourceEpccCustomerRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	var diags diag.Diagnostics

	customerId := d.Id()

	customer, err := epcc.Customers.Get(client, customerId)

	if err != nil {
		return FromAPIError(err)
	}

	if err := d.Set("name", customer.Data.Name); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("email", customer.Data.Email); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceEpccCustomerCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	var diags diag.Diagnostics

	customer := &epcc.Customer{
		Type:  "customer",
		Name:  d.Get("name").(string),
		Email: d.Get("email").(string),
	}

	createdCustomerData, apiError := epcc.Customers.Create(client, customer)

	if apiError != nil {
		return FromAPIError(apiError)
	}

	d.SetId(createdCustomerData.Data.Id)

	resourceEpccCustomerRead(ctx, d, m)

	return diags
}
