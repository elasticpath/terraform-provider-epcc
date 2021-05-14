package provider

import (
	"context"
	"github.com/elasticpath/epcc-terraform-provider/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEpccField() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceEpccFieldRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"field_type": &schema.Schema{ //string, integer, boolean, float, date, or relationship
				Type:     schema.TypeString,
				Computed: true,
			},
			"slug": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"required": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"default": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"order": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"omit_null": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"flow_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceEpccFieldRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*epcc.Client)

	var diags diag.Diagnostics

	fieldId := d.Get("id").(string)

	field, err := epcc.Fields.Get(client, fieldId)

	if err != nil {
		return FromAPIError(err)
	}

	d.Set("type", field.Data.Type)
	d.Set("field_type", field.Data.FieldType)
	d.Set("slug", field.Data.Slug)
	d.Set("name", field.Data.Name)
	d.Set("description", field.Data.Description)
	d.Set("required", field.Data.Required)
	d.Set("default", field.Data.Default)
	d.Set("enabled", field.Data.Enabled)
	d.Set("order", field.Data.Order)
	d.Set("omit_null", field.Data.OmitNull)
	d.Set("flow_id", field.Data.Relationships.Flow.Data.Id)


	d.SetId(field.Data.Id)

	return diags
}
