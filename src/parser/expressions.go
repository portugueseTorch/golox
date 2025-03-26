package parser

import (
	"fmt"
	"golox/src/ast"
	"golox/src/lexer"
	"strconv"
)

func (parser *Parser) expression() (ast.Expr, error) {
	return parser.assignment()
}

func (parser *Parser) assignment() (ast.Expr, error) {
	expr, err := parser.or()
	if err != nil {
		return nil, err
	}

	// if the next token is an equal sign, test for assignment
	if parser.matches(lexer.EQUAL) {
		eq := parser.prev()

		value, err := parser.assignment()
		if err != nil {
			return nil, err
		}

		// if the top level expression is not a variable, this is not a valid assignment
		switch varName := expr.(type) {
		case *ast.Variable:
			assignment := ast.NewAssignment(varName.Name, value)
			return assignment, nil
		default:
			assignmentError := NewParsingError(eq, "invalid assignment operation")
			return nil, assignmentError
		}
	}

	return expr, nil
}

func (parser *Parser) or() (ast.Expr, error) {
	expr, err := parser.and()
	if err != nil {
		return nil, err
	}

	// if next token is an OR, build right side of the expression
	for parser.matches(lexer.OR) {
		op := parser.peek()
		right, err := parser.and()
		if err != nil {
			return nil, err
		}

		expr = ast.NewLogical(expr, op, right)
	}

	return expr, nil
}

func (parser *Parser) and() (ast.Expr, error) {
	expr, err := parser.equality()
	if err != nil {
		return nil, err
	}

	// if next token is an AND, build right side of the expression
	for parser.matches(lexer.AND) {
		op := parser.peek()
		right, err := parser.equality()
		if err != nil {
			return nil, err
		}

		expr = ast.NewLogical(expr, op, right)
	}

	return expr, nil
}

func (parser *Parser) equality() (ast.Expr, error) {
	left, err := parser.comparison()
	if err != nil {
		return nil, err
	}

	// if the next token is a != or ==, evaluate next expression as a comparison
	for parser.matches(lexer.BANG_EQUAL, lexer.EQUAL_EQUAL) {
		operator := parser.prev()
		right, err := parser.comparison()
		if err != nil {
			return nil, err
		}

		left = ast.NewBinary(left, operator, right)
	}

	return left, nil
}

func (parser *Parser) comparison() (ast.Expr, error) {
	left, err := parser.term()
	if err != nil {
		return nil, err
	}

	for parser.matches(lexer.LESS, lexer.LESS_EQUAL, lexer.GREATER, lexer.GREATER_EQUAL) {
		operator := parser.prev()
		right, err := parser.term()
		if err != nil {
			return nil, err
		}

		left = ast.NewBinary(left, operator, right)
	}

	return left, nil
}

func (parser *Parser) term() (ast.Expr, error) {
	left, err := parser.factor()
	if err != nil {
		return nil, err
	}

	for parser.matches(lexer.MINUS, lexer.PLUS) {
		operator := parser.prev()
		right, err := parser.factor()
		if err != nil {
			return nil, err
		}

		left = ast.NewBinary(left, operator, right)
	}

	return left, nil
}

func (parser *Parser) factor() (ast.Expr, error) {
	left, err := parser.unary()
	if err != nil {
		return nil, err
	}

	for parser.matches(lexer.SLASH, lexer.STAR) {
		operator := parser.prev()
		right, err := parser.unary()
		if err != nil {
			return nil, err
		}

		left = ast.NewBinary(left, operator, right)
	}

	return left, nil
}

func (parser *Parser) unary() (ast.Expr, error) {
	// if the next token is a negation operator
	if parser.matches(lexer.BANG, lexer.MINUS) {
		operator := parser.prev()
		expr, err := parser.unary()
		if err != nil {
			return nil, err
		}

		return ast.NewUnary(operator, expr), nil
	}

	return parser.primary()
}

func (parser *Parser) primary() (ast.Expr, error) {
	if parser.matches(lexer.TRUE) {
		return ast.NewLiteral(true), nil
	} else if parser.matches(lexer.FALSE) {
		return ast.NewLiteral(false), nil
	} else if parser.matches(lexer.NIL) {
		return ast.NewLiteral(nil), nil
	} else if parser.matches(lexer.STRING) {
		return ast.NewLiteral(parser.prev().Literal()), nil
	} else if parser.matches(lexer.IDENTIFIER) {
		return ast.NewVariable(parser.prev()), nil
	} else if parser.matches(lexer.NUMBER) {
		num, err := strconv.ParseFloat(parser.prev().Literal(), 64)
		if err != nil {
			return nil, NewParsingError(parser.prev(), "invalid float")
		}
		return ast.NewLiteral(num), nil
	}

	// check if it's a grouping expression
	if parser.matches(lexer.LEFT_PAREN) {
		expr, err := parser.expression()
		if err != nil {
			return nil, err
		}

		// if the next token is not a right paren, something is broken
		if !parser.matches(lexer.RIGHT_PAREN) {
			return nil, NewParsingError(parser.prev(), "expected closing parenthesis ')'")
		}

		return ast.NewGrouping(expr), nil
	}

	return nil, NewParsingError(parser.prev(), "invalid primary expression")
}
