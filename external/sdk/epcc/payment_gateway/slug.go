package payment_gateway

type Slug string

const (
	Manual               Slug = "manual"
	Stripe               Slug = "stripe"
	StripePaymentIntents Slug = "stripe_payment_intents"
	CyberSource          Slug = "cyber_source"
	PaypalExpress        Slug = "paypal_express"
	PayflowExpress       Slug = "payflow_express"
	Adyen                Slug = "adyen"
	Braintree            Slug = "braintree"
	CardConnect          Slug = "card_connect"
	AuthorizeNet         Slug = "authorize_net"
)

func (s Slug) AsString() string {
	return string(s)
}
