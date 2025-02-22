package parser

import (
	"golox/src/ast"
	"golox/src/lexer"
)

func (parser *Parser) statement() (ast.Stmt, error) {
	// --- print statement
	if parser.matches(lexer.PRINT) {
		return parser.printStatement()
	}

	// --- parse regular statement
	return parser.expressionStatement()
}

func (parser *Parser) printStatement() (ast.Stmt, error) {
	expr, err := parser.expression()
	if err != nil {
		return nil, err
	}

	// if the next token is not a semicolon, this is an invalid expression
	if !parser.matches(lexer.SEMICOLON) {
		return nil, NewParsingError(parser.peek(), "invalid token: expected ';'")
	}

	return ast.NewPrintStatement(expr), nil
}

func (parser *Parser) expressionStatement() (ast.Stmt, error) {
	expr, err := parser.expression()
	if err != nil {
		return nil, err
	}

	// if the next token is not a semicolon, this is an invalid expression
	if !parser.matches(lexer.SEMICOLON) {
		return nil, NewParsingError(parser.peek(), "invalid token: expected ';'")
	}

	return ast.NewExpressionStatement(expr), nil
}
