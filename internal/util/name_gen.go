package util

import (
	"fmt"
)

const defaultNamePrefix = "noname"

// NameGenenerator generate sequential names
type NameGenenerator struct {
	id     uint64
	prefix string
}

// NewNameGenerator make a NameGenenerator with name prefix
func NewNameGenerator(prefix string) *NameGenenerator {
	if prefix == "" {
		prefix = defaultNamePrefix
	}
	return &NameGenenerator{
		id:     0,
		prefix: prefix,
	}
}

// Reset initialize the next name sequence
func (g *NameGenenerator) Reset() {
	g.id = 0
}

// NextName generate next sequential name
func (g *NameGenenerator) NextName() string {
	g.id++
	return fmt.Sprintf("{%s#%d}", g.prefix, g.id)
}

//GetOrGenName returns the exists name or generate a new one
func (g *NameGenenerator) GetOrGenName(name string) string {
	if name != "" {
		return name
	}
	return g.NextName()
}
