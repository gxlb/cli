///////////////////////////////////////////////////////////////////
//
// !!!!!!!!!!!! NEVER MODIFY THIS FILE MANUALLY !!!!!!!!!!!!
//
// This file was auto-generated by tool [github.com/gxlb/gogp]
// Last update at: [Mon Feb 22 2021 10:11 CST]
// Generate from:
//   [github.com/gxlb/cli/internal/gp/flag.gp]
//   [github.com/gxlb/cli/flag.gpg] [flag_uint64_slice]
//
// Tool [github.com/gxlb/gogp] info:
// CopyRight 2021 @Ally Dale. All rights reserved.
// Author  : Ally Dale(vipally@gmail.com)
// Site    : https://github.com/vipally
// Version : v3.0.0.final
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

// Uint64SliceSlice wraps []uint64 to satisfy flag.Value
type Uint64Slice struct {
	slice      []uint64
	hasBeenSet bool
}

// NewUint64SliceSlice makes an *Uint64SliceSlice with default values
func NewUint64SliceSlice(defaults ...uint64) *Uint64Slice {
	return &Uint64Slice{
		slice:      append([]uint64{}, defaults...),
		hasBeenSet: false,
	}
}

// clone allocate a copy of self object
func (s *Uint64Slice) clone() *Uint64Slice {
	n := &Uint64Slice{
		slice:      append([]uint64{}, s.slice...),
		hasBeenSet: s.hasBeenSet,
	}
	return n
}

// AppendValues directly append values to the list of values
func (s *Uint64Slice) AppendValues(values ...uint64) {
	s.setValues(false, values)
}

// SetValues directly overite values to the list of values
func (s *Uint64Slice) SetValues(values ...uint64) {
	s.setValues(true, values)
}

func (s *Uint64Slice) setValues(overwrite bool, values []uint64) {
	if !s.hasBeenSet || overwrite {
		s.Reset()
		s.hasBeenSet = true
	}

	s.slice = append(s.slice, values...)
}

// Set parses the value and appends to the list of values
func (s *Uint64Slice) Set(value string) error {

	if strings.HasPrefix(value, impl.SerializedPrefix) {
		// Deserializing assumes overwrite
		_ = json.Unmarshal([]byte(strings.Replace(value, impl.SerializedPrefix, "", 1)), &s.slice)
		s.hasBeenSet = true
		return nil
	}

	//accept multi values for slice flags
	for _, val := range impl.FlagSplitMultiValues(value) {
		value := strings.TrimSpace(val)
		tmp, err := strconv.ParseUint(value, 0, 64)
		if err != nil {
			return err
		}

		if !s.hasBeenSet {
			s.slice = []uint64{}
			s.hasBeenSet = true
		}

		s.slice = append(s.slice, uint64(tmp))
	}

	return nil
}

// Reset clean the last parsed values of this slice
func (s *Uint64Slice) Reset() {
	if s.slice == nil {
		s.slice = []uint64{}
	} else {
		s.slice = s.slice[:0]
	}
	s.hasBeenSet = false
}

// String returns a readable representation of this value (for usage defaults)
func (s *Uint64Slice) String() string {
	return fmt.Sprintf("%#v", s.slice)
}

// Serialize allows Uint64SliceSlice to fulfill Serializer
func (s *Uint64Slice) Serialize() string {
	jsonBytes, _ := json.Marshal(s.slice)
	return fmt.Sprintf("%s%s", impl.SerializedPrefix, string(jsonBytes))
}

// Value returns the slice of ints set by this flag
func (s *Uint64Slice) Value() []uint64 {
	return s.slice
}

// Get returns the slice set by this flag
func (s *Uint64Slice) Get() interface{} {
	return *s
}

// Uint64SliceFlag define a value of type *Uint64Slice
type Uint64SliceFlag struct {
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
	Target      *Uint64Slice // Target value pointer outside
	Default     *Uint64Slice // Default value
	DefaultText string       // Default value display in help info
	Enums       []uint64     // Enumeration of valid values
	Ranges      []uint64     // {[min,max),[min,max),...} ranges of valid values
	//
	////////////////////////////////////////////////////////////////////////////
	//area for parsing
	target *Uint64Slice  // target value pointer(maybe new(*Uint64Slice) if Target not set)
	info   impl.FlagInfo // parsed info of this flag
}

