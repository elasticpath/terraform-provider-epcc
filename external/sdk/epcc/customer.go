package epcc

import (
	"bytes"
	"encoding/json"
	"fmt"
)

var Customers customers

type customers struct{}

type Customer struct {
	Id       string       `json:"id,omitempty"`
	Type     string       `json:"type"`
	Name     string       `json:"name,omitempty"`
	Email    string       `json:"email,omitempty"`
	Password *interface{} `json:"password,omitempty"`
}

func (customers) Get(client *Client, customerId string) (*CustomerData, ApiErrors) {
	path := fmt.Sprintf("/v2/customers/%s", customerId)

	body, apiError := client.DoRequest("GET", path, nil)
	if apiError != nil {
		return nil, apiError
	}

	// TODO Better Manage Parent ID
	var customers CustomerData
	if err := json.Unmarshal(body, &customers); err != nil {
		return nil, FromError(err)
	}

	return &customers, nil
}

// GetAll fetches all customers
func (customers) GetAll(client *Client) (*CustomerList, ApiErrors) {
	path := fmt.Sprintf("/v2/customers")

	body, apiError := client.DoRequest("GET", path, nil)
	if apiError != nil {
		return nil, apiError
	}

	var customers CustomerList
	if err := json.Unmarshal(body, &customers); err != nil {
		return nil, FromError(err)
	}

	return &customers, nil
}

// Create creates a customer
func (customers) Create(client *Client, customer *Customer) (*CustomerData, ApiErrors) {
	customerData := CustomerData{
		Data: *customer,
	}

	jsonPayload, err := json.Marshal(customerData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/v2/customers")

	body, apiError := client.DoRequest("POST", path, bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}

	var newCustomer CustomerData
	if err := json.Unmarshal(body, &newCustomer); err != nil {
		return nil, FromError(err)
	}

	return &newCustomer, nil
}

// Delete deletes a customer.
func (customers) Delete(client *Client, customerID string) ApiErrors {
	path := fmt.Sprintf("/v2/customers/%s", customerID)

	if _, err := client.DoRequest("DELETE", path, nil); err != nil {
		return err
	}

	return nil
}

// Update updates a customer.
func (customers) Update(client *Client, customerID string, customer *Customer) (*CustomerData, ApiErrors) {

	customerData := CustomerData{
		Data: *customer,
	}

	jsonPayload, err := json.Marshal(customerData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/v2/customers/%s", customerID)

	body, apiError := client.DoRequest("PUT", path, bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}

	var updatedCustomer CustomerData
	if err := json.Unmarshal(body, &updatedCustomer); err != nil {
		return nil, FromError(err)
	}

	return &updatedCustomer, nil
}

type CustomerData struct {
	Data Customer `json:"data"`
}

// CustomerMeta contains extra data for an customer
type CustomerMeta struct {
	Timestamps Timestamps `json:"timestamps,omitempty"`
}

type CustomerDataList struct {
}

type CustomerList struct {
	Data []Customer
}
