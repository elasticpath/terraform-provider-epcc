package epcc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
)

var Nodes nodes

type nodes struct{}

type Node struct {
	Id            string              `json:"id,omitempty"`
	Type          string              `json:"type"`
	Attributes    NodeAttributes      `json:"attributes"`
	Relationships *NodesRelationships `json:"relationships,omitempty"`
}

type NodeAttributes struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Slug        string `json:"slug,omitempty"`
}

type NodesRelationships struct {
	Parent *DataForTypeIdRelationship `json:"parent,omitempty"`
}

type DataForTypeIdRelationship struct {
	Data *TypeIdRelationship `json:"data"`
}

type DataForTypeIdRelationshipList struct {
	Data *[]TypeIdRelationship `json:"data"`
}

type TypeIdRelationship struct {
	Id   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
}

func (nodes) Get(ctx *context.Context, client *Client, hierarchyId string, nodeId string) (*NodeData, ApiErrors) {
	path := fmt.Sprintf("/pcm/hierarchies/%s/nodes/%s", hierarchyId, nodeId)

	body, apiError := client.DoRequest(ctx, "GET", path, "", nil)
	if apiError != nil {
		return nil, apiError
	}

	var nodes NodeData
	if err := json.Unmarshal(body, &nodes); err != nil {
		return nil, FromError(err)
	}

	return &nodes, nil
}

// Create creates a node
func (nodes) Create(ctx *context.Context, client *Client, hierarchyId string, node *Node) (*NodeData, ApiErrors) {
	nodeData := NodeData{
		Data: *node,
	}

	jsonPayload, err := json.Marshal(nodeData)
	if err != nil {
		return nil, FromError(err)
	}

	log.Printf("[ERROR] %s\n", jsonPayload)
	path := fmt.Sprintf("/pcm/hierarchies/%s/nodes", hierarchyId)

	body, apiError := client.DoRequest(ctx, "POST", path, "", bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}

	var newNode NodeData
	if err := json.Unmarshal(body, &newNode); err != nil {
		return nil, FromError(err)
	}

	return &newNode, nil
}

// Delete deletes a node.
func (nodes) Delete(ctx *context.Context, client *Client, hierarchyId string, nodeID string) ApiErrors {
	path := fmt.Sprintf("/pcm/hierarchies/%s/nodes/%s", hierarchyId, nodeID)

	if _, err := client.DoRequest(ctx, "DELETE", path, "", nil); err != nil {
		return err
	}

	return nil
}

// Update updates a node.
func (nodes) Update(ctx *context.Context, client *Client, hierarchyId string, nodeID string, node *Node) (*NodeData, ApiErrors) {

	nodeData := NodeData{
		Data: *node,
	}

	// Update the Parent
	if nodeData.Data.Relationships != nil {
		if len(nodeData.Data.Relationships.Parent.Data.Id) == 0 {
			if _, err := client.DoRequest(ctx, "DELETE", fmt.Sprintf("/pcm/hierarchies/%s/nodes/%s/relationships/parent", hierarchyId, nodeID), "", nil); err != nil {
				return nil, FromError(err)
			}
		} else {
			// Update The Node
			jsonPayload, err := json.Marshal(nodeData.Data.Relationships.Parent)
			if err != nil {
				return nil, FromError(err)
			}

			if _, err := client.DoRequest(ctx, "PUT", fmt.Sprintf("/pcm/hierarchies/%s/nodes/%s/relationships/parent", hierarchyId, nodeID), "", bytes.NewBuffer(jsonPayload)); err != nil {
				return nil, FromError(err)
			}
		}
	}

	path := fmt.Sprintf("/pcm/hierarchies/%s/nodes/%s", hierarchyId, nodeID)

	nodeData.Data.Relationships = nil
	// Update The Node
	jsonPayload, err := json.Marshal(nodeData)
	if err != nil {
		return nil, FromError(err)
	}

	log.Printf("[ERROR] %s\n", jsonPayload)

	body, apiError := client.DoRequest(ctx, "PUT", path, "", bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}

	var updatedNode NodeData
	if err := json.Unmarshal(body, &updatedNode); err != nil {
		return nil, FromError(err)
	}

	return &updatedNode, nil
}

func (nodes) CreateNodeProducts(ctx *context.Context, client *Client, hierarchyId string, nodeID string, reference DataForTypeIdRelationshipList) ApiErrors {

	jsonPayload, err := json.Marshal(reference)
	if err != nil {
		return FromError(err)
	}
	log.Printf("jsonPayload: " + string(jsonPayload))

	path := fmt.Sprintf("/pcm/hierarchies/%s/nodes/%s/relationships/products", hierarchyId, nodeID)

	_, apiError := client.DoRequest(ctx, "POST", path, "", bytes.NewBuffer(jsonPayload))

	return apiError
}

func (nodes) UpdateNodeProducts(ctx *context.Context, client *Client, hierarchyId string, nodeID string, reference DataForTypeIdRelationshipList) ApiErrors {

	jsonPayload, err := json.Marshal(reference)
	if err != nil {
		return FromError(err)
	}
	log.Printf("jsonPayload: " + string(jsonPayload))

	path := fmt.Sprintf("/pcm/hierarchies/%s/nodes/%s/relationships/products", hierarchyId, nodeID)

	_, apiError := client.DoRequest(ctx, "PUT", path, "", bytes.NewBuffer(jsonPayload))

	return apiError
}

func (nodes) DeleteNodeProduct(ctx *context.Context, client *Client, hierarchyId string, nodeID string, reference DataForTypeIdRelationshipList) ApiErrors {

	jsonPayload, err := json.Marshal(reference)
	if err != nil {
		return FromError(err)
	}

	path := fmt.Sprintf("/pcm/hierarchies/%s/nodes/%s/relationships/products", hierarchyId, nodeID)

	_, apiError := client.DoRequest(ctx, "DELETE", path, "", bytes.NewBuffer(jsonPayload))

	return apiError
}

func (nodes) GetNodeProducts(ctx *context.Context, client *Client, hierarchyId string, nodeID string) (*DataForTypeIdRelationshipList, ApiErrors) {
	path := fmt.Sprintf("/pcm/hierarchies/%s/nodes/%s/products", hierarchyId, nodeID)

	var offset = 0
	list := make([]TypeIdRelationship, 0)
	for {
		body, apiError := client.DoRequest(ctx, "GET", path, fmt.Sprintf("page%%5Boffset%%5D=%d", offset), nil)
		if apiError != nil {
			return nil, apiError
		}

		var fileRelationships DataForTypeIdRelationshipList
		if err := json.Unmarshal(body, &fileRelationships); err != nil {
			return nil, FromError(err)
		}

		list = append(list, *fileRelationships.Data...)

		count := len(*fileRelationships.Data)
		if count <= 0 {
			break
		}

		if offset > 6000 {
			break
		}
		offset += count
	}

	return &DataForTypeIdRelationshipList{
		Data: &list,
	}, nil

}

type NodeData struct {
	Data Node `json:"data"`
}

type NodeDataList struct {
}

type NodeList struct {
	Data []Node
}
