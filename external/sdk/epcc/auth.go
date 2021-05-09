package epcc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type authResponse struct {
	Expires     int    `json:"expires"`
	ExpiresIn   int    `json:"expires_in"`
	Identifier  string `json:"identifier"`
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
}

//auth returns an AccessToken or an Error
func auth(client Client) (string, error) {
	reqURL, err := url.Parse(client.BaseURL)

	reqURL.Path = fmt.Sprintf("/oauth/access_token")

	values := url.Values{}
	values.Set("client_id", client.Credentials.ClientId)
	values.Set("client_secret", client.Credentials.ClientSecret)
	values.Set("grant_type", "client_credentials")

	body := strings.NewReader(values.Encode())

	req, err := http.NewRequest("POST", reqURL.String(), body)
	if err != nil {
		return "", err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", client.UserAgent)

	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("error: unexpected status %s", resp.Status)
	}

	var buffer bytes.Buffer
	buffer.ReadFrom(resp.Body)
	defer resp.Body.Close()

	var authResponse authResponse
	if err := json.Unmarshal(buffer.Bytes(), &authResponse); err != nil {
		return "", err
	}

	log.Println("authentication successful")
	return authResponse.AccessToken, nil
}
