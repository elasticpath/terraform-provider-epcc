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
