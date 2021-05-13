package epcc

import (
	"bytes"
	"encoding/json"
	"fmt"
)

var Pricebooks pricebooks

type pricebooks struct{}

type Pricebook struct {
	Id         string              `json:"id,omitempty"`
	Type       string              `json:"type"`
	Attributes PricebookAttributes `json:"attributes"`
	Links      PricebookLinks      `json:"links,omitempty"`
}

type PricebookAttributes struct {
	CreatedAt   string `json:"created_at,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

type PricebookLinks struct {
	Self string `json:"self,omitempty"`
}

func (pricebooks) Get(client *Client, pricebookId string) (*PricebookData, ApiErrors) {
	path := fmt.Sprintf("/pcm/pricebooks/%s", pricebookId)

	body, apiError := client.DoRequest("GET", path, nil)
	if apiError != nil {
		return nil, apiError
	}

	// TODO Better Manage Parent ID
	var pricebooks PricebookData
	if err := json.Unmarshal(body, &pricebooks); err != nil {
		return nil, FromError(err)
	}

	return &pricebooks, nil
}

// GetAll fetches all pricebooks
func (pricebooks) GetAll(client *Client) (*PricebookList, ApiErrors) {
	path := fmt.Sprintf("/pcm/pricebooks")

	body, apiError := client.DoRequest("GET", path, nil)
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
func (pricebooks) Create(client *Client, pricebook *Pricebook) (*PricebookData, ApiErrors) {
	pricebookData := PricebookData{
		Data: *pricebook,
	}

	jsonPayload, err := json.Marshal(pricebookData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/pcm/pricebooks")

	body, apiError := client.DoRequest("POST", path, bytes.NewBuffer(jsonPayload))
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
func (pricebooks) Delete(client *Client, pricebookID string) ApiErrors {
	path := fmt.Sprintf("/pcm/pricebooks/%s", pricebookID)

	if _, err := client.DoRequest("DELETE", path, nil); err != nil {
		return err
	}

	return nil
}

// Update updates a pricebook.
func (pricebooks) Update(client *Client, pricebookID string, pricebook *Pricebook) (*PricebookData, ApiErrors) {

	pricebookData := PricebookData{
		Data: *pricebook,
	}

	jsonPayload, err := json.Marshal(pricebookData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/pcm/pricebooks/%s", pricebookID)

	body, apiError := client.DoRequest("PUT", path, bytes.NewBuffer(jsonPayload))
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
