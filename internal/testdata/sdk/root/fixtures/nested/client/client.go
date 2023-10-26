// This file was auto-generated by Fern from our API Definition.

package client

import (
	context "context"
	core "github.com/fern-api/fern-go/internal/testdata/sdk/root/fixtures/core"
	nested "github.com/fern-api/fern-go/internal/testdata/sdk/root/fixtures/nested"
	http "net/http"
)

type Client struct {
	baseURL    string
	httpClient core.HTTPClient
	header     http.Header
}

func NewClient(opts ...core.ClientOption) *Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &Client{
		baseURL:    options.BaseURL,
		httpClient: options.HTTPClient,
		header:     options.ToHeader(),
	}
}

func (c *Client) CreateNested(ctx context.Context, request *nested.GetNestedRequest) (string, error) {
	baseURL := "https://api.foo.io/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "nested"

	var response string
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		&response,
		false,
		c.header,
		nil,
	); err != nil {
		return "", err
	}
	return response, nil
}
