package epcc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

var Pricebooks pricebooks

type pricebooks struct{}

type Pricebook struct {
	Id         string              `json:"id,omitempty"`
	Type       string              `json:"type"`
	Attributes PricebookAttributes `json:"attributes"`
}

type PricebookAttributes struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

func (pricebooks) Get(ctx *context.Context, client *Client, pricebookId string) (*PricebookData, ApiErrors) {
	path := fmt.Sprintf("/pcm/pricebooks/%s", pricebookId)

	body, apiError := client.DoRequest(ctx, "GET", path, "", nil)
	if apiError != nil {
		return nil, apiError
	}

	var pricebook PricebookData
	if err := json.Unmarshal(body, &pricebook); err != nil {
		return nil, FromError(err)
	}

	return &pricebook, nil
}

// GetAll fetches all pricebooks
func (pricebooks) GetAll(ctx *context.Context, client *Client) (*PricebookList, ApiErrors) {
	path := fmt.Sprintf("/pcm/pricebooks")

	body, apiError := client.DoRequest(ctx, "GET", path, "", nil)
	if apiError != nil {
		return nil, apiError
	}

	var pricebooks PricebookList
	if err := json.Unmarshal(body, &pricebooks); err != nil {
		return nil, FromError(err)
	}

	return &pricebooks, nil
}

// Create creates a pricebook
func (pricebooks) Create(ctx *context.Context, client *Client, pricebook *Pricebook) (*PricebookData, ApiErrors) {
	pricebookData := PricebookData{
		Data: *pricebook,
	}

	jsonPayload, err := json.Marshal(pricebookData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/pcm/pricebooks")

	body, apiError := client.DoRequest(ctx, "POST", path, "", bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}

	var newPricebook PricebookData
	if err := json.Unmarshal(body, &newPricebook); err != nil {
		return nil, FromError(err)
	}

	return &newPricebook, nil
}

// Delete deletes a pricebook.
func (pricebooks) Delete(ctx *context.Context, client *Client, pricebookID string) ApiErrors {
	path := fmt.Sprintf("/pcm/pricebooks/%s", pricebookID)

	if _, err := client.DoRequest(ctx, "DELETE", path, "", nil); err != nil {
		return err
	}

	return nil
}

// Update updates a pricebook.
func (pricebooks) Update(ctx *context.Context, client *Client, pricebookID string, pricebook *Pricebook) (*PricebookData, ApiErrors) {

	pricebookData := PricebookData{
		Data: *pricebook,
	}

	jsonPayload, err := json.Marshal(pricebookData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/pcm/pricebooks/%s", pricebookID)

	body, apiError := client.DoRequest(ctx, "PUT", path, "", bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}

	var updatedPricebook PricebookData
	if err := json.Unmarshal(body, &updatedPricebook); err != nil {
		return nil, FromError(err)
	}

	return &updatedPricebook, nil
}

type PricebookData struct {
	Data Pricebook `json:"data"`
}

// PricebookMeta contains extra data for an pricebook
type PricebookMeta struct {
	Timestamps Timestamps `json:"timestamps,omitempty"`
}

type PricebookDataList struct {
}

type PricebookList struct {
	Data []Pricebook
}
