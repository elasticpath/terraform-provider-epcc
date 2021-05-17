package epcc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

var Integrations integrations

type integrations struct{}

type IntegrationData struct {
	Data Integration `json:"data"`
}

type IntegrationConfiguration struct {
	Url       string `json:"url"`
	SecretKey string `json:"secret_key,omitempty"`
}

type integrationObjectType string
type integrationType string

const (
	IntegrationType integrationObjectType = "integration"
	Webhook         integrationType       = "webhook"
)

type Integration struct {
	Id              string                   `json:"id,omitempty"`
	Type            integrationObjectType    `json:"type"`
	IntegrationType integrationType          `json:"integration_type"`
	Name            string                   `json:"name"`
	Description     string                   `json:"description,omitempty"`
	Enabled         bool                     `json:"enabled"`
	Configuration   IntegrationConfiguration `json:"configuration"`
	Observes        []string                 `json:"observes"`
	Meta            IntegrationMeta          `json:"-"`
}

// IntegrationMeta contains extra data for an Integration
type IntegrationMeta struct {
	Timestamps Timestamps `json:"timestamps,omitempty"`
}

type IntegrationList struct {
	Data []Integration
}

// Get fetches one Integration by id
func (integrations) Get(ctx *context.Context, client *Client, id string) (*IntegrationData, ApiErrors) {
	path := fmt.Sprintf("/v2/integrations/%s", id)

	body, apiError := client.DoRequest(ctx, "GET", path, "", nil)
	if apiError != nil {
		return nil, apiError
	}

	var data IntegrationData
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, FromError(err)
	}

	return &data, nil
}

// GetAll fetches IntegrationList
func (integrations) GetAll(ctx *context.Context, client *Client) (*IntegrationList, ApiErrors) {
	path := fmt.Sprintf("/v2/integrations")

	body, apiError := client.DoRequest(ctx, "GET", path, "", nil)
	if apiError != nil {
		return nil, apiError
	}

	var list IntegrationList
	if err := json.Unmarshal(body, &list); err != nil {
		return nil, FromError(err)
	}

	return &list, nil
}

// Create creates an Integration
func (integrations) Create(ctx *context.Context, client *Client, integration *Integration) (*IntegrationData, ApiErrors) {
	data := IntegrationData{
		Data: *integration,
	}

	jsonPayload, err := json.Marshal(data)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/v2/integrations")

	body, apiError := client.DoRequest(ctx, "POST", path, "", bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}

	var newIntegration IntegrationData
	if err := json.Unmarshal(body, &newIntegration); err != nil {
		return nil, FromError(err)
	}

	return &newIntegration, nil
}

// Delete deletes an Integration by id
func (integrations) Delete(ctx *context.Context, client *Client, id string) ApiErrors {
	path := fmt.Sprintf("/v2/integrations/%s", id)

	if _, err := client.DoRequest(ctx, "DELETE", path, "", nil); err != nil {
		return err
	}

	return nil
}

// Update updates a Integration.
func (integrations) Update(ctx *context.Context, client *Client, id string, integration *Integration) (*IntegrationData, ApiErrors) {

	data := IntegrationData{
		Data: *integration,
	}

	jsonPayload, err := json.Marshal(data)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/v2/integrations/%s", id)

	body, apiError := client.DoRequest(ctx, "PUT", path, "", bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}

	var updatedIntegration IntegrationData
	if err := json.Unmarshal(body, &updatedIntegration); err != nil {
		return nil, FromError(err)
	}

	return &updatedIntegration, nil
}
