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

//#GOGP_IFDEF SLICE_TYPE
import (
	//"encoding/json"
	"fmt"
	"strconv"
	//"strings"
)

//#GOGP_ENDIF //SLICE_TYPE

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

//#GOGP_IGNORE_BEGIN
type GOGPElemType = GOGPGlobalNamePrefixSlice //
//#GOGP_IGNORE_END

////////////////////////////////////////////////////////////////////////////////

//#GOGP_IFDEF SLICE_TYPE
// GOGPGlobalNamePrefixSlice wraps []GOGPValueType to satisfy flag.Value
type GOGPGlobalNamePrefixSlice struct {
	slice      []GOGPValueType
	hasBeenSet bool
}

// NewGOGPGlobalNamePrefixSlice makes an *GOGPGlobalNamePrefixSlice with default values
func NewGOGPGlobalNamePrefixSlice(defaults ...GOGPValueType) *GOGPGlobalNamePrefixSlice {
	return &GOGPGlobalNamePrefixSlice{slice: append([]GOGPValueType{}, defaults...)}
}

// TODO: Consistently have specific Set function for Int64 and Float64 ?
// Append directly adds an integer to the list of values
func (s *GOGPGlobalNamePrefixSlice) Append(value ...GOGPValueType) {
	if !s.hasBeenSet {
		s.slice = []GOGPValueType{}
		s.hasBeenSet = true
	}

	s.slice = append(s.slice, value...)
}

// Set parses the value into an integer and appends it to the list of values
func (s *GOGPGlobalNamePrefixSlice) Set(value string) error {
	if !s.hasBeenSet {
		s.slice = []GOGPValueType{}
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

	s.slice = append(s.slice, GOGPValueType(tmp))

	return nil
}

// String returns a readable representation of this value (for usage defaults)
func (s *GOGPGlobalNamePrefixSlice) String() string {
	return fmt.Sprintf("%#v", s.slice)
}

// Serialize allows GOGPGlobalNamePrefixSlice to fulfill Serializer
func (s *GOGPGlobalNamePrefixSlice) Serialize() string {
	//TODO:
	// jsonBytes, _ := json.Marshal(s.slice)
	// return fmt.Sprintf("%s%s", slPfx, string(jsonBytes))
	return ""
}

// Value returns the slice of ints set by this flag
func (s *GOGPGlobalNamePrefixSlice) Value() []GOGPValueType {
	return s.slice
}

// Get returns the slice set by this flag
func (s *GOGPGlobalNamePrefixSlice) Get() interface{} {
	return *s
}

//#GOGP_REPLACE(GOGPElemType, GOGPGlobalNamePrefixSlice)

//#GOGP_ELSE //SLICE_TYPE

//#GOGP_REPLACE(GOGPElemType, GOGPValueType)

//#GOGP_ENDIF //SLICE_TYPE

// GOGPGlobalNamePrefixValue define a value of type GOGPElemType
type GOGPGlobalNamePrefixValue struct {
	Value       GOGPElemType   // The value from ENV of files
	Target      *GOGPElemType  // Target set the outer value pointer
	Default     GOGPElemType   // Default value
	DefaultText string         // Default value help info
	Enums       []GOGPElemType // Enumeration of valid values
	Min, Max    GOGPElemType   // [Min,Max] range of the valid values
	hasBeenSet  bool
}

//#GOGP_FILE_END
//#GOGP_IGNORE_BEGIN ///gogp_file_end
//*/
//#GOGP_IGNORE_END ///gogp_file_end
