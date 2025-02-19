package executor

import (
	"golox/src/ast"
	"golox/src/lexer"
)

// main executor function
func ExecuteAST(expr ast.Expr) (any, error) {
	return exec(expr)
}

func exec(expr ast.Expr) (any, error) {
	switch e := expr.(type) {
	case *ast.Binary:
		return execBinary(*e)
	case *ast.Unary:
		return execUnary(*e)
	case *ast.Grouping:
		return execGrouping(*e)
	case *ast.Literal:
		return execLiteral(*e)
	}

	panic("unreachable")
}

func execBinary(expr ast.Binary) (any, error) {
	left, left_err := exec(expr.Left)
	if left_err != nil {
		return nil, left_err
	}
	right, right_err := exec(expr.Right)
	if right_err != nil {
		return nil, right_err
	}

	switch expr.Operator.TokenType() {
	// arithmetic operators
	case lexer.PLUS:
		return plus(left, right)
	case lexer.MINUS:
		return minus(left, right)
	case lexer.STAR:
		return multiply(left, right)
	case lexer.SLASH:
		return divide(left, right)

	// comparison operators
	case lexer.LESS:
		return lt(left, right)
	case lexer.LESS_EQUAL:
		return lte(left, right)
	case lexer.GREATER:
		return gt(left, right)
	case lexer.GREATER_EQUAL:
		return gte(left, right)
	case lexer.EQUAL_EQUAL:
		return equal(left, right)
	case lexer.BANG_EQUAL:
		return !equal(left, right)
	}

	panic("unreachable")
}

func execLiteral(expr ast.Literal) (any, error) {
	return expr.Value, nil
}

func execGrouping(expr ast.Grouping) (any, error) {
	return exec(expr.Expression)
}

func execUnary(expr ast.Unary) (any, error) {
	child, child_err := exec(expr.Expression)
	if child_err != nil {
		return nil, child_err
	}

	switch expr.Operator.TokenType() {
	case lexer.MINUS:
		// convert expression to float and invert
		exprAsFloat, ok := child.(float64)
		if !ok {
			return nil, NewRuntimeError(expr.Operator, "unary operator should be a number")
		}
		return -exprAsFloat, nil

	case lexer.BANG:
		return !isTruthy(child), nil
	}

	panic("unreachable")
}

// consider falsy to be only <nil> or false
func isTruthy(val any) bool {
	if val == nil {
		return false
	}

	valAsBool, ok := val.(bool)
	if ok {
		return valAsBool
	}

	return true
}
