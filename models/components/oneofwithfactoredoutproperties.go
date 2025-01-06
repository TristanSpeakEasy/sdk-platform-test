// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ericlagergren/decimal"
	"github.com/tristanspeakeasy/sdk-platform-test/internal/utils"
	"github.com/tristanspeakeasy/sdk-platform-test/types"
	"math/big"
	"time"
)

type OneOfWithFactoredOutPropertiesSimpleObject struct {
	// A string property.
	Str string `json:"str"`
	// An extra property.
	AnExtraProperty *string `json:"anExtraProperty,omitempty"`
}

func (o *OneOfWithFactoredOutPropertiesSimpleObject) GetStr() string {
	if o == nil {
		return ""
	}
	return o.Str
}

func (o *OneOfWithFactoredOutPropertiesSimpleObject) GetAnExtraProperty() *string {
	if o == nil {
		return nil
	}
	return o.AnExtraProperty
}

// ExhaustiveObjectIntEnum - An integer enum property.
type ExhaustiveObjectIntEnum int64

const (
	ExhaustiveObjectIntEnumFirst  ExhaustiveObjectIntEnum = 1
	ExhaustiveObjectIntEnumSecond ExhaustiveObjectIntEnum = 2
	ExhaustiveObjectIntEnumThird  ExhaustiveObjectIntEnum = 3
)

func (e ExhaustiveObjectIntEnum) ToPointer() *ExhaustiveObjectIntEnum {
	return &e
}
func (e *ExhaustiveObjectIntEnum) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		*e = ExhaustiveObjectIntEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ExhaustiveObjectIntEnum: %v", v)
	}
}

// ExhaustiveObjectInt32Enum - An int32 enum property.
type ExhaustiveObjectInt32Enum int

const (
	ExhaustiveObjectInt32EnumFiftyFive              ExhaustiveObjectInt32Enum = 55
	ExhaustiveObjectInt32EnumSixtyNine              ExhaustiveObjectInt32Enum = 69
	ExhaustiveObjectInt32EnumOneHundredAndEightyOne ExhaustiveObjectInt32Enum = 181
)

func (e ExhaustiveObjectInt32Enum) ToPointer() *ExhaustiveObjectInt32Enum {
	return &e
}
func (e *ExhaustiveObjectInt32Enum) UnmarshalJSON(data []byte) error {
	var v int
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 55:
		fallthrough
	case 69:
		fallthrough
	case 181:
		*e = ExhaustiveObjectInt32Enum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ExhaustiveObjectInt32Enum: %v", v)
	}
}

type ExhaustiveObjectAnyType string

const (
	ExhaustiveObjectAnyTypeSimpleObject ExhaustiveObjectAnyType = "SimpleObject"
	ExhaustiveObjectAnyTypeStr          ExhaustiveObjectAnyType = "str"
)

type ExhaustiveObjectAny struct {
	SimpleObject *SimpleObject `queryParam:"inline"`
	Str          *string       `queryParam:"inline"`

	Type ExhaustiveObjectAnyType
}

func CreateExhaustiveObjectAnySimpleObject(simpleObject SimpleObject) ExhaustiveObjectAny {
	typ := ExhaustiveObjectAnyTypeSimpleObject

	return ExhaustiveObjectAny{
		SimpleObject: &simpleObject,
		Type:         typ,
	}
}

func CreateExhaustiveObjectAnyStr(str string) ExhaustiveObjectAny {
	typ := ExhaustiveObjectAnyTypeStr

	return ExhaustiveObjectAny{
		Str:  &str,
		Type: typ,
	}
}

func (u *ExhaustiveObjectAny) UnmarshalJSON(data []byte) error {

	var simpleObject SimpleObject = SimpleObject{}
	if err := utils.UnmarshalJSON(data, &simpleObject, "", true, true); err == nil {
		u.SimpleObject = &simpleObject
		u.Type = ExhaustiveObjectAnyTypeSimpleObject
		return nil
	}

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = ExhaustiveObjectAnyTypeStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for ExhaustiveObjectAny", string(data))
}

