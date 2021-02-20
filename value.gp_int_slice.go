///////////////////////////////////////////////////////////////////
//
// !!!!!!!!!!!! NEVER MODIFY THIS FILE MANUALLY !!!!!!!!!!!!
//
// This file was auto-generated by tool [github.com/vipally/gogp]
// Last update at: [Sat Feb 20 2021 10:01 CST]
// Generate from:
//   [github.com/gxlb/cli/internal/gp/value.gp]
//   [github.com/gxlb/cli/value.gpg] [value_int_slice]
//
// Tool [github.com/vipally/gogp] info:
// CopyRight 2021 @Ally Dale. All rights reserved.
// Author  : Ally Dale(vipally@gmail.com)
// Site    : https://github.com/vipally
// Version : 3.0.0.final
// 
///////////////////////////////////////////////////////////////////

package cli

import (
	"flag"

	"fmt"
	"strconv"
	//"encoding/json"
	//"strings"
)

////////////////////////////////////////////////////////////////////////////////

// IntSliceSlice wraps []int to satisfy flag.Value
type IntSliceSlice struct {
	slice      []int
	hasBeenSet bool
}

// NewIntSliceSlice makes an *IntSliceSlice with default values
func NewIntSliceSlice(defaults ...int) *IntSliceSlice {
	return &IntSliceSlice{slice: append([]int{}, defaults...)}
}

// clone allocate a copy of self object
func (s *IntSliceSlice) clone() *IntSliceSlice {
	n := &IntSliceSlice{
		slice:      make([]int, len(s.slice)),
		hasBeenSet: s.hasBeenSet,
	}
	copy(n.slice, s.slice)
	return n
}

// TODO: Consistently have specific Set function for Int64 and Float64 ?
// Append directly adds an integer to the list of values
func (s *IntSliceSlice) Append(value ...int) {
	if !s.hasBeenSet {
		s.slice = []int{}
		s.hasBeenSet = true
	}

	s.slice = append(s.slice, value...)
}

// Set parses the value into an integer and appends it to the list of values
func (s *IntSliceSlice) Set(value string) error {
	if !s.hasBeenSet {
		s.slice = []int{}
		s.hasBeenSet = true
	}

	// if strings.HasPrefix(value, slPfx) {
	// 	// Deserializing assumes overwrite
	// 	_ = json.Unmarshal([]byte(strings.Replace(value, slPfx, "", 1)), &s.slice)
	// 	s.hasBeenSet = true
	// 	return nil
	// }

	tmp, err := strconv.ParseInt(value, 0, 64)
	if err != nil {
		return err
	}

	s.slice = append(s.slice, int(tmp))

	return nil
}

// String returns a readable representation of this value (for usage defaults)
func (s *IntSliceSlice) String() string {
	return fmt.Sprintf("%#v", s.slice)
}

// Serialize allows IntSliceSlice to fulfill Serializer
func (s *IntSliceSlice) Serialize() string {
	//TODO:
	// jsonBytes, _ := json.Marshal(s.slice)
	// return fmt.Sprintf("%s%s", slPfx, string(jsonBytes))
	return ""
}

// Value returns the slice of ints set by this flag
func (s *IntSliceSlice) Value() []int {
	return s.slice
}

// Get returns the slice set by this flag
func (s *IntSliceSlice) Get() interface{} {
	return *s
}

// IntSliceValue define a value of type *IntSliceSlice
type IntSliceValue struct {
	Value       *IntSliceSlice // The value from ENV of files
	Target      *IntSliceSlice // Target set the outer value pointer
	Default     *IntSliceSlice // Default value
	DefaultText string         // Default value in help info
	Enums       []int          // Enumeration of valid values
	Ranges      []int          // [min,max,min,max...] ranges of valid values
	hasBeenSet  bool
}

func (v *IntSliceValue) IsSet() bool {

	return v.Value.hasBeenSet

	return v.hasBeenSet
}

func (v *IntSliceValue) Apply(f *Flag, set *flag.FlagSet) error {
	return nil
}

func (v *IntSliceValue) String() string {
	return ""
}
