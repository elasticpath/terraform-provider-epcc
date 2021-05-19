package epcc

import (
	"context"
	"encoding/json"
	"fmt"
)

var AccountAuthenticationSettings accountAuthenticationSettings

type accountAuthenticationSettings struct{}

type AccountAuthenticationSettingsStruct struct {
	Relationships AccountAuthenticationSettingsRelationshipsStruct `json:"relationships"`
	Meta          AccountAuthenticationSettingsMetaStruct          `json:"meta"`
}

type AccountAuthenticationSettingsRelationshipsStruct struct {
	AuthenticationRealm AccountAuthenticationSettingsRelationshipsAuthRealmStruct `json:"authentication_realm"`
}

type AccountAuthenticationSettingsMetaStruct struct {
	ClientId string `json:"client_id"`
}

type AccountAuthenticationSettingsRelationshipsAuthRealmStruct struct {
	Data AccountAuthenticationSettingsRelationshipsAuthRealmDataStruct `json:"data"`
}

type AccountAuthenticationSettingsRelationshipsAuthRealmDataStruct struct {
	Id string `json:"id"`
}

func (accountAuthenticationSettings) Get(ctx *context.Context, client *Client) (*AccountAuthenticationSettingsData, ApiErrors) {
	path := fmt.Sprintf("/v2/settings/account-authentication")

	body, apiError := client.DoRequest(ctx, "GET", path, "", nil)
	if apiError != nil {
		return nil, apiError
	}

	var accountAuthenticationSettings AccountAuthenticationSettingsData
	if err := json.Unmarshal(body, &accountAuthenticationSettings); err != nil {
		return nil, FromError(err)
	}

	return &accountAuthenticationSettings, nil
}

type AccountAuthenticationSettingsData struct {
	Data AccountAuthenticationSettingsStruct `json:"data"`
}

// AccountAuthenticationSettingsMeta contains extra data for an accountAuthenticationSettings
type AccountAuthenticationSettingsMeta struct {
	Timestamps Timestamps `json:"timestamps,omitempty"`
}

type AccountAuthenticationSettingsDataList struct {
}

type AccountAuthenticationSettingsList struct {
	Data []AccountAuthenticationSettingsStruct
}
