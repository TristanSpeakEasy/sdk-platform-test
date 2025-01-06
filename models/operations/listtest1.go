// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/tristanspeakeasy/sdk-platform-test/internal/utils"
	"github.com/tristanspeakeasy/sdk-platform-test/models/components"
)

const (
	// The only server.
	ListTest1ServerMain string = "Main"
)

var ListTest1ServerList = map[string]string{
	// The only server.
	ListTest1ServerMain: "http://localhost:35123",
}

type ListTest1Globals struct {
	// A long winded, multi-line description
	// for the query parameter number one.
	//
	QueryParam1 *string `queryParam:"style=form,explode=true,name=queryParam1"`
}

func (o *ListTest1Globals) GetQueryParam1() *string {
	if o == nil {
		return nil
	}
	return o.QueryParam1
}

type ListTest1Security struct {
	APIKey string `security:"scheme,type=apiKey,subtype=header,name=api_key,env=sdk_api_key"`
}

func (o *ListTest1Security) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

// QueryParam2 - A enum query parameter.
type QueryParam2 int64

const (
	QueryParam2Zero QueryParam2 = 0
	QueryParam2One  QueryParam2 = 1
	QueryParam2Two  QueryParam2 = 2
)

func (e QueryParam2) ToPointer() *QueryParam2 {
	return &e
}
func (e *QueryParam2) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 0:
		fallthrough
	case 1:
		fallthrough
	case 2:
		*e = QueryParam2(v)
		return nil
	default:
		return fmt.Errorf("invalid value for QueryParam2: %v", v)
	}
}

type ListTest1Request struct {
	// A enum query parameter.
	QueryParam2  QueryParam2 `queryParam:"style=form,explode=true,name=queryParam2"`
	Page         int64       `pathParam:"style=simple,explode=false,name=page"`
	HeaderParam1 string      `header:"style=simple,explode=false,name=headerParam1"`
	QueryParam1  *string     `queryParam:"style=form,explode=true,name=queryParam1"`
}

func (o *ListTest1Request) GetQueryParam2() QueryParam2 {
	if o == nil {
		return QueryParam2(0)
	}
	return o.QueryParam2
}

func (o *ListTest1Request) GetPage() int64 {
	if o == nil {
		return 0
	}
	return o.Page
}

func (o *ListTest1Request) GetHeaderParam1() string {
	if o == nil {
		return ""
	}
	return o.HeaderParam1
}

func (o *ListTest1Request) GetQueryParam1() *string {
	if o == nil {
		return nil
	}
	return o.QueryParam1
}

type ResultArray struct {
	Test1 string `json:"test1"`
}

func (o *ResultArray) GetTest1() string {
	if o == nil {
		return ""
	}
	return o.Test1
}

// ListTest1ResponseBody - OK
type ListTest1ResponseBody struct {
	ResultArray []ResultArray    `json:"resultArray"`
	TotalCount  int64            `json:"totalCount"`
	Type        *components.Enum `default:"First" json:"type"`
}

func (l ListTest1ResponseBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListTest1ResponseBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListTest1ResponseBody) GetResultArray() []ResultArray {
	if o == nil {
		return []ResultArray{}
	}
	return o.ResultArray
}

func (o *ListTest1ResponseBody) GetTotalCount() int64 {
	if o == nil {
		return 0
	}
	return o.TotalCount
}

func (o *ListTest1ResponseBody) GetType() *components.Enum {
	if o == nil {
		return nil
	}
	return o.Type
}

type ListTest1Response struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	Object *ListTest1ResponseBody

	Next func() (*ListTest1Response, error)
}

func (o *ListTest1Response) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListTest1Response) GetObject() *ListTest1ResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
