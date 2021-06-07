package epcc

import (
	"context"
	"encoding/json"
	"fmt"
)

var MerchantRealmMappings merchantRealmMappings

type merchantRealmMappings struct{}

type MerchantRealmMappingsStruct struct {
	ID      string `json:"id"`
	Prefix  string `json:"prefix"`
	RealmID string `json:"realm_id"`
	StoreID string `json:"store_id"`
	Type    string `json:"type"`
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

type MerchantRealmMappingsData struct {
	Data MerchantRealmMappingsStruct `json:"data"`
}
