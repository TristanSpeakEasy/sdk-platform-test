# github.com/tristanspeakeasy/sdk-platform-test

Developer-friendly & type-safe Go SDK specifically catered to leverage *github.com/tristanspeakeasy/sdk-platform-test* API.

<div align="left">
    <a href="https://www.speakeasy.com/?utm_source=github-com/tristanspeakeasy/sdk-platform-test&utm_campaign=go"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-blue.svg" style="width: 100px; height: 28px;" />
    </a>
</div>


<br /><br />
> [!IMPORTANT]
> This SDK is not yet ready for production use. To complete setup please follow the steps outlined in your [workspace](https://app.speakeasy.com/org/tristan-test/sdk-platform-test). Delete this section before > publishing to a package manager.

<!-- Start Summary [summary] -->
## Summary

SDK Review: A test document for reviewing the SDK.

This document will show case as many of our features as possible in as little operations/models as possible.
This will then generate a SDK that we can more easily review than the test SDKs based on uber.yaml spec.

For more information about the API: [Speakeasy Docs](https://speakeasy.com/docs)
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [github.com/tristanspeakeasy/sdk-platform-test](#githubcomtristanspeakeasysdk-platform-test)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Global Parameters](#global-parameters)
  * [Server-sent event streaming](#server-sent-event-streaming)
  * [Pagination](#pagination)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
  * [Special Types](#special-types)
* [Development](#development)
  * [Maturity](#maturity)
  * [Contributions](#contributions)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/tristanspeakeasy/sdk-platform-test
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example 1

```go
package main

import (
	"context"
	sdkplatformtest "github.com/tristanspeakeasy/sdk-platform-test"
	"github.com/tristanspeakeasy/sdk-platform-test/models/components"
	"github.com/tristanspeakeasy/sdk-platform-test/models/operations"
	"log"
	"os"
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
			Content:  content,
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

### Example 2

```go
package main

import (
	"context"
	sdkplatformtest "github.com/tristanspeakeasy/sdk-platform-test"
	"github.com/tristanspeakeasy/sdk-platform-test/models/components"
	"github.com/tristanspeakeasy/sdk-platform-test/types"
	"log"
	"math/big"
)

func main() {
	ctx := context.Background()

	s := sdkplatformtest.New(
		sdkplatformtest.WithQueryParam1("some example query param"),
	)

	res, err := s.TestGroup.Tag2.PostTest(ctx, components.Test2Request{
		Obj: components.ExhaustiveObject{
			Str:        "example",
			Bool:       true,
			Integer:    999999,
			Int32:      1,
			Num:        1.1,
			Float32:    4344.96,
			EnumProp:   components.EnumFirst.ToPointer(),
			Date:       types.MustDateFromString("2024-02-10"),
			DateTime:   types.MustTimeFromString("2020-01-01T00:00:00Z"),
			Anything:   "<value>",
			BoolOpt:    sdkplatformtest.Bool(true),
			IntOptNull: sdkplatformtest.Int64(999999),
			NumOptNull: sdkplatformtest.Float64(1.1),
			IntEnum:    components.IntEnumThird.ToPointer(),
			Int32Enum:  components.Int32EnumSixtyNine,
			Bigint:     big.NewInt(119171),
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
			NullableIntEnum:    components.NullableIntEnumThird.ToPointer(),
			NullableStringEnum: components.NullableStringEnumThird,
			Color:              components.ColorGreen.ToPointer(),
			Icon:               components.IconTick,
			HeroWidth:          components.HeroWidthFourHundredAndEighty.ToPointer(),
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

### A custom readme heading

A custom usage description

```go
package main

import (
	"context"
	sdkplatformtest "github.com/tristanspeakeasy/sdk-platform-test"
	"github.com/tristanspeakeasy/sdk-platform-test/models/operations"
	"log"
	"os"
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
<!-- End SDK Example Usage [usage] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports multiple security scheme combinations globally. You can choose from one of the alternatives by using the `WithSecurity` option when initializing the SDK client instance. The selected option will be used by default to authenticate with the API for all operations that support it.

#### Option1

The `Option1` alternative relies on the following scheme:

| Name                      | Type | Scheme     | Environment Variable              |
| ------------------------- | ---- | ---------- | --------------------------------- |
| `Username`<br/>`Password` | http | HTTP Basic | `SDK_USERNAME`<br/>`SDK_PASSWORD` |

Example:
```go
package main

import (
	"context"
	sdkplatformtest "github.com/tristanspeakeasy/sdk-platform-test"
	"github.com/tristanspeakeasy/sdk-platform-test/models/components"
	"github.com/tristanspeakeasy/sdk-platform-test/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := sdkplatformtest.New(
		sdkplatformtest.WithSecurity(components.Security{
			Option1: &components.SecurityOption1{
				Username: "<USERNAME>",
				Password: "<PASSWORD>",
			},
		}),
		sdkplatformtest.WithQueryParam1("some example query param"),
	)

	content, fileErr := os.Open("example.file")
	if fileErr != nil {
		panic(fileErr)
	}

	res, err := s.PostFile(ctx, operations.PostFileRequestBody{
		File: components.File{
			FileName: "example.file",
			Content:  content,
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

#### Option2

All of the following schemes must be satisfied to use the `Option2` alternative:

| Name         | Type   | Scheme      | Environment Variable |
| ------------ | ------ | ----------- | -------------------- |
| `BearerAuth` | http   | HTTP Bearer | `SDK_BEARER_AUTH`    |
| `APIKey`     | apiKey | API key     | `SDK_API_KEY`        |

Example:
```go
package main

import (
	"context"
	sdkplatformtest "github.com/tristanspeakeasy/sdk-platform-test"
	"github.com/tristanspeakeasy/sdk-platform-test/models/components"
	"github.com/tristanspeakeasy/sdk-platform-test/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := sdkplatformtest.New(
		sdkplatformtest.WithSecurity(components.Security{
			Option2: &components.SecurityOption2{
				BearerAuth: "<YOUR_JWT>",
				APIKey:     os.Getenv("SDK_API_KEY"),
			},
		}),
		sdkplatformtest.WithQueryParam1("some example query param"),
	)

	content, fileErr := os.Open("example.file")
	if fileErr != nil {
		panic(fileErr)
	}

	res, err := s.PostFile(ctx, operations.PostFileRequestBody{
		File: components.File{
			FileName: "example.file",
			Content:  content,
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

#### Option3

The `Option3` alternative relies on the following scheme:

| Name     | Type   | Scheme       | Environment Variable |
| -------- | ------ | ------------ | -------------------- |
| `Oauth2` | oauth2 | OAuth2 token | `SDK_OAUTH2`         |

Example:
```go
package main

import (
	"context"
	sdkplatformtest "github.com/tristanspeakeasy/sdk-platform-test"
	"github.com/tristanspeakeasy/sdk-platform-test/models/components"
	"github.com/tristanspeakeasy/sdk-platform-test/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := sdkplatformtest.New(
		sdkplatformtest.WithSecurity(components.Security{
			Option3: &components.SecurityOption3{
				Oauth2: "Bearer <YOUR_OAUTH2_TOKEN>",
			},
		}),
		sdkplatformtest.WithQueryParam1("some example query param"),
	)

	content, fileErr := os.Open("example.file")
	if fileErr != nil {
		panic(fileErr)
	}

	res, err := s.PostFile(ctx, operations.PostFileRequestBody{
		File: components.File{
			FileName: "example.file",
			Content:  content,
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

#### Option4

The `Option4` alternative relies on the following scheme:

| Name                 | Type | Scheme      | Environment Variable          |
| -------------------- | ---- | ----------- | ----------------------------- |
| `AppID`<br/>`Secret` | http | Custom HTTP | `SDK_APP_ID`<br/>`SDK_SECRET` |

Example:
```go
package main

import (
	"context"
	sdkplatformtest "github.com/tristanspeakeasy/sdk-platform-test"
	"github.com/tristanspeakeasy/sdk-platform-test/models/components"
	"github.com/tristanspeakeasy/sdk-platform-test/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := sdkplatformtest.New(
		sdkplatformtest.WithSecurity(components.Security{
			Option4: &components.SecurityOption4{
				AppID:  "app-speakeasy-123",
				Secret: "MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI",
			},
		}),
		sdkplatformtest.WithQueryParam1("some example query param"),
	)

	content, fileErr := os.Open("example.file")
	if fileErr != nil {
		panic(fileErr)
	}

	res, err := s.PostFile(ctx, operations.PostFileRequestBody{
		File: components.File{
			FileName: "example.file",
			Content:  content,
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

#### Option5

The `Option5` alternative relies on the following scheme:

| Name         | Type   | Scheme       | Environment Variable |
| ------------ | ------ | ------------ | -------------------- |
| `MobileAuth` | oauth2 | OAuth2 token | `SDK_MOBILE_AUTH`    |

Example:
```go
package main

import (
	"context"
	sdkplatformtest "github.com/tristanspeakeasy/sdk-platform-test"
	"github.com/tristanspeakeasy/sdk-platform-test/models/components"
	"github.com/tristanspeakeasy/sdk-platform-test/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := sdkplatformtest.New(
		sdkplatformtest.WithSecurity(components.Security{
			Option5: &components.SecurityOption5{
				MobileAuth: "Bearer <YOUR_OAUTH2_TOKEN>",
			},
		}),
		sdkplatformtest.WithQueryParam1("some example query param"),
	)

	content, fileErr := os.Open("example.file")
	if fileErr != nil {
		panic(fileErr)
	}

	res, err := s.PostFile(ctx, operations.PostFileRequestBody{
		File: components.File{
			FileName: "example.file",
			Content:  content,
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

### Per-Operation Security Schemes

Some operations in this SDK require the security scheme to be specified at the request level. For example:
```go
package main

import (
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
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [SDK](docs/sdks/sdk/README.md)

* [OperationWithLeadingAndTrailingUnderscores](docs/sdks/sdk/README.md#operationwithleadingandtrailingunderscores)
* [PostFile](docs/sdks/sdk/README.md#postfile) - Post File
* [GetPolymorphism](docs/sdks/sdk/README.md#getpolymorphism)
* [GetUnionErrors](docs/sdks/sdk/README.md#getunionerrors)
* [GetRequestBodyFlattenedAway](docs/sdks/sdk/README.md#getrequestbodyflattenedaway)
* [GetFullyFlattenedRequest](docs/sdks/sdk/README.md#getfullyflattenedrequest)
* [TestEndpoint](docs/sdks/sdk/README.md#testendpoint)
* [CreateUser](docs/sdks/sdk/README.md#createuser) - Create User
* [GetUser](docs/sdks/sdk/README.md#getuser) - Get User
* [UpdateUser](docs/sdks/sdk/README.md#updateuser) - Update User
* [DeleteUser](docs/sdks/sdk/README.md#deleteuser) - Delete User
* [Chat](docs/sdks/sdk/README.md#chat)
* [GetBinaryDefaultResponse](docs/sdks/sdk/README.md#getbinarydefaultresponse)

### [Tag1](docs/sdks/tag1/README.md)

* [~~Deprecated1~~](docs/sdks/tag1/README.md#deprecated1) - Deprecated Operation :warning: **Deprecated** Use [GetRequestBodyFlattenedAway](docs/sdks/sdk/README.md#getrequestbodyflattenedaway) instead.
* [ListTest1](docs/sdks/tag1/README.md#listtest1) - Get Test1

### [TestGroup](docs/sdks/testgroup/README.md)


#### [TestGroup.Tag2](docs/sdks/tag2/README.md)

* [PostTest](docs/sdks/tag2/README.md#posttest) - Post Test2

#### [TestGroup.Tag3](docs/sdks/tag3/README.md)

* [PostTest](docs/sdks/tag3/README.md#posttest) - Post Test2

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Global Parameters [global-parameters] -->
## Global Parameters

Certain parameters are configured globally. These parameters may be set on the SDK client instance itself during initialization. When configured as an option during SDK initialization, These global values will be used as defaults on the operations that use them. When such operations are called, there is a place in each to override the global value, if needed.

For example, you can set `queryParam1` to `"some example query param"` at SDK initialization and then you do not have to pass the same value on calls to operations like `GetRequestBodyFlattenedAway`. But if you want to do so you may, which will locally override the global setting. See the example code below for a demonstration.


### Available Globals

The following global parameters are available.
Global parameters can also be set via environment variable.

| Name                  | Type   | Description                                                                        | Environment                 |
| --------------------- | ------ | ---------------------------------------------------------------------------------- | --------------------------- |
| QueryParam1           | string | A long winded, multi-line description<br/>for the query parameter number one.<br/> | SDK_QUERY_PARAM1            |
| DeprecatedQueryParam1 | string | A deprecated description                                                           | SDK_DEPRECATED_QUERY_PARAM1 |
| DeprecatedQueryParam2 | string | The DeprecatedQueryParam2 parameter.                                               | SDK_DEPRECATED_QUERY_PARAM2 |
| LoneQueryParam        | string | The LoneQueryParam parameter.                                                      | SDK_LONE_QUERY_PARAM        |

### Example

```go
package main

import (
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
<!-- End Global Parameters [global-parameters] -->

<!-- Start Server-sent event streaming [eventstream] -->
## Server-sent event streaming

[Server-sent events][mdn-sse] are used to stream content from certain
operations. These operations will expose the stream as an iterable that
can be consumed using a simple `for` loop. The loop will
terminate when the server no longer has any events to send and closes the
underlying connection.

```go
package main

import (
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

[mdn-sse]: https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events
<!-- End Server-sent event streaming [eventstream] -->

<!-- Start Pagination [pagination] -->
## Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
```go
package main

import (
	"context"
	sdkplatformtest "github.com/tristanspeakeasy/sdk-platform-test"
	"github.com/tristanspeakeasy/sdk-platform-test/models/operations"
	"log"
	"os"
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
<!-- End Pagination [pagination] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	sdkplatformtest "github.com/tristanspeakeasy/sdk-platform-test"
	"github.com/tristanspeakeasy/sdk-platform-test/models/components"
	"github.com/tristanspeakeasy/sdk-platform-test/models/operations"
	"github.com/tristanspeakeasy/sdk-platform-test/retry"
	"log"
	"models/operations"
	"os"
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
			Content:  content,
		},
	}, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.File != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	sdkplatformtest "github.com/tristanspeakeasy/sdk-platform-test"
	"github.com/tristanspeakeasy/sdk-platform-test/models/components"
	"github.com/tristanspeakeasy/sdk-platform-test/models/operations"
	"github.com/tristanspeakeasy/sdk-platform-test/retry"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := sdkplatformtest.New(
		sdkplatformtest.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
		sdkplatformtest.WithQueryParam1("some example query param"),
	)

	content, fileErr := os.Open("example.file")
	if fileErr != nil {
		panic(fileErr)
	}

	res, err := s.PostFile(ctx, operations.PostFileRequestBody{
		File: components.File{
			FileName: "example.file",
			Content:  content,
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
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `apierrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `GetUnionErrors` function may return the following errors:

| Error Type                                   | Status Code | Content Type     |
| -------------------------------------------- | ----------- | ---------------- |
| apierrors.Error                              | 404         | application/json |
| apierrors.GetUnionErrorsResponseResponseBody | 500         | application/json |
| apierrors.GetUnionErrorsResponseBody         | 4XX         | application/json |
| apierrors.APIError                           | 5XX         | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	sdkplatformtest "github.com/tristanspeakeasy/sdk-platform-test"
	"github.com/tristanspeakeasy/sdk-platform-test/models/apierrors"
	"log"
)

func main() {
	ctx := context.Background()

	s := sdkplatformtest.New(
		sdkplatformtest.WithQueryParam1("some example query param"),
	)

	res, err := s.GetUnionErrors(ctx, 12)
	if err != nil {

		var e *apierrors.Error
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.GetUnionErrorsResponseResponseBody
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.GetUnionErrorsResponseBody
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex(serverIndex int)` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| #   | Server                     | Variables                           | Default values              |
| --- | -------------------------- | ----------------------------------- | --------------------------- |
| 0   | `http://localhost:35123`   |                                     |                             |
| 1   | `http://{hostname}:{port}` | `hostname string`<br/>`port string` | `"localhost"`<br/>`"35123"` |

If the selected server has variables, you may override their default values using their associated option(s):
 * `WithHostname(hostname string)`
 * `WithPort(port string)`

#### Example

```go
package main

import (
	"context"
	sdkplatformtest "github.com/tristanspeakeasy/sdk-platform-test"
	"github.com/tristanspeakeasy/sdk-platform-test/models/components"
	"github.com/tristanspeakeasy/sdk-platform-test/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := sdkplatformtest.New(
		sdkplatformtest.WithServerIndex(1),
		sdkplatformtest.WithQueryParam1("some example query param"),
	)

	content, fileErr := os.Open("example.file")
	if fileErr != nil {
		panic(fileErr)
	}

	res, err := s.PostFile(ctx, operations.PostFileRequestBody{
		File: components.File{
			FileName: "example.file",
			Content:  content,
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

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	sdkplatformtest "github.com/tristanspeakeasy/sdk-platform-test"
	"github.com/tristanspeakeasy/sdk-platform-test/models/components"
	"github.com/tristanspeakeasy/sdk-platform-test/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := sdkplatformtest.New(
		sdkplatformtest.WithServerURL("http://localhost:35123"),
		sdkplatformtest.WithQueryParam1("some example query param"),
	)

	content, fileErr := os.Open("example.file")
	if fileErr != nil {
		panic(fileErr)
	}

	res, err := s.PostFile(ctx, operations.PostFileRequestBody{
		File: components.File{
			FileName: "example.file",
			Content:  content,
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

### Override Server URL Per-Operation

The server URL can also be overridden on a per-operation basis, provided a server list was specified for the operation. For example:
```go
package main

import (
	"context"
	sdkplatformtest "github.com/tristanspeakeasy/sdk-platform-test"
	"github.com/tristanspeakeasy/sdk-platform-test/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := sdkplatformtest.New(
		sdkplatformtest.WithQueryParam1("some example query param"),
	)

	res, err := s.Tag1.ListTest1(ctx, operations.ListTest1Security{
		APIKey: os.Getenv("SDK_API_KEY"),
	}, operations.QueryParam2One, 100, "some example header param", sdkplatformtest.String("some example query param"), operations.WithServerURL("http://localhost:35123"))
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
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Special Types [types] -->
## Special Types

This SDK defines the following custom types to assist with marshalling and unmarshalling data.

### Date

`types.Date` is a wrapper around time.Time that allows for JSON marshaling a date string formatted as "2006-01-02".

#### Usage

```go
d1 := types.NewDate(time.Now()) // returns *types.Date

d2 := types.DateFromTime(time.Now()) // returns types.Date

d3, err := types.NewDateFromString("2019-01-01") // returns *types.Date, error

d4, err := types.DateFromString("2019-01-01") // returns types.Date, error

d5 := types.MustNewDateFromString("2019-01-01") // returns *types.Date and panics on error

d6 := types.MustDateFromString("2019-01-01") // returns types.Date and panics on error
```
<!-- End Special Types [types] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically. Any manual changes added to internal files will be overwritten on the next generation. 
We look forward to hearing your feedback. Feel free to open a PR or an issue with a proof of concept and we'll do our best to include it in a future release. 

### SDK Created by [Speakeasy](https://www.speakeasy.com/?utm_source=github-com/tristanspeakeasy/sdk-platform-test&utm_campaign=go)
