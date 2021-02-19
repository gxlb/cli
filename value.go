package cli

import (
	"flag"
	"fmt"
)

type Value interface {
	fmt.Stringer
	Apply(*flag.FlagSet) error // Apply Flag settings to the given flag set
	IsSet() bool
}
