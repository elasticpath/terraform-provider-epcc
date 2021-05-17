package epcc

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"time"

	"gopkg.in/retry.v1"
)

// Client is the type used to interface with EPCC API.
type Client struct {
	BaseURL           string
	BetaFeatures      string
	HTTPClient        *http.Client
	RetryStrategy     retry.Strategy
	accessToken       string
	Credentials       *Credentials
	UserAgent         string
	AdditionalHeaders *map[string]string
}

// ClientOptions can be used to configure a new client.
type ClientOptions struct {
	BaseURL           string // BaseURL is the where requests will be made to.
	BetaFeatures      string
	ClientTimeout     time.Duration // ClientTimeout is how long the client waits for a response before timing out.
	RetryLimitTimeout time.Duration // RetryLimitTimeout is how long requests will be retried for status codes 429, 500, 503 & 504
	Credentials       *Credentials
	UserAgent         string
	AdditionalHeaders *map[string]string
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
		UserAgent:         "go-epcc-client",
		AdditionalHeaders: &map[string]string{},
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
				UserAgent:         options[i].UserAgent,
				AdditionalHeaders: options[i].AdditionalHeaders,
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
func (c *Client) DoRequest(ctx *context.Context, method string, path string, query string, payload io.Reader) (body []byte, error ApiErrors) {
	var teeBuf bytes.Buffer
	tee := io.TeeReader(payload, &teeBuf)
	var requestBody = "(n/a)"
	if payload != nil {
		requestBodyBytes, _ := ioutil.ReadAll(tee)
		requestBody = string(requestBodyBytes)
	}
	return c.doRequestInternal(ctx, method, "application/json", path, query, bytes.NewReader(teeBuf.Bytes()), requestBody)
}

func (c *Client) DoFileRequest(ctx *context.Context, path string, payload io.Reader, contentType string) (body []byte, error ApiErrors) {
	return c.doRequestInternal(ctx, "POST", contentType, path, "", payload, "(multipart data)")
}

func (c *Client) logToDisk(method string, path string, requestBytes []byte, responseBytes []byte) {

}
func (c *Client) logErrorToDiag(ctx *context.Context, resp *http.Response, method string, path string, requestBody string) {

	var buffer bytes.Buffer
	_, _ = buffer.ReadFrom(resp.Body)
	diagnostics := (*ctx).Value("diags").(*diag.Diagnostics)
	diagnosticsAppended := append(*diagnostics, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "HTTP Request Failure",
		Detail: fmt.Sprintf("Request Method:%s\nRequest Path:%s\nRequest Body:%s\nResponse Status:%s\nResponse Body:%s",
			method, path, requestBody, strconv.Itoa(resp.StatusCode), buffer.String())})

	*diagnostics = diagnosticsAppended

}

// DoRequest makes a html request to the EPCC API and handles the response.
func (c *Client) doRequestInternal(ctx *context.Context, method string, contentType string, path string, query string, payload io.Reader, requestBody string) (body []byte, error ApiErrors) {
	reqURL, err := url.Parse(c.BaseURL)
	if err != nil {
		return nil, FromError(err)
	}

	reqURL.Path = path
	reqURL.RawQuery = query

	req, err := http.NewRequest(method, reqURL.String(), payload)
	if err != nil {
		return nil, FromError(err)
	}

	if len(c.accessToken) > 0 {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.accessToken))
	}

	req.Header.Add("Content-Type", contentType)
	req.Header.Add("User-Agent", c.UserAgent)

	for header, value := range *c.AdditionalHeaders {
		req.Header.Add(header, value)
	}

	if len(c.BetaFeatures) > 0 {
		req.Header.Add("EP-Beta-Features", c.BetaFeatures)
	}

	for r := retry.Start(c.RetryStrategy, nil); r.Next(); {

		reqDump, _ := httputil.DumpRequestOut(req, true)

		resp, err := c.HTTPClient.Do(req)

		respDump, _ := httputil.DumpResponse(resp, true)

		c.logToDisk(method, path, reqDump, respDump)
		if err != nil {
			return nil, FromError(err)
		}
		defer resp.Body.Close()

		switch resp.StatusCode {
		case 429, 500, 502, 503, 504:
			c.logErrorToDiag(ctx, resp, method, path, requestBody)
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

			var jsonApiError ErrorList

			c.logErrorToDiag(ctx, resp, method, path, requestBody)

			if err := json.Unmarshal(buffer.Bytes(), &jsonApiError); err != nil {
				return nil, FromError(err)
			}

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

// https://stackoverflow.com/questions/20205796/post-data-using-the-content-type-multipart-form-data
func EncodeForm(values map[string]string, filename string, paramName string, fileContents []byte) (byteBuf *bytes.Buffer, contentType string, err error) {

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	for key, val := range values {
		_ = writer.WriteField(key, val)
	}

	if len(paramName) > 0 {
		part, err := writer.CreateFormFile(paramName, filename)

		if err != nil {
			return nil, "", err
		}

		part.Write(fileContents)
	}

	err = writer.Close()
	if err != nil {
		return nil, "", err
	}

	return body, writer.FormDataContentType(), nil
}
