package epcc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
)

var Hierarchies hierarchies

type hierarchies struct{}

type Hierarchy struct {
	Id         string              `json:"id,omitempty"`
	Type       string              `json:"type"`
	Attributes HierarchyAttributes `json:"attributes"`
}

type HierarchyAttributes struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Slug        string `json:"slug,omitempty"`
}

func (hierarchies) Get(ctx *context.Context, client *Client, hierarchyId string) (*HierarchyData, ApiErrors) {
	path := fmt.Sprintf("/pcm/hierarchies/%s", hierarchyId)

	body, apiError := client.DoRequest(ctx, "GET", path, "", nil)
	if apiError != nil {
		return nil, apiError
	}

	var hierarchies HierarchyData
	if err := json.Unmarshal(body, &hierarchies); err != nil {
		return nil, FromError(err)
	}

	return &hierarchies, nil
}

// GetAll fetches all hierarchies
func (hierarchies) GetAll(ctx *context.Context, client *Client) (*HierarchyList, ApiErrors) {
	path := fmt.Sprintf("/pcm/hierarchies")

	body, apiError := client.DoRequest(ctx, "GET", path, "", nil)
	if apiError != nil {
		return nil, apiError
	}

	var hierarchies HierarchyList
	if err := json.Unmarshal(body, &hierarchies); err != nil {
		return nil, FromError(err)
	}

	return &hierarchies, nil
}

// Create creates a hierarchy
func (hierarchies) Create(ctx *context.Context, client *Client, hierarchy *Hierarchy) (*HierarchyData, ApiErrors) {
	hierarchyData := HierarchyData{
		Data: *hierarchy,
	}

	jsonPayload, err := json.Marshal(hierarchyData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/pcm/hierarchies")

	body, apiError := client.DoRequest(ctx, "POST", path, "", bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}

	var newHierarchy HierarchyData
	if err := json.Unmarshal(body, &newHierarchy); err != nil {
		return nil, FromError(err)
	}

	return &newHierarchy, nil
}

// Delete deletes a hierarchy.
func (hierarchies) Delete(ctx *context.Context, client *Client, hierarchyID string) ApiErrors {
	path := fmt.Sprintf("/pcm/hierarchies/%s", hierarchyID)

	if _, err := client.DoRequest(ctx, "DELETE", path, "", nil); err != nil {
		return err
	}

	return nil
}

// Update updates a hierarchy.
func (hierarchies) Update(ctx *context.Context, client *Client, hierarchyID string, hierarchy *Hierarchy) (*HierarchyData, ApiErrors) {

	hierarchyData := HierarchyData{
		Data: *hierarchy,
	}

	jsonPayload, err := json.Marshal(hierarchyData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/pcm/hierarchies/%s", hierarchyID)

	body, apiError := client.DoRequest(ctx, "PUT", path, "", bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}

	var updatedHierarchy HierarchyData
	if err := json.Unmarshal(body, &updatedHierarchy); err != nil {
		return nil, FromError(err)
	}

	return &updatedHierarchy, nil
}

type HierarchyData struct {
	Data Hierarchy `json:"data"`
}

type HierarchyDataList struct {
}

type HierarchyList struct {
	Data []Hierarchy
}
