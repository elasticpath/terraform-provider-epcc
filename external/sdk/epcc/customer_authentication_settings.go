package epcc

import (
	"context"
	"encoding/json"
	"fmt"
)

var CustomerAuthenticationSettings customerAuthenticationSettings

type customerAuthenticationSettings struct{}

type CustomerAuthenticationSettingsStruct struct {
	Relationships CustomerAuthenticationSettingsRelationshipsStruct `json:"relationships"`
	Meta          CustomerAuthenticationSettingsMetaStruct          `json:"meta"`
}

type CustomerAuthenticationSettingsRelationshipsStruct struct {
	AuthenticationRealm CustomerAuthenticationSettingsRelationshipsAuthRealmStruct `json:"authentication-realm"`
}

type CustomerAuthenticationSettingsMetaStruct struct {
	ClientId string `json:"client_id"`
}

type CustomerAuthenticationSettingsRelationshipsAuthRealmStruct struct {
	Data CustomerAuthenticationSettingsRelationshipsAuthRealmDataStruct `json:"data"`
}

type CustomerAuthenticationSettingsRelationshipsAuthRealmDataStruct struct {
	Id string `json:"id"`
}

func (customerAuthenticationSettings) Get(ctx *context.Context, client *Client) (*CustomerAuthenticationSettingsData, ApiErrors) {
	path := fmt.Sprintf("/v2/settings/customer-authentication")

	body, apiError := client.DoRequest(ctx, "GET", path, "", nil)
	if apiError != nil {
		return nil, apiError
	}

	var customerAuthenticationSettings CustomerAuthenticationSettingsData
	if err := json.Unmarshal(body, &customerAuthenticationSettings); err != nil {
		return nil, FromError(err)
	}

	return &customerAuthenticationSettings, nil
}

type CustomerAuthenticationSettingsData struct {
	Data CustomerAuthenticationSettingsStruct `json:"data"`
}

// CustomerAuthenticationSettingsMeta contains extra data for an customerAuthenticationSettings
type CustomerAuthenticationSettingsMeta struct {
	Timestamps Timestamps `json:"timestamps,omitempty"`
}

type CustomerAuthenticationSettingsDataList struct {
}

type CustomerAuthenticationSettingsList struct {
	Data []CustomerAuthenticationSettingsStruct
}
