// Package object defines what the value of parsed tokens will evaluate to in
// the Monkey language.
package object

import "fmt"

// List of different objects supported in Monkey.
const (
	INTEGER_OBJ = "INTEGER"
)

// ObjectType represents a value.  All types are represented as Objects.
type ObjectType string

// Object represents all values.
type Object interface {
	Type() ObjectType
	Inspect() string
}

// Integer represents all integers.
type Integer struct {
	Value int64
}

// Type returns the type of Integer.
func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

// Inspect returns a string representation of the Integer type.
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}
