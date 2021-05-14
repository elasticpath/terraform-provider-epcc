package epcc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/elasticpath/epcc-terraform-provider/external/sdk/epcc/payment_gateway"
)

var PaymentGateways paymentGateways

type paymentGateways struct{}

type PaymentGatewayData struct {
	Data PaymentGateway `json:"data"`
}

type PaymentGateway interface {
	Type() payment_gateway.Slug
	Base() *PaymentGatewayBase
}

type PaymentGatewayBase struct {
	Name    string `json:"name,omitempty"`
	Slug    string `json:"slug"`
	Enabled bool   `json:"enabled"`
	Test    bool   `json:"test"`
}

func (p *PaymentGatewayBase) Type() payment_gateway.Slug {
	return payment_gateway.Slug(p.Slug)
}

func (p *PaymentGatewayBase) Base() *PaymentGatewayBase {
	return p
}

type ManualPaymentGateway struct {
	PaymentGatewayBase
}

type StripePaymentGateway struct {
	PaymentGatewayBase
	Login string `json:"login"`
}

type CyberSourcePaymentGateway struct {
	PaymentGatewayBase
	Login    string `json:"login"`
	Password string `json:"password"`
}

type PayPalExpressPaymentGateway struct {
	PaymentGatewayBase
	Login     string `json:"login"`
	Password  string `json:"password"`
	Signature string `json:"signature"`
}

type PayPalPayflowPaymentGateway struct {
	PaymentGatewayBase
	Partner  string `json:"partner"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type AdyenPaymentGateway struct {
	PaymentGatewayBase
	MerchantAccount string `json:"merchant_account"`
	Username        string `json:"username"`
	Password        string `json:"password"`
}

type BraintreePaymentGateway struct {
	PaymentGatewayBase
	MerchantId  string `json:"merchant_id"`
	PrivateKey  string `json:"private_key"`
	PublicKey   string `json:"public_key"`
	Environment string `json:"environment"`
}

type CartConnectPaymentGateway struct {
	PaymentGatewayBase
	MerchantId string `json:"merchant_id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

type AuthorizeNetPaymentGateway struct {
	PaymentGatewayBase
	Login    string `json:"login"`
	Password string `json:"password"`
}

type PaymentGatewayList struct {
	Data []PaymentGateway
}

// Get fetches one PaymentGateway by id
func (paymentGateways) Get(client *Client, slug payment_gateway.Slug) (*PaymentGatewayData, ApiErrors) {
	path := fmt.Sprintf("/v2/gateways/%s", slug)

	body, apiError := client.DoRequest("GET", path, nil)
	if apiError != nil {
		return nil, apiError
	}

	var data PaymentGatewayData
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, FromError(err)
	}

	return &data, nil
}

// GetAll fetches PaymentGatewayList
func (paymentGateways) GetAll(client *Client) (*PaymentGatewayList, ApiErrors) {
	path := fmt.Sprintf("/v2/gateways")

	body, apiError := client.DoRequest("GET", path, nil)
	if apiError != nil {
		return nil, apiError
	}

	var list PaymentGatewayList
	if err := json.Unmarshal(body, &list); err != nil {
		return nil, FromError(err)
	}

	return &list, nil
}

// Update updates a Integration.
func (paymentGateways) Update(client *Client, slug payment_gateway.Slug, obj *PaymentGateway) (*PaymentGatewayData, ApiErrors) {
	data := PaymentGatewayData{
		Data: *obj,
	}

	jsonPayload, err := json.Marshal(data)
	if err != nil {
		return nil, FromError(err)
	}

	path := fmt.Sprintf("/v2/gateways/%s", slug)

	body, apiError := client.DoRequest("PUT", path, bytes.NewBuffer(jsonPayload))
	if apiError != nil {
		return nil, apiError
	}

	var updated PaymentGatewayData
	if err := json.Unmarshal(body, &updated); err != nil {
		return nil, FromError(err)
	}

	return &updated, nil
}

func (d *PaymentGatewayData) UnmarshalJSON(body []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(body, &objMap)
	if err != nil {
		return err
	}

	data, ok := objMap["data"]
	if !ok {
		return fmt.Errorf("no data found in the response body")
	}

	var base PaymentGatewayBase
	err = json.Unmarshal(*data, &base)
	if err != nil {
		return err
	}

	switch base.Type() {
	case payment_gateway.Manual:
		var concrete ManualPaymentGateway
		err = json.Unmarshal(*data, &concrete)
		if err != nil {
			return err
		}
		d.Data = &concrete
	case payment_gateway.Stripe:
		fallthrough
	case payment_gateway.StripePaymentIntents:
		var concrete StripePaymentGateway
		err = json.Unmarshal(*data, &concrete)
		if err != nil {
			return err
		}
		d.Data = &concrete
	case payment_gateway.CyberSource:
		var concrete CyberSourcePaymentGateway
		err = json.Unmarshal(*data, &concrete)
		if err != nil {
			return err
		}
		d.Data = &concrete
	case payment_gateway.PaypalExpress:
		var concrete PayPalExpressPaymentGateway
		err = json.Unmarshal(*data, &concrete)
		if err != nil {
			return err
		}
		d.Data = &concrete
	case payment_gateway.PayflowExpress:
		var concrete PayPalPayflowPaymentGateway
		err = json.Unmarshal(*data, &concrete)
		if err != nil {
			return err
		}
		d.Data = &concrete
	case payment_gateway.Adyen:
		var concrete AdyenPaymentGateway
		err = json.Unmarshal(*data, &concrete)
		if err != nil {
			return err
		}
		d.Data = &concrete
	case payment_gateway.Braintree:
		var concrete BraintreePaymentGateway
		err = json.Unmarshal(*data, &concrete)
		if err != nil {
			return err
		}
		d.Data = &concrete
	case payment_gateway.CardConnect:
		var concrete CartConnectPaymentGateway
		err = json.Unmarshal(*data, &concrete)
		if err != nil {
			return err
		}
		d.Data = &concrete
	case payment_gateway.AuthorizeNet:
		var concrete AuthorizeNetPaymentGateway
		err = json.Unmarshal(*data, &concrete)
		if err != nil {
			return err
		}
		d.Data = &concrete
	default:
		panic("unknown payment gateway type " + base.Type())
	}

	return nil
}
