package cli

import (
	"cli/internal/impl"
)

// Flag is interface of a flag
type Flag = impl.Flag

// FlagInfo is parsed info of a flag
type FlagInfo = impl.FlagInfo

// Timestamp wrap to satisfy golang's flag interface.
type TimestampValue = impl.TimestampValue

// GenericValue is a generic parseable type identified by a specific flag
type GenericValue = impl.GenericValue

// GenericValue is alias of time.Duration
type DurationValue = impl.DurationValue

// NewTimestampValue is Timestamp constructor
var NewTimestampValue = impl.NewTimestampValue
