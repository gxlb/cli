package cli

import (
	"flag"
	"fmt"
	"strings"
)

const maxSliceLen = 100 //define max slice len to avoid low effect range of slice

// Flag
type Flag interface {
	fmt.Stringer               // Show flag help info
	Init() error               // init parsing of this flag
	Apply(*flag.FlagSet) error // Apply Flag settings to the given flag set
	IsSet() bool               // check if the flag value was set
	Info() *FlagInfo           // parsed info of this flag
	Reset()                    // reset the flag value
}

// FlagInfo
type FlagInfo struct {
	LogicName  string   // logic name of the flag
	Names      []string // name+aliases of the flag
	Usage      string   // usage string
	Required   bool     // if required
	Hidden     bool     // hidden this flag
	EnvVars    []string // environment values
	FilePath   string   // file path
	Flag       Flag     // value reference of this flag
	HasBeenSet bool     // if the value was set
}

func flagSplitMultiValues(val string) []string {
	return strings.Split(val, ",")
}
