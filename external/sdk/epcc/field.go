package epcc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elasticpath/epcc-terraform-provider/external/sdk/epcc/field"
)

var Fields fields

type fields struct{}

type FieldData struct {
	Data Field `json:"data"`
}

type FieldList struct {
	Data []Field
}

type Field struct {
	Id              string                     `json:"id,omitempty"`
	Type            string                     `json:"type"`
	FieldType       string                     `json:"field_type"`
	Slug            string                     `json:"slug"`
	Name            string                     `json:"name"`
	Description     string                     `json:"description"`
	Required        bool                       `json:"required"`
	Default         string                     `json:"default,omitempty"`
	Enabled         bool                       `json:"enabled"`
	Order           int                        `json:"order,omitempty"`
	OmitNull        bool                       `json:"omit_null,omitempty"`
	ValidationRules []ValidationRuleAttribute  `json:"validation_rules,omitempty"`
	Relationships   *FlowRelationshipAttribute `json:"relationships"`
}

type ValidationRuleAttribute interface {
	ValidationType() field.ValidationType
}

type ValidationRuleAttributeBasic struct {
	Type string `json:"type"`
}

func (p *ValidationRuleAttributeBasic) ValidationType() field.ValidationType {
	return field.ValidationType(p.Type)
}

type ValidationRuleStringEnumAttribute struct {
	ValidationRuleAttributeBasic
	Options []string `json:"options"`
}

type ValidationRuleIntegerEnumAttribute struct {
	ValidationRuleAttributeBasic
	Options []int `json:"options"`
}

type ValidationRuleFloatEnumAttribute struct {
	ValidationRuleAttributeBasic
	Options []float64 `json:"options"`
}

type ValidationRuleRelationshipAttribute struct {
	ValidationRuleAttributeBasic
	To string `json:"to"`
}

type ValidationRuleBetweenIntegersAttribute struct {
	ValidationRuleAttributeBasic
	Options BetweenIntegers `json:"options"`
}

type BetweenIntegers struct {
	From int `json:"from"`
	To   int `json:"to"`
}

type ValidationRuleBetweenFloatsAttribute struct {
	ValidationRuleAttributeBasic
	Options BetweenFloats `json:"options"`
}

type BetweenFloats struct {
	From float64 `json:"from"`
	To   float64 `json:"to"`
}

type FlowRelationshipAttribute struct {
	Flow *FlowRelationshipAttributeData `json:"flow"`
}

type FlowRelationshipAttributeData struct {
	Data *FlowRelationship `json:"data,omitempty"`
}

type FlowRelationship struct {
	Id   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
}

func (fields) Get(ctx *context.Context, client *Client, id string) (*FieldData, ApiErrors) {
	path := fmt.Sprintf("/v2/fields/%s", id)

	body, apiError := client.DoRequest(ctx, "GET", path, nil)
	if apiError != nil {
		return nil, apiError
	}

	var data FieldData
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, FromError(err)
	}

	return &data, nil
}

func (fields) GetAll(ctx *context.Context, client *Client) (*FieldList, ApiErrors) {
	path := fmt.Sprintf("/v2/fields")

	body, apiError := client.DoRequest(ctx, "GET", path, nil)
	if apiError != nil {
		return nil, apiError
	}

	var list FieldList
	if err := json.Unmarshal(body, &list); err != nil {
		return nil, FromError(err)
	}

	return &list, nil
}

// GetFieldsForFlow Find all the fields that extend flow slug (slug is a resource name)
func (fields) GetFieldsForFlow(ctx *context.Context, client *Client, flowSlug string) (*FieldList, ApiErrors) {
	path := fmt.Sprintf("/v2/flows/%s/fields", flowSlug)

	body, apiError := client.DoRequest(ctx, "GET", path, nil)
	if apiError != nil {
		return nil, apiError
	}

	var fields FieldList
	if err := json.Unmarshal(body, &fields); err != nil {
		return nil, FromError(err)
	}

	return &fields, nil
}

