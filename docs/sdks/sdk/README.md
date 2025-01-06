# SDK

## Overview

This document will show case as many of our features as possible in as little operations/models as possible.
This will then generate a SDK that we can more easily review than the test SDKs based on uber.yaml spec.

Speakeasy Docs
<https://speakeasy.com/docs>

### Available Operations

* [OperationWithLeadingAndTrailingUnderscores](#operationwithleadingandtrailingunderscores)
* [PostFile](#postfile) - Post File
* [GetPolymorphism](#getpolymorphism)
* [GetUnionErrors](#getunionerrors)
* [GetRequestBodyFlattenedAway](#getrequestbodyflattenedaway)
* [GetFullyFlattenedRequest](#getfullyflattenedrequest)
* [TestEndpoint](#testendpoint)
* [CreateUser](#createuser) - Create User
* [GetUser](#getuser) - Get User
* [UpdateUser](#updateuser) - Update User
* [DeleteUser](#deleteuser) - Delete User
* [Chat](#chat)
* [GetBinaryDefaultResponse](#getbinarydefaultresponse)

## OperationWithLeadingAndTrailingUnderscores

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

    res, err := s.OperationWithLeadingAndTrailingUnderscores(ctx)
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

**[*operations.OperationWithLeadingAndTrailingUnderscoresResponse](../../models/operations/operationwithleadingandtrailingunderscoresresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## PostFile

This is a test endpoint.
It has a description.

### Example Usage

```go
package main

import(
	"context"
	sdkplatformtest "github.com/tristanspeakeasy/sdk-platform-test"
	"os"
	"github.com/tristanspeakeasy/sdk-platform-test/models/components"
	"github.com/tristanspeakeasy/sdk-platform-test/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkplatformtest.New(
        sdkplatformtest.WithQueryParam1("some example query param"),
    )

    content, fileErr := os.Open("example.file")
    if fileErr != nil {
        panic(fileErr)
    }


    res, err := s.PostFile(ctx, operations.PostFileRequestBody{
        File: components.File{
            FileName: "example.file",
            Content: content,
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.File != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.PostFileRequestBody](../../models/operations/postfilerequestbody.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |

### Response

**[*operations.PostFileResponse](../../models/operations/postfileresponse.md), error**

### Errors

| Error Type       | Status Code      | Content Type     |
| ---------------- | ---------------- | ---------------- |
| apierrors.Error  | 415, 4XX, 5XX    | application/json |

## GetPolymorphism

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

    res, err := s.GetPolymorphism(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
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

**[*operations.GetPolymorphismResponse](../../models/operations/getpolymorphismresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetUnionErrors

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

    res, err := s.GetUnionErrors(ctx, 12)
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

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `page`                                                   | *int64*                                                  | :heavy_check_mark:                                       | N/A                                                      | 12                                                       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetUnionErrorsResponse](../../models/operations/getunionerrorsresponse.md), error**

### Errors

| Error Type                                   | Status Code                                  | Content Type                                 |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| apierrors.Error                              | 404                                          | application/json                             |
| apierrors.GetUnionErrorsResponseResponseBody | 500                                          | application/json                             |
| apierrors.GetUnionErrorsResponseBody         | 4XX                                          | application/json                             |
| apierrors.APIError                           | 5XX                                          | \*/\*                                        |

## GetRequestBodyFlattenedAway

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

    res, err := s.GetRequestBodyFlattenedAway(ctx)
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

**[*operations.GetRequestBodyFlattenedAwayResponse](../../models/operations/getrequestbodyflattenedawayresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetFullyFlattenedRequest

### Example Usage

```go
package main

import(
	"context"
	sdkplatformtest "github.com/tristanspeakeasy/sdk-platform-test"
	"github.com/tristanspeakeasy/sdk-platform-test/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkplatformtest.New(
        sdkplatformtest.WithQueryParam1("some example query param"),
    )

    res, err := s.GetFullyFlattenedRequest(ctx, "en", operations.GetFullyFlattenedRequestRequestBody{
        Name: "<value>",
    }, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `security`                                                                                                       | [operations.GetFullyFlattenedRequestSecurity](../../models/operations/getfullyflattenedrequestsecurity.md)       | :heavy_check_mark:                                                                                               | The security requirements to use for the request.                                                                |
| `lang`                                                                                                           | *string*                                                                                                         | :heavy_check_mark:                                                                                               | N/A                                                                                                              |
| `requestBody`                                                                                                    | [operations.GetFullyFlattenedRequestRequestBody](../../models/operations/getfullyflattenedrequestrequestbody.md) | :heavy_check_mark:                                                                                               | N/A                                                                                                              |
| `maxLength`                                                                                                      | **int64*                                                                                                         | :heavy_minus_sign:                                                                                               | N/A                                                                                                              |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.GetFullyFlattenedRequestResponse](../../models/operations/getfullyflattenedrequestresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## TestEndpoint

### Example Usage

```go
package main

import(
	"context"
	sdkplatformtest "github.com/tristanspeakeasy/sdk-platform-test"
	"github.com/tristanspeakeasy/sdk-platform-test/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkplatformtest.New(
        sdkplatformtest.WithQueryParam1("some example query param"),
    )

    res, err := s.TestEndpoint(ctx, "<value>", operations.TestEndpointRequestBody{
        Test: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `testName`                                                                               | *string*                                                                                 | :heavy_check_mark:                                                                       | N/A                                                                                      |
| `requestBody`                                                                            | [operations.TestEndpointRequestBody](../../models/operations/testendpointrequestbody.md) | :heavy_check_mark:                                                                       | N/A                                                                                      |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.TestEndpointResponse](../../models/operations/testendpointresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateUser

Create User

### Example Usage

```go
package main

import(
	"context"
	sdkplatformtest "github.com/tristanspeakeasy/sdk-platform-test"
	"github.com/tristanspeakeasy/sdk-platform-test/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkplatformtest.New(
        sdkplatformtest.WithQueryParam1("some example query param"),
    )

    res, err := s.CreateUser(ctx, components.BaseUser{
        ID: sdkplatformtest.String("8ffac18c-7d88-4879-b057-e5f45b9ce7de"),
        Email: "Creola.Kutch71@yahoo.com",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |
| `request`                                                  | [components.BaseUser](../../models/components/baseuser.md) | :heavy_check_mark:                                         | The request object to use for the request.                 |
| `opts`                                                     | [][operations.Option](../../models/operations/option.md)   | :heavy_minus_sign:                                         | The options for this request.                              |

### Response

**[*operations.CreateUserResponse](../../models/operations/createuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetUser

Get User

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

    res, err := s.GetUser(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetUserResponse](../../models/operations/getuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateUser

Update User

### Example Usage

```go
package main

import(
	"context"
	sdkplatformtest "github.com/tristanspeakeasy/sdk-platform-test"
	"github.com/tristanspeakeasy/sdk-platform-test/models/components"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkplatformtest.New(
        sdkplatformtest.WithQueryParam1("some example query param"),
    )

    res, err := s.UpdateUser(ctx, "<id>", components.User{
        ID: "8ffac18c-7d88-4879-b057-e5f45b9ce7de",
        Email: "Micah.Stracke@gmail.com",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.User != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `user`                                                   | [components.User](../../models/components/user.md)       | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.UpdateUserResponse](../../models/operations/updateuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteUser

Delete User

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

    res, err := s.DeleteUser(ctx, "<id>")
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
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteUserResponse](../../models/operations/deleteuserresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Chat

### Example Usage

```go
package main

import(
	"context"
	sdkplatformtest "github.com/tristanspeakeasy/sdk-platform-test"
	"github.com/tristanspeakeasy/sdk-platform-test/models/operations"
	"log"
)

func main() {
    ctx := context.Background()
    
    s := sdkplatformtest.New(
        sdkplatformtest.WithQueryParam1("some example query param"),
    )

    res, err := s.Chat(ctx, operations.ChatRequestBody{
        Prompt: sdkplatformtest.String("What is the largest city in the world?"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ChatStream != nil {
        defer res.ChatStream.Close()

        for res.ChatStream.Next() {
            event := res.ChatStream.Value()
            log.Print(event)
            // Handle the event
	      }
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.ChatRequestBody](../../models/operations/chatrequestbody.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*operations.ChatResponse](../../models/operations/chatresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetBinaryDefaultResponse

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

    res, err := s.GetBinaryDefaultResponse(ctx)
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

**[*operations.GetBinaryDefaultResponseResponse](../../models/operations/getbinarydefaultresponseresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |