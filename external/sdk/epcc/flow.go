package epcc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

var Flows flows

type flows struct{}

type Flow struct {
	Id          string `json:"id,omitempty"`
	Type        string `json:"type"`
	Name        string `json:"name,omitempty"`
	Slug        string `json:"slug,omitempty"`
	Description string `json:"description,omitempty"`
	Enabled     bool   `json:"enabled"`
}

func (flows) Get(ctx *context.Context, client *Client, flowId string) (*FlowData, ApiErrors) {
	path := fmt.Sprintf("/v2/flows/%s", flowId)

	body, apiError := client.DoRequest(ctx, "GET", path, nil)
	if apiError != nil {
		return nil, apiError
	}

	// TODO Better Manage Parent ID
	var flows FlowData
	if err := json.Unmarshal(body, &flows); err != nil {
		return nil, FromError(err)
	}

	return &flows, nil
}

// GetAll fetches all flows
func (flows) GetAll(ctx *context.Context, client *Client) (*FlowList, ApiErrors) {
	path := fmt.Sprintf("/v2/flows")

	body, apiError := client.DoRequest(ctx, "GET", path, nil)
	if apiError != nil {
		return nil, apiError
	}

	var flows FlowList
	if err := json.Unmarshal(body, &flows); err != nil {
		return nil, FromError(err)
	}

	return &flows, nil
}

// Create creates a flow
func (flows) Create(ctx *context.Context, client *Client, flow *Flow) (*FlowData, ApiErrors) {
	flowData := FlowData{
		Data: *flow,
	}

	jsonPayload, err := json.Marshal(flowData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/v2/flows")

	body, apiError := client.DoRequest(ctx, "POST", path, bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}

	var newFlow FlowData
	if err := json.Unmarshal(body, &newFlow); err != nil {
		return nil, FromError(err)
	}

	return &newFlow, nil
}

// Delete deletes a flow.
func (flows) Delete(ctx *context.Context, client *Client, flowID string) ApiErrors {
	path := fmt.Sprintf("/v2/flows/%s", flowID)

	if _, err := client.DoRequest(ctx, "DELETE", path, nil); err != nil {
		return err
	}

	return nil
}

// Update updates a flow.
func (flows) Update(ctx *context.Context, client *Client, flowID string, flow *Flow) (*FlowData, ApiErrors) {

	flowData := FlowData{
		Data: *flow,
	}

	jsonPayload, err := json.Marshal(flowData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/v2/flows/%s", flowID)

	body, apiError := client.DoRequest(ctx, "PUT", path, bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}

	var updatedFlow FlowData
	if err := json.Unmarshal(body, &updatedFlow); err != nil {
		return nil, FromError(err)
	}

	return &updatedFlow, nil
}

type FlowData struct {
	Data Flow `json:"data"`
}

// FlowMeta contains extra data for a flow
type FlowMeta struct {
	Timestamps Timestamps `json:"timestamps,omitempty"`
}

type FlowDataList struct {
}

type FlowList struct {
	Data []Flow
}
