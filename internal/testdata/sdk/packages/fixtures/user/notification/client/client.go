// This file was auto-generated by Fern from our API Definition.

package client

import (
	context "context"
	fmt "fmt"
	core "github.com/fern-api/fern-go/internal/testdata/sdk/packages/fixtures/core"
	option "github.com/fern-api/fern-go/internal/testdata/sdk/packages/fixtures/option"
	notification "github.com/fern-api/fern-go/internal/testdata/sdk/packages/fixtures/user/notification"
	notificationnotification "github.com/fern-api/fern-go/internal/testdata/sdk/packages/fixtures/user/notification/notification"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header

	Notification *notificationnotification.Client
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	return &Client{
		baseURL:      options.BaseURL,
		caller:       core.NewCaller(options.HTTPClient),
		header:       options.ToHeader(),
		Notification: notificationnotification.NewClient(opts...),
	}
}

func (c *Client) GetUserNotification(
	ctx context.Context,
	userId string,
	notificationId string,
	opts ...option.RequestOption,
) (*notification.Notification, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.foo.io/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"users/%v/notifications/%v", userId, notificationId)

	headers := c.header.Clone()
	for key, values := range options.HTTPHeader {
		for _, value := range values {
			headers.Add(key, value)
		}
	}

	var response *notification.Notification
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