func (u ExhaustiveObjectAny) MarshalJSON() ([]byte, error) {
	if u.SimpleObject != nil {
		return utils.MarshalJSON(u.SimpleObject, "", true)
	}

	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	return nil, errors.New("could not marshal union type ExhaustiveObjectAny: all fields are null")
}

type ExhaustiveObjectNullableIntEnum int64

const (
	ExhaustiveObjectNullableIntEnumFirst  ExhaustiveObjectNullableIntEnum = 1
	ExhaustiveObjectNullableIntEnumSecond ExhaustiveObjectNullableIntEnum = 2
	ExhaustiveObjectNullableIntEnumThird  ExhaustiveObjectNullableIntEnum = 3
)

func (e ExhaustiveObjectNullableIntEnum) ToPointer() *ExhaustiveObjectNullableIntEnum {
	return &e
}
func (e *ExhaustiveObjectNullableIntEnum) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		*e = ExhaustiveObjectNullableIntEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ExhaustiveObjectNullableIntEnum: %v", v)
	}
}

type ExhaustiveObjectNullableStringEnum string

const (
	ExhaustiveObjectNullableStringEnumFirst  ExhaustiveObjectNullableStringEnum = "First"
	ExhaustiveObjectNullableStringEnumSecond ExhaustiveObjectNullableStringEnum = "Second"
	ExhaustiveObjectNullableStringEnumThird  ExhaustiveObjectNullableStringEnum = "Third"
)

func (e ExhaustiveObjectNullableStringEnum) ToPointer() *ExhaustiveObjectNullableStringEnum {
	return &e
}
func (e *ExhaustiveObjectNullableStringEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "First":
		fallthrough
	case "Second":
		fallthrough
	case "Third":
		*e = ExhaustiveObjectNullableStringEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ExhaustiveObjectNullableStringEnum: %v", v)
	}
}

type ExhaustiveObjectIcon string

const (
	ExhaustiveObjectIconTick     ExhaustiveObjectIcon = "tick"
	ExhaustiveObjectIconThumbsUp ExhaustiveObjectIcon = "thumbs-up"
	ExhaustiveObjectIconFire     ExhaustiveObjectIcon = "fire"
)

func (e ExhaustiveObjectIcon) ToPointer() *ExhaustiveObjectIcon {
	return &e
}
func (e *ExhaustiveObjectIcon) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "tick":
		fallthrough
	case "thumbs-up":
		fallthrough
	case "fire":
		*e = ExhaustiveObjectIcon(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ExhaustiveObjectIcon: %v", v)
	}
}

type ExhaustiveObjectHeroWidth int64

const (
	ExhaustiveObjectHeroWidthFourHundredAndEighty  ExhaustiveObjectHeroWidth = 480
	ExhaustiveObjectHeroWidthSevenHundredAndTwenty ExhaustiveObjectHeroWidth = 720
	ExhaustiveObjectHeroWidthOneThousandAndEighty  ExhaustiveObjectHeroWidth = 1080
)

func (e ExhaustiveObjectHeroWidth) ToPointer() *ExhaustiveObjectHeroWidth {
	return &e
}

