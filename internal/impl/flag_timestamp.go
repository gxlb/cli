package impl

import (
	"fmt"
	"time"
)

// Timestamp wrap to satisfy golang's flag interface.
type Timestamp struct {
	timestamp  *time.Time
	hasBeenSet bool
	layout     string
}

// Timestamp constructor
func NewTimestamp(timestamp time.Time) *Timestamp {
	return &Timestamp{timestamp: &timestamp}
}

// Set the timestamp value directly
func (t *Timestamp) SetTimestamp(value time.Time) {
	if !t.hasBeenSet {
		t.timestamp = &value
		t.hasBeenSet = true
	}
}

// Set the timestamp string layout for future parsing
func (t *Timestamp) SetLayout(layout string) {
	t.layout = layout
}

// Parses the string value to timestamp
func (t *Timestamp) Set(value string) error {
	timestamp, err := time.Parse(t.layout, value)
	if err != nil {
		return err
	}

	t.timestamp = &timestamp
	t.hasBeenSet = true
	return nil
}

// String returns a readable representation of this value (for usage defaults)
func (t *Timestamp) String() string {
	return fmt.Sprintf("%#v", t.timestamp)
}

// Value returns the timestamp value stored in the flag
func (t *Timestamp) Value() *time.Time {
	return t.timestamp
}

// Get returns the flag structure
func (t *Timestamp) Get() interface{} {
	return *t
}

// Generic is a generic parseable type identified by a specific flag
type Generic interface {
	Set(value string) error
	String() string
}
