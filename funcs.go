package cli

import (
	"cli/internal/impl"
)

// BashCompleteFunc is an action to execute when the shell completion flag is set
type BashCompleteFunc = impl.BashCompleteFunc

// BeforeFunc is an action to execute before any subcommands are run, but after
// the context is ready if a non-nil error is returned, no subcommands are run
type BeforeFunc = impl.BeforeFunc

// AfterFunc is an action to execute after any subcommands are run, but after the
// subcommand has finished it is run even if Action() panics
type AfterFunc = impl.AfterFunc

// ActionFunc is the action to execute when no subcommands are specified
type ActionFunc = impl.ActionFunc

// CommandNotFoundFunc is executed if the proper command cannot be found
type CommandNotFoundFunc = impl.CommandNotFoundFunc

// OnUsageErrorFunc is executed if an usage error occurs. This is useful for displaying
// customized usage error messages.  This function is able to replace the
// original error messages.  If this function is not set, the "Incorrect usage"
// is displayed and the execution is interrupted.
type OnUsageErrorFunc = impl.OnUsageErrorFunc

// ExitErrHandlerFunc is executed if provided in order to handle exitError values
// returned by Actions and Before/After functions.
type ExitErrHandlerFunc = impl.ExitErrHandlerFunc

// FlagStringFunc is used by the help generation to display a flag, which is
// expected to be a single line.
type FlagStringFunc = impl.FlagStringFunc

// FlagNamePrefixFunc is used by the default FlagStringFunc to create prefix
// text for a flag's full name.
type FlagNamePrefixFunc = impl.FlagNamePrefixFunc

// FlagEnvHintFunc is used by the default FlagStringFunc to annotate flag help
// with the environment variable details.
type FlagEnvHintFunc = impl.FlagEnvHintFunc

// FlagFileHintFunc is used by the default FlagStringFunc to annotate flag help
// with the file path details.
type FlagFileHintFunc = impl.FlagFileHintFunc
