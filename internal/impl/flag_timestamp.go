// MIT License
//
// Copyright (c) 2021 Ally Dale <vipally@gamil.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package impl

import (
	"fmt"
	"time"
)

// DefaultTimestampLayout defines the default time format
const DefaultTimestampLayout = "2006-01-02T15:04:05 MST"

// TimestampString represent a timestamp string with format **DefaultTimestampLayout**
type TimestampString string

// Less judge if t less than other
func (t TimestampString) Less(other TimestampString) bool {
	l, _ := time.Parse(DefaultTimestampLayout, t.String())
	r, _ := time.Parse(DefaultTimestampLayout, other.String())
	return l.Before(r)
}

// Equal judge if t equal with other
func (t TimestampString) Equal(other TimestampString) bool {
	l, _ := time.Parse(DefaultTimestampLayout, t.String())
	r, _ := time.Parse(DefaultTimestampLayout, other.String())
	return l.Equal(r)
}

// String returns the timestamp as string
func (t TimestampString) String() string {
	return string(t)
}

// TimestampValue wrap to satisfy golang's flag interface.
type TimestampValue struct {
	timestamp  time.Time
	hasBeenSet bool
	layout     string
}

// Timestamp constructor
func NewTimestampValue(timestamp time.Time, layout string) *TimestampValue {
	_layout := layout
	if layout == "" {
		_layout = DefaultTimestampLayout
	}
	return &TimestampValue{
		timestamp:  timestamp,
		layout:     _layout,
		hasBeenSet: false,
	}
}

// Set the timestamp value directly
func (t *TimestampValue) SetTimestamp(value time.Time) {
	if !t.hasBeenSet {
		t.timestamp = value
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

	t.timestamp = timestamp
	t.hasBeenSet = true
	return nil
}

// String returns a readable representation of this value (for usage defaults)
func (t *TimestampValue) String() string {
	return fmt.Sprintf("%#v", t.timestamp)
}

// Value returns the timestamp value stored in the flag
func (t *TimestampValue) Value() time.Time {
	return t.timestamp
}

// Get returns the flag structure
func (t *TimestampValue) Get() interface{} {
	return *t
}
