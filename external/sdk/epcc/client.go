package epcc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"gopkg.in/retry.v1"
)

// Client is the type used to interface with EPCC API.
type Client struct {
	BaseURL       string
	BetaFeatures  string
	HTTPClient    *http.Client
	RetryStrategy retry.Strategy
	accessToken   string
	Credentials   *Credentials
	UserAgent     string
}

// ClientOptions can be used to configure a new client.
type ClientOptions struct {
	BaseURL           string // BaseURL is the where requests will be made to.
	BetaFeatures      string
	ClientTimeout     time.Duration // ClientTimeout is how long the client waits for a response before timing out.
	RetryLimitTimeout time.Duration // RetryLimitTimeout is how long requests will be retried for status codes 429, 500, 503 & 504
	Credentials       *Credentials
	UserAgent         string
}

type Credentials struct {
	ClientId     string
	ClientSecret string
}

// NewClient creates a new instance of a Client.
func NewClient(options ...ClientOptions) *Client {
	exp := retry.Exponential{
		Initial: 10 * time.Millisecond,
		Factor:  1.5,
		Jitter:  true,
	}

	strategy := retry.LimitTime(cfg.RetryLimitTimeout, exp)

	defaultClient := Client{
		BaseURL:      cfg.BaseURL,
		BetaFeatures: cfg.BetaFeatures,
		HTTPClient: &http.Client{
			Timeout: cfg.ClientTimeout,
		},
		RetryStrategy: strategy,
		Credentials: &Credentials{
			ClientId:     cfg.Credentials.ClientID,
			ClientSecret: cfg.Credentials.ClientSecret,
		},
		UserAgent: "go-epcc-client",
	}

	// If no configuration options are provided, return the default client.
	if len(options) == 0 {
		return &defaultClient
	}

	// Otherwise configure a client with custom options.
	for i := range options {
		if i == 0 {
			strategy := retry.LimitTime(options[i].RetryLimitTimeout, exp)
			customClient := Client{
				BaseURL:      options[i].BaseURL,
				BetaFeatures: options[i].BetaFeatures,
				HTTPClient: &http.Client{
					Timeout: options[i].ClientTimeout,
				},
				RetryStrategy: strategy,
				Credentials: &Credentials{
					ClientId:     options[i].Credentials.ClientId,
					ClientSecret: options[i].Credentials.ClientSecret,
				},
				UserAgent: options[i].UserAgent,
			}

			if len(customClient.BaseURL) == 0 {
				customClient.BaseURL = defaultClient.BaseURL
			}

			if len(customClient.BetaFeatures) == 0 {
				customClient.BetaFeatures = defaultClient.BetaFeatures
			}

			if len(customClient.Credentials.ClientId) == 0 {
				customClient.Credentials.ClientId = defaultClient.Credentials.ClientId
			}

			if len(customClient.Credentials.ClientSecret) == 0 {
				customClient.Credentials.ClientSecret = defaultClient.Credentials.ClientSecret
			}

			return &customClient
		}
	}

	return nil
}

//Authenticate attempts to generate an access token and save it on the client.
func (c *Client) Authenticate() error {
	token, err := auth(*c)
	if err != nil {
		return err
	}

	c.accessToken = token
	return nil
}

// DoRequest makes a html request to the EPCC API and handles the response.
func (c *Client) DoRequest(method string, path string, payload io.Reader) (body []byte, error ApiErrors) {
	reqURL, err := url.Parse(c.BaseURL)
	if err != nil {
		return nil, FromError(err)
	}

	reqURL.Path = path

	req, err := http.NewRequest(method, reqURL.String(), payload)
	if err != nil {
		return nil, FromError(err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", c.UserAgent)

	if len(c.BetaFeatures) > 0 {
		req.Header.Add("EP-Beta-Features", c.BetaFeatures)
	}

	for r := retry.Start(c.RetryStrategy, nil); r.Next(); {
		resp, err := c.HTTPClient.Do(req)

		if err != nil {
			return nil, FromError(err)
		}
		defer resp.Body.Close()

		switch resp.StatusCode {
		case 429, 500, 503, 504:
			log.Printf("Response Status %d Retrying request", resp.StatusCode)
			continue

		case 200, 201:
			var buffer bytes.Buffer
			if _, err := buffer.ReadFrom(resp.Body); err != nil {
				return nil, FromError(err)
			}

			return buffer.Bytes(), nil

		case 204:
			return nil, nil

		default:
			var buffer bytes.Buffer
			if _, err := buffer.ReadFrom(resp.Body); err != nil {
				return nil, FromError(err)
			}

			// TODO Better Manage Parent ID

			var jsonApiError ErrorList

			bytes := buffer.Bytes()

			if err := json.Unmarshal(bytes, &jsonApiError); err != nil {
				return nil, FromError(err)
			}

			log.Printf("response: %s", buffer.String())
			err := fmt.Errorf("status code %d is not ok", resp.StatusCode)

			return nil, &ApiErrorResult{
				errorString:    err.Error(),
				apiErrors:      &jsonApiError,
				httpMethod:     method,
				httpPath:       path,
				httpStatusCode: uint16(resp.StatusCode),
			}
		}
	}

	err = errors.New("retry timeout error")
	return nil, FromError(err)
}
