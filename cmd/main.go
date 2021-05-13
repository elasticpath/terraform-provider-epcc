package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/elasticpath/epcc-terraform-provider/external/sdk/epcc"
)

type TokenResponse struct {
	Expires      int64
	Identifier   string
	Expires_in   int64
	Access_token string
	Token_type   string
}

func main() {
	// Using the Config value, create the Client
	client := epcc.NewClient()

	client.Authenticate()

	//exercisePricebook(client)
	exerciseProduct(client)

}

func main2() {
	fmt.Println("Hello, World")

	resp, err := http.PostForm("https://api.moltin.com/oauth/access_token", url.Values{"key": {"Value"}, "client_id": {"fksQxPnjYROsHMFbsrk6AQmshvPoOcD8HNF5qafqi0"}, "grant_type": {"implicit"}})

	if err != nil {
		log.Fatalf("Error %s", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalf("Error %s", err)
	}

	log.Printf("%s", body)

	foo := TokenResponse{}

	err2 := json.Unmarshal(body, &foo)
	log.Printf("%s\n", string(body))
	log.Printf("%s\n", foo.Access_token)

	if err2 != nil {
		log.Fatalf("Error %s %T\n", err2, err2)
	} else {
		log.Printf("Hrm %s\n", string(body))
	}
}

func exercisePricebook(client *epcc.Client) {

	newPricebook := &epcc.Pricebook{
		Type: "pricebook",
		Attributes: epcc.PricebookAttributes{
			Name:        "pricebook10",
			Description: "My 10 pricebook",
		},
	}

	result, err := epcc.Pricebooks.Create(client, newPricebook)

	log.Printf("Created! %s, %s", result, err)

	// Subsequent requests need the ID included in the body
	newPricebook.Id = result.Data.Id

	// Change the Name
	newPricebook.Attributes.Name = "pricebook22"
	result, err = epcc.Pricebooks.Update(client, result.Data.Id, newPricebook)
	log.Printf("Updated! %s, %s", result, err)

	// Delete
	err = epcc.Pricebooks.Delete(client, result.Data.Id)
	log.Printf("Deleted! %s, %s", result, err)
}

func exerciseProduct(client *epcc.Client) {

	newProduct := &epcc.Product{
		Type: "product",
		Attributes: epcc.ProductAttributes{
			Name:          "Product1",
			CommodityType: "physical",
			Sku:           "product-1",
			Description:   "My 1 product",
		},
	}

	result, err := epcc.Products.Create(client, newProduct)

	log.Printf("Created! %s, %s", result, err)

	// Subsequent requests need the ID included in the body
	newProduct.Id = result.Data.Id

	// Change the Name
	newProduct.Attributes.Name = "Product1b"
	result, err = epcc.Products.Update(client, result.Data.Id, newProduct)
	log.Printf("Updated! %s, %s", result, err)

	// Delete
	err = epcc.Products.Delete(client, result.Data.Id)
	log.Printf("Deleted! %s, %s", result, err)
}
