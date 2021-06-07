package epcc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

var MerchantRealmMappingPrefixes merchantRealmMappingPrefixes

type merchantRealmMappingPrefixes struct{}

type MerchantRealmMappingPrefix struct {
	Id     string `json:"id,omitempty"`
	Type   string `json:"type"`
	Prefix string `json:"prefix"`
}

func (merchantRealmMappingPrefixes) Get(ctx *context.Context, client *Client) (*MerchantRealmMappingPrefixData, ApiErrors) {
	path := fmt.Sprintf("/v2/merchant-realm-mappings")

	body, apiError := client.DoRequest(ctx, "GET", path, "", nil)
	if apiError != nil {
		return nil, apiError
	}

	var merchantRealmMappingPrefixes MerchantRealmMappingPrefixData
	if err := json.Unmarshal(body, &merchantRealmMappingPrefixes); err != nil {
		return nil, FromError(err)
	}

	return &merchantRealmMappingPrefixes, nil
}

// Update updates a merchantRealmMappingPrefix.
func (merchantRealmMappingPrefixes) Update(ctx *context.Context, client *Client, id, prefix *string) (*MerchantRealmMappingPrefixData, ApiErrors) {

	merchantRealmMappingPrefixData := MerchantRealmMappingPrefixData{
		Data: MerchantRealmMappingPrefix{
			Type:   "merchant-realm-mappings",
			Prefix: *prefix,
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
	var updatedMerchantRealmMappingPrefix MerchantRealmMappingPrefixData
	if err := json.Unmarshal(body, &updatedMerchantRealmMappingPrefix); err != nil {
		return nil, FromError(err)
	}

	return &updatedMerchantRealmMappingPrefix, nil
}

type MerchantRealmMappingPrefixData struct {
	Data MerchantRealmMappingPrefix `json:"data"`
}
