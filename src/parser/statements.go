package parser

import (
	"fmt"
	"golox/src/ast"
	"golox/src/lexer"
)

func (parser *Parser) declaration() (ast.Stmt, error) {
	if parser.isAtEnd() {
		return nil, nil
	}

	var stmt ast.Stmt = nil
	var err error = nil

	// --- if next token is var, attempt to parse a variable declaration
	if parser.matches(lexer.VAR) {
		stmt, err = parser.variableDeclaration()
	} else {
		stmt, err = parser.statement()
	}

	if err != nil {
		fmt.Printf("%s", err)
		parser.synchronize()
		return nil, nil
	}

	return stmt, nil
}

func (parser *Parser) variableDeclaration() (ast.Stmt, error) {
	// --- if next next token is not an initializer, not a valid variable declaration
	if parser.peek().TokenType() != lexer.IDENTIFIER {
		return nil, NewParsingError(parser.peek(), "invalid variable declaration: expected variable name")
	}

	ident := parser.next()
	var init *ast.Expr = nil
	// --- if next token is an equal, parse expression
	if parser.matches(lexer.EQUAL) {
		expr, err := parser.expression()
		if err != nil {
			return nil, err
		}

		init = &expr
	}

	// if the next token is not a semicolon, this is an invalid expression
	if !parser.matches(lexer.SEMICOLON) {
		return nil, NewParsingError(parser.peek(), "invalid token: expected ';'")
	}

	return ast.NewVariableStatement(ident, init), nil
}

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
