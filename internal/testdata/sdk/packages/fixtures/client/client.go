// This file was auto-generated by Fern from our API Definition.

package client

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	errors "errors"
	fixtures "github.com/fern-api/fern-go/internal/testdata/sdk/packages/fixtures"
	configclient "github.com/fern-api/fern-go/internal/testdata/sdk/packages/fixtures/config/client"
	core "github.com/fern-api/fern-go/internal/testdata/sdk/packages/fixtures/core"
	organizationclient "github.com/fern-api/fern-go/internal/testdata/sdk/packages/fixtures/organization/client"
	userclient "github.com/fern-api/fern-go/internal/testdata/sdk/packages/fixtures/user/client"
	io "io"
	http "net/http"
)

type Client struct {
	baseURL    string
	httpClient core.HTTPClient
	header     http.Header

	User         *userclient.Client
	Config       *configclient.Client
	Organization *organizationclient.Client
}

func NewClient(opts ...core.ClientOption) *Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &Client{
		baseURL:      options.BaseURL,
		httpClient:   options.HTTPClient,
		header:       options.ToHeader(),
		User:         userclient.NewClient(opts...),
		Config:       configclient.NewClient(opts...),
		Organization: organizationclient.NewClient(opts...),
	}
}

func (c *Client) GetFoo(ctx context.Context) ([]*fixtures.Foo, error) {
	baseURL := "https://api.foo.io/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "foo"

	var response []*fixtures.Foo
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
		nil,
		&response,
		false,
		c.header,
		nil,
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) PostFoo(ctx context.Context, request *fixtures.Foo) (*fixtures.Foo, error) {
	baseURL := "https://api.foo.io/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "foo"

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 409:
			value := new(fixtures.ConflictError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		case 422:
			value := new(fixtures.UnprocessableEntityError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *fixtures.Foo
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		&response,
		false,
		c.header,
		errorDecoder,
	); err != nil {
		return nil, err
	}
	return response, nil
}