// OneOfWithFactoredOutPropertiesExhaustiveObject - A simple object that uses all our supported primitive types and enums and has optional properties.
//
// https://speakeasy.com - A link to the external docs.
type OneOfWithFactoredOutPropertiesExhaustiveObject struct {
	// A string property.
	Str string `json:"str"`
	// A boolean property.
	Bool bool `json:"bool"`
	// An integer property.
	Integer int64 `json:"int"`
	// An int32 property.
	Int32 int `json:"int32"`
	// A number property.
	Num float64 `json:"num"`
	// A float32 property.
	Float32  float32 `json:"float32"`
	EnumProp *Enum   `default:"First" json:"enumProp"`
	// A date property.
	Date types.Date `json:"date"`
	// A date-time property.
	DateTime time.Time `json:"dateTime"`
	// An any property.
	Anything any `json:"anything"`
	// An optional string property.
	StrOpt *string `json:"strOpt,omitempty"`
	// An optional boolean property.
	BoolOpt *bool `json:"boolOpt,omitempty"`
	// An optional integer property will be null for tests.
	IntOptNull *int64 `json:"intOptNull,omitempty"`
	// An optional number property will be null for tests.
	NumOptNull *float64 `json:"numOptNull,omitempty"`
	// An integer enum property.
	IntEnum *ExhaustiveObjectIntEnum `default:"2" json:"intEnum"`
	// An int32 enum property.
	Int32Enum  ExhaustiveObjectInt32Enum `json:"int32Enum"`
	Bigint     *big.Int                  `json:"bigint"`
	BigintStr  *big.Int                  `default:"12345678901234567890" bigint:"string" json:"bigintStr"`
	decimal    *decimal.Big              `const:"3.141592653589" decimal:"number" json:"decimal,omitempty"`
	DecimalStr *decimal.Big              `decimal:"number" json:"decimalStr"`
	Obj        SimpleObject              `json:"obj"`
	Map        map[string]SimpleObject   `json:"map"`
	Arr        []SimpleObject            `json:"arr"`
	Any        ExhaustiveObjectAny       `json:"any"`
	Type       *string                   `default:"0" json:"type"`
	// A property with dots.
	SomePropertyWithDots *string                            `json:"some.property.with.dots,omitempty"`
	NullableIntEnum      *ExhaustiveObjectNullableIntEnum   `default:"2" json:"nullableIntEnum"`
	NullableStringEnum   ExhaustiveObjectNullableStringEnum `json:"nullableStringEnum"`
	Color                *Color                             `default:"blue" json:"color"`
	Icon                 ExhaustiveObjectIcon               `json:"icon"`
	HeroWidth            *ExhaustiveObjectHeroWidth         `json:"heroWidth,omitempty"`
	// An extra property.
	AnExtraProperty *string `json:"anExtraProperty,omitempty"`
}

func (o OneOfWithFactoredOutPropertiesExhaustiveObject) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OneOfWithFactoredOutPropertiesExhaustiveObject) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *OneOfWithFactoredOutPropertiesExhaustiveObject) GetStr() string {
	if o == nil {
		return ""
	}
	return o.Str
}

func (o *OneOfWithFactoredOutPropertiesExhaustiveObject) GetBool() bool {
	if o == nil {
		return false
	}
	return o.Bool
}

func (o *OneOfWithFactoredOutPropertiesExhaustiveObject) GetInteger() int64 {
	if o == nil {
		return 0
	}
	return o.Integer
}

func (o *OneOfWithFactoredOutPropertiesExhaustiveObject) GetInt32() int {
	if o == nil {
		return 0
	}
	return o.Int32
}

func (o *OneOfWithFactoredOutPropertiesExhaustiveObject) GetNum() float64 {
	if o == nil {
		return 0.0
	}
	return o.Num
}

func (o *OneOfWithFactoredOutPropertiesExhaustiveObject) GetFloat32() float32 {
	if o == nil {
		return 0.0
	}
	return o.Float32
}

func (o *OneOfWithFactoredOutPropertiesExhaustiveObject) GetEnumProp() *Enum {
	if o == nil {
		return nil
	}
	return o.EnumProp
}

func (o *OneOfWithFactoredOutPropertiesExhaustiveObject) GetDate() types.Date {
	if o == nil {
		return types.Date{}
	}
	return o.Date
}

func (o *OneOfWithFactoredOutPropertiesExhaustiveObject) GetDateTime() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.DateTime
}

func (o *OneOfWithFactoredOutPropertiesExhaustiveObject) GetAnything() any {
	if o == nil {
		return nil
	}
	return o.Anything
}

func (o *OneOfWithFactoredOutPropertiesExhaustiveObject) GetStrOpt() *string {
	if o == nil {
		return nil
	}
	return o.StrOpt
}

func (o *OneOfWithFactoredOutPropertiesExhaustiveObject) GetBoolOpt() *bool {
	if o == nil {
		return nil
	}
	return o.BoolOpt
}

func (o *OneOfWithFactoredOutPropertiesExhaustiveObject) GetIntOptNull() *int64 {
	if o == nil {
		return nil
	}
	return o.IntOptNull
}

