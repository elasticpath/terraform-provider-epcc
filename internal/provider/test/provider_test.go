package test

import (
	"github.com/elasticpath/terraform-provider-epcc/internal/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"os"
	"testing"
)

// providerFactories are used to instantiate a provider during acceptance testing.
// The factory function will be invoked for every Terraform CLI command executed
// to create a provider server to which the CLI can reattach.
var providerFactories = map[string]func() (*schema.Provider, error){
	"epcc": func() (*schema.Provider, error) {
		return provider.New("dev")(), nil
	},
}

func TestProvider(t *testing.T) {

	provider := provider.New("dev")()

	if err := provider.InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}

	resourcesWithNoDescription := make([]string, 0)
	for key, resource := range provider.ResourcesMap {

		if resource.Description == "" {
			resourcesWithNoDescription = append(resourcesWithNoDescription, key)
		}

	}

	dataSourcesWithNoDescription := make([]string, 0)
	for key, dataSource := range provider.DataSourcesMap {
		if dataSource.Description == "" {
			dataSourcesWithNoDescription = append(dataSourcesWithNoDescription, key)
		}

	}

	if len(resourcesWithNoDescription) > 0 || len(dataSourcesWithNoDescription) > 0 {

		badObjects := len(resourcesWithNoDescription) + len(dataSourcesWithNoDescription)
		t.Fatalf("%d object's don't have descriptions:\n\tResources:%s\nData Sources:%s\n", badObjects, resourcesWithNoDescription, dataSourcesWithNoDescription)
	}

}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("EPCC_CLIENT_ID"); v == "" {
		t.Fatal("EPCC_CLIENT_ID must be set for acceptance tests")
	}
	if v := os.Getenv("EPCC_CLIENT_SECRET"); v == "" {
		t.Fatal("EPCC_CLIENT_SECRET must be set for acceptance tests")
	}
}
