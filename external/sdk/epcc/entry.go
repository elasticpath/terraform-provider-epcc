package epcc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

var Entries entries

type entries struct{}

type EntryData struct {
	Data Entry `json:"data"`
}

type EntryList struct {
	Data []Entry
}

type Entry struct {
	Id       string             `json:"id,omitempty"`
	Type     string             `json:"type,omitempty"`
	Strings  map[string]string  `json:"-"`
	Numbers  map[string]float64 `json:"-"`
	Booleans map[string]bool    `json:"-"`
}

func (entries) Get(ctx *context.Context, client *Client, flowSlug string, entryID string) (*EntryData, ApiErrors) {
	if flowSlug == "" {
		return nil, FromError(fmt.Errorf("slug is required"))
	}
	if entryID == "" {
		return nil, FromError(fmt.Errorf("id is required"))
	}

	path := fmt.Sprintf("/v2/flows/%s/entries/%s", flowSlug, entryID)

	body, apiError := client.DoRequest(ctx, "GET", path, "", nil)
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
	if flowSlug == "" {
		return nil, FromError(fmt.Errorf("slug is required"))
	}
	path := fmt.Sprintf("/v2/flows/%s/entries", flowSlug)

	body, apiError := client.DoRequest(ctx, "GET", path, "", nil)
	if apiError != nil {
		return nil, apiError
	}

	var list EntryList
	if err := json.Unmarshal(body, &list); err != nil {
		return nil, FromError(err)
	}

	return &list, nil
}

func (entries) Create(ctx *context.Context, client *Client, flowSlug string, entry *Entry) (*EntryData, ApiErrors) {
	entryData := EntryData{
		Data: *entry,
	}

	jsonPayload, err := json.Marshal(entryData)
	if err != nil {
		return nil, FromError(err)
	}

	if flowSlug == "" {
		return nil, FromError(fmt.Errorf("slug is required"))
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
	if flowSlug == "" {
		return FromError(fmt.Errorf("slug is required"))
	}
	if entryID == "" {
		return FromError(fmt.Errorf("id is required"))
	}
	path := fmt.Sprintf("/v2/flows/%s/entries/%s", flowSlug, entryID)

	if _, err := client.DoRequest(ctx, "DELETE", path, "", nil); err != nil {
		return err
	}

	return nil
}

func (entries) Update(ctx *context.Context, client *Client, flowSlug string, entry *Entry) (*EntryData, ApiErrors) {
	entryData := EntryData{
		Data: *entry,
	}

	jsonPayload, err := json.Marshal(entryData)
	if err != nil {
		return nil, FromError(err)
	}

	if flowSlug == "" {
		return nil, FromError(fmt.Errorf("slug is required"))
	}
	if entry.Id == "" {
		return nil, FromError(fmt.Errorf("id is required"))
	}
	path := fmt.Sprintf("/v2/flows/%s/entries/%s", flowSlug, entry.Id)

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

func (e Entry) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	if e.Id != "" {
		out["id"] = e.Id
	}
	out["type"] = e.Type
	for k, v := range e.Strings {
		out[k] = v
	}
	for k, v := range e.Numbers {
		out[k] = v
	}
	for k, v := range e.Booleans {
		out[k] = v
	}
	return json.Marshal(out)
}

func (e *Entry) UnmarshalJSON(body []byte) error {
	var fieldRawMap map[string]*json.RawMessage
	err := json.Unmarshal(body, &fieldRawMap)
	if err != nil {
		return err
	}

	if err = unmarshalRaw(fieldRawMap, "id", &e.Id); err != nil {
		return err
	}
	if err = unmarshalRaw(fieldRawMap, "type", &e.Type); err != nil {
		return err
	}

	e.Strings = map[string]string{}
	e.Numbers = map[string]float64{}
	e.Booleans = map[string]bool{}

	for key, message := range fieldRawMap {
		if key == "id" || key == "type" || key == "meta" || key == "links" {
			continue
		}

		if message == nil {
			continue
		}

		var s string
		if err = json.Unmarshal(*message, &s); err == nil {
			e.Strings[key] = s
			continue
		}

		var f float64
		if err = json.Unmarshal(*message, &f); err == nil {
			e.Numbers[key] = f
			continue
		}

		var b bool
		if err = json.Unmarshal(*message, &b); err == nil {
			e.Booleans[key] = b
			continue
		}
	}

	return nil
}
