package provider

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"os"
	"testing"
	"time"
)

// providerFactories are used to instantiate a provider during acceptance testing.
// The factory function will be invoked for every Terraform CLI command executed
// to create a provider server to which the CLI can reattach.
var providerFactories = map[string]func() (*schema.Provider, error){
	"epcc": func() (*schema.Provider, error) {
		return New("dev")(), nil
	},
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestProvider(t *testing.T) {

	provider := New("dev")()

	if err := provider.InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}

	resourcesWithNoDescription := make([]string, 0)

	resourceAttributesWithNoDescription := make([]string, 0)

	totalBadCount := 0
	for key, resource := range provider.ResourcesMap {

		if resource.Description == "" {
			resourcesWithNoDescription = append(resourcesWithNoDescription, key)
		}

		count := 0
		for _, schema := range resource.Schema {

			if schema.Description == "" {
				count++
				totalBadCount++
			}
		}

		if count > 0 {
			resourceAttributesWithNoDescription = append(resourceAttributesWithNoDescription, fmt.Sprintf("%s => %d", key, count))
		}

	}

	dataSourcesWithNoDescription := make([]string, 0)
	dataSourceAttributesWithNoDescription := make([]string, 0)
	for key, dataSource := range provider.DataSourcesMap {
		if dataSource.Description == "" {
			dataSourcesWithNoDescription = append(dataSourcesWithNoDescription, key)
		}

		count := 0
		for _, schema := range dataSource.Schema {
			if schema.Description == "" {
				count++
				totalBadCount++
			}
		}
		if count > 0 {
			dataSourceAttributesWithNoDescription = append(dataSourceAttributesWithNoDescription, fmt.Sprintf("%s => %d", key, count))
		}
	}

	dataSourceAndResourcesMissingDescriptions := len(resourcesWithNoDescription) + len(dataSourcesWithNoDescription)

	if dataSourceAndResourcesMissingDescriptions > 0 {
		t.Fatalf("%d object's don't have descriptions:\n\tResources:%s\nData Sources:%s\n", dataSourceAndResourcesMissingDescriptions, resourcesWithNoDescription, dataSourcesWithNoDescription)
	}

	currentDay := int(time.Now().Unix() / 86400)
	currentTarget := max(19110-currentDay, 0)

	if totalBadCount+10 > currentTarget {
		t.Fatalf("%d object's don't have descriptions\n\tWe have a lot of technical debt so this tests permits a non zero value but over time decreases the number of descriptions needed by 1 per day, so just go and get below this number: %d\n\tResources:%s\nData Sources:%s\n", totalBadCount, currentTarget, resourceAttributesWithNoDescription, dataSourceAttributesWithNoDescription)
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
