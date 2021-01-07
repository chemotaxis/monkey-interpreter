// Package evaluator provides evaluation step of turning ast.Nodes into values.
package evaluator

import (
	"monkey/ast"
	"monkey/object"
)

// Eval evaluates a node and returns the node's value
func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}

	return nil
}
