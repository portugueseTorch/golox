package executor

import (
	"fmt"
	"golox/src/ast"
	"golox/src/lexer"
)

func (exec *Executor) execCall(call *ast.Call) (any, error) {
	callee, err := exec.execExpr(call.Callee)
	if err != nil {
		return nil, err
	}

	// --- if callee is not callable, runtime error
	callable, ok := callee.(Callable)
	if !ok {
		return nil, NewRuntimeError(call.Paren, "expression is not callable")
	}

	// --- check arity
	if callable.arity() != len(call.Args) {
		return nil, NewRuntimeError(call.Paren, fmt.Sprintf("invalid number of arguments: expected %d but got %d\n", callable.arity(), len(call.Args)))
	}

	args := make([]any, 0)
	for _, arg := range call.Args {
		val, err := exec.execExpr(arg)
		if err != nil {
			return nil, err
		}

		args = append(args, val)
	}

	return callable.call(exec, args)
}

func (exec *Executor) execAssignment(expr *ast.Assignment) (any, error) {
	value, err := exec.execExpr(expr.Value)
	if err != nil {
		return nil, err
	}

	level, ok := exec.locals[expr]
	if ok {
		return exec.setAt(level, expr.Name, value)
	} else {
		return exec.env.Assign(expr.Name, value)
	}
}

func (exec *Executor) execLogical(expr *ast.Logical) (any, error) {
	left, left_err := exec.execExpr(expr.Left)
	if left_err != nil {
		return nil, left_err
	}

	// short circuit, if appropriate
	leftIsTruthy := isTruthy(left)
	if expr.Operator.TokenType() == lexer.AND && !leftIsTruthy {
		return left, nil
	} else if expr.Operator.TokenType() == lexer.OR && leftIsTruthy {
		return left, nil
	}

	return exec.execExpr(expr.Right)
}

func (exec *Executor) execBinary(expr *ast.Binary) (any, error) {
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

func (exec *Executor) execLiteral(expr *ast.Literal) (any, error) {
	return expr.Value, nil
}

func (exec *Executor) execGrouping(expr *ast.Grouping) (any, error) {
	return exec.execExpr(expr.Expression)
}

func (exec *Executor) execUnary(expr *ast.Unary) (any, error) {
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

func (exec *Executor) execVariable(expr *ast.Variable) (any, error) {
	level, ok := exec.locals[expr]
	if ok {
		return exec.getAt(level, expr.Name)
	} else {
		return exec.global.Get(expr.Name)
	}
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
