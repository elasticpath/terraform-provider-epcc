package epcc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

var Accounts accounts

type accounts struct{}

type Account struct {
	Id             string `json:"id,omitempty"`
	Type           string `json:"type"`
	Name           string `json:"name,omitempty"`
	LegalName      string `json:"legal_name"`
	RegistrationId string `json:"registration_id,omitempty"`
	ParentId       string `json:"parent_id,omitempty"`
}

func (accounts) Get(ctx *context.Context, client *Client, accountId string) (*AccountData, ApiErrors) {
	path := fmt.Sprintf("/v2/accounts/%s", accountId)
	if accountId == "" {
		return nil, FromError(fmt.Errorf("account id should not be empty [%s]", accountId))
	}

	body, apiError := client.DoRequest(ctx, "GET", path, "", nil)
	if apiError != nil {
		return nil, apiError
	}

	// TODO Better Manage Parent ID
	var accounts AccountData
	if err := json.Unmarshal(body, &accounts); err != nil {
		return nil, FromError(err)
	}

	return &accounts, nil
}

// GetAll fetches all accounts
func (accounts) GetAll(ctx *context.Context, client *Client) (*AccountList, ApiErrors) {
	path := fmt.Sprintf("/v2/accounts")

	body, apiError := client.DoRequest(ctx, "GET", path, "", nil)
	if apiError != nil {
		return nil, apiError
	}

	var accounts AccountList
	if err := json.Unmarshal(body, &accounts); err != nil {
		return nil, FromError(err)
	}

	return &accounts, nil
}

// Create creates a account
func (accounts) Create(ctx *context.Context, client *Client, account *Account) (*AccountData, ApiErrors) {
	accountData := AccountData{
		Data: *account,
	}

	jsonPayload, err := json.Marshal(accountData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/v2/accounts")

	body, apiError := client.DoRequest(ctx, "POST", path, "", bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}

	var newAccount AccountData
	if err := json.Unmarshal(body, &newAccount); err != nil {
		return nil, FromError(err)
	}

	return &newAccount, nil
}

// Delete deletes a account.
func (accounts) Delete(ctx *context.Context, client *Client, accountID string) ApiErrors {
	path := fmt.Sprintf("/v2/accounts/%s", accountID)

	if _, err := client.DoRequest(ctx, "DELETE", path, "", nil); err != nil {
		return err
	}

	return nil
}

// Update updates a account.
func (accounts) Update(ctx *context.Context, client *Client, accountID string, account *Account) (*AccountData, ApiErrors) {

	accountData := AccountData{
		Data: *account,
	}

	jsonPayload, err := json.Marshal(accountData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/v2/accounts/%s", accountID)

	body, apiError := client.DoRequest(ctx, "PUT", path, "", bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}

	var updatedAccount AccountData
	if err := json.Unmarshal(body, &updatedAccount); err != nil {
		return nil, FromError(err)
	}

	return &updatedAccount, nil
}

type AccountData struct {
	Data Account `json:"data"`
}

// AccountMeta contains extra data for an account
type AccountMeta struct {
	Timestamps Timestamps `json:"timestamps,omitempty"`
}

type AccountDataList struct {
}

type AccountList struct {
	Data []Account
}
