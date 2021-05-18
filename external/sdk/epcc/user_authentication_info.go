package epcc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

var UserAuthenticationInfos userAuthenticationInfos

type userAuthenticationInfos struct{}

type UserAuthenticationInfo struct {
	Id      string `json:"id"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	RealmId string `json:"-"`
}

func (userAuthenticationInfos) Get(ctx *context.Context, client *Client, realmId, userAuthenticationInfoId string) (*UserAuthenticationInfoData, ApiErrors) {
	path := fmt.Sprintf("/v2/authentication-realms/%s/user-authentication-info/%s", realmId, userAuthenticationInfoId)

	body, apiError := client.DoRequest(ctx, "GET", path, "", nil)
	if apiError != nil {
		return nil, apiError
	}

	var userAuthenticationInfos UserAuthenticationInfoData
	if err := json.Unmarshal(body, &userAuthenticationInfos); err != nil {
		return nil, FromError(err)
	}

	return &userAuthenticationInfos, nil
}

// GetAll fetches all userAuthenticationInfos
func (userAuthenticationInfos) GetAll(ctx *context.Context, client *Client, realmId string) (*UserAuthenticationInfoList, ApiErrors) {
	path := fmt.Sprintf("/v2/authentication-realms/%s/user-authentication-info", realmId)

	body, apiError := client.DoRequest(ctx, "GET", path, "", nil)
	if apiError != nil {
		return nil, apiError
	}

	var userAuthenticationInfos UserAuthenticationInfoList
	if err := json.Unmarshal(body, &userAuthenticationInfos); err != nil {
		return nil, FromError(err)
	}

	return &userAuthenticationInfos, nil
}

// Create creates a userAuthenticationInfo
func (userAuthenticationInfos) Create(ctx *context.Context, client *Client, userAuthenticationInfo *UserAuthenticationInfo) (*UserAuthenticationInfoData, ApiErrors) {
	userAuthenticationInfoData := UserAuthenticationInfoData{
		Data: *userAuthenticationInfo,
	}

	jsonPayload, err := json.Marshal(userAuthenticationInfoData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/v2/authentication-realms/%s/user-authentication-info/", userAuthenticationInfo.RealmId)

	body, apiError := client.DoRequest(ctx, "POST", path, "", bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}
	var newUserAuthenticationInfo UserAuthenticationInfoData
	if err := json.Unmarshal(body, &newUserAuthenticationInfo); err != nil {
		return nil, FromError(err)
	}

	return &newUserAuthenticationInfo, nil
}

// Delete deletes a userAuthenticationInfo.
func (userAuthenticationInfos) Delete(ctx *context.Context, client *Client, userAuthenticationInfoID, realmId string) ApiErrors {
	path := fmt.Sprintf("/v2/authentication-realms/%s/user-authentication-info/%s", realmId, userAuthenticationInfoID)

	if _, err := client.DoRequest(ctx, "DELETE", path, "", nil); err != nil {
		return err
	}

	return nil
}

// Update updates a userAuthenticationInfo.
func (userAuthenticationInfos) Update(ctx *context.Context, client *Client, userAuthenticationInfo *UserAuthenticationInfo) (*UserAuthenticationInfoData, ApiErrors) {

	userAuthenticationInfoData := UserAuthenticationInfoData{
		Data: *userAuthenticationInfo,
	}

	jsonPayload, err := json.Marshal(userAuthenticationInfoData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/v2/authentication-realms/%s/user-authentication-info/%s", userAuthenticationInfo.RealmId, userAuthenticationInfo.Id)

	body, apiError := client.DoRequest(ctx, "PUT", path, "", bytes.NewBuffer(jsonPayload))

	if apiError != nil {
		return nil, apiError
	}
	var updatedUserAuthenticationInfo UserAuthenticationInfoData
	if err := json.Unmarshal(body, &updatedUserAuthenticationInfo); err != nil {
		return nil, FromError(err)
	}

	return &updatedUserAuthenticationInfo, nil
}

type UserAuthenticationInfoData struct {
	Data UserAuthenticationInfo `json:"data"`
}

// UserAuthenticationInfoMeta contains extra data for an userAuthenticationInfo
type UserAuthenticationInfoMeta struct {
	Timestamps Timestamps `json:"timestamps,omitempty"`
}

type UserAuthenticationInfoDataList struct {
}

type UserAuthenticationInfoList struct {
	Data []UserAuthenticationInfo
}
