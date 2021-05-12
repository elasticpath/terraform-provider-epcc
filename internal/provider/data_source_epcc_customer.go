package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gitlab.elasticpath.com/Steve.Ramage/epcc-terraform-provider/external/sdk/epcc"
)

func dataSourceEpccCustomer() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceEpccCustomerRead,
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

func dataSourceEpccCustomerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*epcc.Client)

	var diags diag.Diagnostics

	customerId := d.Get("id").(string)

	customer, err := epcc.Customers.Get(client, customerId)

	if err != nil {
		return FromAPIError(err)
	}

	d.Set("name", customer.Data.Name)
	d.Set("type", customer.Data.Type)
	d.Set("email", customer.Data.Type)

	d.SetId(customer.Data.Id)

	return diags
}
