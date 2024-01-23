// This file was auto-generated by Fern from our API Definition.

package user

import (
	context "context"
	base64 "encoding/base64"
	fmt "fmt"
	fixtures "github.com/fern-api/fern-go/internal/testdata/sdk/query-params-complex/fixtures"
	core "github.com/fern-api/fern-go/internal/testdata/sdk/query-params-complex/fixtures/core"
	option "github.com/fern-api/fern-go/internal/testdata/sdk/query-params-complex/fixtures/option"
	http "net/http"
	url "net/url"
	time "time"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	return &Client{
		baseURL: options.BaseURL,
		caller:  core.NewCaller(options.HTTPClient),
		header:  options.ToHeader(),
	}
}

func (c *Client) GetUsername(
	ctx context.Context,
	request *fixtures.GetUsersRequest,
	opts ...option.RequestOption,
) (*fixtures.User, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := ""
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/" + "user"

	queryParams := make(url.Values)
	queryParams.Add("id", fmt.Sprintf("%v", request.Id))
	queryParams.Add("date", fmt.Sprintf("%v", request.Date.Format("2006-01-02")))
	queryParams.Add("deadline", fmt.Sprintf("%v", request.Deadline.Format(time.RFC3339)))
	queryParams.Add("bytes", fmt.Sprintf("%v", base64.StdEncoding.EncodeToString(request.Bytes)))
	if request.OptionalId != nil {
		queryParams.Add("optionalId", fmt.Sprintf("%v", *request.OptionalId))
	}
	if request.OptionalDate != nil {
		queryParams.Add("optionalDate", fmt.Sprintf("%v", request.OptionalDate.Format("2006-01-02")))
	}
	if request.OptionalDeadline != nil {
		queryParams.Add("optionalDeadline", fmt.Sprintf("%v", request.OptionalDeadline.Format(time.RFC3339)))
	}
	if request.OptionalBytes != nil {
		queryParams.Add("optionalBytes", fmt.Sprintf("%v", base64.StdEncoding.EncodeToString(*request.OptionalBytes)))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	headers := c.header.Clone()
	for key, values := range options.HTTPHeader {
		for _, value := range values {
			headers.Add(key, value)
		}
	}

	var response *fixtures.User
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
