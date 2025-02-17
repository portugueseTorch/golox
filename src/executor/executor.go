package executor

import (
	"golox/src/ast"
	"golox/src/lexer"
)

// main executor function
func ExecuteAST(expr ast.Expr) any {
	return exec(expr)
}

func exec(expr ast.Expr) any {
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

func execBinary(expr ast.Binary) any {
	left := exec(expr.Left)
	right := exec(expr.Right)

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

func execLiteral(expr ast.Literal) any {
	return expr.Value
}

func execGrouping(expr ast.Grouping) any {
	return exec(expr.Expression)
}

func execUnary(expr ast.Unary) any {
	child := exec(expr.Expression)

	switch expr.Operator.TokenType() {
	case lexer.MINUS:
		// convert expression to float and invert
		exprAsFloat, ok := child.(float64)
		if !ok {
			panic("TODO: graceful error handling")
		}

		return -exprAsFloat
	case lexer.BANG:
		return !isTruthy(child)
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