func (o *OneOfWithFactoredOutPropertiesExhaustiveObject) GetNumOptNull() *float64 {
	if o == nil {
		return nil
	}
	return o.NumOptNull
}

func (o *OneOfWithFactoredOutPropertiesExhaustiveObject) GetIntEnum() *ExhaustiveObjectIntEnum {
	if o == nil {
		return nil
	}
	return o.IntEnum
}

func (o *OneOfWithFactoredOutPropertiesExhaustiveObject) GetInt32Enum() ExhaustiveObjectInt32Enum {
	if o == nil {
		return ExhaustiveObjectInt32Enum(0)
	}
	return o.Int32Enum
}

func (o *OneOfWithFactoredOutPropertiesExhaustiveObject) GetBigint() *big.Int {
	if o == nil {
		return big.NewInt(0)
	}
	return o.Bigint
}

func (o *OneOfWithFactoredOutPropertiesExhaustiveObject) GetBigintStr() *big.Int {
	if o == nil {
		return nil
	}
	return o.BigintStr
}

func (o *OneOfWithFactoredOutPropertiesExhaustiveObject) GetDecimal() *decimal.Big {
	return types.MustNewDecimalFromString("3.141592653589")
}

func (o *OneOfWithFactoredOutPropertiesExhaustiveObject) GetDecimalStr() *decimal.Big {
	if o == nil {
		return new(decimal.Big).SetFloat64(0.0)
	}
	return o.DecimalStr
}

func (o *OneOfWithFactoredOutPropertiesExhaustiveObject) GetObj() SimpleObject {
	if o == nil {
		return SimpleObject{}
	}
	return o.Obj
}

func (o *OneOfWithFactoredOutPropertiesExhaustiveObject) GetMap() map[string]SimpleObject {
	if o == nil {
		return map[string]SimpleObject{}
	}
	return o.Map
}

func (o *OneOfWithFactoredOutPropertiesExhaustiveObject) GetArr() []SimpleObject {
	if o == nil {
		return []SimpleObject{}
	}
	return o.Arr
}

func (o *OneOfWithFactoredOutPropertiesExhaustiveObject) GetAny() ExhaustiveObjectAny {
	if o == nil {
		return ExhaustiveObjectAny{}
	}
	return o.Any
}

func (o *OneOfWithFactoredOutPropertiesExhaustiveObject) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *OneOfWithFactoredOutPropertiesExhaustiveObject) GetSomePropertyWithDots() *string {
	if o == nil {
		return nil
	}
	return o.SomePropertyWithDots
}

func (o *OneOfWithFactoredOutPropertiesExhaustiveObject) GetNullableIntEnum() *ExhaustiveObjectNullableIntEnum {
	if o == nil {
		return nil
	}
	return o.NullableIntEnum
}

func (o *OneOfWithFactoredOutPropertiesExhaustiveObject) GetNullableStringEnum() ExhaustiveObjectNullableStringEnum {
	if o == nil {
		return ExhaustiveObjectNullableStringEnum("")
	}
	return o.NullableStringEnum
}

func (o *OneOfWithFactoredOutPropertiesExhaustiveObject) GetColor() *Color {
	if o == nil {
		return nil
	}
	return o.Color
}

func (o *OneOfWithFactoredOutPropertiesExhaustiveObject) GetIcon() ExhaustiveObjectIcon {
	if o == nil {
		return ExhaustiveObjectIcon("")
	}
	return o.Icon
}

func (o *OneOfWithFactoredOutPropertiesExhaustiveObject) GetHeroWidth() *ExhaustiveObjectHeroWidth {
	if o == nil {
		return nil
	}
	return o.HeroWidth
}

func (o *OneOfWithFactoredOutPropertiesExhaustiveObject) GetAnExtraProperty() *string {
	if o == nil {
		return nil
	}
	return o.AnExtraProperty
}

type OneOfWithFactoredOutPropertiesType string

