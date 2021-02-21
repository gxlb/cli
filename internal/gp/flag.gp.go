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
type GOGPREPElemType = GOGPGlobalNamePrefixSlice
type GOGPREPRawElemType = GOGPGlobalNamePrefixSlice

var GOGPREPSingleValue GOGPValueType
var GOGPREPSliceValue GOGPGlobalNamePrefixSlice

func GOGPREPParseString(string) (a GOGPValueType, e error) { return }

////////////////////////////////////////////////////////////////////////////////
//#GOGP_IGNORE_END //fake defines

//#GOGP_IFDEF SLICE_TYPE
var _ = (*strconv.NumError)(nil) //avoid compile error

// GOGPGlobalNamePrefixSlice wraps []GOGPValueType to satisfy flag.Value
type GOGPGlobalNamePrefixSlice struct {
	slice      []GOGPValueType
	hasBeenSet bool
}

// NewGOGPGlobalNamePrefixSlice makes an *GOGPGlobalNamePrefixSlice with default values
func NewGOGPGlobalNamePrefixSlice(defaults ...GOGPValueType) *GOGPGlobalNamePrefixSlice {
	return &GOGPGlobalNamePrefixSlice{
		slice:      append([]GOGPValueType{}, defaults...),
		hasBeenSet: false,
	}
}

// clone allocate a copy of self object
func (s *GOGPGlobalNamePrefixSlice) clone() *GOGPGlobalNamePrefixSlice {
	n := &GOGPGlobalNamePrefixSlice{
		slice:      make([]GOGPValueType, len(s.slice)),
		hasBeenSet: s.hasBeenSet,
	}
	copy(n.slice, s.slice)
	return n
}

// Append directly append values to the list of values
func (s *GOGPGlobalNamePrefixSlice) Append(values ...GOGPValueType) {
	s.setValues(false, values)
}

// Append directly overite values to the list of values
func (s *GOGPGlobalNamePrefixSlice) SetValues(values ...GOGPValueType) {
	s.setValues(true, values)
}

// Append directly adds values to the list of values
func (s *GOGPGlobalNamePrefixSlice) setValues(overwrite bool, values []GOGPValueType) {
	if !s.hasBeenSet || overwrite {
		s.Reset()
		s.hasBeenSet = true
	}

	s.slice = append(s.slice, values...)
}

// Set parses the value and appends it to the list of values
func (s *GOGPGlobalNamePrefixSlice) Set(value string) error {

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
func (s *GOGPGlobalNamePrefixSlice) Reset() {
	if s.slice == nil {
		s.slice = []GOGPValueType{}
	} else {
		s.slice = s.slice[:0]
	}
	s.hasBeenSet = false
}

// String returns a readable representation of this value (for usage defaults)
func (s *GOGPGlobalNamePrefixSlice) String() string {
	return fmt.Sprintf("%#v", s.slice)
}

// Serialize allows GOGPGlobalNamePrefixSlice to fulfill Serializer
func (s *GOGPGlobalNamePrefixSlice) Serialize() string {
	jsonBytes, _ := json.Marshal(s.slice)
	return fmt.Sprintf("%s%s", impl.SerializedPrefix, string(jsonBytes))
}

// Value returns the slice of ints set by this flag
func (s *GOGPGlobalNamePrefixSlice) Value() []GOGPValueType {
	return s.slice
}

// Get returns the slice set by this flag
func (s *GOGPGlobalNamePrefixSlice) Get() interface{} {
	return *s
}

//#GOGP_REPLACE(*GOGPREPElemType, *GOGPGlobalNamePrefixSlice)
//#GOGP_REPLACE(GOGPREPElemType, *GOGPGlobalNamePrefixSlice)
//#GOGP_REPLACE(GOGPREPParseString(value), GOGPParseString)
//#GOGP_REPLACE(GOGPREPSliceValue, v.target)
//#GOGP_REPLACE(GOGPREPRawElemType, GOGPGlobalNamePrefixSlice)

//#GOGP_ELSE //SLICE_TYPE

//#GOGP_REPLACE(GOGPREPSingleValue, values)
//#GOGP_REPLACE(GOGPREPElemType, GOGPValueType)
//#GOGP_REPLACE(GOGPREPRawElemType, GOGPValueType)

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

// Init verify and init the value by ower flag
func (v *GOGPGlobalNamePrefixFlag) Init(namegen *util.NameGenenerator) error {
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
		//#GOGP_IFDEF SLICE_TYPE
		v.target = NewGOGPGlobalNamePrefixSlice()
		//#GOGP_ELSE
		v.target = new(GOGPREPRawElemType)
		//#GOGP_ENDIF //SLICE_TYPE
	}
	return nil
}

// IsSet check if value was set
func (v *GOGPGlobalNamePrefixFlag) IsSet() bool {
	//#GOGP_IFDEF SLICE_TYPE
	return v.target.hasBeenSet
	//#GOGP_ELSE
	return v.info.HasBeenSet
	//#GOGP_ENDIF //SLICE_TYPE
}

// Apply coordinate the value to flagset
func (v *GOGPGlobalNamePrefixFlag) Apply(set *flag.FlagSet) error {
	return nil
}

// String return the value for view
func (v *GOGPGlobalNamePrefixFlag) String() string {
	return ""
}

// ValidateValues verify if all values was valid
func (v *GOGPGlobalNamePrefixFlag) ValidateValues() error {
	//#GOGP_IFDEF SLICE_TYPE
	return v.validateValues(GOGPREPSliceValue)
	//#GOGP_ELSE
	return v.validateValues(*v.target)
	//#GOGP_ENDIF //SLICE_TYPE

}

// Info returns parsed info of this flag
func (v *GOGPGlobalNamePrefixFlag) Info() *impl.FlagInfo {
	return &v.info
}

// Reset clean the last parsed value of this flag
func (v *GOGPGlobalNamePrefixFlag) Reset() {
	//#GOGP_IFDEF SLICE_TYPE
	v.target.Reset()
	//#GOGP_ELSE
	var t GOGPREPElemType
	*v.target = t
	//#GOGP_ENDIF //SLICE_TYPE
	v.info.HasBeenSet = false
}

// for default value verify
func (v *GOGPGlobalNamePrefixFlag) validateValues(values GOGPREPElemType) error {
	//#GOGP_IFDEF SLICE_TYPE
	for _, val := range values.slice {
		if err := v.validValue(val); err != nil {
			return err
		}
	}
	return nil
	//#GOGP_ELSE
	return v.validValue(GOGPREPSingleValue)
	//#GOGP_ENDIF //SLICE_TYPE
}

// check if value if valid for this flag
func (v *GOGPGlobalNamePrefixFlag) validValue(value GOGPValueType) error {
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

var _ impl.Flag = (*GOGPGlobalNamePrefixFlag)(nil) //for interface verification only

//#GOGP_FILE_END
//#GOGP_IGNORE_BEGIN ///gogp_file_end
//*/
//#GOGP_IGNORE_END ///gogp_file_end
