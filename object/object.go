// Package object defines what the value of parsed tokens will evaluate to in
// the Monkey language.
package object

import "fmt"

// List of different objects supported in Monkey.
const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ    = "NULL"
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

// Boolean represents true or false values.
type Boolean struct {
	Value bool
}

// Type returns the type of Boolean
func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}

// Inspect returns a string representation of the Boolean type.
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

// Null represents a missing value.  As mentioned in the text, the reason for
// defining Null in the first place isn't so we can use it.  It's so that we
// think twice before using it.
type Null struct{}

// Type returns the Null type
func (n *Null) Type() ObjectType {
	return NULL_OBJ
}

// Inspect returns the string representation of the Null type.
func (n *Null) Inspect() string {
	return "null"
}
