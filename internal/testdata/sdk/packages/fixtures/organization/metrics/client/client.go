// This file was auto-generated by Fern from our API Definition.

package client

import (
	context "context"
	fmt "fmt"
	core "github.com/fern-api/fern-go/internal/testdata/sdk/packages/fixtures/core"
	option "github.com/fern-api/fern-go/internal/testdata/sdk/packages/fixtures/option"
	organization "github.com/fern-api/fern-go/internal/testdata/sdk/packages/fixtures/organization"
	metrics "github.com/fern-api/fern-go/internal/testdata/sdk/packages/fixtures/organization/metrics"
	tag "github.com/fern-api/fern-go/internal/testdata/sdk/packages/fixtures/organization/metrics/tag"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header

	Tag *tag.Client
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	return &Client{
		baseURL: options.BaseURL,
		caller:  core.NewCaller(options.HTTPClient),
		header:  options.ToHeader(),
		Tag:     tag.NewClient(opts...),
	}
}

func (c *Client) CreateMetricsTag(
	ctx context.Context,
	request *organization.CreateMetricsTagRequest,
	opts ...option.RequestOption,
) (*metrics.Tag, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.foo.io/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/" + "metrics"

	headers := c.header.Clone()
	for key, values := range options.HTTPHeader {
		for _, value := range values {
			headers.Add(key, value)
		}
	}

	var response *metrics.Tag
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodPost,
			Headers:  headers,
			Client:   options.HTTPClient,
			Request:  request,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

func (c *Client) GetMetricsTag(
	ctx context.Context,
	id string,
	opts ...option.RequestOption,
) (*metrics.Tag, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.foo.io/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"metrics/%v", id)

	headers := c.header.Clone()
	for key, values := range options.HTTPHeader {
		for _, value := range values {
			headers.Add(key, value)
		}
	}

	var response *metrics.Tag
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodGet,
			Headers:  headers,
			Client:   options.HTTPClient,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
