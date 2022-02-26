package epcc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

var MerchantRealmMappings merchantRealmMappings

type merchantRealmMappings struct{}

type MerchantRealmMappingsStruct struct {
	ID      string  `json:"id"`
	Prefix  *string `json:"prefix"`
	RealmID string  `json:"realm_id,omitempty"`
	StoreID string  `json:"store_id,omitempty"`
	Type    string  `json:"type"`
}

func (merchantRealmMappings) Get(ctx *context.Context, client *Client) (*MerchantRealmMappingsData, ApiErrors) {
	path := fmt.Sprintf("/v2/merchant-realm-mappings")

	body, apiError := client.DoRequest(ctx, "GET", path, "", nil)
	if apiError != nil {
		return nil, apiError
	}

	var merchantRealmMappings MerchantRealmMappingsData
	if err := json.Unmarshal(body, &merchantRealmMappings); err != nil {
		return nil, FromError(err)
	}

	return &merchantRealmMappings, nil
}

func (merchantRealmMappings) Update(ctx *context.Context, client *Client, id, prefix *string) (*MerchantRealmMappingsData, ApiErrors) {

	merchantRealmMappingPrefixData := MerchantRealmMappingsData{
		Data: MerchantRealmMappingsStruct{
			Type:   "merchant-realm-mappings",
			Prefix: prefix,
		},
	}

	jsonPayload, err := json.Marshal(merchantRealmMappingPrefixData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/v2/merchant-realm-mappings/%s", *id)

	body, apiError := client.DoRequest(ctx, "PUT", path, "", bytes.NewBuffer(jsonPayload))

	if apiError != nil {
		return nil, apiError
	}
	var updatedMerchantRealmMappingPrefix MerchantRealmMappingsData
	if err := json.Unmarshal(body, &updatedMerchantRealmMappingPrefix); err != nil {
		return nil, FromError(err)
	}

	return &updatedMerchantRealmMappingPrefix, nil
}

type MerchantRealmMappingsData struct {
	Data MerchantRealmMappingsStruct `json:"data"`
}
