// Package object defines what the value of parsed tokens will evaluate to in
// the Monkey language.
package object

// ObjectType represents a value.  All types are represented as Objects.
type ObjectType string

// Object represents all values.
type Object interface {
	Type() ObjectType
	Inspect() string
}
