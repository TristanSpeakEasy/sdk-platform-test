// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/tristanspeakeasy/sdk-platform-test/models/components"
)

type GetUserRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetUserRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetUserResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	User *components.User
}

func (o *GetUserResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetUserResponse) GetUser() *components.User {
	if o == nil {
		return nil
	}
	return o.User
}
