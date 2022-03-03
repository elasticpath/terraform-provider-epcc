package epcc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

var OidcProfiles oidcProfiles

type oidcProfiles struct{}

type OidcProfile struct {
	Id           string `json:"id"`
	Type         string `json:"type"`
	Name         string `json:"name"`
	DiscoveryUrl string `json:"discovery_url"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RealmId      string `json:"-"`
}

type OidcProfileLinks struct {
	CallbackEndpoint      string `json:"callback-endpoint,omitempty"`
	AuthorizationEndpoint string `json:"authorization-endpoint,omitempty"`
	ClientDiscoveryUrl    string `json:"client-discovery-url,omitempty"`
}

func (oidcProfiles) Get(ctx *context.Context, client *Client, realmId, profileId string) (*ProfileData, ApiErrors) {
	path := fmt.Sprintf("/v2/authentication-realms/%s/oidc-profiles/%s", realmId, profileId)

	body, apiError := client.DoRequest(ctx, "GET", path, "", nil)
	if apiError != nil {
		return nil, apiError
	}

	var profiles ProfileData
	if err := json.Unmarshal(body, &profiles); err != nil {
		return nil, FromError(err)
	}

	return &profiles, nil
}

// GetAll fetches all oidcProfiles
func (oidcProfiles) GetAll(ctx *context.Context, client *Client, realmId string) (*ProfileList, ApiErrors) {
	path := fmt.Sprintf("/v2/authentication-realms/%s/oidc-profiles", realmId)

	body, apiError := client.DoRequest(ctx, "GET", path, "", nil)
	if apiError != nil {
		return nil, apiError
	}

	var profiles ProfileList
	if err := json.Unmarshal(body, &profiles); err != nil {
		return nil, FromError(err)
	}

	return &profiles, nil
}

// Create creates a profile
func (oidcProfiles) Create(ctx *context.Context, client *Client, profile *OidcProfile) (*ProfileData, ApiErrors) {
	profileData := ProfileData{
		Data: *profile,
	}

	jsonPayload, err := json.Marshal(profileData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/v2/authentication-realms/%s/oidc-profiles/", profile.RealmId)

	body, apiError := client.DoRequest(ctx, "POST", path, "", bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}
	var newProfile ProfileData
	if err := json.Unmarshal(body, &newProfile); err != nil {
		return nil, FromError(err)
	}

	return &newProfile, nil
}

// Delete deletes a profile.
func (oidcProfiles) Delete(ctx *context.Context, client *Client, profileID, realmId string) ApiErrors {
	path := fmt.Sprintf("/v2/authentication-realms/%s/oidc-profiles/%s", realmId, profileID)

	if _, err := client.DoRequest(ctx, "DELETE", path, "", nil); err != nil {
		return err
	}

	return nil
}

// Update updates a profile.
func (oidcProfiles) Update(ctx *context.Context, client *Client, profile *OidcProfile) (*ProfileData, ApiErrors) {

	profileData := ProfileData{
		Data: *profile,
	}

	jsonPayload, err := json.Marshal(profileData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/v2/authentication-realms/%s/oidc-profiles/%s", profile.RealmId, profile.Id)

	body, apiError := client.DoRequest(ctx, "PUT", path, "", bytes.NewBuffer(jsonPayload))

	if apiError != nil {
		return nil, apiError
	}
	var updatedProfile ProfileData
	if err := json.Unmarshal(body, &updatedProfile); err != nil {
		return nil, FromError(err)
	}

	return &updatedProfile, nil
}

type ProfileData struct {
	Data  OidcProfile       `json:"data"`
	Links *OidcProfileLinks `json:"links,omitempty"`
}

// ProfileMeta contains extra data for an profile
type ProfileMeta struct {
	Timestamps Timestamps `json:"timestamps,omitempty"`
}

type ProfileDataList struct {
}

type ProfileList struct {
	Data []OidcProfile
}
