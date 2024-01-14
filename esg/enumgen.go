// Code generated by "goki generate -add-types"; DO NOT EDIT.

package esg

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"goki.dev/enums"
)

var _CondElsValues = []CondEls{0, 1, 2, 3, 4}

// CondElsN is the highest valid value
// for type CondEls, plus one.
const CondElsN CondEls = 5

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _CondElsNoOp() {
	var x [1]struct{}
	_ = x[CRule-(0)]
	_ = x[And-(1)]
	_ = x[Or-(2)]
	_ = x[Not-(3)]
	_ = x[SubCond-(4)]
}

var _CondElsNameToValueMap = map[string]CondEls{
	`CRule`:   0,
	`crule`:   0,
	`And`:     1,
	`and`:     1,
	`Or`:      2,
	`or`:      2,
	`Not`:     3,
	`not`:     3,
	`SubCond`: 4,
	`subcond`: 4,
}

var _CondElsDescMap = map[CondEls]string{
	0: `CRule means Rule is name of a rule to evaluate truth value`,
	1: ``,
	2: ``,
	3: ``,
	4: `SubCond is a sub-condition expression`,
}

var _CondElsMap = map[CondEls]string{
	0: `CRule`,
	1: `And`,
	2: `Or`,
	3: `Not`,
	4: `SubCond`,
}

// String returns the string representation
// of this CondEls value.
func (i CondEls) String() string {
	if str, ok := _CondElsMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the CondEls value from its
// string representation, and returns an
// error if the string is invalid.
func (i *CondEls) SetString(s string) error {
	if val, ok := _CondElsNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	if val, ok := _CondElsNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type CondEls")
}

// Int64 returns the CondEls value as an int64.
func (i CondEls) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the CondEls value from an int64.
func (i *CondEls) SetInt64(in int64) {
	*i = CondEls(in)
}

// Desc returns the description of the CondEls value.
func (i CondEls) Desc() string {
	if str, ok := _CondElsDescMap[i]; ok {
		return str
	}
	return i.String()
}

// CondElsValues returns all possible values
// for the type CondEls.
func CondElsValues() []CondEls {
	return _CondElsValues
}

// Values returns all possible values
// for the type CondEls.
func (i CondEls) Values() []enums.Enum {
	res := make([]enums.Enum, len(_CondElsValues))
	for i, d := range _CondElsValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type CondEls.
func (i CondEls) IsValid() bool {
	_, ok := _CondElsMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i CondEls) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *CondEls) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println("CondEls.UnmarshalText:", err)
	}
	return nil
}

var _ElementsValues = []Elements{0, 1}

// ElementsN is the highest valid value
// for type Elements, plus one.
const ElementsN Elements = 2

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _ElementsNoOp() {
	var x [1]struct{}
	_ = x[RuleEl-(0)]
	_ = x[TokenEl-(1)]
}

var _ElementsNameToValueMap = map[string]Elements{
	`RuleEl`:  0,
	`ruleel`:  0,
	`TokenEl`: 1,
	`tokenel`: 1,
}

var _ElementsDescMap = map[Elements]string{
	0: `RuleEl means Value is name of a rule`,
	1: `TokenEl means Value is a token to emit`,
}

var _ElementsMap = map[Elements]string{
	0: `RuleEl`,
	1: `TokenEl`,
}

// String returns the string representation
// of this Elements value.
func (i Elements) String() string {
	if str, ok := _ElementsMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the Elements value from its
// string representation, and returns an
// error if the string is invalid.
func (i *Elements) SetString(s string) error {
	if val, ok := _ElementsNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	if val, ok := _ElementsNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type Elements")
}

// Int64 returns the Elements value as an int64.
func (i Elements) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the Elements value from an int64.
func (i *Elements) SetInt64(in int64) {
	*i = Elements(in)
}

// Desc returns the description of the Elements value.
func (i Elements) Desc() string {
	if str, ok := _ElementsDescMap[i]; ok {
		return str
	}
	return i.String()
}

// ElementsValues returns all possible values
// for the type Elements.
func ElementsValues() []Elements {
	return _ElementsValues
}

// Values returns all possible values
// for the type Elements.
func (i Elements) Values() []enums.Enum {
	res := make([]enums.Enum, len(_ElementsValues))
	for i, d := range _ElementsValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type Elements.
func (i Elements) IsValid() bool {
	_, ok := _ElementsMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i Elements) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *Elements) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println("Elements.UnmarshalText:", err)
	}
	return nil
}

