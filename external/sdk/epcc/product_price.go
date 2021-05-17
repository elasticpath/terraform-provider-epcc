package epcc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
)

var ProductPrices productPrices

type productPrices struct{}

type ProductPrice struct {
	Type       string                 `json:"type"`
	Id         string                 `json:"id,omitempty"`
	Attributes ProductPriceAttributes `json:"attributes"`
}

type ProductPriceInCurrency struct {
	Amount      int  `json:"amount"`
	IncludesTax bool `json:"includes_tax,omitempty"`
}

type ProductPriceAttributes struct {
	Sku        string                            `json:"sku"`
	Currencies map[string]ProductPriceInCurrency `json:"currencies"`
}

func (productPrices) Get(ctx *context.Context, client *Client, productPricebookId string, productPriceId string) (*ProductPriceData, ApiErrors) {
	path := fmt.Sprintf("/pcm/pricebooks/%s/prices/%s", productPricebookId, productPriceId)

	body, apiError := client.DoRequest(ctx, "GET", path, "", nil)
	if apiError != nil {
		return nil, apiError
	}

	var productPrice ProductPriceData
	if err := json.Unmarshal(body, &productPrice); err != nil {
		return nil, FromError(err)
	}

	return &productPrice, nil
}

// GetAll fetches all productPrices
func (productPrices) GetAll(ctx *context.Context, client *Client, productPricebookId string) (*ProductPriceList, ApiErrors) {
	path := fmt.Sprintf("/pcm/pricebooks/%s/prices", productPricebookId)

	body, apiError := client.DoRequest(ctx, "GET", path, "", nil)
	if apiError != nil {
		return nil, apiError
	}

	var productPrices ProductPriceList
	if err := json.Unmarshal(body, &productPrices); err != nil {
		return nil, FromError(err)
	}

	return &productPrices, nil
}

// Create creates a productPrice
func (productPrices) Create(ctx *context.Context, client *Client, productPricebookId string, productProductPrice *ProductPrice) (*ProductPriceData, ApiErrors) {
	productPriceData := ProductPriceData{
		Data: *productProductPrice,
	}

	jsonPayload, err := json.Marshal(productPriceData)
	if err != nil {
		return nil, FromError(err)
	}
	log.Printf("jsonPayload: " + string(jsonPayload))

	path := fmt.Sprintf("/pcm/pricebooks/%s/prices", productPricebookId)

	body, apiError := client.DoRequest(ctx, "POST", path, "", bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}

	var newProductPrice ProductPriceData
	if err := json.Unmarshal(body, &newProductPrice); err != nil {
		return nil, FromError(err)
	}

	return &newProductPrice, nil
}

// Delete deletes a productPrice.
func (productPrices) Delete(ctx *context.Context, client *Client, productPricebookId string, productPriceId string) ApiErrors {
	path := fmt.Sprintf("/pcm/pricebooks/%s/prices/%s", productPricebookId, productPriceId)

	if _, err := client.DoRequest(ctx, "DELETE", path, "", nil); err != nil {
		return err
	}

	return nil
}

// Update updates a productPrice.
func (productPrices) Update(ctx *context.Context, client *Client, productPricebookId string, productPriceId string, productProductPrice *ProductPrice) (*ProductPriceData, ApiErrors) {

	productPriceData := ProductPriceData{
		Data: *productProductPrice,
	}

	jsonPayload, err := json.Marshal(productPriceData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/pcm/pricebooks/%s/prices/%s", productPricebookId, productPriceId)

	body, apiError := client.DoRequest(ctx, "PUT", path, "", bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}

	var updatedProductPrice ProductPriceData
	if err := json.Unmarshal(body, &updatedProductPrice); err != nil {
		return nil, FromError(err)
	}

	return &updatedProductPrice, nil
}

type ProductPriceData struct {
	Data ProductPrice `json:"data"`
}

// ProductPriceMeta contains extra data for an productPrice
type ProductPriceMeta struct {
	Timestamps Timestamps `json:"timestamps,omitempty"`
}

type ProductPriceDataList struct {
}

type ProductPriceList struct {
	Data []ProductPrice
}
