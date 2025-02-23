package executor

import (
	"golox/src/lexer"
	"strconv"
)

func handleArithmetic(op lexer.Token, left, right any) (any, error) {
	l, l_ok := left.(float64)
	if !l_ok {
		return nil, NewRuntimeError(op, "left side of arithmetic operation is not a number")
	}

	r, r_ok := right.(float64)
	if !r_ok {
		return nil, NewRuntimeError(op, "right side of arithmetic operation is not a number")
	}

	switch op.TokenType() {
	case lexer.MINUS:
		return l - r, nil
	case lexer.STAR:
		return l * r, nil
	case lexer.SLASH:
		if r == 0 {
			return nil, NewRuntimeError(op, "right side of division can not be is zero")
		}
		return l / r, nil
	}

	panic("unreachable")
}

func handleComparison(op lexer.Token, left, right any) (any, error) {
	l, l_ok := left.(float64)
	if !l_ok {
		return nil, NewRuntimeError(op, "left side of comparison operation is not a number")
	}

	r, r_ok := right.(float64)
	if !r_ok {
		return nil, NewRuntimeError(op, "right side of comparison operation is not a number")
	}

	switch op.TokenType() {
	case lexer.LESS:
		return l < r, nil
	case lexer.LESS_EQUAL:
		return l <= r, nil
	case lexer.GREATER:
		return l > r, nil
	case lexer.GREATER_EQUAL:
		return l >= r, nil
	}

	panic("unreachable")
}

func handleEquality(op lexer.Token, left, right any) (any, error) {
	switch l := left.(type) {
	case bool:
		if r, ok := right.(bool); ok {
			switch op.TokenType() {
			case lexer.EQUAL_EQUAL:
				return l == r, nil
			case lexer.BANG_EQUAL:
				return l != r, nil
			}
		}
	case float64:
		if r, ok := right.(float64); ok {
			switch op.TokenType() {
			case lexer.EQUAL_EQUAL:
				return l == r, nil
			case lexer.BANG_EQUAL:
				return l != r, nil
			}
		}
	}

	return nil, NewRuntimeError(op, "sides of equality operation need to be both booleans or numbers")
}

func handlePlus(op lexer.Token, left, right any) (any, error) {
	switch l := left.(type) {
	case float64:
		// if right is not a number, error
		if r, ok := right.(float64); ok {
			return l + r, nil
		}

	case string:
		// if right is not a string, error
		if r, ok := right.(string); ok {
			return l + r, nil
		}

		// if right is a number, attempt to convert it to a float
		switch r := right.(type) {
		case int:
			rightAsString := strconv.Itoa(r)
			return l + rightAsString, nil
		case float64:
			rightAsFloat := strconv.FormatFloat(r, 'f', -1, 64)
			return l + rightAsFloat, nil
		}

	default:
		return nil, NewRuntimeError(op, "left side of addition must either be a number or a string")
	}

	return nil, NewRuntimeError(op, "right side of addition must either be a number or a string")
}
