package epcc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

var Products products

type products struct{}

type Product struct {
	Type       string            `json:"type"`
	Id         string            `json:"id,omitempty"`
	Attributes ProductAttributes `json:"attributes"`
}

type ProductAttributes struct {
	Name          string `json:"name"`
	CommodityType string `json:"commodity_type"` // Valid values: physical or digital
	Sku           string `json:"sku"`
	Slug          string `json:"slug,omitempty"`
	Description   string `json:"description,omitempty"`
	Mpn           string `json:"mpn,omitempty"`
	Status        string `json:"status,omitempty"`
	UpcEan        string `json:"upc_ean,omitempty"`
}

type ProductRelationships struct {
	Files     ProductRelationshipsChild `json:"files"`
	Templates ProductRelationshipsChild `json:"templates"`
}

type ProductRelationshipsChild struct {
	Data  string                         `json:"files"`
	Links ProductRelationshipsChildLinks `json:"links,omitempty"`
}

type ProductRelationshipsChildLinks struct {
	Self string `json:"self,omitempty"`
}

func (products) Get(client *Client, productId string) (*ProductData, ApiErrors) {
	path := fmt.Sprintf("/pcm/products/%s", productId)

	body, apiError := client.DoRequest("GET", path, nil)
	if apiError != nil {
		return nil, apiError
	}

	var product ProductData
	if err := json.Unmarshal(body, &product); err != nil {
		return nil, FromError(err)
	}

	return &product, nil
}

// GetAll fetches all products
func (products) GetAll(client *Client) (*ProductList, ApiErrors) {
	path := fmt.Sprintf("/pcm/products")

	body, apiError := client.DoRequest("GET", path, nil)
	if apiError != nil {
		return nil, apiError
	}

	var products ProductList
	if err := json.Unmarshal(body, &products); err != nil {
		return nil, FromError(err)
	}

	return &products, nil
}

// Create creates a product
func (products) Create(client *Client, product *Product) (*ProductData, ApiErrors) {
	productData := ProductData{
		Data: *product,
	}

	jsonPayload, err := json.Marshal(productData)
	if err != nil {
		return nil, FromError(err)
	}
	log.Printf("jsonPayload: " + string(jsonPayload))

	path := fmt.Sprintf("/pcm/products")

	body, apiError := client.DoRequest("POST", path, bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}

	var newProduct ProductData
	if err := json.Unmarshal(body, &newProduct); err != nil {
		return nil, FromError(err)
	}

	return &newProduct, nil
}

// Delete deletes a product.
func (products) Delete(client *Client, productID string) ApiErrors {
	path := fmt.Sprintf("/pcm/products/%s", productID)

	if _, err := client.DoRequest("DELETE", path, nil); err != nil {
		return err
	}

	return nil
}

// Update updates a product.
func (products) Update(client *Client, productID string, product *Product) (*ProductData, ApiErrors) {

	productData := ProductData{
		Data: *product,
	}

	jsonPayload, err := json.Marshal(productData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/pcm/products/%s", productID)

	body, apiError := client.DoRequest("PUT", path, bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}

	var updatedProduct ProductData
	if err := json.Unmarshal(body, &updatedProduct); err != nil {
		return nil, FromError(err)
	}

	return &updatedProduct, nil
}

type ProductData struct {
	Data Product `json:"data"`
}

// ProductMeta contains extra data for an product
type ProductMeta struct {
	Timestamps Timestamps `json:"timestamps,omitempty"`
}

type ProductDataList struct {
}

type ProductList struct {
	Data []Product
}