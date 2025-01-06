# Tag2
(*TestGroup.Tag2*)

## Overview

### Available Operations

* [PostTest](#posttest) - Post Test2

## PostTest

This is a test endpoint.
It has a description.

### Example Usage

```go
package main

import(
	"context"
	sdkplatformtest "github.com/tristanspeakeasy/sdk-platform-test"
	"github.com/tristanspeakeasy/sdk-platform-test/models/components"
	"github.com/tristanspeakeasy/sdk-platform-test/types"
	"math/big"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkplatformtest.New(
        sdkplatformtest.WithQueryParam1("some example query param"),
    )

    res, err := s.TestGroup.Tag2.PostTest(ctx, components.Test2Request{
        Obj: components.ExhaustiveObject{
            Str: "example",
            Bool: true,
            Integer: 999999,
            Int32: 1,
            Num: 1.1,
            Float32: 4344.96,
            EnumProp: components.EnumFirst.ToPointer(),
            Date: types.MustDateFromString("2024-02-10"),
            DateTime: types.MustTimeFromString("2020-01-01T00:00:00Z"),
            Anything: "<value>",
            BoolOpt: sdkplatformtest.Bool(true),
            IntOptNull: sdkplatformtest.Int64(999999),
            NumOptNull: sdkplatformtest.Float64(1.1),
            IntEnum: components.IntEnumThird.ToPointer(),
            Int32Enum: components.Int32EnumSixtyNine,
            Bigint: big.NewInt(119171),
            DecimalStr: types.MustNewDecimalFromString("4560.33"),
            Obj: components.SimpleObject{
                Str: "example",
            },
            Map: map[string]components.SimpleObject{
                "key": components.SimpleObject{
                    Str: "example",
                },
                "key1": components.SimpleObject{
                    Str: "example",
                },
                "key2": components.SimpleObject{
                    Str: "example",
                },
            },
            Arr: []components.SimpleObject{
                components.SimpleObject{
                    Str: "example",
                },
                components.SimpleObject{
                    Str: "example",
                },
                components.SimpleObject{
                    Str: "example",
                },
            },
            Any: components.CreateAnyStr(
                "<value>",
            ),
            NullableIntEnum: components.NullableIntEnumThird.ToPointer(),
            NullableStringEnum: components.NullableStringEnumThird,
            Color: components.ColorGreen.ToPointer(),
            Icon: components.IconTick,
            HeroWidth: components.HeroWidthFourHundredAndEighty.ToPointer(),
        },
        Type: components.TypeSuperType1.ToPointer(),
    }, nil, sdkplatformtest.String("some example query param"))
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