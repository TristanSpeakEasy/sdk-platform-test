# Tag2
(*TestGroup.Tag2*)

## Overview

### Available Operations

* [PostTest](#posttest) - Post Test2

## PostTest

This is a test endpoint.
It has a description.

### Example Usage

<!-- UsageSnippet language="go" operationID="postTest2" method="post" path="/test2" -->
```go
package main

import(
	"context"
	sdkplatformtest "github.com/tristanspeakeasy/sdk-platform-test"
	"github.com/tristanspeakeasy/sdk-platform-test/types"
	"github.com/tristanspeakeasy/sdk-platform-test/models/components"
	"math/big"
	"log"
)

func main() {
    ctx := context.Background()

    s := sdkplatformtest.New(
        sdkplatformtest.WithDeprecatedQueryParam1("some example query param"),
        sdkplatformtest.WithDeprecatedQueryParam2("some example query param"),
    )

    res, err := s.TestGroup.Tag2.PostTest(ctx, components.Test2Request{
        Obj: components.ExhaustiveObject{
            Str: "example",
            Bool: true,
            Integer: 999999,
            Int32: 1,
            Num: 1.1,
            Float32: 8499.3,
            Date: types.MustDateFromString("2024-10-12"),
            DateTime: types.MustTimeFromString("2020-01-01T00:00:00Z"),
            Anything: "<value>",
            BoolOpt: sdkplatformtest.Pointer(true),
            IntOptNull: sdkplatformtest.Pointer[int64](999999),
            NumOptNull: sdkplatformtest.Pointer[float64](1.1),
            IntEnum: components.IntEnumThird.ToPointer(),
            Int32Enum: components.Int32EnumSixtyNine,
            Bigint: big.NewInt(702830),
            DecimalStr: types.MustNewDecimalFromString("3858.6"),
            Obj: components.SimpleObject{
                Str: "example",
            },
            Map: map[string]components.SimpleObject{

            },
            Arr: []components.SimpleObject{
                components.SimpleObject{
                    Str: "example",
                },
            },
            Any: components.CreateAnySimpleObject(
                components.SimpleObject{
                    Str: "example",
                },
            ),
            NullableIntEnum: components.NullableIntEnumThird.ToPointer(),
            NullableStringEnum: components.NullableStringEnumSecond,
            Color: components.ColorGreen.ToPointer(),
            Icon: components.IconTick,
            HeroWidth: components.HeroWidthFourHundredAndEighty.ToPointer(),
        },
        Type: components.TypeSuperType1.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Body != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                               | Type                                                                                                                    | Required                                                                                                                | Description                                                                                                             | Example                                                                                                                 |
| ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                                   | :heavy_check_mark:                                                                                                      | The context to use for the request.                                                                                     |                                                                                                                         |
| `test2Request`                                                                                                          | [components.Test2Request](../../models/components/test2request.md)                                                      | :heavy_check_mark:                                                                                                      | N/A                                                                                                                     |                                                                                                                         |
| `deprecatedQueryParam1`                                                                                                 | **string*                                                                                                               | :heavy_minus_sign:                                                                                                      | : warning: ** DEPRECATED **: This will be removed in a future release, please migrate away from it as soon as possible. | some example query param                                                                                                |
| `deprecatedQueryParam2`                                                                                                 | **string*                                                                                                               | :heavy_minus_sign:                                                                                                      | : warning: ** DEPRECATED **: This will be removed in a future release, please migrate away from it as soon as possible. | some example query param                                                                                                |
| `opts`                                                                                                                  | [][operations.Option](../../models/operations/option.md)                                                                | :heavy_minus_sign:                                                                                                      | The options for this request.                                                                                           |                                                                                                                         |

### Response

**[*operations.PostTest2Response](../../models/operations/posttest2response.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| apierrors.BadRequestResponse | 400                          | application/json             |
| apierrors.Error              | 404                          | application/json             |
| apierrors.Test2Response      | 500                          | application/json             |
| apierrors.APIError           | 4XX, 5XX                     | \*/\*                        |