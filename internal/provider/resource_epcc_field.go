package provider

import (
	"context"
	"github.com/elasticpath/epcc-terraform-provider/external/sdk/epcc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEpccField() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceEpccFieldCreate,
		ReadContext:   resourceEpccFieldRead,
		UpdateContext: resourceEpccFieldUpdate,
		DeleteContext: resourceEpccFieldDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"field_type": &schema.Schema{ //string, integer, boolean, float, date, or relationship
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"required": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"default": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"omit_null": &schema.Schema{
				Type:     schema.TypeBool,
				Required: true,
			},
			"flow_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}

}

func resourceEpccFieldDelete(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	var diags diag.Diagnostics

	fieldID := d.Id()

	err := epcc.Fields.Delete(client, fieldID)

	if err != nil {
		FromAPIError(err)
	}

	d.SetId("")

	return diags
}

func resourceEpccFieldUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	fieldID := d.Id()
	field := constructField(d)

	createdFieldData, apiError := epcc.Fields.Update(client, fieldID, field)

	if apiError != nil {
		return FromAPIError(apiError)
	}

	d.SetId(createdFieldData.Data.Id)

	return resourceEpccFieldRead(ctx, d, m)
}

func resourceEpccFieldRead(_ context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	var diags diag.Diagnostics

	fieldID := d.Id()

	field, err := epcc.Fields.Get(client, fieldID)

	if err != nil {
		return FromAPIError(err)
	}

	if err := d.Set("field_type", field.Data.FieldType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("slug", field.Data.Slug); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", field.Data.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", field.Data.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("required", field.Data.Required); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("default", field.Data.Default); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("enabled", field.Data.Enabled); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("order", field.Data.Order); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("omit_null", field.Data.OmitNull); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("flow_id", field.Data.Relationships.Flow.Data.Id); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("type", field.Data.Type); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceEpccFieldCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	var diags diag.Diagnostics

	field := constructField(d)

	createdFieldData, apiError := epcc.Fields.Create(client, field)

	if apiError != nil {
		return FromAPIError(apiError)
	}

	d.SetId(createdFieldData.Data.Id)

	resourceEpccFieldRead(ctx, d, m)

	return diags
}

func constructField(d *schema.ResourceData) *epcc.Field {
	flowRelationship := &epcc.RelationshipsAttribute{
		Flow: &epcc.RelationshipsAttributeFlow{
			Data: &epcc.RelationshipsAttributeFlowData{
				Id:   d.Get("flow_id").(string),
				Type: "flow",
			},
		},
	}

	field := &epcc.Field{
		Type:          "field",
		FieldType:     d.Get("field_type").(string),
		Slug:          d.Get("slug").(string),
		Name:          d.Get("name").(string),
		Description:   d.Get("description").(string),
		Required:      d.Get("required").(bool),
		Default:       d.Get("default").(string),
		Enabled:       true,
		Order:         1,
		OmitNull:      d.Get("omit_null").(bool),
		Relationships: flowRelationship,
	}
	return field
}
