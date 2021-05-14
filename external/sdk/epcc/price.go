package epcc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

var Prices prices

type prices struct{}

type Price struct {
	Type       string          `json:"type"`
	Id         string          `json:"id,omitempty"`
	Attributes PriceAttributes `json:"attributes"`
	Sku        string          `json:"sku"`
}

type PriceInCurrency struct {
	Amount      string `json:"amount"`
	IncludesTax bool   `json:"includes_tax,omitempty"`
}

type PriceAttributes struct {
	Currencies map[string]PriceInCurrency `json:"currencies"`
}

func (prices) Get(client *Client, pricebookId string, priceId string) (*PriceData, ApiErrors) {
	path := fmt.Sprintf("/pcm/pricebooks/%s/prices/%s", pricebookId, priceId)

	body, apiError := client.DoRequest("GET", path, nil)
	if apiError != nil {
		return nil, apiError
	}

	var price PriceData
	if err := json.Unmarshal(body, &price); err != nil {
		return nil, FromError(err)
	}

	return &price, nil
}

// GetAll fetches all prices
func (prices) GetAll(client *Client, pricebookId string) (*PriceList, ApiErrors) {
	path := fmt.Sprintf("/pcm/pricebooks/%s/prices", pricebookId)

	body, apiError := client.DoRequest("GET", path, nil)
	if apiError != nil {
		return nil, apiError
	}

	var prices PriceList
	if err := json.Unmarshal(body, &prices); err != nil {
		return nil, FromError(err)
	}

	return &prices, nil
}

// Create creates a price
func (prices) Create(client *Client, pricebookId string, productPrice *Price) (*PriceData, ApiErrors) {
	priceData := PriceData{
		Data: *productPrice,
	}

	jsonPayload, err := json.Marshal(priceData)
	if err != nil {
		return nil, FromError(err)
	}
	log.Printf("jsonPayload: " + string(jsonPayload))

	path := fmt.Sprintf("/pcm/pricebooks/%s/prices", pricebookId)

	body, apiError := client.DoRequest("POST", path, bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}

	var newPrice PriceData
	if err := json.Unmarshal(body, &newPrice); err != nil {
		return nil, FromError(err)
	}

	return &newPrice, nil
}

// Delete deletes a price.
func (prices) Delete(client *Client, pricebookId string, priceId string) ApiErrors {
	path := fmt.Sprintf("/pcm/pricebooks/%s/prices/%s", pricebookId, priceId)

	if _, err := client.DoRequest("DELETE", path, nil); err != nil {
		return err
	}

	return nil
}

// Update updates a price.
func (prices) Update(client *Client, pricebookId string, priceId string, productPrice *Price) (*PriceData, ApiErrors) {

	priceData := PriceData{
		Data: *productPrice,
	}

	jsonPayload, err := json.Marshal(priceData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/pcm/pricebooks/%s/prices/%s", pricebookId, priceId)

	body, apiError := client.DoRequest("PUT", path, bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}

	var updatedPrice PriceData
	if err := json.Unmarshal(body, &updatedPrice); err != nil {
		return nil, FromError(err)
	}

	return &updatedPrice, nil
}

type PriceData struct {
	Data Price `json:"data"`
}

// PriceMeta contains extra data for an price
type PriceMeta struct {
	Timestamps Timestamps `json:"timestamps,omitempty"`
}

type PriceDataList struct {
}

type PriceList struct {
	Data []Price
}
