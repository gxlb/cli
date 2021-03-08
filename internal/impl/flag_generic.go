package impl

import (
	"time"
)

// GenericValue is a generic parseable type identified by a specific flag
type GenericValue interface {
	Set(value string) error
	String() string
}

// GenericValue is alias of time.Duration
type DurationValue = time.Duration
