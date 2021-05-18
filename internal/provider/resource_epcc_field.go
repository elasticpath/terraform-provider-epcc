package provider

import (
	"context"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc/field"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEpccField() *schema.Resource {
	return &schema.Resource{
		Description:   "Represents the EPCC API [Fields Object](https://documentation.elasticpath.com/commerce-cloud/docs/api/advanced/custom-data/fields/index.html).",
		CreateContext: addDiagToContext(resourceEpccFieldCreate),
		ReadContext:   addDiagToContext(resourceEpccFieldRead),
		UpdateContext: addDiagToContext(resourceEpccFieldUpdate),
		DeleteContext: addDiagToContext(resourceEpccFieldDelete),
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"field_type": {
				Type: schema.TypeString,
				ValidateDiagFunc: func(value interface{}, path cty.Path) diag.Diagnostics {
					v := value.(string)
					err := field.Type(v).Validate()
					if err != nil {
						return diag.FromErr(err)
					}
					return nil
				},
				Required: true,
			},
			"slug": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"required": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"default": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"order": {
				Type:     schema.TypeInt,
				Default:  1,
				Optional: true,
			},
			"omit_null": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"flow_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"valid_string_format": {
				Type: schema.TypeString,
				ValidateDiagFunc: func(value interface{}, path cty.Path) diag.Diagnostics {
					v := value.(string)
					fieldType := field.ValidationType(v)
					if fieldType != field.Email && fieldType != field.Slug && fieldType != field.Uuid {
						return diag.Errorf("only \"email\", \"slug\" and \"uuid\" are allowed")
					}
					return nil
				},
				Optional: true,
				ConflictsWith: []string{
					"valid_int_enum",
					"valid_float_enum",
					"valid_int_range",
					"valid_float_range",
					"relationship_to_one",
					"relationship_to_many",
				},
			},
			"valid_string_enum": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				ConflictsWith: []string{
					"valid_int_enum",
					"valid_float_enum",
					"valid_int_range",
					"valid_float_range",
					"relationship_to_one",
					"relationship_to_many",
				},
			},
			"valid_int_enum": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Optional: true,
				ConflictsWith: []string{
					"valid_string_format",
					"valid_string_enum",
					"valid_float_enum",
					"valid_float_range",
					"relationship_to_one",
					"relationship_to_many",
				},
			},
			"valid_int_range": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeMap,
					Elem: &schema.Schema{
						Type: schema.TypeInt,
					},
					ValidateDiagFunc: func(value interface{}, path cty.Path) diag.Diagnostics {
						v := value.(map[string]interface{})
						_, fromPresent := v["from"]
						_, toPresent := v["to"]
						if !fromPresent || !toPresent {
							return diag.Errorf("both \"from\" and \"to\" values are expected")
						}
						if len(v) > 2 {
							return diag.Errorf("only \"from\" and \"to\" fields are allowed")
						}
						return nil
					},
				},
				Optional: true,
				ConflictsWith: []string{
					"valid_string_format",
					"valid_string_enum",
					"valid_float_enum",
					"valid_float_range",
					"relationship_to_one",
					"relationship_to_many",
				},
			},
			"valid_float_enum": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeFloat,
				},
				Optional: true,
				ConflictsWith: []string{
					"valid_string_format",
					"valid_string_enum",
					"valid_int_enum",
					"valid_int_range",
					"relationship_to_one",
					"relationship_to_many",
				},
			},
			"valid_float_range": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeMap,
					Elem: &schema.Schema{
						Type: schema.TypeFloat,
					},
					ValidateDiagFunc: func(value interface{}, path cty.Path) diag.Diagnostics {
						v := value.(map[string]interface{})
						_, fromPresent := v["from"]
						_, toPresent := v["to"]
						if !fromPresent || !toPresent {
							return diag.Errorf("both \"from\" and \"to\" values are expected")
						}
						if len(v) > 2 {
							return diag.Errorf("only \"from\" and \"to\" fields are allowed")
						}
						return nil
					},
				},
				Optional: true,
				ConflictsWith: []string{
					"valid_string_format",
					"valid_string_enum",
					"valid_int_enum",
					"valid_int_range",
					"relationship_to_one",
					"relationship_to_many",
				},
			},
			"relationship_to_one": {
				Type:     schema.TypeString,
				Optional: true,
				ConflictsWith: []string{
					"valid_string_format",
					"valid_string_enum",
					"valid_int_enum",
					"valid_float_enum",
					"valid_int_range",
					"valid_float_range",
					"relationship_to_many",
				},
			},
			"relationship_to_many": {
				Type:     schema.TypeString,
				Optional: true,
				ConflictsWith: []string{
					"valid_string_format",
					"valid_string_enum",
					"valid_int_enum",
					"valid_float_enum",
					"valid_int_range",
					"valid_float_range",
					"relationship_to_one",
				},
			},
		},
	}
}