var _RuleTypesValues = []RuleTypes{0, 1, 2, 3, 4}

// RuleTypesN is the highest valid value
// for type RuleTypes, plus one.
const RuleTypesN RuleTypes = 5

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the enumgen command to generate them again.
func _RuleTypesNoOp() {
	var x [1]struct{}
	_ = x[UniformItems-(0)]
	_ = x[ProbItems-(1)]
	_ = x[CondItems-(2)]
	_ = x[SequentialItems-(3)]
	_ = x[PermutedItems-(4)]
}

var _RuleTypesNameToValueMap = map[string]RuleTypes{
	`UniformItems`:    0,
	`uniformitems`:    0,
	`ProbItems`:       1,
	`probitems`:       1,
	`CondItems`:       2,
	`conditems`:       2,
	`SequentialItems`: 3,
	`sequentialitems`: 3,
	`PermutedItems`:   4,
	`permuteditems`:   4,
}

var _RuleTypesDescMap = map[RuleTypes]string{
	0: `UniformItems is the default mutually exclusive items chosen at uniform random`,
	1: `ProbItems has specific probabilities for each item`,
	2: `CondItems has conditionals for each item, indicated by ?`,
	3: `SequentialItems progresses through items in sequential order, indicated by |`,
	4: `PermutedItems progresses through items in permuted order, indicated by $`,
}

var _RuleTypesMap = map[RuleTypes]string{
	0: `UniformItems`,
	1: `ProbItems`,
	2: `CondItems`,
	3: `SequentialItems`,
	4: `PermutedItems`,
}

// String returns the string representation
// of this RuleTypes value.
func (i RuleTypes) String() string {
	if str, ok := _RuleTypesMap[i]; ok {
		return str
	}
	return strconv.FormatInt(int64(i), 10)
}

// SetString sets the RuleTypes value from its
// string representation, and returns an
// error if the string is invalid.
func (i *RuleTypes) SetString(s string) error {
	if val, ok := _RuleTypesNameToValueMap[s]; ok {
		*i = val
		return nil
	}
	if val, ok := _RuleTypesNameToValueMap[strings.ToLower(s)]; ok {
		*i = val
		return nil
	}
	return errors.New(s + " is not a valid value for type RuleTypes")
}

// Int64 returns the RuleTypes value as an int64.
func (i RuleTypes) Int64() int64 {
	return int64(i)
}

// SetInt64 sets the RuleTypes value from an int64.
func (i *RuleTypes) SetInt64(in int64) {
	*i = RuleTypes(in)
}

// Desc returns the description of the RuleTypes value.
func (i RuleTypes) Desc() string {
	if str, ok := _RuleTypesDescMap[i]; ok {
		return str
	}
	return i.String()
}

// RuleTypesValues returns all possible values
// for the type RuleTypes.
func RuleTypesValues() []RuleTypes {
	return _RuleTypesValues
}

// Values returns all possible values
// for the type RuleTypes.
func (i RuleTypes) Values() []enums.Enum {
	res := make([]enums.Enum, len(_RuleTypesValues))
	for i, d := range _RuleTypesValues {
		res[i] = d
	}
	return res
}

// IsValid returns whether the value is a
// valid option for type RuleTypes.
func (i RuleTypes) IsValid() bool {
	_, ok := _RuleTypesMap[i]
	return ok
}

// MarshalText implements the [encoding.TextMarshaler] interface.
func (i RuleTypes) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the [encoding.TextUnmarshaler] interface.
func (i *RuleTypes) UnmarshalText(text []byte) error {
	if err := i.SetString(string(text)); err != nil {
		log.Println("RuleTypes.UnmarshalText:", err)
	}
	return nil
}
