package epcc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

var Catalogs catalogs

type catalogs struct{}

type Catalog struct {
	Id         string            `json:"id,omitempty"`
	Type       string            `json:"type"`
	Attributes CatalogAttributes `json:"attributes"`
}

type CatalogAttributes struct {
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	Hierarchies []string `json:"hierarchy_ids,omitempty"`
	PriceBook   string   `json:"pricebook_id,omitempty"`
}

func (catalogs) Get(ctx *context.Context, client *Client, catalogId string) (*CatalogData, ApiErrors) {
	path := fmt.Sprintf("/pcm/catalogs/%s", catalogId)

	body, apiError := client.DoRequest(ctx, "GET", path, "", nil)
	if apiError != nil {
		return nil, apiError
	}

	var catalogs CatalogData
	if err := json.Unmarshal(body, &catalogs); err != nil {
		return nil, FromError(err)
	}

	return &catalogs, nil
}

// GetAll fetches all catalogs
func (catalogs) GetAll(ctx *context.Context, client *Client) (*CatalogList, ApiErrors) {
	path := fmt.Sprintf("/pcm/catalogs")

	body, apiError := client.DoRequest(ctx, "GET", path, "", nil)
	if apiError != nil {
		return nil, apiError
	}

	var catalogs CatalogList
	if err := json.Unmarshal(body, &catalogs); err != nil {
		return nil, FromError(err)
	}

	return &catalogs, nil
}

// Create creates a catalog
func (catalogs) Create(ctx *context.Context, client *Client, catalog *Catalog) (*CatalogData, ApiErrors) {
	catalogData := CatalogData{
		Data: *catalog,
	}

	jsonPayload, err := json.Marshal(catalogData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/pcm/catalogs")

	body, apiError := client.DoRequest(ctx, "POST", path, "", bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}

	var newCatalog CatalogData
	if err := json.Unmarshal(body, &newCatalog); err != nil {
		return nil, FromError(err)
	}

	return &newCatalog, nil
}

// Delete deletes a catalog.
func (catalogs) Delete(ctx *context.Context, client *Client, catalogID string) ApiErrors {
	path := fmt.Sprintf("/pcm/catalogs/%s", catalogID)

	if _, err := client.DoRequest(ctx, "DELETE", path, "", nil); err != nil {
		return err
	}

	return nil
}

// Update updates a catalog.
func (catalogs) Update(ctx *context.Context, client *Client, catalogID string, catalog *Catalog) (*CatalogData, ApiErrors) {

	catalogData := CatalogData{
		Data: *catalog,
	}

	jsonPayload, err := json.Marshal(catalogData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/pcm/catalogs/%s", catalogID)

	body, apiError := client.DoRequest(ctx, "PUT", path, "", bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}

	var updatedCatalog CatalogData
	if err := json.Unmarshal(body, &updatedCatalog); err != nil {
		return nil, FromError(err)
	}

	return &updatedCatalog, nil
}

type CatalogData struct {
	Data Catalog `json:"data"`
}

// CatalogMeta contains extra data for an catalog
type CatalogMeta struct {
	Timestamps Timestamps `json:"timestamps,omitempty"`
}

type CatalogDataList struct {
}

type CatalogList struct {
	Data []Catalog
}
