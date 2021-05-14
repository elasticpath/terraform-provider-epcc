package epcc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

var Currencies currencies

type currencies struct{}

type Currency struct {
	Id                string `json:"id"`
	Type              string `json:"type"`
	Code              string `json:"code"`
	ExchangeRate      int    `json:"exchange_rate"`
	Format            string `json:"format"`
	DecimalPoint      string `json:"decimal_point"`
	ThousandSeparator string `json:"thousand_separator"`
	DecimalPlaces     int    `json:"decimal_places"`
	Default           bool   `json:"default"`
	Enabled           bool   `json:"enabled"`
}

func (currencies) Get(ctx *context.Context, client *Client, currencyId string) (*CurrencyData, ApiErrors) {
	path := fmt.Sprintf("/v2/currencies/%s", currencyId)

	body, apiError := client.DoRequest(ctx, "GET", path, nil)
	if apiError != nil {
		return nil, apiError
	}

	var currencies CurrencyData
	if err := json.Unmarshal(body, &currencies); err != nil {
		return nil, FromError(err)
	}

	return &currencies, nil
}

// GetAll fetches all currencies
func (currencies) GetAll(ctx *context.Context, client *Client) (*CurrencyList, ApiErrors) {
	path := fmt.Sprintf("/v2/currencies")

	body, apiError := client.DoRequest(ctx, "GET", path, nil)
	if apiError != nil {
		return nil, apiError
	}

	var currencies CurrencyList
	if err := json.Unmarshal(body, &currencies); err != nil {
		return nil, FromError(err)
	}

	return &currencies, nil
}

// Create creates a currency
func (currencies) Create(ctx *context.Context, client *Client, currency *Currency) (*CurrencyData, ApiErrors) {
	currencyData := CurrencyData{
		Data: *currency,
	}

	jsonPayload, err := json.Marshal(currencyData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/v2/currencies")

	body, apiError := client.DoRequest(ctx, "POST", path, bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}
	var newCurrency CurrencyData
	if err := json.Unmarshal(body, &newCurrency); err != nil {
		return nil, FromError(err)
	}

	return &newCurrency, nil
}

// Delete deletes a currency.
func (currencies) Delete(ctx *context.Context, client *Client, currencyID string) ApiErrors {
	path := fmt.Sprintf("/v2/currencies/%s", currencyID)

	if _, err := client.DoRequest(ctx, "DELETE", path, nil); err != nil {
		return err
	}

	return nil
}

// Update updates a currency.
func (currencies) Update(ctx *context.Context, client *Client, currencyID string, currency *Currency) (*CurrencyData, ApiErrors) {

	currencyData := CurrencyData{
		Data: *currency,
	}

	jsonPayload, err := json.Marshal(currencyData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/v2/currencies/%s", currencyID)

	body, apiError := client.DoRequest(ctx, "PUT", path, bytes.NewBuffer(jsonPayload))

	if apiError != nil {
		return nil, apiError
	}
	var updatedCurrency CurrencyData
	if err := json.Unmarshal(body, &updatedCurrency); err != nil {
		return nil, FromError(err)
	}

	return &updatedCurrency, nil
}

type CurrencyData struct {
	Data Currency `json:"data"`
}

// CurrencyMeta contains extra data for an currency
type CurrencyMeta struct {
	Timestamps Timestamps `json:"timestamps,omitempty"`
}

type CurrencyDataList struct {
}

type CurrencyList struct {
	Data []Currency
}
