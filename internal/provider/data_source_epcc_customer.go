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
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
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
