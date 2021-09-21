package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEpccCustomer() *schema.Resource {
	return &schema.Resource{
		ReadContext: addDiagToContext(dataSourceEpccCustomerRead),
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique identifier for this customer.",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The `name` of the customer.",
			},
			"email": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The `email` of the customer.",
			},
		},
	}
}

func dataSourceEpccCustomerRead(ctx context.Context, d *schema.ResourceData, m interface{}) {

	client := m.(*epcc.Client)
	customerId := d.Get("id").(string)

	customer, err := epcc.Customers.Get(&ctx, client, customerId)

	if err != nil {
		ReportAPIError(ctx, err)
	} else {
		d.Set("name", customer.Data.Name)
		d.Set("type", customer.Data.Type)
		d.Set("email", customer.Data.Email)
		d.SetId(customer.Data.Id)
	}
}
