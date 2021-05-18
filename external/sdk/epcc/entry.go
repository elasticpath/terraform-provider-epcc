package epcc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

var Entries entries

type entries struct{}

type Entry struct {
	Id   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
}

func (entries) Get(ctx *context.Context, client *Client, flowSlug string, entryID string) (*EntryData, ApiErrors) {
	path := fmt.Sprintf("/v2/flows/%s/entries/%s", flowSlug, entryID)

	body, apiError := client.DoRequest(ctx, "GET", path, "",nil)
	if apiError != nil {
		return nil, apiError
	}

	var entry EntryData
	if err := json.Unmarshal(body, &entry); err != nil {
		return nil, FromError(err)
	}

	return &entry, nil
}

func (entries) GetAll(ctx *context.Context, client *Client, flowSlug string) (*EntryList, ApiErrors) {
	path := fmt.Sprintf("/v2/flows/%s/entries", flowSlug)

	body, apiError := client.DoRequest(ctx, "GET", path,  "",nil)
	if apiError != nil {
		return nil, apiError
	}

	var entries EntryList
	if err := json.Unmarshal(body, &entries); err != nil {
		return nil, FromError(err)
	}

	return &entries, nil
}

func (entries) Create(ctx *context.Context, client *Client, flowSlug string, entry *Entry, payload map[string]interface{}) (*EntryData, ApiErrors) {
	entryData := EntryData{
		Data: *entry,
	}

	jsonPayload, err := ToJSON(entryData, payload)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/v2/flows/%s/entries", flowSlug)

	body, apiError := client.DoRequest(ctx, "POST", path, "", bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}

	var newEntry EntryData
	if err := json.Unmarshal(body, &newEntry); err != nil {
		return nil, FromError(err)
	}

	return &newEntry, nil
}

func (entries) Delete(ctx *context.Context, client *Client, flowSlug string, entryID string) ApiErrors {
	path := fmt.Sprintf("/v2/flows/%s/entries/%s", flowSlug, entryID)

	if _, err := client.DoRequest(ctx, "DELETE", path,  "",nil); err != nil {
		return err
	}

	return nil
}

func (entries) Update(ctx *context.Context, client *Client, flowSlug string, entryID string, payload map[string]interface{}) (*EntryData, ApiErrors) {

	entryData := EntryData{
		Data: Entry{
			Type: "entry",
		},
	}

	jsonPayload, err := ToJSON(entryData, payload)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/v2/flows/%s/entries/%s", flowSlug, entryID)

	body, apiError := client.DoRequest(ctx, "PUT", path, "", bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}

	var updatedEntry EntryData
	if err := json.Unmarshal(body, &updatedEntry); err != nil {
		return nil, FromError(err)
	}

	return &updatedEntry, nil
}

type EntryData struct {
	Data Entry `json:"data"`
}

type EntryMeta struct {
	Timestamps Timestamps `json:"timestamps,omitempty"`
}

type EntryDataList struct {
}

type EntryList struct {
	Data [] Entry
}
