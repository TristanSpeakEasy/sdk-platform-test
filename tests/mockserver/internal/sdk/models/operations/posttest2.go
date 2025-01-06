// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type PostTest2Globals struct {
	// A deprecated description
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	DeprecatedQueryParam1 *string `queryParam:"style=form,explode=true,name=deprecatedQueryParam1"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	DeprecatedQueryParam2 *string `queryParam:"style=form,explode=true,name=deprecatedQueryParam2"`
}

func (o *PostTest2Globals) GetDeprecatedQueryParam1() *string {
	if o == nil {
		return nil
	}
	return o.DeprecatedQueryParam1
}

func (o *PostTest2Globals) GetDeprecatedQueryParam2() *string {
	if o == nil {
		return nil
	}
	return o.DeprecatedQueryParam2
}

type PostTest2Request struct {
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	DeprecatedQueryParam1 *string `queryParam:"style=form,explode=true,name=deprecatedQueryParam1"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	DeprecatedQueryParam2 *string                 `queryParam:"style=form,explode=true,name=deprecatedQueryParam2"`
	Test2Request          components.Test2Request `request:"mediaType=application/json"`
}

func (o *PostTest2Request) GetDeprecatedQueryParam1() *string {
	if o == nil {
		return nil
	}
	return o.DeprecatedQueryParam1
}

func (o *PostTest2Request) GetDeprecatedQueryParam2() *string {
	if o == nil {
		return nil
	}
	return o.DeprecatedQueryParam2
}

func (o *PostTest2Request) GetTest2Request() components.Test2Request {
	if o == nil {
		return components.Test2Request{}
	}
	return o.Test2Request
}

type PostTest2Response struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Body     []byte
}

func (o *PostTest2Response) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PostTest2Response) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}
