///////////////////////////////////////////////////////////////////
//
// !!!!!!!!!!!! NEVER MODIFY THIS FILE MANUALLY !!!!!!!!!!!!
//
// This file was auto-generated by tool [github.com/gxlb/gogp]
// Last update at: [Sun Feb 21 2021 20:18 CST]
// Generate from:
//   [github.com/gxlb/cli/internal/gp/flag.gp]
//   [github.com/gxlb/cli/flag.gpg] [flag_int64_slice]
//
// Tool [github.com/gxlb/gogp] info:
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

	"cli/internal/impl"
	"cli/internal/util"
	"encoding/json"
	"strconv"
	"strings"
)

var _ = (*strconv.NumError)(nil) //avoid compile error

// Int64SliceSlice wraps []int64 to satisfy flag.Value
type Int64SliceSlice struct {
	slice      []int64
	hasBeenSet bool
}

// NewInt64SliceSlice makes an *Int64SliceSlice with default values
func NewInt64SliceSlice(defaults ...int64) *Int64SliceSlice {
	return &Int64SliceSlice{
		slice:      append([]int64{}, defaults...),
		hasBeenSet: false,
	}
}

// clone allocate a copy of self object
func (s *Int64SliceSlice) clone() *Int64SliceSlice {
	n := &Int64SliceSlice{
		slice:      make([]int64, len(s.slice)),
		hasBeenSet: s.hasBeenSet,
	}
	copy(n.slice, s.slice)
	return n
}

// Append directly append values to the list of values
func (s *Int64SliceSlice) Append(values ...int64) {
	s.setValues(false, values)
}

// Append directly overite values to the list of values
func (s *Int64SliceSlice) SetValues(values ...int64) {
	s.setValues(true, values)
}

// Append directly adds values to the list of values
func (s *Int64SliceSlice) setValues(overwrite bool, values []int64) {
	if !s.hasBeenSet || overwrite {
		s.Reset()
		s.hasBeenSet = true
	}

	s.slice = append(s.slice, values...)
}

// Set parses the value and appends it to the list of values
func (s *Int64SliceSlice) Set(value string) error {

	if strings.HasPrefix(value, impl.SerializedPrefix) {
		// Deserializing assumes overwrite
		_ = json.Unmarshal([]byte(strings.Replace(value, impl.SerializedPrefix, "", 1)), &s.slice)
		s.hasBeenSet = true
		return nil
	}

	//accept multi values for slice flags
	for _, val := range impl.FlagSplitMultiValues(value) {
		value := strings.TrimSpace(val)
		tmp, err := strconv.ParseInt(value, 0, 64)
		if err != nil {
			return err
		}

		if !s.hasBeenSet {
			s.slice = []int64{}
			s.hasBeenSet = true
		}

		s.slice = append(s.slice, int64(tmp))
	}

	return nil
}

// Reset clean the last parsed values of this slice
func (s *Int64SliceSlice) Reset() {
	if s.slice == nil {
		s.slice = []int64{}
	} else {
		s.slice = s.slice[:0]
	}
	s.hasBeenSet = false
}

// String returns a readable representation of this value (for usage defaults)
func (s *Int64SliceSlice) String() string {
	return fmt.Sprintf("%#v", s.slice)
}

// Serialize allows Int64SliceSlice to fulfill Serializer
func (s *Int64SliceSlice) Serialize() string {
	jsonBytes, _ := json.Marshal(s.slice)
	return fmt.Sprintf("%s%s", impl.SerializedPrefix, string(jsonBytes))
}

// Value returns the slice of ints set by this flag
func (s *Int64SliceSlice) Value() []int64 {
	return s.slice
}

// Get returns the slice set by this flag
func (s *Int64SliceSlice) Get() interface{} {
	return *s
}

// Int64SliceFlag define a value of type *Int64SliceSlice
type Int64SliceFlag struct {
	//
	//name related area
	//
	LogicName string   // logic name of the flag
	Name      string   // name of the flag
	Aliases   []string // aliases of the flag
	Usage     string   // usage string
	Required  bool     // if required
	Hidden    bool     // hidden this flag
	EnvVars   []string // environment values
	FilePath  string   // file path
	//
	//value related area
	//
	Target      *Int64SliceSlice // Target value pointer outside
	Default     *Int64SliceSlice // Default value
	DefaultText string           // Default value display in help info
	Enums       []int64          // Enumeration of valid values
	Ranges      []int64          // {[min,max),[min,max),...} ranges of valid values
	//
	////////////////////////////////////////////////////////////////////////////
	//area for parsing
	target *Int64SliceSlice // target value pointer(maybe new(*Int64SliceSlice) if Target not set)
	info   impl.FlagInfo    // parsed info of this flag
}

