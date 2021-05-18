package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEpccCustomer() *schema.Resource {
	return &schema.Resource{
		Description:   "Represents the EPCC API [Customer Object](https://documentation.elasticpath.com/commerce-cloud/docs/api/orders-and-customers/customers/index.html#the-customer-object).",
		CreateContext: addDiagToContext(resourceEpccCustomerCreate),
		ReadContext:   addDiagToContext(resourceEpccCustomerRead),
		UpdateContext: addDiagToContext(resourceEpccCustomerUpdate),
		DeleteContext: addDiagToContext(resourceEpccCustomerDelete),
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

func resourceEpccCustomerDelete(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	customerID := d.Id()

	err := epcc.Customers.Delete(&ctx, client, customerID)

	if err != nil {
		ReportAPIError(ctx, err)
	}

	d.SetId("")
}

func resourceEpccCustomerUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	customerId := d.Id()

	customer := &epcc.Customer{
		Type:  "customer",
		Name:  d.Get("name").(string),
		Email: d.Get("email").(string),
	}

	updatedCustomerData, apiError := epcc.Customers.Update(&ctx, client, customerId, customer)

	if apiError != nil {
		ReportAPIError(ctx, apiError)
		return
	}

	d.SetId(updatedCustomerData.Data.Id)

	resourceEpccCustomerRead(ctx, d, m)
}

func resourceEpccCustomerRead(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	customerId := d.Id()

	customer, err := epcc.Customers.Get(&ctx, client, customerId)

	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	if err := d.Set("name", customer.Data.Name); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}

	if err := d.Set("email", customer.Data.Email); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
}

func resourceEpccCustomerCreate(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)
	customer := &epcc.Customer{
		Type:  "customer",
		Name:  d.Get("name").(string),
		Email: d.Get("email").(string),
	}

	createdCustomerData, apiError := epcc.Customers.Create(&ctx, client, customer)

	if apiError != nil {
		ReportAPIError(ctx, apiError)
		return
	}

	d.SetId(createdCustomerData.Data.Id)

	resourceEpccCustomerRead(ctx, d, m)
}
