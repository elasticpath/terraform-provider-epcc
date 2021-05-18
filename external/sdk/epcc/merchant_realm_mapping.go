package epcc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

var MerchantRealmMappings merchantRealmMappings

type merchantRealmMappings struct{}

type MerchantRealmMapping struct {
	Id     string `json:"id,omitempty"`
	Prefix string `json:"prefix"`
	RealmId  string `json:"realm_id,omitempty"`
	StoreId  string `json:"store_id,omitempty"`
	Type   string `json:"type,omitempty"`
}

func (merchantRealmMappings) Get(ctx *context.Context, client *Client, merchantRealmMappingId string) (*MerchantRealmMappingData, ApiErrors) {
	path := fmt.Sprintf("/v2/merchant-realm-mappings/%s", merchantRealmMappingId)

	body, apiError := client.DoRequest(ctx, "GET", path, "", nil)
	if apiError != nil {
		return nil, apiError
	}

	// TODO Better Manage Parent ID
	var merchantRealmMappings MerchantRealmMappingData
	if err := json.Unmarshal(body, &merchantRealmMappings); err != nil {
		return nil, FromError(err)
	}

	return &merchantRealmMappings, nil
}

// GetAll fetches all merchantRealmMappings
func (merchantRealmMappings) GetAll(ctx *context.Context, client *Client) (*MerchantRealmMappingList, ApiErrors) {
	path := fmt.Sprintf("/v2/merchant-realm-mappings")

	body, apiError := client.DoRequest(ctx, "GET", path, "", nil)
	if apiError != nil {
		return nil, apiError
	}

	var merchantRealmMappings MerchantRealmMappingList
	if err := json.Unmarshal(body, &merchantRealmMappings); err != nil {
		return nil, FromError(err)
	}

	return &merchantRealmMappings, nil
}

// Create creates a merchantRealmMapping
func (merchantRealmMappings) Create(ctx *context.Context, client *Client, merchantRealmMapping *MerchantRealmMapping) (*MerchantRealmMappingData, ApiErrors) {
	merchantRealmMappingData := MerchantRealmMappingData{
		Data: *merchantRealmMapping,
	}

	jsonPayload, err := json.Marshal(merchantRealmMappingData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/v2/merchant-realm-mappings")

	body, apiError := client.DoRequest(ctx, "POST", path, "", bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}

	var newMerchantRealmMapping MerchantRealmMappingData
	if err := json.Unmarshal(body, &newMerchantRealmMapping); err != nil {
		return nil, FromError(err)
	}

	return &newMerchantRealmMapping, nil
}

// Delete deletes a merchantRealmMapping.
func (merchantRealmMappings) Delete(ctx *context.Context, client *Client, merchantRealmMappingID string) ApiErrors {
	path := fmt.Sprintf("/v2/merchant-realm-mappings/%s", merchantRealmMappingID)

	if _, err := client.DoRequest(ctx, "DELETE", path, "", nil); err != nil {
		return err
	}

	return nil
}

// Update updates a merchantRealmMapping.
func (merchantRealmMappings) Update(ctx *context.Context, client *Client, merchantRealmMappingID string, merchantRealmMapping *MerchantRealmMapping) (*MerchantRealmMappingData, ApiErrors) {

	merchantRealmMappingData := MerchantRealmMappingData{
		Data: *merchantRealmMapping,
	}

	jsonPayload, err := json.Marshal(merchantRealmMappingData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/v2/merchant-realm-mappings/%s", merchantRealmMappingID)

	body, apiError := client.DoRequest(ctx, "PUT", path, "", bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}

	var updatedMerchantRealmMapping MerchantRealmMappingData
	if err := json.Unmarshal(body, &updatedMerchantRealmMapping); err != nil {
		return nil, FromError(err)
	}

	return &updatedMerchantRealmMapping, nil
}

type MerchantRealmMappingData struct {
	Data MerchantRealmMapping `json:"data"`
}

// MerchantRealmMappingMeta contains extra data for a merchantRealmMapping
type MerchantRealmMappingMeta struct {
	Timestamps Timestamps `json:"timestamps,omitempty"`
}

type MerchantRealmMappingDataList struct {
}

type MerchantRealmMappingList struct {
	Data []MerchantRealmMapping
}
