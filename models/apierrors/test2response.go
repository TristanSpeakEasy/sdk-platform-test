// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package apierrors

import (
	"encoding/json"
)

type Test2Response struct {
	Test2     *string `json:"test2,omitempty"`
	SomeConst *string `const:"someConstValue" json:"someConst,omitempty"`
}

var _ error = &Test2Response{}

func (e *Test2Response) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
