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