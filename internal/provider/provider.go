package provider

import (
	"context"
	"fmt"
	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"
	"github.com/hashicorp/go-cty/cty"
	"math"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

/*
		ClientID     string `envconfig:"EPCC_CLIENT_ID"`
		ClientSecret string `envconfig:"EPCC_CLIENT_SECRET"`
	}
	BaseURL           string `envconfig:"EPCC_API_BASE_URL"`
	BetaFeatures	  string `envconfig:"EPCC_BETA_API_FEATURES"`
*/
func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"client_id": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("EPCC_CLIENT_ID", nil),
					Description: "The **Client ID** API key for the store, this value is available in Commerce Manager under \"Your API keys\"",
				},
				"client_secret": {
					Type:        schema.TypeString,
					Optional:    true,
					Sensitive:   true,
					DefaultFunc: schema.EnvDefaultFunc("EPCC_CLIENT_SECRET", nil),
					Description: "The **Client Secret** API key for the store, this value is available in Commerce Manager under \"Your API keys\"",
				},
				"api_base_url": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("EPCC_API_BASE_URL", "https://api.moltin.com/"),
					Description: "The **API base URL** for the store, this value is available in Commerce Manager under \"Your API keys\"",
				},
				// TODO Change this to an array maybe that would be cleaner.
				"beta_features": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("EPCC_BETA_API_FEATURES", ""),
					Description: "The value to use for the `EP-Beta_Features` header value which controls access to [Beta APIs](https://documentation.elasticpath.com/commerce-cloud/docs/api/basics/api-contract.html#beta-apis)",
				},
				"enable_authentication": {
					Type:        schema.TypeBool,
					Optional:    true,
					Default:     true,
					Description: "Controls whether or not to authenticate before making a request. Disabling this may be appropriate if you are using additional_headers to supply an authentication token.",
				},
				"additional_headers": {
					Type:     schema.TypeMap,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Description: "A set of additional HTTP Headers to send on all requests",
				},
				"rate_limit": {
					Type:        schema.TypeInt,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("EPCC_RATE_LIMIT", 25),
					ValidateDiagFunc: func(i interface{}, path cty.Path) diag.Diagnostics {
						integer := i.(int)

						if integer < 1 || integer > math.MaxUint16 {
							return diag.FromErr(fmt.Errorf("rate_limit value `%d` is not in the range [%d, %d]", integer, 1, math.MaxUint16))
						}

						return nil

					},
					Description: "Controls the maximum number of requests this provider will make per second, which conforms to the [Rate Limits](https://documentation.elasticpath.com/commerce-cloud/docs/api/basics/rate-limits.html) of EPCC.",
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				"epcc_account":                          dataSourceEpccAccount(),
				"epcc_account_authentication_settings":  dataSourceAccountAuthenticationSettings(),
				"epcc_authentication_realm":             dataSourceEpccAuthenticationRealm(),
				"epcc_catalog":                          dataSourceEpccCatalog(),
				"epcc_catalog_rule":                     dataSourceEpccCatalogRule(),
				"epcc_currency":                         dataSourceEpccCurrency(),
				"epcc_customer":                         dataSourceEpccCustomer(),
				"epcc_customer_authentication_settings": dataSourceCustomerAuthenticationSettings(),
				"epcc_entry":                            EntryDataSourceProvider{}.DataSource(),
				"epcc_field":                            dataSourceEpccField(),
				"epcc_file":                             dataSourceEpccFile(),
				"epcc_flow":                             dataSourceEpccFlow(),
				"epcc_hierarchy":                        dataSourceEpccHierarchy(),
				"epcc_integration":                      IntegrationDataSourceProvider{}.DataSource(),
				"epcc_merchant_realm_mappings":          dataSourceMerchantRealmMappings(),
				"epcc_node":                             dataSourceEpccNode(),
				"epcc_node_product":                     dataSourceEpccNodeProduct(),
				"epcc_oidc_profile":                     dataSourceEpccOidcProfile(),
				"epcc_payment_gateway":                  PaymentGatewayDataSourceProvider{}.DataSource(),
				"epcc_pricebook":                        dataSourceEpccPricebook(),
				"epcc_product":                          dataSourceEpccProduct(),
				"epcc_product_price":                    dataSourceEpccProductPrice(),
				"epcc_promotion":                        dataSourceEpccPromotion(),
				"epcc_settings":                         dataSourceEpccSettings(),
				"epcc_user_authentication_info":         dataSourceEpccUserAuthenticationInfo(),
			},

			ResourcesMap: map[string]*schema.Resource{
				"epcc_account":                  resourceEpccAccount(),
				"epcc_authentication_realm":     resourceEpccAuthenticationRealm(),
				"epcc_catalog":                  resourceEpccCatalog(),
				"epcc_catalog_rule":             resourceEpccCatalogRule(),
				"epcc_currency":                 resourceEpccCurrency(),
				"epcc_customer":                 resourceEpccCustomer(),
				"epcc_entry":                    EntryResourceProvider{}.Resource(),
				"epcc_field":                    resourceEpccField(),
				"epcc_file":                     resourceEpccFile(),
				"epcc_flow":                     resourceEpccFlow(),
				"epcc_hierarchy":                resourceEpccHierarchy(),
				"epcc_integration":              IntegrationResourceProvider{}.Resource(),
				"epcc_merchant_realm_mapping":   resourceEpccMerchantRealmMapping(),
				"epcc_node":                     resourceEpccNode(),
				"epcc_node_product":             resourceEpccNodeProduct(),
				"epcc_oidc_profile":             resourceEpccOidcProfile(),
				"epcc_payment_gateway":          PaymentGatewayResourceProvider{}.Resource(),
				"epcc_pricebook":                resourceEpccPricebook(),
				"epcc_product":                  resourceEpccProduct(),
				"epcc_product_price":            resourceEpccProductPrice(),
				"epcc_promotion":                resourceEpccPromotion(),
				"epcc_settings":                 resourceEpccSettings(),
				"epcc_user_authentication_info": resourceEpccUserAuthenticationInfo(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		for key, dataSource := range p.DataSourcesMap {

			if dataSource.Description == "" {
				resource := p.ResourcesMap[key]

				if resource != nil {
					dataSource.Description = resource.Description
				}
			}
		}
		return p
	}
}

