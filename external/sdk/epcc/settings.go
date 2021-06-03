package epcc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

var SettingsVar settings

type settings struct{}

type Settings struct {
	Id                  string        `json:"id"`
	Type                string        `json:"type"`
	PageLength          int           `json:"page_length"`
	ListChildProducts   bool          `json:"list_child_products"`
	AdditionalLanguages []interface{} `json:"additional_languages"`
	CalculationMethod   string        `json:"calculation_method"`
}

func (settings) Get(ctx *context.Context, client *Client) (*SettingsData, ApiErrors) {
	path := fmt.Sprintf("/v2/settings")

	body, apiError := client.DoRequest(ctx, "GET", path, "", nil)
	if apiError != nil {
		return nil, apiError
	}

	var settings SettingsData
	if err := json.Unmarshal(body, &settings); err != nil {
		return nil, FromError(err)
	}

	return &settings, nil
}

//// GetAll fetches all profiles
//func (profiles) GetAll(ctx *context.Context, client *Client, realmId string) (*ProfileList, ApiErrors) {
//	path := fmt.Sprintf("/v2/authentication-realms/%s/oidc-profiles", realmId)
//
//	body, apiError := client.DoRequest(ctx, "GET", path, "", nil)
//	if apiError != nil {
//		return nil, apiError
//	}
//
//	var profiles ProfileList
//	if err := json.Unmarshal(body, &profiles); err != nil {
//		return nil, FromError(err)
//	}
//
//	return &profiles, nil
//}
////
////// Create creates a profile
////func (profiles) Create(ctx *context.Context, client *Client, profile *Profile) (*ProfileData, ApiErrors) {
////	profileData := ProfileData{
////		Data: *profile,
////	}
////
////	jsonPayload, err := json.Marshal(profileData)
////	if err != nil {
////		return nil, FromError(err)
////	}
////
////	path := fmt.Sprintf("/v2/authentication-realms/%s/oidc-profiles/", profile.RealmId)
////
////	body, apiError := client.DoRequest(ctx, "POST", path, "", bytes.NewBuffer(jsonPayload))
////	if apiError != nil {
////		return nil, apiError
////	}
////	var newProfile ProfileData
////	if err := json.Unmarshal(body, &newProfile); err != nil {
////		return nil, FromError(err)
////	}
////
////	return &newProfile, nil
////}
//
//// Delete deletes a profile.
//func (profiles) Delete(ctx *context.Context, client *Client, profileID, realmId string) ApiErrors {
//	path := fmt.Sprintf("/v2/authentication-realms/%s/oidc-profiles/%s", realmId, profileID)
//
//	if _, err := client.DoRequest(ctx, "DELETE", path, "", nil); err != nil {
//		return err
//	}
//
//	return nil
//}

// Update updates a profile.
func (settings) Update(ctx *context.Context, client *Client, settings Settings) (*SettingsData, ApiErrors) {

	jsonPayload, err := json.Marshal(SettingsData{Data: settings})
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/v2/settings")

	body, apiError := client.DoRequest(ctx, "PUT", path, "", bytes.NewBuffer(jsonPayload))

	if apiError != nil {
		return nil, apiError
	}
	var updatedSettings SettingsData
	if err := json.Unmarshal(body, &updatedSettings); err != nil {
		return nil, FromError(err)
	}

	return &updatedSettings, nil
}

type SettingsData struct {
	Data Settings `json:"data"`
}

//
//// ProfileMeta contains extra data for an profile
//type ProfileMeta struct {
//	Timestamps Timestamps `json:"timestamps,omitempty"`
//}
//
//type ProfileDataList struct {
//}
//
//type ProfileList struct {
//	Data []Profile
//}
