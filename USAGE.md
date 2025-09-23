<!-- Start SDK Example Usage [usage] -->
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

	s := sdkplatformtest.New()

	example, fileErr := os.Open("example.file")
	if fileErr != nil {
		panic(fileErr)
	}

	res, err := s.PostFile(ctx, operations.PostFileRequestBody{
		File: components.File{
			FileName: "example.file",
			Content:  example,
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
		sdkplatformtest.WithDeprecatedQueryParam1("some example query param"),
		sdkplatformtest.WithDeprecatedQueryParam2("some example query param"),
	)

	res, err := s.TestGroup.Tag2.PostTest(ctx, components.Test2Request{
		Obj: components.ExhaustiveObject{
			Str:        "example",
			Bool:       true,
			Integer:    999999,
			Int32:      1,
			Num:        1.1,
			Float32:    8499.3,
			Date:       types.MustDateFromString("2024-10-12"),
			DateTime:   types.MustTimeFromString("2020-01-01T00:00:00Z"),
			Anything:   "<value>",
			BoolOpt:    sdkplatformtest.Pointer(true),
			IntOptNull: sdkplatformtest.Pointer[int64](999999),
			NumOptNull: sdkplatformtest.Pointer[float64](1.1),
			IntEnum:    components.IntEnumThird.ToPointer(),
			Int32Enum:  components.Int32EnumSixtyNine,
			Bigint:     big.NewInt(702830),
			DecimalStr: types.MustNewDecimalFromString("3858.6"),
			Obj: components.SimpleObject{
				Str: "example",
			},
			Map: map[string]components.SimpleObject{},
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
			NullableIntEnum:    components.NullableIntEnumThird.ToPointer(),
			NullableStringEnum: components.NullableStringEnumSecond,
			Color:              components.ColorGreen.ToPointer(),
			Icon:               components.IconTick,
			HeroWidth:          components.HeroWidthFourHundredAndEighty.ToPointer(),
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
	}, operations.QueryParam2One, 100, "some example header param")
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