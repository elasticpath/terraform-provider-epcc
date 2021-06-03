package epcc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

var CatalogRules catalogRules

type catalogRules struct{}

type CatalogRule struct {
	Id         string                 `json:"id,omitempty"`
	Type       string                 `json:"type"`
	Attributes CatalogRulesAttributes `json:"attributes"`
}

type CatalogRulesAttributes struct {
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	Catalog     string   `json:"catalog_id,omitempty"`
	Customers   []string `json:"customer_ids,omitempty"`
	// TODO in future add channels and tags. They are in beta version from my understanding
}

func (catalogRules) Get(ctx *context.Context, client *Client, catalogRuleId string) (*CatalogRuleData, ApiErrors) {
	path := fmt.Sprintf("/pcm/catalogs/rules/%s", catalogRuleId)

	body, apiError := client.DoRequest(ctx, "GET", path, "", nil)
	if apiError != nil {
		return nil, apiError
	}

	var catalogRule CatalogRuleData
	if err := json.Unmarshal(body, &catalogRule); err != nil {
		return nil, FromError(err)
	}

	return &catalogRule, nil
}

func (catalogRules) GetAll(ctx *context.Context, client *Client) (*CatalogRuleList, ApiErrors) {
	path := fmt.Sprintf("/pcm/catalogs/rules")

	body, apiError := client.DoRequest(ctx, "GET", path, "", nil) // TODO query params for pagination
	if apiError != nil {
		return nil, apiError
	}

	var catalogRules CatalogRuleList
	if err := json.Unmarshal(body, &catalogRules); err != nil {
		return nil, FromError(err)
	}

	return &catalogRules, nil
}

func (catalogRules) Create(ctx *context.Context, client *Client, catalogRule *CatalogRule) (*CatalogRuleData, ApiErrors) {
	catalogRuleData := CatalogRuleData{
		Data: *catalogRule,
	}

	jsonPayload, err := json.Marshal(catalogRuleData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/pcm/catalogs/rules")

	body, apiError := client.DoRequest(ctx, "POST", path, "", bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}

	var newCatalogRule CatalogRuleData
	if err := json.Unmarshal(body, &newCatalogRule); err != nil {
		return nil, FromError(err)
	}

	return &newCatalogRule, nil
}

func (catalogRules) Delete(ctx *context.Context, client *Client, catalogRuleId string) ApiErrors {
	path := fmt.Sprintf("/pcm/catalogs/rules/%s", catalogRuleId)

	if _, err := client.DoRequest(ctx, "DELETE", path, "", nil); err != nil {
		return err
	}

	return nil
}

func (catalogRules) Update(ctx *context.Context, client *Client, catalogRuleId string, catalogRule *CatalogRule) (*CatalogRuleData, ApiErrors) {

	catalogRuleData := CatalogRuleData{
		Data: *catalogRule,
	}

	jsonPayload, err := json.Marshal(catalogRuleData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/pcm/catalogs/rules/%s", catalogRuleId)

	body, apiError := client.DoRequest(ctx, "PUT", path, "", bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}

	var updatedCatalogRule CatalogRuleData
	if err := json.Unmarshal(body, &updatedCatalogRule); err != nil {
		return nil, FromError(err)
	}

	return &updatedCatalogRule, nil
}

type CatalogRuleData struct {
	Data CatalogRule `json:"data"`
}

type CatalogRuleMeta struct {
	Timestamps Timestamps `json:"timestamps,omitempty"`
}

type CatalogRuleDataList struct {
}

type CatalogRuleList struct {
	Data []CatalogRule
}
