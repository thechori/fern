// Generated by Fern. Do not edit.

package api

import (
	context "context"
	fmt "fmt"
	core "github.com/fern-api/fern-go/internal/testdata/sdk/path-params/fixtures/core"
	http "net/http"
	strings "strings"
)

type UserClient interface {
	GetUser(ctx context.Context, userId string) (string, error)
	GetUserV2(ctx context.Context, userId string) (string, error)
	GetUserV3(ctx context.Context, userId string, infoId string) (string, error)
}

func NewUserClient(baseURL string, httpClient core.HTTPClient, opts ...core.ClientOption) UserClient {
	options := new(core.ClientOptions)
	for _, opt := range opts {
		opt(options)
	}
	return &userClient{
		baseURL:    strings.TrimRight(baseURL, "/"),
		httpClient: httpClient,
		header:     options.ToHeader(),
	}
}

type userClient struct {
	baseURL    string
	httpClient core.HTTPClient
	header     http.Header
}

func (u *userClient) GetUser(ctx context.Context, userId string) (string, error) {
	endpointURL := fmt.Sprintf(u.baseURL+"/"+"/users/%v", userId)

	var response string
	if err := core.DoRequest(
		ctx,
		u.httpClient,
		endpointURL,
		http.MethodGet,
		nil,
		&response,
		u.header,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}

func (u *userClient) GetUserV2(ctx context.Context, userId string) (string, error) {
	endpointURL := fmt.Sprintf(u.baseURL+"/"+"/users/get/%v/info", userId)

	var response string
	if err := core.DoRequest(
		ctx,
		u.httpClient,
		endpointURL,
		http.MethodGet,
		nil,
		&response,
		u.header,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}

func (u *userClient) GetUserV3(ctx context.Context, userId string, infoId string) (string, error) {
	endpointURL := fmt.Sprintf(u.baseURL+"/"+"/users/get/%v/info/%v", userId, infoId)

	var response string
	if err := core.DoRequest(
		ctx,
		u.httpClient,
		endpointURL,
		http.MethodGet,
		nil,
		&response,
		u.header,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}