const (
	OneOfWithFactoredOutPropertiesTypeOneOfWithFactoredOutPropertiesExhaustiveObject OneOfWithFactoredOutPropertiesType = "OneOfWithFactoredOutProperties_ExhaustiveObject"
	OneOfWithFactoredOutPropertiesTypeOneOfWithFactoredOutPropertiesSimpleObject     OneOfWithFactoredOutPropertiesType = "OneOfWithFactoredOutProperties_SimpleObject"
)

// OneOfWithFactoredOutProperties - A union of two types with factored out properties.
type OneOfWithFactoredOutProperties struct {
	OneOfWithFactoredOutPropertiesExhaustiveObject *OneOfWithFactoredOutPropertiesExhaustiveObject `queryParam:"inline"`
	OneOfWithFactoredOutPropertiesSimpleObject     *OneOfWithFactoredOutPropertiesSimpleObject     `queryParam:"inline"`

	Type OneOfWithFactoredOutPropertiesType
}

func CreateOneOfWithFactoredOutPropertiesOneOfWithFactoredOutPropertiesExhaustiveObject(oneOfWithFactoredOutPropertiesExhaustiveObject OneOfWithFactoredOutPropertiesExhaustiveObject) OneOfWithFactoredOutProperties {
	typ := OneOfWithFactoredOutPropertiesTypeOneOfWithFactoredOutPropertiesExhaustiveObject

	return OneOfWithFactoredOutProperties{
		OneOfWithFactoredOutPropertiesExhaustiveObject: &oneOfWithFactoredOutPropertiesExhaustiveObject,
		Type: typ,
	}
}

func CreateOneOfWithFactoredOutPropertiesOneOfWithFactoredOutPropertiesSimpleObject(oneOfWithFactoredOutPropertiesSimpleObject OneOfWithFactoredOutPropertiesSimpleObject) OneOfWithFactoredOutProperties {
	typ := OneOfWithFactoredOutPropertiesTypeOneOfWithFactoredOutPropertiesSimpleObject

	return OneOfWithFactoredOutProperties{
		OneOfWithFactoredOutPropertiesSimpleObject: &oneOfWithFactoredOutPropertiesSimpleObject,
		Type: typ,
	}
}

func (u *OneOfWithFactoredOutProperties) UnmarshalJSON(data []byte) error {

	var oneOfWithFactoredOutPropertiesSimpleObject OneOfWithFactoredOutPropertiesSimpleObject = OneOfWithFactoredOutPropertiesSimpleObject{}
	if err := utils.UnmarshalJSON(data, &oneOfWithFactoredOutPropertiesSimpleObject, "", true, true); err == nil {
		u.OneOfWithFactoredOutPropertiesSimpleObject = &oneOfWithFactoredOutPropertiesSimpleObject
		u.Type = OneOfWithFactoredOutPropertiesTypeOneOfWithFactoredOutPropertiesSimpleObject
		return nil
	}

	var oneOfWithFactoredOutPropertiesExhaustiveObject OneOfWithFactoredOutPropertiesExhaustiveObject = OneOfWithFactoredOutPropertiesExhaustiveObject{}
	if err := utils.UnmarshalJSON(data, &oneOfWithFactoredOutPropertiesExhaustiveObject, "", true, true); err == nil {
		u.OneOfWithFactoredOutPropertiesExhaustiveObject = &oneOfWithFactoredOutPropertiesExhaustiveObject
		u.Type = OneOfWithFactoredOutPropertiesTypeOneOfWithFactoredOutPropertiesExhaustiveObject
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for OneOfWithFactoredOutProperties", string(data))
}

func (u OneOfWithFactoredOutProperties) MarshalJSON() ([]byte, error) {
	if u.OneOfWithFactoredOutPropertiesExhaustiveObject != nil {
		return utils.MarshalJSON(u.OneOfWithFactoredOutPropertiesExhaustiveObject, "", true)
	}

	if u.OneOfWithFactoredOutPropertiesSimpleObject != nil {
		return utils.MarshalJSON(u.OneOfWithFactoredOutPropertiesSimpleObject, "", true)
	}

	return nil, errors.New("could not marshal union type OneOfWithFactoredOutProperties: all fields are null")
}
