package util

import (
	"fmt"
)

// NameFilter filt duplicate names
type NameFilter struct {
	mp map[string]struct{}
}

// NewNameFilter creates new NameFilter object
func NewNameFilter() *NameFilter {
	return &NameFilter{
		mp: make(map[string]struct{}),
	}
}

// Add try to add a new name, it return false if name already exists
func (f *NameFilter) Add(name string) bool {
	return f.add(name, false)
}

// AddBatch try to add a batch name, it return error if one name already exists
func (f *NameFilter) AddBatch(names []string) error {
	for _, name := range names {
		if !f.Add(name) {
			return fmt.Errorf(`duplicate <%s>`, name)
		}
	}
	return nil
}

// Exists check if name was already exists
func (f *NameFilter) Exists(name string) bool {
	return !f.add(name, true)
}

// Reset cleanup the name set
func (f *NameFilter) Reset() {
	f.mp = make(map[string]struct{})
}

// Count return the number of names in set
func (f *NameFilter) Count() int {
	return len(f.mp)
}

func (f *NameFilter) add(name string, checkOnly bool) bool {
	if _, ok := f.mp[name]; !ok {
		if !checkOnly {
			f.mp[name] = struct{}{}
		}
		return true
	}
	return false
}

// FiltNames checks if names has duplicate ones
func FiltNames(names []string) error {
	var f NameFilter
	f.Reset()
	return f.AddBatch(names)
}
