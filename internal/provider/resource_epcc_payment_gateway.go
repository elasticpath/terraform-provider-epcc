package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc/payment_gateway"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type PaymentGatewayResourceProvider struct {
}

func (r PaymentGatewayResourceProvider) Resource() *schema.Resource {
	return &schema.Resource{
		Description:   "Payment gateway connectivity configuration",
		CreateContext: addDiagToContext(r.update),
		ReadContext:   addDiagToContext(r.read),
		UpdateContext: addDiagToContext(r.update),
		DeleteContext: addDiagToContext(r.delete),
		Schema: map[string]*schema.Schema{
			"slug": {
				Type:     schema.TypeString,
				Required: true,
			},
			"enabled": {
				Type:        schema.TypeBool,
				Description: "Should the gateway process payments. Default: `false`",
				Optional:    true,
			},
			"test": {
				Type:        schema.TypeBool,
				Description: "Is this a sandbox environment. Default: `false`",
				Optional:    true,
			},
			"options": {
				Description: "Parameters specific to concrete payment provider",
				Type:        schema.TypeMap,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, data *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				err := data.Set("slug", data.Id())
				return []*schema.ResourceData{data}, err
			},
		},
	}
}

func (r PaymentGatewayResourceProvider) update(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	result, diagErr := r.updatePaymentGateway(ctx, client, data)
	if diagErr != nil {
		return diagErr
	}

	diagErr = r.parseResourceData(ctx, result, data)
	if diagErr != nil {
		return diagErr
	}

	data.SetId(result.Data.Type().AsString())

	return nil
}

func (r PaymentGatewayResourceProvider) parseResourceData(ctx context.Context, result *epcc.PaymentGatewayData, data *schema.ResourceData) diag.Diagnostics {
	base := result.Data.Base()
	if err := data.Set("enabled", base.Enabled); err != nil {
		return diag.FromErr(err)
	}
	if err := data.Set("test", base.Test); err != nil {
		return diag.FromErr(err)
	}

	optionsBytes, jsonErr := json.Marshal(result.Data)
	if jsonErr != nil {
		return diag.FromErr(jsonErr)
	}
	var options map[string]interface{}
	jsonErr = json.Unmarshal(optionsBytes, &options)
	if jsonErr != nil {
		return diag.FromErr(jsonErr)
	}
	delete(options, "slug")
	delete(options, "enabled")
	delete(options, "test")
	if err := data.Set("options", &options); err != nil {
		return diag.FromErr(err)
	}
	return nil
}

func (r PaymentGatewayResourceProvider) updatePaymentGateway(ctx context.Context, client *epcc.Client, data *schema.ResourceData) (*epcc.PaymentGatewayData, diag.Diagnostics) {

	base := epcc.PaymentGatewayBase{
		Slug:    data.Get("slug").(string),
		Enabled: data.Get("enabled").(bool),
		Test:    data.Get("test").(bool),
	}
	options := data.Get("options").(map[string]interface{})

	var obj epcc.PaymentGateway = &base
	switch base.Type() {
	case payment_gateway.Manual:
		obj = &epcc.ManualPaymentGateway{
			PaymentGatewayBase: base,
		}
	case payment_gateway.Stripe:
		fallthrough
	case payment_gateway.StripePaymentIntents:
		obj = &epcc.StripePaymentGateway{
			PaymentGatewayBase: base,
			Login:              mapStringValue(options, "login"),
		}
	case payment_gateway.CyberSource:
		obj = &epcc.CyberSourcePaymentGateway{
			PaymentGatewayBase: base,
			Login:              mapStringValue(options, "login"),
			Password:           mapStringValue(options, "password"),
		}
	case payment_gateway.PaypalExpress:
		obj = &epcc.PayPalExpressPaymentGateway{
			PaymentGatewayBase: base,
			Login:              mapStringValue(options, "login"),
			Password:           mapStringValue(options, "password"),
			Signature:          mapStringValue(options, "signature"),
		}
	case payment_gateway.PayflowExpress:
		obj = &epcc.PayPalPayflowPaymentGateway{
			PaymentGatewayBase: base,
			Partner:            mapStringValue(options, "partner"),
			Login:              mapStringValue(options, "login"),
			Password:           mapStringValue(options, "password"),
		}
	case payment_gateway.Adyen:
		obj = &epcc.AdyenPaymentGateway{
			PaymentGatewayBase: base,
			MerchantAccount:    mapStringValue(options, "merchant_account"),
			Username:           mapStringValue(options, "username"),
			Password:           mapStringValue(options, "password"),
		}
	case payment_gateway.Braintree:
		if base.Test {
			return nil, diag.FromErr(fmt.Errorf("test parameter is not supported by Braintree"))
		}
		obj = &epcc.BraintreePaymentGateway{
			PaymentGatewayBase: base,
			MerchantId:         mapStringValue(options, "merchant_id"),
			PrivateKey:         mapStringValue(options, "private_key"),
			PublicKey:          mapStringValue(options, "public_key"),
			Environment:        mapStringValue(options, "environment"),
		}
	case payment_gateway.CardConnect:
		obj = &epcc.CartConnectPaymentGateway{
			PaymentGatewayBase: base,
			MerchantId:         mapStringValue(options, "merchant_id"),
			Username:           mapStringValue(options, "username"),
			Password:           mapStringValue(options, "password"),
		}
	case payment_gateway.AuthorizeNet:
		obj = &epcc.AuthorizeNetPaymentGateway{
			PaymentGatewayBase: base,
			Login:              mapStringValue(options, "login"),
			Password:           mapStringValue(options, "password"),
		}
	}

	result, apiError := epcc.PaymentGateways.Update(&ctx, client, obj.Type(), &obj)
	if apiError != nil {
		return nil, FromAPIError(apiError)
	}

	return result, nil
}

func mapStringValue(m map[string]interface{}, key string) string {
	value, ok := m[key]
	if !ok {
		return ""
	}
	return value.(string)
}

func (r PaymentGatewayResourceProvider) read(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	slug := payment_gateway.Slug(data.Get("slug").(string))
	result, err := epcc.PaymentGateways.Get(&ctx, client, slug)
	if err != nil {
		return FromAPIError(err)
	}

	base := result.Data.Base()
	if err := data.Set("enabled", base.Enabled); err != nil {
		return diag.FromErr(err)
	}
	if err := data.Set("test", base.Test); err != nil {
		return diag.FromErr(err)
	}

	data.SetId(result.Data.Type().AsString())

	return nil
}

func (r PaymentGatewayResourceProvider) delete(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*epcc.Client)

	if err := data.Set("options", map[string]interface{}{}); err != nil {
		return diag.FromErr(err)
	}
	if err := data.Set("enabled", false); err != nil {
		return diag.FromErr(err)
	}
	if err := data.Set("test", false); err != nil {
		return diag.FromErr(err)
	}

	_, diagErr := r.updatePaymentGateway(ctx, client, data)
	if diagErr != nil {
		return diagErr
	}

	data.SetId("")

	return nil
}
