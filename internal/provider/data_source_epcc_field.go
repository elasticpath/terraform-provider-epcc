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
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique identifier for this field.",
			},
			"field_type": {
				Type:     schema.TypeString,
				Computed: true,
				Description: "Specifies the type of field, such as string, integer, boolean, float, date, relationship.",
			},
			"slug": {
				Type:     schema.TypeString,
				Computed: true,
				Description: "A unique slug identifier for the field.",
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: "The name of the field.",
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
				Description: "Any description for this field.",
			},
			"required": {
				Type:     schema.TypeBool,
				Computed: true,
				Description: "true if required on input, false if not. Always false if the field_type is a relationship.",
			},
			"default": {
				Type:     schema.TypeString,
				Computed: true,
				Description: "A default value if none is supplied and field is not required.",
			},
			"enabled": {
				Type:     schema.TypeBool,
				Computed: true,
				Description: "If this field is enabled on the flow this should be true, otherwise false.",
			},
			"order": {
				Type:     schema.TypeInt,
				Computed: true,
				Description: "Denotes the order in which this field is returned relative to the rest of the flow fields.",
			},
			"omit_null": {
				Type:     schema.TypeBool,
				Computed: true,
				Description: "Hide this field from responses if the value is null.",
			},
			"flow_id": {
				Type:     schema.TypeString,
				Computed: true,
				Description: "The id of the flow that this field applies to.",
			},
			"valid_string_format": {
				Type:     schema.TypeString,
				Computed: true,
				Description: "Corresponds to the field validation rules for string, one of \"email\", \"slug\", or \"uuid\".",
			},
			"valid_string_enum": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed: true,
				Description: "A predefined collection of strings that represent the allowed value for this string field.",
			},
			"valid_int_enum": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Computed: true,
				Description: "A predefined collection of integer that represent the allowed value for this integer field.",
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
				Description: "A list of integers specified with from= and to= that represent the range of this value",
			},
			"valid_float_enum": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeFloat,
				},
				Computed: true,
				Description: "A predefined collection of floats that represent the allowed value for this float field.",
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

func dataSourceEpccFieldRead(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	id := d.Get("id").(string)
	result, err := epcc.Fields.Get(&ctx, client, id)

	if err != nil {
		ReportAPIError(ctx, err)
		return
	}

	if err := d.Set("field_type", result.Data.FieldType); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("slug", result.Data.Slug); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("name", result.Data.Name); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("description", result.Data.Description); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("required", result.Data.Required); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("default", result.Data.Default); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("enabled", result.Data.Enabled); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("order", result.Data.Order); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("omit_null", result.Data.OmitNull); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
	}
	if err := d.Set("flow_id", result.Data.Relationships.Flow.Data.Id); err != nil {
		addToDiag(ctx, diag.FromErr(err))
		return
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
				addToDiag(ctx, diag.FromErr(err))
				return
			}
		case field.Enum:
			switch fieldType {
			case field.Date:
				fallthrough
			case field.String:
				if err := d.Set("valid_string_enum", validationRule.(*epcc.ValidationRuleStringEnumAttribute).Options); err != nil {
					addToDiag(ctx, diag.FromErr(err))
					return
				}
			case field.Integer:
				if err := d.Set("valid_int_enum", validationRule.(*epcc.ValidationRuleIntegerEnumAttribute).Options); err != nil {
					addToDiag(ctx, diag.FromErr(err))
					return
				}
			case field.Float:
				if err := d.Set("valid_float_enum", validationRule.(*epcc.ValidationRuleFloatEnumAttribute).Options); err != nil {
					addToDiag(ctx, diag.FromErr(err))
					return
				}
			default:
				addToDiag(ctx, diag.Errorf("unknown enum for field type %v", fieldType))
				return
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
					addToDiag(ctx, diag.FromErr(err))
					return
				}
			case field.Float:
				attribute := validationRule.(*epcc.ValidationRuleBetweenFloatsAttribute)
				validFloatRanges = append(validFloatRanges, map[string]float64{
					"from": attribute.Options.From,
					"to":   attribute.Options.To,
				})
				if err := d.Set("valid_float_range", validFloatRanges); err != nil {
					addToDiag(ctx, diag.FromErr(err))
					return
				}
			default:
				addToDiag(ctx, diag.Errorf("unknown range for field type %v", fieldType))
				return
			}
		case field.OneToMany:
			if err := d.Set("relationship_to_many", validationRule.(*epcc.ValidationRuleRelationshipAttribute).To); err != nil {
				addToDiag(ctx, diag.FromErr(err))
				return
			}
		case field.OneToOne:
			if err := d.Set("relationship_to_one", validationRule.(*epcc.ValidationRuleRelationshipAttribute).To); err != nil {
				addToDiag(ctx, diag.FromErr(err))
				return
			}
		}
	}

	d.SetId(result.Data.Id)
}
