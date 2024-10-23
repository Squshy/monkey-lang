package evaluator

import (
	"monkey/ast"
	"monkey/object"
)

var (
	// Only keep one reference to a boolean object.
	// All trues will be the same, same with all falses
	// We don't need to create a new `object.Boolean` each time we encounter
	// a boolean value, we can just reference these two.
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
	NULL  = &object.Null{}
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	// Statements
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	// Expressions
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		if node.Value == true {
			return TRUE
		}

		return FALSE
	}

	return nil
}

func evalStatements(statements []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range statements {
		result = Eval(statement)
	}

	return result
}