// init verify and init the flag info
func (f *Uint64SliceFlag) init(namegen *util.NameGenenerator) error {
	f.info.Flag = f
	f.info.EnvVars = append([]string{}, f.EnvVars...)
	f.info.Usage = f.Usage
	f.info.DefaultText = f.DefaultText
	f.info.Required = f.Required
	f.info.Hidden = f.Hidden
	f.info.FilePath = f.FilePath
	f.info.HasBeenSet = false
	f.info.Name = namegen.GetOrGenName(f.Name)
	f.info.NonameFlag = f.info.Name != f.Name
	f.info.LogicName = impl.FlagLogicName(f.Name, f.LogicName)
	f.info.ValueName = impl.FlagValueName(f.LogicName)
	impl.MergeNames(f.info.Name, f.Aliases, &f.info.Names) //use v.info.Name to enable auto-generated name

	//make the target pointer
	if f.Target != nil {
		f.target = f.Target
	} else {
		f.target = NewUint64SliceSlice()
	}

	maxSliceLen := impl.MaxSliceLen
	if f.Name == "" && f.LogicName == "" { // Name & LogicName cannot both missing
		return fmt.Errorf("flag missing both Name & LogicName: %v", f)
	}
	if f.Name == "" && len(f.Aliases) > 0 { // Noname ones must without Aliases
		return fmt.Errorf("flag %s missing name, but has Aliases %v", f.info.LogicName, f.Aliases)
	}
	if l := len(f.Enums); l > 0 { // Enums length check
		if l > maxSliceLen {
			return fmt.Errorf("flag %s.Enums too long: %d/%d", f.info.LogicName, l, maxSliceLen)
		}

		if l > 1 {
			var filter = make(map[uint64]struct{})
			for _, v := range f.Enums {
				if i, ok := filter[v]; !ok {
					filter[v] = struct{}{}
				} else {
					return fmt.Errorf("flag %s.Enums error: duplicate %v at %d", f.info.LogicName, v, i)
				}
			}
		}
	}
	if l := len(f.Ranges); l > 0 { // Ranges length check and [min,max) pair check
		if l > maxSliceLen {
			return fmt.Errorf("flag %s.Ranges too long: %d/%d", f.info.LogicName, l, maxSliceLen)
		}
		if l%2 != 0 {
			return fmt.Errorf("flag %s.Ranges doesn't match [min,max) pairs: %d", f.info.LogicName, l)
		}
		for i := 0; i < l; i += 2 {
			min, max := f.Ranges[i], f.Ranges[i+1]
			if valid := min <= max; !valid {
				return fmt.Errorf("flag %s.Ranges doesn't match [min,max): (%v,%v)", f.info.LogicName, min, max)
			}
			for j := 0; j < i; j += 2 { //check range overlapping
				m, n := f.Ranges[j], f.Ranges[j+1]
				if m >= min && m < max || n >= min && n < max {
					return fmt.Errorf("flag %s.Ranges %d~[%v,%v) overlapping %d~[%v,%v) ", f.info.LogicName, i, min, max, j, m, n)
				}
			}
		}
	}
	if err := f.validateValues(f.Default); err != nil { // verify default values
		return fmt.Errorf("default value invalid: %s", err.Error())
	}
	if err := util.FiltNames(f.info.Names); err != nil { // verify name duplicate
		return fmt.Errorf("flag %s.Names error: %s", f.info.LogicName, err.Error())
	}
	if err := util.FiltNames(f.info.EnvVars); err != nil { // verify EnvVars duplicate
		return fmt.Errorf("flag %s.EnvVars error: %s", f.info.LogicName, err.Error())
	}
	return nil
}

// IsSet check if value was set
func (f *Uint64SliceFlag) IsSet() bool {
	return f.target.hasBeenSet
}

//GetLogicName returns the logic name of the falg
func (f *Uint64SliceFlag) GetLogicName() string {
	return f.info.LogicName
}

//GetValueName returns the value name of the falg
func (f *Uint64SliceFlag) GetValueName() string {
	return f.info.ValueName
}

// Names returns the names of the flag
func (f *Uint64SliceFlag) Names() []string {
	return f.info.Names
}

// IsRequired returns whether or not the flag is required
func (f *Uint64SliceFlag) IsRequired() bool {
	return f.Required
}

// TakesValue returns true of the flag takes a value, otherwise false
func (f *Uint64SliceFlag) TakesValue() bool {
	return false
}

// GetUsage returns the usage string for the flag
func (f *Uint64SliceFlag) GetUsage() string {
	return f.info.Usage
}

// GetValue returns the flags value as string representation and an empty
// string if the flag takes no value at all.
func (f *Uint64SliceFlag) GetValue() string {
	return ""
}

// Apply coordinate the value to flagset
func (f *Uint64SliceFlag) Apply(set *flag.FlagSet) error {
	return nil
}

// String return the value for view
func (f *Uint64SliceFlag) String() string {
	return ""
}

// ValidateValues verify if all values was valid
func (f *Uint64SliceFlag) ValidateValues() error {
	return f.validateValues(f.target)
}

// Info returns parsed info of this flag, the returned object must READ-ONLY
func (v *Uint64SliceFlag) Info() *impl.FlagInfo {
	return &v.info
}

// Reset clean the last parsed value of this flag
func (f *Uint64SliceFlag) Reset() {
	f.target.Reset()
	f.info.HasBeenSet = false
}

// for default value verify
func (f *Uint64SliceFlag) validateValues(values *Uint64Slice) error {
	for _, val := range values.slice {
		if err := f.validValue(val); err != nil {
			return err
		}
	}
	return nil
}

// check if value if valid for this flag
func (f *Uint64SliceFlag) validValue(value uint64) error {
	if len(f.Enums) > 0 {
		found := false
		for _, v := range f.Enums {
			if value == v {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("flag %s value %v out of Enums: %v", f.info.LogicName, value, f.Enums)
		}
	}
	if len(f.Ranges) > 0 {
		found := false
		for i := 0; i < len(f.Ranges); i += 2 {
			min, max := f.Ranges[i], f.Ranges[i+1]
			if value >= min && value < max {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("flag %s value %v out of Ranges: %v", f.info.LogicName, value, f.Enums)
		}
	}
	return nil
}

// // Uint64Slice looks up the value of a local Uint64SliceFlag, returns
// // nil if not found
// func (c *Context) Uint64Slice(name string) []uint64 {
// 	if fs := lookupFlagSet(name, c); fs != nil {
// 		return lookupUint64Slice(name, fs)
// 	}
// 	return nil
// }

// func lookupUint64Slice(name string, set *flag.FlagSet) []uint64 {
// 	f := set.Lookup(name)
// 	if f != nil {
// 		if slice, ok := f.Value.(*Uint64Slice); ok {
// 			return slice.Value()
// 		}
// 	}
// 	return nil
// }

var _ impl.Flag = (*Uint64SliceFlag)(nil) //for interface verification only
