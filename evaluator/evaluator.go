// Package evaluator provides evaluation step of turning ast.Nodes into values.
package evaluator

import (
	"monkey/ast"
	"monkey/object"
)

var (
	// Create "singleton" boolean values.  The book explains it as a small
	// performance improvement since we only declare the boolean objects once.
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
	NULL  = &object.Null{}
)

// Eval evaluates a node and returns the node's value or traverses to the next
// expression to be evaluated.
func Eval(node ast.Node) object.Object {
	switch node := node.(type) {

	case *ast.Program:
		return evalStatements(node.Statements)

	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	case *ast.Boolean:
		if node.Value {
			return TRUE
		}
		return FALSE

	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)
	}

	return nil
}

// evalStatements recursively evaluates each statement.  It returns the value of
// the last statement in the program.
func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range stmts {
		result = Eval(statement)
	}

	return result
}

// evalPrefixExpression evaluates bang (!) and minus (-) operators.  Currently,
// it only supports (!).
func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	default:
		return NULL
	}
}

// evalBangOperatorExpression defines the behavior of the bang (!) operator.
func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}
