package impl

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

// MaxSliceLen limits max length of a flag config value
const MaxSliceLen = 100

// SerializedPrefix is the string prefix of a flag serialized value
var SerializedPrefix = fmt.Sprintf("sl:::%d:::", time.Now().UTC().UnixNano())

const (
	defaultPlaceholder = "value"
)

// Flag is interface of a flag
type Flag interface {
	fmt.Stringer // Show flag help info

	Apply(*flag.FlagSet) error // Apply Flag settings to the given flag set
	IsSet() bool               // check if the flag value was set

	//init(namegen *util.NameGenenerator) error // init parsing of this flag
	Info() *FlagInfo       // parsed info of this flag, the returned object must READ-ONLY
	Reset()                // reset the flag value
	ValidateValues() error // validate set values
	GetLogicName() string  // GetLogicName returns the logic name of the falg
	GetValueName() string  // GetValueName returns the value name of the falg

	//RequiredFlag
	IsRequired() bool // IsRequired returns whether or not the flag is required

	//DocGenerationFlag
	// TakesValue returns true if the flag takes a value, otherwise false
	TakesValue() bool
	// GetUsage returns the usage string for the flag
	GetUsage() string
	// GetValue returns the flags value as string representation and an empty
	// string if the flag takes no value at all.
	GetValue() string
}

// FlagInfo is parsed info of a flag
type FlagInfo struct {
	LogicName   string   // logic name of the flag
	Name        string   // Name of this flag(auto generate if not defined)
	ValueName   string   // ValueName is name of the flag value
	Names       []string // name+aliases of the flag
	Usage       string   // usage string
	Required    bool     // if required
	Hidden      bool     // hidden this flag
	EnvVars     []string // environment values
	FilePath    string   // file path
	Flag        Flag     // value reference of this flag
	HasBeenSet  bool     // if the value was set
	NonameFlag  bool     // if this is a noname flag
	DefaultText string   // Default value in help info
}

//FlagSplitMultiValues splits multi values separated with comma
func FlagSplitMultiValues(val string) []string {
	return strings.Split(val, ",")
}

// MergeNames merge name+aliases into out
func MergeNames(name string, aliases []string, out *[]string) bool {
	if name != "" { //ignore noname merge
		a := *out
		if len(a) == 0 || a[0] != name {
			*out = append([]string{name}, aliases...)
			return true
		}
	}
	return false
}

// FlagValueName returns value name of a flag, 1:logicName, 2:"value"
func FlagValueName(logicName string) string {
	if logicName != "" {
		return logicName
	}
	return defaultPlaceholder
}

// FlagLogicName returns logic name of a flag, 1:logicName, 2:name
func FlagLogicName(name string, logicName string) string {
	if logicName != "" {
		return logicName
	}
	return name
}

// Serializer is used to circumvent the limitations of flag.FlagSet.Set
type Serializer interface {
	Serialize() string
}
