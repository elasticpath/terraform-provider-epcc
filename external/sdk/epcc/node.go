package epcc

import (
	"bytes"
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

func (nodes) Get(client *Client, hierarchyId string, nodeId string) (*NodeData, ApiErrors) {
	path := fmt.Sprintf("/pcm/hierarchies/%s/nodes/%s", hierarchyId, nodeId)

	body, apiError := client.DoRequest("GET", path, nil)
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
func (nodes) Create(client *Client, hierarchyId string, node *Node) (*NodeData, ApiErrors) {
	nodeData := NodeData{
		Data: *node,
	}

	jsonPayload, err := json.Marshal(nodeData)
	if err != nil {
		return nil, FromError(err)
	}

	log.Printf("[ERROR] %s\n", jsonPayload)
	path := fmt.Sprintf("/pcm/hierarchies/%s/nodes", hierarchyId)

	body, apiError := client.DoRequest("POST", path, bytes.NewBuffer(jsonPayload))
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
func (nodes) Delete(client *Client, hierarchyId string, nodeID string) ApiErrors {
	path := fmt.Sprintf("/pcm/hierarchies/%s/nodes/%s", hierarchyId, nodeID)

	if _, err := client.DoRequest("DELETE", path, nil); err != nil {
		return err
	}

	return nil
}

// Update updates a node.
func (nodes) Update(client *Client, hierarchyId string, nodeID string, node *Node) (*NodeData, ApiErrors) {

	nodeData := NodeData{
		Data: *node,
	}

	// Update the Parent
	if nodeData.Data.Relationships != nil {
		if len(nodeData.Data.Relationships.Parent.Data.Id) == 0 {
			if _, err := client.DoRequest("DELETE", fmt.Sprintf("/pcm/hierarchies/%s/nodes/%s/relationships/parent", hierarchyId, nodeID), nil); err != nil {
				return nil, FromError(err)
			}
		} else {
			// Update The Node
			jsonPayload, err := json.Marshal(nodeData.Data.Relationships.Parent)
			if err != nil {
				return nil, FromError(err)
			}

			if _, err := client.DoRequest("PUT", fmt.Sprintf("/pcm/hierarchies/%s/nodes/%s/relationships/parent", hierarchyId, nodeID), bytes.NewBuffer(jsonPayload)); err != nil {
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

	body, apiError := client.DoRequest("PUT", path, bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}

	var updatedNode NodeData
	if err := json.Unmarshal(body, &updatedNode); err != nil {
		return nil, FromError(err)
	}

	return &updatedNode, nil
}

type NodeData struct {
	Data Node `json:"data"`
}

type NodeDataList struct {
}

type NodeList struct {
	Data []Node
}
