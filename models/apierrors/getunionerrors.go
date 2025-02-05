// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package apierrors

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tristanspeakeasy/sdk-platform-test/internal/utils"
	"github.com/tristanspeakeasy/sdk-platform-test/models/components"
)

type GetUnionErrorsResponseBodyType string

const (
	GetUnionErrorsResponseBodyTypeTag1 GetUnionErrorsResponseBodyType = "tag1"
	GetUnionErrorsResponseBodyTypeTag2 GetUnionErrorsResponseBodyType = "tag2"
)

// GetUnionErrorsResponseBody - Something went wrong
type GetUnionErrorsResponseBody struct {
	TaggedError1 *components.TaggedError1 `queryParam:"inline"`
	TaggedError2 *components.TaggedError2 `queryParam:"inline"`

	Type GetUnionErrorsResponseBodyType

	HTTPMeta components.HTTPMetadata `json:"-"`
}

var _ error = &GetUnionErrorsResponseBody{}

func CreateGetUnionErrorsResponseBodyTag1(tag1 components.TaggedError1) GetUnionErrorsResponseBody {
	typ := GetUnionErrorsResponseBodyTypeTag1

	typStr := components.Tag(typ)
	tag1.Tag = typStr

	return GetUnionErrorsResponseBody{
		TaggedError1: &tag1,
		Type:         typ,
	}
}

func CreateGetUnionErrorsResponseBodyTag2(tag2 components.TaggedError2) GetUnionErrorsResponseBody {
	typ := GetUnionErrorsResponseBodyTypeTag2

	return GetUnionErrorsResponseBody{
		TaggedError2: &tag2,
		Type:         typ,
	}
}

func (u *GetUnionErrorsResponseBody) UnmarshalJSON(data []byte) error {

	type discriminator struct {
		Tag string `json:"tag"`
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.Tag {
	case "tag1":
		taggedError1 := new(components.TaggedError1)
		if err := utils.UnmarshalJSON(data, &taggedError1, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (Tag == tag1) type components.TaggedError1 within GetUnionErrorsResponseBody: %w", string(data), err)
		}

		u.TaggedError1 = taggedError1
		u.Type = GetUnionErrorsResponseBodyTypeTag1
		return nil
	case "tag2":
		taggedError2 := new(components.TaggedError2)
		if err := utils.UnmarshalJSON(data, &taggedError2, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (Tag == tag2) type components.TaggedError2 within GetUnionErrorsResponseBody: %w", string(data), err)
		}

		u.TaggedError2 = taggedError2
		u.Type = GetUnionErrorsResponseBodyTypeTag2
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for GetUnionErrorsResponseBody", string(data))
}

func (u GetUnionErrorsResponseBody) MarshalJSON() ([]byte, error) {
	if u.TaggedError1 != nil {
		return utils.MarshalJSON(u.TaggedError1, "", true)
	}

	if u.TaggedError2 != nil {
		return utils.MarshalJSON(u.TaggedError2, "", true)
	}

	return nil, errors.New("could not marshal union type GetUnionErrorsResponseBody: all fields are null")
}

func (u GetUnionErrorsResponseBody) Error() string {
	switch u.Type {
	case GetUnionErrorsResponseBodyTypeTag1:
		data, _ := json.Marshal(u.TaggedError1)
		return string(data)
	case GetUnionErrorsResponseBodyTypeTag2:
		data, _ := json.Marshal(u.TaggedError2)
		return string(data)
	default:
		return "unknown error"
	}
}

type GetUnionErrorsResponseResponseBodyType string

const (
	GetUnionErrorsResponseResponseBodyTypeErrorType1 GetUnionErrorsResponseResponseBodyType = "ErrorType1"
	GetUnionErrorsResponseResponseBodyTypeErrorType2 GetUnionErrorsResponseResponseBodyType = "ErrorType2"
)

// GetUnionErrorsResponseResponseBody - Internal Server Error
type GetUnionErrorsResponseResponseBody struct {
	ErrorType1 *components.ErrorType1 `queryParam:"inline"`
	ErrorType2 *components.ErrorType2 `queryParam:"inline"`

	Type GetUnionErrorsResponseResponseBodyType

	HTTPMeta components.HTTPMetadata `json:"-"`
}

var _ error = &GetUnionErrorsResponseResponseBody{}

func CreateGetUnionErrorsResponseResponseBodyErrorType1(errorType1 components.ErrorType1) GetUnionErrorsResponseResponseBody {
	typ := GetUnionErrorsResponseResponseBodyTypeErrorType1

	return GetUnionErrorsResponseResponseBody{
		ErrorType1: &errorType1,
		Type:       typ,
	}
}

func CreateGetUnionErrorsResponseResponseBodyErrorType2(errorType2 components.ErrorType2) GetUnionErrorsResponseResponseBody {
	typ := GetUnionErrorsResponseResponseBodyTypeErrorType2

	return GetUnionErrorsResponseResponseBody{
		ErrorType2: &errorType2,
		Type:       typ,
	}
}

func (u *GetUnionErrorsResponseResponseBody) UnmarshalJSON(data []byte) error {

	var errorType2 components.ErrorType2 = components.ErrorType2{}
	if err := utils.UnmarshalJSON(data, &errorType2, "", true, true); err == nil {
		u.ErrorType2 = &errorType2
		u.Type = GetUnionErrorsResponseResponseBodyTypeErrorType2
		return nil
	}

	var errorType1 components.ErrorType1 = components.ErrorType1{}
	if err := utils.UnmarshalJSON(data, &errorType1, "", true, true); err == nil {
		u.ErrorType1 = &errorType1
		u.Type = GetUnionErrorsResponseResponseBodyTypeErrorType1
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for GetUnionErrorsResponseResponseBody", string(data))
}

func (u GetUnionErrorsResponseResponseBody) MarshalJSON() ([]byte, error) {
	if u.ErrorType1 != nil {
		return utils.MarshalJSON(u.ErrorType1, "", true)
	}

	if u.ErrorType2 != nil {
		return utils.MarshalJSON(u.ErrorType2, "", true)
	}

	return nil, errors.New("could not marshal union type GetUnionErrorsResponseResponseBody: all fields are null")
}

func (u GetUnionErrorsResponseResponseBody) Error() string {
	switch u.Type {
	case GetUnionErrorsResponseResponseBodyTypeErrorType1:
		data, _ := json.Marshal(u.ErrorType1)
		return string(data)
	case GetUnionErrorsResponseResponseBodyTypeErrorType2:
		data, _ := json.Marshal(u.ErrorType2)
		return string(data)
	default:
		return "unknown error"
	}
}
