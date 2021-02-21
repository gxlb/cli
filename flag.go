package cli

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"syscall"
	"time"

	"cli/internal/namegen"
)

const (
	defaultPlaceholder = "value"
	maxSliceLen        = 100 //define max slice len to avoid low effect range of slice
)

var (
	slPfx           = fmt.Sprintf("sl:::%d:::", time.Now().UTC().UnixNano())
	commaWhitespace = regexp.MustCompile("[, ]+.*")
)

// Flag
type Flag interface {
	fmt.Stringer                                 // Show flag help info
	Init(namegen *namegen.NameGenenerator) error // init parsing of this flag
	Apply(*flag.FlagSet) error                   // Apply Flag settings to the given flag set
	IsSet() bool                                 // check if the flag value was set
	Info() *FlagInfo                             // parsed info of this flag
	Reset()                                      // reset the flag value
}

// FlagInfo
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

func flagFromEnvOrFile(envVars []string, filePath string) (val string, ok bool) {
	for _, envVar := range envVars {
		envVar = strings.TrimSpace(envVar)
		if val, ok := syscall.Getenv(envVar); ok {
			return val, true
		}
	}
	for _, fileVar := range strings.Split(filePath, ",") {
		if data, err := ioutil.ReadFile(fileVar); err == nil {
			return string(data), true
		}
	}
	return "", false
}

func flagSplitMultiValues(val string) []string {
	return strings.Split(val, ",")
}

// merge name+aliases into out
func mergeNames(name string, aliases []string, out *[]string) bool {
	if name != "" { //ignore noname merge
		a := *out
		if len(a) == 0 || a[0] != name {
			*out = append([]string{name}, aliases...)
			return true
		}
	}
	return false
}

func logicName(logicName string) string {
	if logicName != "" {
		return logicName
	}
	return defaultPlaceholder
}
