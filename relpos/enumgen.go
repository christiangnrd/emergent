// Code generated by "goki generate"; DO NOT EDIT.

package relpos

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"goki.dev/enums"
)

var _RelationsValues = []Relations{0, 1, 2, 3, 4, 5, 6}

// RelationsN is the highest valid value
// for type Relations, plus one.
const RelationsN Relations = 7

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _RelationsNoOp() {
	var x [1]struct{}
	_ = x[NoRel-(0)]
	_ = x[RightOf-(1)]
	_ = x[LeftOf-(2)]
	_ = x[Behind-(3)]
	_ = x[FrontOf-(4)]
	_ = x[Above-(5)]
	_ = x[Below-(6)]
}

var _RelationsNameToValueMap = map[string]Relations{
	`NoRel`:   0,
	`norel`:   0,
	`RightOf`: 1,
	`rightof`: 1,
	`LeftOf`:  2,
	`leftof`:  2,
	`Behind`:  3,
	`behind`:  3,
	`FrontOf`: 4,
	`frontof`: 4,
	`Above`:   5,
	`above`:   5,
	`Below`:   6,
	`below`:   6,
}

var _RelationsDescMap = map[Relations]string{
	0: ``,
	1: ``,
	2: ``,
	3: ``,
	4: ``,
	5: ``,
	6: ``,
}

var _RelationsMap = map[Relations]string{
	0: `NoRel`,
	1: `RightOf`,
	2: `LeftOf`,
	3: `Behind`,
	4: `FrontOf`,
	5: `Above`,
	6: `Below`,
}

// String returns the string representation
// of this Relations value.
func (i Relations) String() string {
	if str, ok := _RelationsMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the Relations value from its
// string representation, and returns an
// error if the string is invalid.
func (i *Relations) SetString(s string) error {
	if val, ok := _RelationsNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	if val, ok := _RelationsNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type Relations")
}

// Int64 returns the Relations value as an int64.
func (i Relations) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the Relations value from an int64.
func (i *Relations) SetInt64(in int64) {
	*i = Relations(in)
}

// Desc returns the description of the Relations value.
func (i Relations) Desc() string {
	if str, ok := _RelationsDescMap[i]; ok {
		return str
	}
	return i.String()
}

// RelationsValues returns all possible values
// for the type Relations.
func RelationsValues() []Relations {
	return _RelationsValues
}

// Values returns all possible values
// for the type Relations.
func (i Relations) Values() []enums.Enum {
	res := make([]enums.Enum, len(_RelationsValues))
	for i, d := range _RelationsValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type Relations.
func (i Relations) IsValid() bool {
	_, ok := _RelationsMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Relations) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Relations) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println(err)
	}
	return nil
}

var _XAlignsValues = []XAligns{0, 1, 2}

// XAlignsN is the highest valid value
// for type XAligns, plus one.
const XAlignsN XAligns = 3

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _XAlignsNoOp() {
	var x [1]struct{}
	_ = x[Left-(0)]
	_ = x[Middle-(1)]
	_ = x[Right-(2)]
}

var _XAlignsNameToValueMap = map[string]XAligns{
	`Left`:   0,
	`left`:   0,
	`Middle`: 1,
	`middle`: 1,
	`Right`:  2,
	`right`:  2,
}

var _XAlignsDescMap = map[XAligns]string{
	0: ``,
	1: ``,
	2: ``,
}

var _XAlignsMap = map[XAligns]string{
	0: `Left`,
	1: `Middle`,
	2: `Right`,
}

// String returns the string representation
// of this XAligns value.
func (i XAligns) String() string {
	if str, ok := _XAlignsMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the XAligns value from its
// string representation, and returns an
// error if the string is invalid.
func (i *XAligns) SetString(s string) error {
	if val, ok := _XAlignsNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	if val, ok := _XAlignsNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type XAligns")
}

// Int64 returns the XAligns value as an int64.
func (i XAligns) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the XAligns value from an int64.
func (i *XAligns) SetInt64(in int64) {
	*i = XAligns(in)
}

// Desc returns the description of the XAligns value.
func (i XAligns) Desc() string {
	if str, ok := _XAlignsDescMap[i]; ok {
		return str
	}
	return i.String()
}

// XAlignsValues returns all possible values
// for the type XAligns.
func XAlignsValues() []XAligns {
	return _XAlignsValues
}

// Values returns all possible values
// for the type XAligns.
func (i XAligns) Values() []enums.Enum {
	res := make([]enums.Enum, len(_XAlignsValues))
	for i, d := range _XAlignsValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type XAligns.
func (i XAligns) IsValid() bool {
	_, ok := _XAlignsMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i XAligns) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *XAligns) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println(err)
	}
	return nil
}

var _YAlignsValues = []YAligns{0, 1, 2}

// YAlignsN is the highest valid value
// for type YAligns, plus one.
const YAlignsN YAligns = 3

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _YAlignsNoOp() {
	var x [1]struct{}
	_ = x[Front-(0)]
	_ = x[Center-(1)]
	_ = x[Back-(2)]
}

var _YAlignsNameToValueMap = map[string]YAligns{
	`Front`:  0,
	`front`:  0,
	`Center`: 1,
	`center`: 1,
	`Back`:   2,
	`back`:   2,
}

var _YAlignsDescMap = map[YAligns]string{
	0: ``,
	1: ``,
	2: ``,
}

var _YAlignsMap = map[YAligns]string{
	0: `Front`,
	1: `Center`,
	2: `Back`,
}

// String returns the string representation
// of this YAligns value.
func (i YAligns) String() string {
	if str, ok := _YAlignsMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the YAligns value from its
// string representation, and returns an
// error if the string is invalid.
func (i *YAligns) SetString(s string) error {
	if val, ok := _YAlignsNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	if val, ok := _YAlignsNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type YAligns")
}

// Int64 returns the YAligns value as an int64.
func (i YAligns) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the YAligns value from an int64.
func (i *YAligns) SetInt64(in int64) {
	*i = YAligns(in)
}

// Desc returns the description of the YAligns value.
func (i YAligns) Desc() string {
	if str, ok := _YAlignsDescMap[i]; ok {
		return str
	}
	return i.String()
}

// YAlignsValues returns all possible values
// for the type YAligns.
func YAlignsValues() []YAligns {
	return _YAlignsValues
}

// Values returns all possible values
// for the type YAligns.
func (i YAligns) Values() []enums.Enum {
	res := make([]enums.Enum, len(_YAlignsValues))
	for i, d := range _YAlignsValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type YAligns.
func (i YAligns) IsValid() bool {
	_, ok := _YAlignsMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i YAligns) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *YAligns) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println(err)
	}
	return nil
}