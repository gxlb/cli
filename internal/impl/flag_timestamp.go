package impl

import (
	"fmt"
	"time"
)

// Timestamp wrap to satisfy golang's flag interface.
type TimestampValue struct {
	timestamp  *time.Time
	hasBeenSet bool
	layout     string
}

// Timestamp constructor
func NewTimestampValue(timestamp time.Time) *TimestampValue {
	return &TimestampValue{timestamp: &timestamp}
}

// Set the timestamp value directly
func (t *TimestampValue) SetTimestamp(value time.Time) {
	if !t.hasBeenSet {
		t.timestamp = &value
		t.hasBeenSet = true
	}
}

// Set the timestamp string layout for future parsing
func (t *TimestampValue) SetLayout(layout string) {
	t.layout = layout
}

// Parses the string value to timestamp
func (t *TimestampValue) Set(value string) error {
	timestamp, err := time.Parse(t.layout, value)
	if err != nil {
		return err
	}

	t.timestamp = &timestamp
	t.hasBeenSet = true
	return nil
}

// String returns a readable representation of this value (for usage defaults)
func (t *TimestampValue) String() string {
	return fmt.Sprintf("%#v", t.timestamp)
}

// Value returns the timestamp value stored in the flag
func (t *TimestampValue) Value() *time.Time {
	return t.timestamp
}

// Get returns the flag structure
func (t *TimestampValue) Get() interface{} {
	return *t
}
