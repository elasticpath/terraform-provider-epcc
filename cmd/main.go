package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"gitlab.elasticpath.com/Steve.Ramage/epcc-terraform-provider/external/sdk/epcc"
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

	newPricebook := &epcc.Pricebook{
		Type: "pricebook",
		Attributes: epcc.PricebookAttributes{
			Name:        "pricebook10",
			Description: "My 10 pricebook",
		},
	}

	result, err := epcc.Pricebooks.Create(client, newPricebook)

	log.Printf("Created! #{result}, #{err}", result, err)

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
	log.Printf("%s\n", foo)
	log.Printf("%s\n", foo.Access_token)

	if err2 != nil {
		log.Fatalf("Error %s %T\n", err2, err2)
	} else {
		log.Printf("Hrm %s\n", foo)
	}
}
