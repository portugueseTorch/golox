package executor

import (
	"fmt"
	"golox/src/ast"
	"golox/src/lexer"
)

func (exec *Executor) execBinary(expr ast.Binary) (any, error) {
	left, left_err := exec.execExpr(expr.Left)
	if left_err != nil {
		return nil, left_err
	}
	right, right_err := exec.execExpr(expr.Right)
	if right_err != nil {
		return nil, right_err
	}

	switch expr.Operator.TokenType() {
	// arithmetic operators
	case lexer.MINUS, lexer.STAR, lexer.SLASH:
		return handleArithmetic(expr.Operator, left, right)

	// comparison operators
	case lexer.LESS, lexer.LESS_EQUAL, lexer.GREATER, lexer.GREATER_EQUAL:
		return handleComparison(expr.Operator, left, right)

	// equality operators
	case lexer.EQUAL_EQUAL, lexer.BANG_EQUAL:
		return handleEquality(expr.Operator, left, right)

	// special case
	case lexer.PLUS:
		return handlePlus(expr.Operator, left, right)

	}

	panic("unreachable")
}

func (exec *Executor) execLiteral(expr ast.Literal) (any, error) {
	return expr.Value, nil
}

func (exec *Executor) execGrouping(expr ast.Grouping) (any, error) {
	return exec.execExpr(expr.Expression)
}

func (exec *Executor) execUnary(expr ast.Unary) (any, error) {
	child, child_err := exec.execExpr(expr.Expression)
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

func (exec *Executor) execVariable(expr ast.Variable) (any, error) {
	val, ok := exec.env.Get(expr.Name.Literal())
	if !ok {
		return nil, NewRuntimeError(expr.Name, fmt.Sprintf("undefined variable '%s'", expr.Name.Literal()))
	}

	return val, nil
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
