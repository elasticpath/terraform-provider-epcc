package epcc

import (
	"bytes"
	"context"
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
	Data string `json:"files"`
}

func (products) Get(ctx *context.Context, client *Client, productId string) (*ProductData, ApiErrors) {
	path := fmt.Sprintf("/pcm/products/%s", productId)

	body, apiError := client.DoRequest(ctx, "GET", path, nil)
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
func (products) GetAll(ctx *context.Context, client *Client) (*ProductList, ApiErrors) {
	path := fmt.Sprintf("/pcm/products")

	body, apiError := client.DoRequest(ctx, "GET", path, nil)
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
func (products) Create(ctx *context.Context, client *Client, product *Product) (*ProductData, ApiErrors) {
	productData := ProductData{
		Data: *product,
	}

	jsonPayload, err := json.Marshal(productData)
	if err != nil {
		return nil, FromError(err)
	}
	log.Printf("jsonPayload: " + string(jsonPayload))

	path := fmt.Sprintf("/pcm/products")

	body, apiError := client.DoRequest(ctx, "POST", path, bytes.NewBuffer(jsonPayload))
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
func (products) Delete(ctx *context.Context, client *Client, productID string) ApiErrors {
	path := fmt.Sprintf("/pcm/products/%s", productID)

	if _, err := client.DoRequest(ctx, "DELETE", path, nil); err != nil {
		return err
	}

	return nil
}

// Update updates a product.
func (products) Update(ctx *context.Context, client *Client, productID string, product *Product) (*ProductData, ApiErrors) {

	productData := ProductData{
		Data: *product,
	}

	jsonPayload, err := json.Marshal(productData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/pcm/products/%s", productID)

	body, apiError := client.DoRequest(ctx, "PUT", path, bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}

	var updatedProduct ProductData
	if err := json.Unmarshal(body, &updatedProduct); err != nil {
		return nil, FromError(err)
	}

	return &updatedProduct, nil
}

// Create Product Files creates file relationships for products
func (products) CreateProductFile(ctx *context.Context, client *Client, productId string, reference DataForTypeIdRelationshipList) ApiErrors {

	jsonPayload, err := json.Marshal(reference)
	if err != nil {
		return FromError(err)
	}
	log.Printf("jsonPayload: " + string(jsonPayload))

	path := fmt.Sprintf("/pcm/products/%s/relationships/files", productId)

	_, apiError := client.DoRequest(ctx, "POST", path, bytes.NewBuffer(jsonPayload))

	return apiError
}

// Update Product Files creates file relationships for products
func (products) UpdateProductFile(ctx *context.Context, client *Client, productId string, reference DataForTypeIdRelationshipList) ApiErrors {

	jsonPayload, err := json.Marshal(reference)
	if err != nil {
		return FromError(err)
	}
	log.Printf("jsonPayload: " + string(jsonPayload))

	path := fmt.Sprintf("/pcm/products/%s/relationships/files", productId)

	_, apiError := client.DoRequest(ctx, "PUT", path, bytes.NewBuffer(jsonPayload))

	return apiError
}

func (products) GetProductFiles(ctx *context.Context, client *Client, productId string) (*DataForTypeIdRelationshipList, ApiErrors) {
	path := fmt.Sprintf("/pcm/products/%s/relationships/files", productId)

	body, apiError := client.DoRequest(ctx, "GET", path, nil)
	if apiError != nil {
		return nil, apiError
	}

	var fileRelationships DataForTypeIdRelationshipList
	if err := json.Unmarshal(body, &fileRelationships); err != nil {
		return nil, FromError(err)
	}

	return &fileRelationships, nil
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
