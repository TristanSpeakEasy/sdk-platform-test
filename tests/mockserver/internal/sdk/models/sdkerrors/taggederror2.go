// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"mockserver/internal/sdk/models/components"
)

type TaggedError2Error struct {
	Message string `json:"message"`
}

func (o *TaggedError2Error) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

type TaggedError2 struct {
	Tag      string                  `const:"tag2" json:"tag"`
	Error_   TaggedError2Error       `json:"error"`
	HTTPMeta components.HTTPMetadata `json:"-"`
}

var _ error = &TaggedError2{}

func (e *TaggedError2) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
