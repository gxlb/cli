package impl

import (
	"flag"
	"fmt"
	"strings"
	"time"

	"cli/internal/util"
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
	fmt.Stringer                              // Show flag help info
	Init(namegen *util.NameGenenerator) error // init parsing of this flag
	Apply(*flag.FlagSet) error                // Apply Flag settings to the given flag set
	IsSet() bool                              // check if the flag value was set
	Info() *FlagInfo                          // parsed info of this flag
	Reset()                                   // reset the flag value
}

// FlagInfo is parsed info of a flag
type FlagInfo struct {
	DispName    string   // DispName is display name of the flag
	Name        string   // Name of this flag(auto generate if not defined)
	LogicName   string   // logic name of the flag
	Names       []string // name+aliases of the flag
	Usage       string   // usage string
	Required    bool     // if required
	Hidden      bool     // hidden this flag
	EnvVars     []string // environment values
	FilePath    string   // file path
	Flag        Flag     // value reference of this flag
	HasBeenSet  bool     // if the value was set
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

// FlagLogicName returns logic name if set(default "value")
func FlagLogicName(logicName string) string {
	if logicName != "" {
		return logicName
	}
	return defaultPlaceholder
}

// FlagDispName returns display name of a flag
func FlagDispName(name, logicName string) string {
	if logicName != "" {
		return logicName
	}
	return name
}
