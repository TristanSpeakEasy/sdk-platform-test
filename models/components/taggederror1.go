// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type Tag string

const (
	TagTag1 Tag = "tag1"
)

func (e Tag) ToPointer() *Tag {
	return &e
}
func (e *Tag) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "tag1":
		*e = Tag(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Tag: %v", v)
	}
}

type TaggedError1 struct {
	Tag   Tag    `json:"tag"`
	Error string `json:"error"`
}

func (o *TaggedError1) GetTag() Tag {
	if o == nil {
		return Tag("")
	}
	return o.Tag
}

func (o *TaggedError1) GetError() string {
	if o == nil {
		return ""
	}
	return o.Error
}