func resourceEpccFieldDelete(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	err := epcc.Fields.Delete(&ctx, client, d.Id())
	if err != nil {
		ReportAPIError(ctx, err)
	}

	d.SetId("")
}

func resourceEpccFieldUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	f, diagErr := constructField(d)
	if diagErr != nil {
		addToDiag(ctx, diagErr)
	}

	updated, apiError := epcc.Fields.Update(&ctx, client, d.Id(), f)
	if apiError != nil {
		ReportAPIError(ctx, apiError)
		return
	}

	d.SetId(updated.Data.Id)

	resourceEpccFieldRead(ctx, d, m)
}

func resourceEpccFieldRead(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	result, err := epcc.Fields.Get(&ctx, client, d.Id())
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
				addToDiag(ctx,diag.Errorf("unknown enum for field type %v", fieldType))
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
}

func resourceEpccFieldCreate(ctx context.Context, d *schema.ResourceData, m interface{}) {
	client := m.(*epcc.Client)

	f, diagErr := constructField(d)
	if diagErr != nil {
		addToDiag(ctx,  diagErr)
		return
	}

	created, apiError := epcc.Fields.Create(&ctx, client, f)
	if apiError != nil {
		ReportAPIError(ctx, apiError)
		return
	}

	d.SetId(created.Data.Id)
}

func constructField(d *schema.ResourceData) (*epcc.Field, diag.Diagnostics) {
	validationRules, diagErr := parseValidationRules(d)
	if diagErr != nil {
		return nil, diagErr
	}

	return &epcc.Field{
		Type:            "field",
		FieldType:       d.Get("field_type").(string),
		Slug:            d.Get("slug").(string),
		Name:            d.Get("name").(string),
		Description:     d.Get("description").(string),
		Required:        d.Get("required").(bool),
		Default:         d.Get("default").(string),
		Enabled:         d.Get("enabled").(bool),
		Order:           d.Get("order").(int),
		OmitNull:        d.Get("omit_null").(bool),
		ValidationRules: validationRules,
		Relationships: &epcc.FlowRelationshipAttribute{
			Flow: &epcc.FlowRelationshipAttributeData{
				Data: &epcc.FlowRelationship{
					Id:   d.Get("flow_id").(string),
					Type: "flow",
				},
			},
		},
	}, nil
}

