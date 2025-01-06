# Tag1
(*Tag1*)

## Overview

The first tag.

### Available Operations

* [~~Deprecated1~~](#deprecated1) - Deprecated Operation :warning: **Deprecated** Use [GetRequestBodyFlattenedAway](docs/sdks/sdk/README.md#getrequestbodyflattenedaway) instead.
* [ListTest1](#listtest1) - Get Test1

## ~~Deprecated1~~

Deprecated Operation

> :warning: **DEPRECATED**: This endpoint is deprecated.. Use `GetRequestBodyFlattenedAway` instead.

### Example Usage

```go
package main

import(
	"context"
	sdkplatformtest "github.com/tristanspeakeasy/sdk-platform-test"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkplatformtest.New(
        sdkplatformtest.WithQueryParam1("some example query param"),
    )

    res, err := s.Tag1.Deprecated1(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.Deprecated1Response](../../models/operations/deprecated1response.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListTest1

This is a {{test}} endpoint.
It has a description.

### Example Usage

```go
package main

import(
	"context"
	sdkplatformtest "github.com/tristanspeakeasy/sdk-platform-test"
	"github.com/tristanspeakeasy/sdk-platform-test/models/operations"
	"os"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkplatformtest.New(
        sdkplatformtest.WithQueryParam1("some example query param"),
    )

    res, err := s.Tag1.ListTest1(ctx, operations.ListTest1Security{
        APIKey: os.Getenv("SDK_API_KEY"),
    }, operations.QueryParam2One, 100, "some example header param", sdkplatformtest.String("some example query param"))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        for {
            // handle items

            res, err = res.Next()

            if err != nil {
                // handle error
            }

            if res == nil {
                break
            }
        }
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  | Example                                                                      |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |                                                                              |
| `security`                                                                   | [operations.ListTest1Security](../../models/operations/listtest1security.md) | :heavy_check_mark:                                                           | The security requirements to use for the request.                            |                                                                              |
| `queryParam2`                                                                | [operations.QueryParam2](../../models/operations/queryparam2.md)             | :heavy_check_mark:                                                           | A enum query parameter.                                                      | 1                                                                            |
| `page`                                                                       | *int64*                                                                      | :heavy_check_mark:                                                           | N/A                                                                          | 100                                                                          |
| `headerParam1`                                                               | *string*                                                                     | :heavy_check_mark:                                                           | N/A                                                                          | some example header param                                                    |
| `queryParam1`                                                                | **string*                                                                    | :heavy_minus_sign:                                                           | N/A                                                                          | some example query param                                                     |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |                                                                              |

### Response

**[*operations.ListTest1Response](../../models/operations/listtest1response.md), error**

### Errors

| Error Type                   | Status Code                  | Content Type                 |
| ---------------------------- | ---------------------------- | ---------------------------- |
| apierrors.BadRequestResponse | 400                          | application/json             |
| apierrors.Error              | 500                          | application/json             |
| apierrors.APIError           | 4XX, 5XX                     | \*/\*                        |