func configure(version string, p *schema.Provider) func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {

		// Warning or errors can be collected in a slice type
		var diags diag.Diagnostics

		epccClientId := d.Get("client_id").(string)

		if epccClientId == "" {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create EPCC Client",
				Detail:   "To resolve this, ensure that you have properly set `client_id` in the provider configuration or the `EPCC_CLIENT_ID` environment variable.",
			})
			return nil, diags
		}

		epccClientSecret := d.Get("client_secret").(string)

		if epccClientSecret == "" {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create EPCC Client",
				Detail:   "To resolve this, ensure that you  have properly set `client_secret` in the provider configuration or set the `EPCC_CLIENT_SECRET` environment variable.",
			})
			return nil, diags
		}

		epccApiBaseUrl := d.Get("api_base_url").(string)
		epccBetaFeatures := d.Get("beta_features").(string)

		additionalHeaders := d.Get("additional_headers").(map[string]interface{})

		stringAdditionalHeaders := make(map[string]string)

		if additionalHeaders != nil {
			for key, val := range additionalHeaders {
				(stringAdditionalHeaders)[key] = fmt.Sprintf("%s", val)
			}
		}

		client := epcc.NewClient(epcc.ClientOptions{
			BaseURL:      epccApiBaseUrl,
			BetaFeatures: epccBetaFeatures,
			Credentials: &epcc.Credentials{
				ClientId:     epccClientId,
				ClientSecret: epccClientSecret,
			},
			UserAgent:                  "terraform-provider-epcc / " + version,
			AdditionalHeaders:          &stringAdditionalHeaders,
			RetryLimitTimeout:          120 * time.Second,
			RateLimitRequestsPerSecond: uint16(d.Get("rate_limit").(int)),
		})

		enableAuthentication := d.Get("enable_authentication").(bool)

		if enableAuthentication {
			err := client.Authenticate()
			if err != nil {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Unable to create EPCC Client",
					Detail:   "Unable to authenticate against the EPCC API: " + err.Error(),
				})
				return nil, diags
			}
		}
		return client, diags
	}
}

func ReportAPIError(ctx context.Context, err epcc.ApiErrors) {
	if err == nil {
		return
	}
	diagnostics := ctx.Value("diags").(*diag.Diagnostics)

	diagnosticsAppended := append(*diagnostics, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  err.Error(),
		Detail:   fmt.Sprintf("API Error Response [%s %s => %d]\n%s", err.HttpMethod(), err.HttpPath(), err.HttpStatusCode(), strings.ReplaceAll("\n"+err.ListOfErrors().String(), "\n", "\n\t")),
	})
	*diagnostics = diagnosticsAppended
}

func addToDiag(ctx context.Context, diagnostics diag.Diagnostics) {
	ctxDiags := ctx.Value("diags").(*diag.Diagnostics)
	diagnosticsAppended := append(*ctxDiags, diagnostics...)
	*ctxDiags = diagnosticsAppended
}