func parseValidationRules(d *schema.ResourceData) ([]epcc.ValidationRuleAttribute, diag.Diagnostics) {
	var validationRules []epcc.ValidationRuleAttribute
	var diags diag.Diagnostics

	fieldType := field.Type(d.Get("field_type").(string))

	validStringFormat, hasValidStringFormat := d.GetOk("valid_string_format")
	if hasValidStringFormat {
		if fieldType == field.String {
			validationRules = append(validationRules, &epcc.ValidationRuleAttributeBasic{
				Type: validStringFormat.(string),
			})
		} else {
			diags = append(diags, diag.Errorf("valid_string_format can only be used with string field_type")...)
		}
	}
	validStringEnum, hasValidStringEnum := d.GetOk("valid_string_enum")
	if hasValidStringEnum {
		if fieldType == field.String || fieldType == field.Date {
			validationRules = append(validationRules, &epcc.ValidationRuleStringEnumAttribute{
				ValidationRuleAttributeBasic: epcc.ValidationRuleAttributeBasic{
					Type: field.Enum.AsString(),
				},
				Options: convertArrayToStringSlice(validStringEnum.([]interface{})),
			})
		} else {
			diags = append(diags, diag.Errorf("valid_string_enum can only be used with string or date field_type")...)
		}
	}

	validIntEnum, hasValidIntEnum := d.GetOk("valid_int_enum")
	validIntRange, hasValidIntRange := d.GetOk("valid_int_range")
	if hasValidIntEnum || hasValidIntRange {
		if fieldType != field.Integer {
			diags = append(diags, diag.Errorf("valid_int_enum, valid_int_range can only be used with integer field_type")...)
		} else {
			if hasValidIntEnum {
				validationRules = append(validationRules, &epcc.ValidationRuleIntegerEnumAttribute{
					ValidationRuleAttributeBasic: epcc.ValidationRuleAttributeBasic{
						Type: field.Enum.AsString(),
					},
					Options: convertArrayToIntSlice(validIntEnum.([]interface{})),
				})
			}
			if hasValidIntRange {
				elements := validIntRange.([]interface{})
				for _, element := range elements {
					intRange := element.(map[string]interface{})
					validationRules = append(validationRules, &epcc.ValidationRuleBetweenIntegersAttribute{
						ValidationRuleAttributeBasic: epcc.ValidationRuleAttributeBasic{
							Type: field.Between.AsString(),
						},
						Options: epcc.BetweenIntegers{
							From: intRange["from"].(int),
							To:   intRange["to"].(int),
						},
					})
				}
			}
		}
	}

	validFloatEnum, hasValidFloatEnum := d.GetOk("valid_float_enum")
	validFloatRange, hasValidFloatRange := d.GetOk("valid_float_range")
	if hasValidFloatEnum || hasValidFloatRange {
		if fieldType != field.Float {
			diags = append(diags, diag.Errorf("valid_float_enum, valid_float_range can only be used with float field_type")...)
		} else {
			if hasValidFloatEnum {
				validationRules = append(validationRules, &epcc.ValidationRuleFloatEnumAttribute{
					ValidationRuleAttributeBasic: epcc.ValidationRuleAttributeBasic{
						Type: field.Enum.AsString(),
					},
					Options: convertArrayToFloatSlice(validFloatEnum.([]interface{})),
				})
			}
			if hasValidFloatRange {
				elements := validFloatRange.([]interface{})
				for _, element := range elements {
					floatRange := element.(map[string]interface{})
					validationRules = append(validationRules, &epcc.ValidationRuleBetweenFloatsAttribute{
						ValidationRuleAttributeBasic: epcc.ValidationRuleAttributeBasic{
							Type: field.Between.AsString(),
						},
						Options: epcc.BetweenFloats{
							From: floatRange["from"].(float64),
							To:   floatRange["to"].(float64),
						},
					})
				}
			}
		}
	}

	relToOne, hasRelToOne := d.GetOk("relationship_to_one")
	relToMany, hasRelToMany := d.GetOk("relationship_to_many")
	if hasRelToOne || hasRelToMany {
		if fieldType != field.Relationship {
			diags = append(diags, diag.Errorf("relationship_to_one, relationship_to_many can only be used with relationship field_type")...)
		} else {
			if hasRelToOne {
				validationRules = append(validationRules, &epcc.ValidationRuleRelationshipAttribute{
					ValidationRuleAttributeBasic: epcc.ValidationRuleAttributeBasic{
						Type: field.OneToOne.AsString(),
					},
					To: relToOne.(string),
				})
			} else {
				validationRules = append(validationRules, &epcc.ValidationRuleRelationshipAttribute{
					ValidationRuleAttributeBasic: epcc.ValidationRuleAttributeBasic{
						Type: field.OneToMany.AsString(),
					},
					To: relToMany.(string),
				})
			}
		}
	}

	return validationRules, diags
}
