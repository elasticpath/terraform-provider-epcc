package main

import (
	"encoding/json"
	"fmt"
	"gitlab.elasticpath.com/Steve.Ramage/epcc-terraform-provider/external/sdk/epcc"
	"io"
	"log"
	"net/url"
)

import "net/http"

type TokenResponse struct {
	Expires      int64
	Identifier   string
	Expires_in   int64
	Access_token string
	Token_type   string
}

func main() {
	// Using the Config value, create the Account Client
	client := epcc.NewClient()

	client.Authenticate()

	newAccount := &epcc.Account{
		Type:      "accunt",
		Name:      "Steve's Account",
		LegalName: "Legal Name",
	}

	result, err := epcc.Accounts.Create(client, newAccount)

	log.Printf("Yay! %s, %s", result, err)

	result4, err4 := epcc.Accounts.Create(client, newAccount)

	log.Printf("Yay! %s, %s", result4, err4)

	newAccount.Name = "Steve's Other Account"

	result, err = epcc.Accounts.Update(client, result.Data.Id, newAccount)

	log.Printf("Yay! %s, %s", result, err)

	result2, err2 := epcc.Accounts.GetAll(client)

	log.Printf("Yay! %s, %s", result2, err2)

	err = epcc.Accounts.Delete(client, result.Data.Id)

	log.Printf("Yay! %s", err)

	result3, err3 := epcc.Accounts.GetAll(client)

	log.Printf("Yay! %s, %s", result3, err3)

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
