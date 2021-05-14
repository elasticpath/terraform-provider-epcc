package epcc

import (
	"bytes"
	"encoding/json"
	"fmt"
)

var Fields fields

type fields struct{}

type Field struct {
	Id          string `json:"id,omitempty"`
	Type        string `json:"type,omitempty"`
	FieldType   string `json:"field_type,omitempty"`
	Slug        string `json:"slug,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Required    bool   `json:"required,omitempty"`
	Default     string `json:"default,omitempty"`
	Enabled     bool   `json:"enabled,omitempty"`
	Order       int    `json:"order,omitempty"`
	OmitNull    bool   `json:"omit_null,omitempty"`
	//ValidationRules ValidationRulesAttribute    `json:"validation_rules,omitempty"` // TODO support at some point
	Relationships *RelationshipsAttribute `json:"relationships,omitempty"`
}

type RelationshipsAttribute struct {
	Flow *RelationshipsAttributeFlow `json:"flow,omitempty"`
}

type RelationshipsAttributeFlow struct {
	Data *RelationshipsAttributeFlowData `json:"data,omitempty"`
}

type RelationshipsAttributeFlowData struct {
	Id              string `json:"id,omitempty"`
	Type            string `json:"type,omitempty"`
}

func (fields) Get(client *Client, fieldId string) (*FieldData, ApiErrors) {
	path := fmt.Sprintf("/v2/fields/%s", fieldId)

	body, apiError := client.DoRequest("GET", path, nil)
	if apiError != nil {
		return nil, apiError
	}

	// TODO Better Manage Parent ID
	var field FieldData
	if err := json.Unmarshal(body, &field); err != nil {
		return nil, FromError(err)
	}

	return &field, nil
}

func (fields) GetAll(client *Client) (*FieldList, ApiErrors) {
	path := fmt.Sprintf("/v2/fields")

	body, apiError := client.DoRequest("GET", path, nil)
	if apiError != nil {
		return nil, apiError
	}

	var fields FieldList
	if err := json.Unmarshal(body, &fields); err != nil {
		return nil, FromError(err)
	}

	return &fields, nil
}

// Get all the fields that extend flow slug (slug is a resource name)
func (fields) GetAllFromFlowSlug(client *Client, flowSlug string) (*FieldList, ApiErrors) {
	path := fmt.Sprintf("/v2/flows/%s/fields", flowSlug)

	body, apiError := client.DoRequest("GET", path, nil)
	if apiError != nil {
		return nil, apiError
	}

	var fields FieldList
	if err := json.Unmarshal(body, &fields); err != nil {
		return nil, FromError(err)
	}

	return &fields, nil
}

func (fields) Create(client *Client, field *Field) (*FieldData, ApiErrors) {
	fieldsData := FieldData{
		Data: *field,
	}

	jsonPayload, err := json.Marshal(fieldsData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/v2/fields")

	body, apiError := client.DoRequest("POST", path, bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}

	var newField FieldData
	if err := json.Unmarshal(body, &newField); err != nil {
		return nil, FromError(err)
	}

	return &newField, nil
}

func (fields) Delete(client *Client, fieldID string) ApiErrors {
	path := fmt.Sprintf("/v2/fields/%s", fieldID)

	if _, err := client.DoRequest("DELETE", path, nil); err != nil {
		return err
	}

	return nil
}

func (fields) Update(client *Client, fieldID string, field *Field) (*FieldData, ApiErrors) {

	fieldData := FieldData{
		Data: *field,
	}

	jsonPayload, err := json.Marshal(fieldData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/v2/fields/%s", fieldID)

	body, apiError := client.DoRequest("PUT", path, bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}

	var updatedField FieldData
	if err := json.Unmarshal(body, &updatedField); err != nil {
		return nil, FromError(err)
	}

	return &updatedField, nil
}

type FieldData struct {
	Data Field `json:"data"`
}

type FieldMeta struct {
	Timestamps Timestamps `json:"timestamps,omitempty"`
}

type FiledDataList struct {
}

type FieldList struct {
	Data [] Field
}
