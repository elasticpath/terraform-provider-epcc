package epcc

import (
	"bytes"
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

func (accounts) Get(client *Client, accountId string) (*AccountData, ApiErrors) {
	path := fmt.Sprintf("/v2/accounts/%s", accountId)

	body, apiError := client.DoRequest("GET", path, nil)
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
func (accounts) GetAll(client *Client) (*AccountList, ApiErrors) {
	path := fmt.Sprintf("/v2/accounts")

	body, apiError := client.DoRequest("GET", path, nil)
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
func (accounts) Create(client *Client, account *Account) (*AccountData, ApiErrors) {
	accountData := AccountData{
		Data: *account,
	}

	jsonPayload, err := json.Marshal(accountData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/v2/accounts")

	body, apiError := client.DoRequest("POST", path, bytes.NewBuffer(jsonPayload))
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
func (accounts) Delete(client *Client, accountID string) ApiErrors {
	path := fmt.Sprintf("/v2/accounts/%s", accountID)

	if _, err := client.DoRequest("DELETE", path, nil); err != nil {
		return err
	}

	return nil
}

// Update updates a account.
func (accounts) Update(client *Client, accountID string, account *Account) (*AccountData, ApiErrors) {

	accountData := AccountData{
		Data: *account,
	}

	jsonPayload, err := json.Marshal(accountData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/v2/accounts/%s", accountID)

	body, apiError := client.DoRequest("PUT", path, bytes.NewBuffer(jsonPayload))
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
