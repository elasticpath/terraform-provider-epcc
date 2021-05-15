package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc/field"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEpccField() *schema.Resource {
	return &schema.Resource{
		ReadContext: addDiagToContext(dataSourceEpccFieldRead),
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"field_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"slug": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"required": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"default": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"order": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"omit_null": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"flow_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"valid_string_format": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"valid_string_enum": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed: true,
			},
			"valid_int_enum": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Computed: true,
			},
			"valid_int_range": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeMap,
					Elem: &schema.Schema{
						Type: schema.TypeInt,
					},
				},
				Computed: true,
			},
			"valid_float_enum": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeFloat,
				},
				Computed: true,
			},
			"valid_float_range": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeMap,
					Elem: &schema.Schema{
						Type: schema.TypeFloat,
					},
				},
				Computed: true,
			},
			"relationship_to_one": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"relationship_to_many": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceEpccFieldRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	id := d.Get("id").(string)
	result, err := epcc.Fields.Get(&ctx, client, id)

	if err != nil {
		return FromAPIError(err)
	}

	if err := d.Set("field_type", result.Data.FieldType); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("slug", result.Data.Slug); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", result.Data.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", result.Data.Description); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("required", result.Data.Required); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("default", result.Data.Default); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("enabled", result.Data.Enabled); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("order", result.Data.Order); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("omit_null", result.Data.OmitNull); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("flow_id", result.Data.Relationships.Flow.Data.Id); err != nil {
		return diag.FromErr(err)
	}

	fieldType := field.Type(result.Data.FieldType)

	var validIntRanges []map[string]int
	var validFloatRanges []map[string]float64

	for _, validationRule := range result.Data.ValidationRules {
		switch validationRule.ValidationType() {
		case field.Email:
			fallthrough
		case field.Slug:
			fallthrough
		case field.Uuid:
			if err := d.Set("valid_string_format", validationRule.ValidationType().AsString()); err != nil {
				return diag.FromErr(err)
			}
		case field.Enum:
			switch fieldType {
			case field.Date:
				fallthrough
			case field.String:
				if err := d.Set("valid_string_enum", validationRule.(*epcc.ValidationRuleStringEnumAttribute).Options); err != nil {
					return diag.FromErr(err)
				}
			case field.Integer:
				if err := d.Set("valid_int_enum", validationRule.(*epcc.ValidationRuleIntegerEnumAttribute).Options); err != nil {
					return diag.FromErr(err)
				}
			case field.Float:
				if err := d.Set("valid_float_enum", validationRule.(*epcc.ValidationRuleFloatEnumAttribute).Options); err != nil {
					return diag.FromErr(err)
				}
			default:
				return diag.Errorf("unknown enum for field type %v", fieldType)
			}
		case field.Between:
			switch fieldType {
			case field.Integer:
				attribute := validationRule.(*epcc.ValidationRuleBetweenIntegersAttribute)
				validIntRanges = append(validIntRanges, map[string]int{
					"from": attribute.Options.From,
					"to":   attribute.Options.To,
				})
				if err := d.Set("valid_int_range", validIntRanges); err != nil {
					return diag.FromErr(err)
				}
			case field.Float:
				attribute := validationRule.(*epcc.ValidationRuleBetweenFloatsAttribute)
				validFloatRanges = append(validFloatRanges, map[string]float64{
					"from": attribute.Options.From,
					"to":   attribute.Options.To,
				})
				if err := d.Set("valid_float_range", validFloatRanges); err != nil {
					return diag.FromErr(err)
				}
			default:
				return diag.Errorf("unknown range for field type %v", fieldType)
			}
		case field.OneToMany:
			if err := d.Set("relationship_to_many", validationRule.(*epcc.ValidationRuleRelationshipAttribute).To); err != nil {
				return diag.FromErr(err)
			}
		case field.OneToOne:
			if err := d.Set("relationship_to_one", validationRule.(*epcc.ValidationRuleRelationshipAttribute).To); err != nil {
				return diag.FromErr(err)
			}
		}
	}

	d.SetId(result.Data.Id)

	return *ctx.Value("diags").(*diag.Diagnostics)
}
