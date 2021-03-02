// MIT License
//
// Copyright (c) 2021 Ally Dale <vipally@gamil.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package gp

//#GOGP_FILE_BEGIN
//#GOGP_IGNORE_BEGIN ///gogp_file_begin
//
/*   //This line can be uncommented to disable all this file, and it doesn't effect to the .gp file
//	 //If test or change .gp file required, comment it to modify and compile as normal go file
//
// This is a fake go code file
// It is used to generate .gp file by gogp tool
// Real go code file will be generated from .gp file
//
//#GOGP_IGNORE_END ///gogp_file_begin

import (
	"flag"
	"fmt"

	"cli/internal/impl"
	"cli/internal/util"

	//#GOGP_IFDEF SLICE_TYPE
	"encoding/json"
	"strconv"
	"strings"
	//#GOGP_ENDIF //SLICE_TYPE
)

//#GOGP_REQUIRE(github.com/gxlb/gogp/lib/fakedef,_)
//#GOGP_IGNORE_BEGIN ///require begin from(github.com/gxlb/gogp/lib/fakedef)
//#GOGP_IFDEF GOGP_DO_NOT_HAS_THIS_DEFINE__
//this is a fake types for other gp file
//#GOGP_ENDIF

//these defines are used to make sure this fake go file can be compiled correctlly
//and they will be removed from real go files
//vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv
//#GOGP_IFDEF KEY_TYPE
type GOGPKeyType int                             //
func (this GOGPKeyType) Less(o GOGPKeyType) bool { return this < o }
func (this GOGPKeyType) Show() string            { return "" } //
//#GOGP_ENDIF //KEY_TYPE

type GOGPValueType int                               //
func (this GOGPValueType) Less(o GOGPValueType) bool { return this < o }
func (this GOGPValueType) Show() string              { return "" } //
//^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
//#GOGP_IGNORE_END ///require end from(github.com/gxlb/gogp/lib/fakedef)

//#GOGP_IGNORE_BEGIN //fake defines
////////////////////////////////////////////////////////////////////////////////
type GOGPREPElemType = GOGPGlobalNamePrefix
type GOGPREPRawElemType = GOGPGlobalNamePrefix
type Context int
type GOGPREValueType = []GOGPValueType

var GOGPREPSingleValue GOGPValueType
var GOGPREPSliceValue GOGPGlobalNamePrefix

func GOGPREPParseString(string) (a GOGPValueType, e error) { return }

////////////////////////////////////////////////////////////////////////////////
//#GOGP_IGNORE_END //fake defines

//#GOGP_IFDEF SLICE_TYPE

// GOGPGlobalNamePrefixSlice wraps []GOGPValueType to satisfy flag.Value
type GOGPGlobalNamePrefix struct {
	slice      []GOGPValueType
	hasBeenSet bool
}

// NewGOGPGlobalNamePrefix makes an *NewGOGPGlobalNamePrefix with default values
func NewGOGPGlobalNamePrefix(defaults ...GOGPValueType) *GOGPGlobalNamePrefix {
	return &GOGPGlobalNamePrefix{
		slice:      append([]GOGPValueType{}, defaults...),
		hasBeenSet: false,
	}
}

// clone allocate a copy of self object
func (s *GOGPGlobalNamePrefix) clone() *GOGPGlobalNamePrefix {
	n := &GOGPGlobalNamePrefix{
		slice:      append([]GOGPValueType{}, s.slice...),
		hasBeenSet: s.hasBeenSet,
	}
	return n
}

// AppendValues directly append values to the list of values
func (s *GOGPGlobalNamePrefix) AppendValues(values ...GOGPValueType) {
	s.setValues(false, values)
}

// SetValues directly overite values to the list of values
func (s *GOGPGlobalNamePrefix) SetValues(values ...GOGPValueType) {
	s.setValues(true, values)
}

func (s *GOGPGlobalNamePrefix) setValues(overwrite bool, values []GOGPValueType) {
	if !s.hasBeenSet || overwrite {
		s.Reset()
		s.hasBeenSet = true
	}

	s.slice = append(s.slice, values...)
}

// Set parses the value and appends to the list of values
func (s *GOGPGlobalNamePrefix) Set(value string) error {

	if strings.HasPrefix(value, impl.SerializedPrefix) {
		// Deserializing assumes overwrite
		_ = json.Unmarshal([]byte(strings.Replace(value, impl.SerializedPrefix, "", 1)), &s.slice)
		s.hasBeenSet = true
		return nil
	}

	//accept multi values for slice flags
	for _, val := range impl.FlagSplitMultiValues(value) {
		value := strings.TrimSpace(val)
		tmp, err := GOGPREPParseString(value)
		if err != nil {
			return err
		}

		if !s.hasBeenSet {
			s.slice = []GOGPValueType{}
			s.hasBeenSet = true
		}

		s.slice = append(s.slice, GOGPValueType(tmp))
	}

	return nil
}

// Reset clean the last parsed values of this slice
func (s *GOGPGlobalNamePrefix) Reset() {
	if s.slice == nil {
		s.slice = []GOGPValueType{}
	} else {
		s.slice = s.slice[:0]
	}
	s.hasBeenSet = false
}

// String returns a readable representation of this value (for usage defaults)
func (s *GOGPGlobalNamePrefix) String() string {
	return fmt.Sprintf("%#v", s.slice)
}

// Serialize allows GOGPGlobalNamePrefixSlice to fulfill Serializer
func (s *GOGPGlobalNamePrefix) Serialize() string {
	jsonBytes, _ := json.Marshal(s.slice)
	return fmt.Sprintf("%s%s", impl.SerializedPrefix, string(jsonBytes))
}

// Value returns the slice of ints set by this flag
func (s *GOGPGlobalNamePrefix) Value() []GOGPValueType {
	return s.slice
}

// Get returns the slice set by this flag
func (s *GOGPGlobalNamePrefix) Get() interface{} {
	return *s
}

//#GOGP_REPLACE(*GOGPREPElemType, *GOGPGlobalNamePrefix)
//#GOGP_REPLACE(GOGPREPElemType, *GOGPGlobalNamePrefix)
//#GOGP_REPLACE(GOGPREPParseString(value), GOGPParseString)
//#GOGP_REPLACE(GOGPREPSliceValue, f.target)
//#GOGP_REPLACE(GOGPREPRawElemType, GOGPGlobalNamePrefix)
//#GOGP_REPLACE(GOGPREValueType, []GOGPValueType)

//#GOGP_ELSE //SLICE_TYPE

//#GOGP_REPLACE(GOGPREPSingleValue, values)
//#GOGP_REPLACE(GOGPREPElemType, GOGPValueType)
//#GOGP_REPLACE(GOGPREPRawElemType, GOGPValueType)
//#GOGP_REPLACE(GOGPREValueType, GOGPValueType)

//#GOGP_ENDIF //SLICE_TYPE

// GOGPGlobalNamePrefixFlag define a value of type GOGPREPElemType
type GOGPGlobalNamePrefixFlag struct {
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
	Target      *GOGPREPElemType // Target value pointer outside
	Default     GOGPREPElemType  // Default value
	DefaultText string           // Default value display in help info
	Enums       []GOGPValueType  // Enumeration of valid values
	Ranges      []GOGPValueType  // {[min,max),[min,max),...} ranges of valid values
	//
	////////////////////////////////////////////////////////////////////////////
	//area for parsing
	target *GOGPREPElemType // target value pointer(maybe new(GOGPREPElemType) if Target not set)
	info   impl.FlagInfo    // parsed info of this flag
}

// init verify and init the flag info
func (f *GOGPGlobalNamePrefixFlag) init(namegen *util.NameGenenerator) error {
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
		//#GOGP_IFDEF SLICE_TYPE
		f.target = NewGOGPGlobalNamePrefix()
		//#GOGP_ELSE
		f.target = new(GOGPREPRawElemType)
		//#GOGP_ENDIF //SLICE_TYPE
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
			var filter = make(map[GOGPValueType]struct{})
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
func (f *GOGPGlobalNamePrefixFlag) IsSet() bool {
	//#GOGP_IFDEF SLICE_TYPE
	return f.target.hasBeenSet
	//#GOGP_ELSE
	return f.info.HasBeenSet
	//#GOGP_ENDIF //SLICE_TYPE
}

//GetLogicName returns the logic name of the falg
func (f *GOGPGlobalNamePrefixFlag) GetLogicName() string {
	return f.info.LogicName
}

//GetValueName returns the value name of the falg
func (f *GOGPGlobalNamePrefixFlag) GetValueName() string {
	return f.info.ValueName
}

// Names returns the names of the flag
func (f *GOGPGlobalNamePrefixFlag) Names() []string {
	return f.info.Names
}

// IsRequired returns whether or not the flag is required
func (f *GOGPGlobalNamePrefixFlag) IsRequired() bool {
	return f.Required
}

// TakesValue returns true of the flag takes a value, otherwise false
func (f *GOGPGlobalNamePrefixFlag) TakesValue() bool {
	return false
}

// GetUsage returns the usage string for the flag
func (f *GOGPGlobalNamePrefixFlag) GetUsage() string {
	return f.info.Usage
}

// GetValue returns the flags value as string representation.
func (f *GOGPGlobalNamePrefixFlag) GetValue() string {
	return ""
}

// Apply coordinate the value to flagset
func (f *GOGPGlobalNamePrefixFlag) Apply(set *flag.FlagSet) error {
	return nil
}

// String return the value for view
func (f *GOGPGlobalNamePrefixFlag) String() string {
	return ""
}

// ValidateValues verify if all values are valid
func (f *GOGPGlobalNamePrefixFlag) ValidateValues() error {
	//#GOGP_IFDEF SLICE_TYPE
	return f.validateValues(GOGPREPSliceValue)
	//#GOGP_ELSE
	return f.validateValues(*f.target)
	//#GOGP_ENDIF //SLICE_TYPE
}

// Info returns parsed info of this flag, the returned object must READ-ONLY
func (v *GOGPGlobalNamePrefixFlag) Info() *impl.FlagInfo {
	return &v.info
}

// Reset clean the last parsed value of this flag
func (f *GOGPGlobalNamePrefixFlag) Reset() {
	//#GOGP_IFDEF SLICE_TYPE
	f.target.Reset()
	//#GOGP_ELSE
	var t GOGPREPElemType
	*f.target = t
	//#GOGP_ENDIF //SLICE_TYPE
	f.info.HasBeenSet = false
}

// for default value verify
func (f *GOGPGlobalNamePrefixFlag) validateValues(values GOGPREPElemType) error {
	//#GOGP_IFDEF SLICE_TYPE
	for _, val := range values.slice {
		if err := f.validValue(val); err != nil {
			return err
		}
	}
	return nil
	//#GOGP_ELSE
	return f.validValue(GOGPREPSingleValue)
	//#GOGP_ENDIF //SLICE_TYPE
}

// check if value if valid for this flag
func (f *GOGPGlobalNamePrefixFlag) validValue(value GOGPValueType) error {
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

// // GOGPGlobalNamePrefix looks up the value of a local GOGPGlobalNamePrefixFlag
// func (c *Context) GOGPGlobalNamePrefix(name string) GOGPREValueType {
// 	if fs := lookupFlagSet(name, c); fs != nil {
// 		return lookupGOGPREPRawElemType(name, fs)
// 	}
// 	return nil
// }

// func lookupGOGPGlobalNamePrefix(name string, set *flag.FlagSet) GOGPREValueType {
// 	f := set.Lookup(name)
// 	if f != nil {
// 		if slice, ok := f.Value.(*GOGPGlobalNamePrefix); ok {
// 			return slice.Value()
// 		}
// 	}
// 	return nil
// }

var _ impl.Flag = (*GOGPGlobalNamePrefixFlag)(nil) //for interface verification only
//#GOGP_IFDEF SLICE_TYPE
var _ = (*strconv.NumError)(nil) //avoid compile error
//#GOGP_ENDIF //SLICE_TYPE

//#GOGP_FILE_END
//#GOGP_IGNORE_BEGIN ///gogp_file_end
//*/
//#GOGP_IGNORE_END ///gogp_file_end
