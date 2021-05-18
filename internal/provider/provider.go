package provider

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/elasticpath/terraform-provider-epcc/external/sdk/epcc"

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
				"client_id": &schema.Schema{
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("EPCC_CLIENT_ID", nil),
				},
				"client_secret": &schema.Schema{
					Type:        schema.TypeString,
					Optional:    true,
					Sensitive:   true,
					DefaultFunc: schema.EnvDefaultFunc("EPCC_CLIENT_SECRET", nil),
				},
				"api_base_url": &schema.Schema{
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("EPCC_API_BASE_URL", "https://api.moltin.com/"),
				},
				// TODO Change this to an array maybe that would be cleaner.
				"beta_features": &schema.Schema{
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("EPCC_BETA_API_FEATURES", ""),
				},
				"enable_authentication": &schema.Schema{
					Type:     schema.TypeBool,
					Optional: true,
					Default:  true,
				},
				"additional_headers": &schema.Schema{
					Type:     schema.TypeMap,
					Optional: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				"epcc_account":                  dataSourceEpccAccount(),
				"epcc_catalog":                  dataSourceEpccCatalog(),
				"epcc_currency":                 dataSourceEpccCurrency(),
				"epcc_customer":                 dataSourceEpccCustomer(),
				"epcc_entry":                    dataSourceEpccEntry(),
				"epcc_field":                    dataSourceEpccField(),
				"epcc_file":                     dataSourceEpccFile(),
				"epcc_flow":                     dataSourceEpccFlow(),
				"epcc_hierarchy":                dataSourceEpccHierarchy(),
				"epcc_integration":              IntegrationDataSourceProvider{}.DataSource(),
				"epcc_node":                     dataSourceEpccNode(),
				"epcc_node_product":             dataSourceEpccNodeProduct(),
				"epcc_payment_gateway":          PaymentGatewayDataSourceProvider{}.DataSource(),
				"epcc_pricebook":                dataSourceEpccPricebook(),
				"epcc_product":                  dataSourceEpccProduct(),
				"epcc_product_price":            dataSourceEpccProductPrice(),
				"epcc_promotion":                dataSourceEpccPromotion(),
				"epcc_realm":                    dataSourceEpccRealm(),
				"epcc_profile":                  dataSourceEpccProfile(),
				"epcc_user_authentication_info": dataSourceEpccUserAuthenticationInfo(),
			},

			ResourcesMap: map[string]*schema.Resource{
				"epcc_account":                  resourceEpccAccount(),
				"epcc_catalog":                  resourceEpccCatalog(),
				"epcc_currency":                 resourceEpccCurrency(),
				"epcc_customer":                 resourceEpccCustomer(),
				"epcc_entry":                    resourceEpccEntry(),
				"epcc_field":                    resourceEpccField(),
				"epcc_file":                     resourceEpccFile(),
				"epcc_flow":                     resourceEpccFlow(),
				"epcc_hierarchy":                resourceEpccHierarchy(),
				"epcc_integration":              IntegrationResourceProvider{}.Resource(),
				"epcc_node":                     resourceEpccNode(),
				"epcc_node_product":             resourceEpccNodeProduct(),
				"epcc_payment_gateway":          PaymentGatewayResourceProvider{}.Resource(),
				"epcc_pricebook":                resourceEpccPricebook(),
				"epcc_product":                  resourceEpccProduct(),
				"epcc_product_price":            resourceEpccProductPrice(),
				"epcc_promotion":                resourceEpccPromotion(),
				"epcc_realm":                    resourceEpccRealm(),
				"epcc_profile":                  resourceEpccProfile(),
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
		epccClientSecret := d.Get("client_secret").(string)
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
			UserAgent:         "terraform-provider-epcc / " + version,
			AdditionalHeaders: &stringAdditionalHeaders,
			RetryLimitTimeout: 120 * time.Second,
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

func FromAPIError(err epcc.ApiErrors) diag.Diagnostics {
	if err == nil {
		return nil
	}

	return diag.Diagnostics{
		diag.Diagnostic{
			Severity: diag.Error,
			Summary:  err.Error(),
			Detail:   fmt.Sprintf("API Error Response [%s %s => %d]\n%s", err.HttpMethod(), err.HttpPath(), err.HttpStatusCode(), strings.ReplaceAll("\n"+err.ListOfErrors().String(), "\n", "\n\t")),
		},
	}
}
