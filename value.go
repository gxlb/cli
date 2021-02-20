package cli

import (
	"flag"
	"fmt"
)

const maxSliceLen = 100 //define max slice len to avoid low effect range of slice

type Value interface {
	fmt.Stringer
	Init(*Flag) error
	Apply(*flag.FlagSet) error // Apply Flag settings to the given flag set
	IsSet() bool
}
