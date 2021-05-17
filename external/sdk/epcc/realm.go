package epcc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

var Realms realms

type realms struct{}

type Realm struct {
	Id                   string              `json:"id"`
	Type                 string              `json:"type"`
	Name                 string              `json:"name"`
	RedirectUris         []interface{}       `json:"redirect_uris"`
	DuplicateEmailPolicy string              `json:"duplicate_email_policy"`
	Relationships        *RealmRelationships `json:"relationships"`
}

type RealmRelationships struct {
	Origin *RealmRelationshipsOrigin `json:"origin"`
}
type RealmRelationshipsOrigin struct {
	Data *RealmRelationshipsOriginData `json:"data"`
}
type RealmRelationshipsOriginData struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

func (realms) Get(ctx *context.Context, client *Client, realmId string) (*RealmData, ApiErrors) {
	path := fmt.Sprintf("/v2/authentication-realms/%s", realmId)

	body, apiError := client.DoRequest(ctx, "GET", path, "", nil)
	if apiError != nil {
		return nil, apiError
	}

	var realms RealmData
	if err := json.Unmarshal(body, &realms); err != nil {
		return nil, FromError(err)
	}

	return &realms, nil
}

// GetAll fetches all realms
func (realms) GetAll(ctx *context.Context, client *Client) (*RealmList, ApiErrors) {
	path := fmt.Sprintf("/v2/authentication-realms")

	body, apiError := client.DoRequest(ctx, "GET", path, "", nil)
	if apiError != nil {
		return nil, apiError
	}

	var realms RealmList
	if err := json.Unmarshal(body, &realms); err != nil {
		return nil, FromError(err)
	}

	return &realms, nil
}

// Create creates a realm
func (realms) Create(ctx *context.Context, client *Client, realm *Realm) (*RealmData, ApiErrors) {
	realmData := RealmData{
		Data: *realm,
	}

	jsonPayload, err := json.Marshal(realmData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/v2/authentication-realms")

	body, apiError := client.DoRequest(ctx, "POST", path, "", bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}
	var newRealm RealmData
	if err := json.Unmarshal(body, &newRealm); err != nil {
		return nil, FromError(err)
	}

	return &newRealm, nil
}

// Delete deletes a realm.
func (realms) Delete(ctx *context.Context, client *Client, realmID string) ApiErrors {
	path := fmt.Sprintf("/v2/authentication-realms/%s", realmID)

	if _, err := client.DoRequest(ctx, "DELETE", path, "", nil); err != nil {
		return err
	}

	return nil
}

// Update updates a realm.
func (realms) Update(ctx *context.Context, client *Client, realmID string, realm *Realm) (*RealmData, ApiErrors) {

	realmData := RealmData{
		Data: *realm,
	}

	jsonPayload, err := json.Marshal(realmData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/v2/authentication-realms/%s", realmID)

	body, apiError := client.DoRequest(ctx, "PUT", path, "", bytes.NewBuffer(jsonPayload))

	if apiError != nil {
		return nil, apiError
	}
	var updatedRealm RealmData
	if err := json.Unmarshal(body, &updatedRealm); err != nil {
		return nil, FromError(err)
	}

	return &updatedRealm, nil
}

type RealmData struct {
	Data Realm `json:"data"`
}

// RealmMeta contains extra data for an realm
type RealmMeta struct {
	Timestamps Timestamps `json:"timestamps,omitempty"`
}

type RealmDataList struct {
}

type RealmList struct {
	Data []Realm
}
