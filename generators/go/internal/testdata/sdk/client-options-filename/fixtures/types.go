// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/fern-api/fern-go/internal/testdata/sdk/client-options-filename/fixtures/core"
)

type ClientOptions struct {
	Value string `json:"value" url:"value"`

	_rawJSON json.RawMessage
}

func (c *ClientOptions) UnmarshalJSON(data []byte) error {
	type unmarshaler ClientOptions
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ClientOptions(value)
	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *ClientOptions) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}