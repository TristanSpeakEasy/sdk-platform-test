// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type Type string

const (
	TypeSuperType1 Type = "type1"
	TypeSuperType2 Type = "type2"
)

func (e Type) ToPointer() *Type {
	return &e
}
func (e *Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "type1":
		fallthrough
	case "type2":
		*e = Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Type: %v", v)
	}
}

type Test2Request struct {
	// A simple object that uses all our supported primitive types and enums and has optional properties.
	Obj  ExhaustiveObject `json:"obj"`
	Type *Type            `json:"type,omitempty"`
}

func (o *Test2Request) GetObj() ExhaustiveObject {
	if o == nil {
		return ExhaustiveObject{}
	}
	return o.Obj
}

func (o *Test2Request) GetType() *Type {
	if o == nil {
		return nil
	}
	return o.Type
}
