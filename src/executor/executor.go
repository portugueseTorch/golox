package executor

import (
	"golox/src/ast"
	"golox/src/lexer"
)

// main executor function
func ExecuteAST(expr ast.Expr) {
	ret := exec(expr)
}

func exec(expr ast.Expr) any {
	return true
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

	// TODO: graceful error handling
	panic("UNREACHEABLE")
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