// Init verify and init the value by ower flag
func (v *Int64SliceFlag) Init(namegen *util.NameGenenerator) error {
	v.info.Flag = v
	v.info.EnvVars = v.EnvVars
	v.info.Usage = v.Usage
	v.info.DefaultText = v.DefaultText
	v.info.Required = v.Required
	v.info.Hidden = v.Hidden
	v.info.FilePath = v.FilePath
	v.info.LogicName = impl.FlagLogicName(v.LogicName)
	v.info.Name = namegen.GetOrGenName(v.Name)
	v.info.HasBeenSet = false
	v.info.DispName = impl.FlagDispName(v.Name, v.LogicName)
	impl.MergeNames(v.Name, v.Aliases, &v.info.Names)

	maxSliceLen := impl.MaxSliceLen
	if l := len(v.Enums); l > maxSliceLen {
		return fmt.Errorf("flag %s.Enums too long: %d/%d", v.info.DispName, l, maxSliceLen)
	}
	if l := len(v.Ranges); l > 0 {
		if l > maxSliceLen {
			return fmt.Errorf("flag %s.Ranges too long: %d/%d", v.info.DispName, l, maxSliceLen)
		}
		if l%2 != 0 {
			return fmt.Errorf("flag %s.Ranges doesn't match [min,max) pairs: %d", v.info.DispName, l)
		}
		for i := 0; i < l; i += 2 {
			min, max := v.Ranges[i], v.Ranges[i+1]
			if valid := min <= max; !valid {
				return fmt.Errorf("flag %s.Ranges doesn't match [min,max): (%d,%d)", v.info.DispName, min, max)
			}
		}
	}
	if v.Name == "" && v.LogicName == "" {
		return fmt.Errorf("flag missing both Name & LogicName: %v", v)
	}
	if v.Name == "" && len(v.Aliases) > 0 {
		return fmt.Errorf("flag %s missing name, but has Aliases %v", v.info.DispName, v.Aliases)
	}

	if err := v.validateValues(v.Default); err != nil {
		return fmt.Errorf("default value invalid: %s", err.Error())
	}
	if v.Target != nil {
		v.target = v.Target
	} else {
		v.target = NewInt64SliceSlice()
	}
	return nil
}

// IsSet check if value was set
func (v *Int64SliceFlag) IsSet() bool {
	return v.target.hasBeenSet
}

// Apply coordinate the value to flagset
func (v *Int64SliceFlag) Apply(set *flag.FlagSet) error {
	return nil
}

// String return the value for view
func (v *Int64SliceFlag) String() string {
	return ""
}

// ValidateValues verify if all values was valid
func (v *Int64SliceFlag) ValidateValues() error {
	return v.validateValues(v.target)
}

// Info returns parsed info of this flag
func (v *Int64SliceFlag) Info() *impl.FlagInfo {
	return &v.info
}

// Reset clean the last parsed value of this flag
func (v *Int64SliceFlag) Reset() {
	v.target.Reset()
	v.info.HasBeenSet = false
}

// for default value verify
func (v *Int64SliceFlag) validateValues(values *Int64SliceSlice) error {
	for _, val := range values.slice {
		if err := v.validValue(val); err != nil {
			return err
		}
	}
	return nil
}

// check if value if valid for this flag
func (v *Int64SliceFlag) validValue(value int64) error {
	f := &v.info
	if len(v.Enums) > 0 {
		found := false
		for _, v := range v.Enums {
			if value == v {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("flag %s value %v out of Enums: %v", f.LogicName, value, v.Enums)
		}
	}
	if len(v.Ranges) > 0 {
		found := false
		for i := 0; i < len(v.Ranges); i += 2 {
			min, max := v.Ranges[i], v.Ranges[i+1]
			if value >= min && value < max {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("flag %s value %v out of Ranges: %v", f.LogicName, value, v.Enums)
		}
	}
	return nil
}

var _ impl.Flag = (*Int64SliceFlag)(nil) //for interface verification only
