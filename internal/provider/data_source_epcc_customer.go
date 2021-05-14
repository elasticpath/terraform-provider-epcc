package provider

import (
	"context"
	"github.com/elasticpath/epcc-terraform-provider/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
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

func dataSourceEpccCustomerRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*epcc.Client)
	customerId := d.Get("id").(string)

	customer, err := epcc.Customers.Get(&ctx, client, customerId)

	if err != nil {
		return FromAPIError(err)
	}

	d.Set("name", customer.Data.Name)
	d.Set("type", customer.Data.Type)
	d.Set("email", customer.Data.Email)

	d.SetId(customer.Data.Id)

	return *ctx.Value("diags").(*diag.Diagnostics)
}
