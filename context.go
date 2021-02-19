package cli

import (
	"context"
	"flag"
)

// Context is a type that is passed through to
// each Handler action in a cli application. Context
// can be used to retrieve context-specific args and
// parsed command-line options.
type Context struct {
	context.Context
	App           *App
	Command       *Command
	shellComplete bool
	flagSet       *flag.FlagSet
	parentContext *Context
}