func (fields) Create(ctx *context.Context, client *Client, data *Field) (*FieldData, ApiErrors) {
	fieldsData := FieldData{
		Data: *data,
	}

	jsonPayload, err := json.Marshal(fieldsData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/v2/fields")

	body, apiError := client.DoRequest(ctx, "POST", path, bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}

	var newField FieldData
	if err := json.Unmarshal(body, &newField); err != nil {
		return nil, FromError(err)
	}

	return &newField, nil
}

func (fields) Delete(ctx *context.Context, client *Client, id string) ApiErrors {
	path := fmt.Sprintf("/v2/fields/%s", id)

	if _, err := client.DoRequest(ctx, "DELETE", path, nil); err != nil {
		return err
	}

	return nil
}

func (fields) Update(ctx *context.Context, client *Client, fieldID string, data *Field) (*FieldData, ApiErrors) {
	fieldData := FieldData{
		Data: *data,
	}

	jsonPayload, err := json.Marshal(fieldData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/v2/fields/%s", fieldID)

	body, apiError := client.DoRequest(ctx, "PUT", path, bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}

	var updatedField FieldData
	if err := json.Unmarshal(body, &updatedField); err != nil {
		return nil, FromError(err)
	}

	return &updatedField, nil
}

func (f *Field) UnmarshalJSON(body []byte) error {
	var fieldRawMap map[string]*json.RawMessage
	err := json.Unmarshal(body, &fieldRawMap)
	if err != nil {
		return err
	}

	if err = unmarshalRaw(fieldRawMap, "id", &f.Id); err != nil {
		return err
	}
	if err = unmarshalRaw(fieldRawMap, "type", &f.Type); err != nil {
		return err
	}
	if err = unmarshalRaw(fieldRawMap, "field_type", &f.FieldType); err != nil {
		return err
	}
	if err = unmarshalRaw(fieldRawMap, "slug", &f.Slug); err != nil {
		return err
	}
	if err = unmarshalRaw(fieldRawMap, "name", &f.Name); err != nil {
		return err
	}
	if err = unmarshalRaw(fieldRawMap, "description", &f.Description); err != nil {
		return err
	}
	if err = unmarshalRaw(fieldRawMap, "required", &f.Required); err != nil {
		return err
	}
	if err = unmarshalRaw(fieldRawMap, "default", &f.Default); err != nil {
		return err
	}
	if err = unmarshalRaw(fieldRawMap, "enabled", &f.Enabled); err != nil {
		return err
	}
	if err = unmarshalRaw(fieldRawMap, "order", &f.Order); err != nil {
		return err
	}
	if err = unmarshalRaw(fieldRawMap, "omit_null", &f.OmitNull); err != nil {
		return err
	}
	if err = unmarshalRaw(fieldRawMap, "relationships", &f.Relationships); err != nil {
		return err
	}

	validationRulesData, ok := fieldRawMap["validation_rules"]
	if !ok {
		return nil
	}

	var ruleRawMap []*json.RawMessage
	err = json.Unmarshal(*validationRulesData, &ruleRawMap)
	if err != nil {
		return err
	}

	var bases []ValidationRuleAttributeBasic
	err = json.Unmarshal(*validationRulesData, &bases)
	if err != nil {
		return err
	}

	fieldType := field.Type(f.FieldType)

	for i, base := range bases {
		switch base.ValidationType() {
		case field.Email:
			fallthrough
		case field.Slug:
			fallthrough
		case field.Uuid:
			var concrete ValidationRuleAttributeBasic
			err = json.Unmarshal(*ruleRawMap[i], &concrete)
			if err != nil {
				return err
			}
			f.ValidationRules = append(f.ValidationRules, &concrete)
		case field.Enum:
			if fieldType == field.String {
				var concrete ValidationRuleStringEnumAttribute
				err = json.Unmarshal(*ruleRawMap[i], &concrete)
				if err != nil {
					return err
				}
				f.ValidationRules = append(f.ValidationRules, &concrete)
			} else if fieldType == field.Integer {
				var concrete ValidationRuleIntegerEnumAttribute
				err = json.Unmarshal(*ruleRawMap[i], &concrete)
				if err != nil {
					return err
				}
				f.ValidationRules = append(f.ValidationRules, &concrete)
			} else if fieldType == field.Float {
				var concrete ValidationRuleFloatEnumAttribute
				err = json.Unmarshal(*ruleRawMap[i], &concrete)
				if err != nil {
					return err
				}
				f.ValidationRules = append(f.ValidationRules, &concrete)
			} else {
				return fmt.Errorf("unknown enum field type %v", fieldType)
			}
		case field.Between:
			if fieldType == field.Integer {
				var concrete ValidationRuleBetweenIntegersAttribute
				err = json.Unmarshal(*ruleRawMap[i], &concrete)
				if err != nil {
					return err
				}
				f.ValidationRules = append(f.ValidationRules, &concrete)
			} else if fieldType == field.Float {
				var concrete ValidationRuleBetweenFloatsAttribute
				err = json.Unmarshal(*ruleRawMap[i], &concrete)
				if err != nil {
					return err
				}
				f.ValidationRules = append(f.ValidationRules, &concrete)
			} else {
				return fmt.Errorf("unknown from-to field type %v", fieldType)
			}
		case field.OneToMany:
			fallthrough
		case field.OneToOne:
			var concrete ValidationRuleRelationshipAttribute
			err = json.Unmarshal(*ruleRawMap[i], &concrete)
			if err != nil {
				return err
			}
			f.ValidationRules = append(f.ValidationRules, &concrete)
		default:
			return fmt.Errorf("unknown validation type %v", base.ValidationType())
		}
	}

	return nil
}

func unmarshalRaw(m map[string]*json.RawMessage, key string, target interface{}) error {
	data, ok := m[key]
	if ok && data != nil {
		return json.Unmarshal(*data, &target)
	}
	return nil
}
