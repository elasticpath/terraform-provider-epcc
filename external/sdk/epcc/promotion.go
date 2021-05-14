package epcc

import (
	"bytes"
	"encoding/json"
	"fmt"
)

var Promotions promotions

type promotions struct{}

type Promotion struct {
	Id               string      `json:"id"`
	Type             string      `json:"type"`
	Name             string      `json:"name"`
	Description      string      `json:"description"`
	Enabled          bool        `json:"enabled"`
	Automatic        bool        `json:"automatic"`
	PromotionType    string      `json:"promotion_type"`
	Start            string      `json:"start"`
	End              string      `json:"end"`
	Schema           interface{} `json:"schema"`
	MinCartValue     interface{} `json:"min_cart_value"`
	MaxDiscountValue interface{} `json:"max_discount_value"`
}

func (promotions) Get(client *Client, promotionId string) (*PromotionData, ApiErrors) {
	path := fmt.Sprintf("/v2/promotions/%s", promotionId)

	body, apiError := client.DoRequest("GET", path, nil)
	if apiError != nil {
		return nil, apiError
	}

	var promotions PromotionData
	if err := json.Unmarshal(body, &promotions); err != nil {
		return nil, FromError(err)
	}

	return &promotions, nil
}

// GetAll fetches all promotions
func (promotions) GetAll(client *Client) (*PromotionList, ApiErrors) {
	path := fmt.Sprintf("/v2/promotions")

	body, apiError := client.DoRequest("GET", path, nil)
	if apiError != nil {
		return nil, apiError
	}

	var promotions PromotionList
	if err := json.Unmarshal(body, &promotions); err != nil {
		return nil, FromError(err)
	}

	return &promotions, nil
}

// Create creates a promotion
func (promotions) Create(client *Client, promotion *Promotion) (*PromotionData, ApiErrors) {
	promotionData := PromotionData{
		Data: *promotion,
	}

	jsonPayload, err := json.Marshal(promotionData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/v2/promotions")

	body, apiError := client.DoRequest("POST", path, bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}
	var newPromotion PromotionData
	if err := json.Unmarshal(body, &newPromotion); err != nil {
		return nil, FromError(err)
	}

	return &newPromotion, nil
}

// Delete deletes a promotion.
func (promotions) Delete(client *Client, promotionID string) ApiErrors {
	path := fmt.Sprintf("/v2/promotions/%s", promotionID)

	if _, err := client.DoRequest("DELETE", path, nil); err != nil {
		return err
	}

	return nil
}

// Update updates a promotion.
func (promotions) Update(client *Client, promotionID string, promotion *Promotion) (*PromotionData, ApiErrors) {

	promotionData := PromotionData{
		Data: *promotion,
	}

	jsonPayload, err := json.Marshal(promotionData)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/v2/promotions/%s", promotionID)

	body, apiError := client.DoRequest("PUT", path, bytes.NewBuffer(jsonPayload))

	if apiError != nil {
		return nil, apiError
	}
	var updatedPromotion PromotionData
	if err := json.Unmarshal(body, &updatedPromotion); err != nil {
		return nil, FromError(err)
	}

	return &updatedPromotion, nil
}

type PromotionData struct {
	Data Promotion `json:"data"`
}

// PromotionMeta contains extra data for an promotion
type PromotionMeta struct {
	Timestamps Timestamps `json:"timestamps,omitempty"`
}

type PromotionDataList struct {
}

type PromotionList struct {
	Data []Promotion
}
