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
	} else if parser.matches(lexer.FUN) {
		stmt, err = parser.function()
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

func (parser *Parser) function() (ast.Stmt, error) {
	if !parser.matches(lexer.IDENTIFIER) {
		return nil, NewParsingError(parser.peek(), fmt.Sprintf("expected IDENTIFIER, got %s\n", parser.peek().TokenType()))
	}
	name := parser.prev()

	if !parser.matches(lexer.LEFT_PAREN) {
		return nil, NewParsingError(parser.peek(), fmt.Sprintf("expected '(', got %s\n", parser.peek().TokenType()))
	}
	params := make([]lexer.Token, 0)
	for parser.matches(lexer.COMMA) {
		if len(params) >= 255 {
			return nil, NewParsingError(parser.peek(), "functions cannot take more than 255 parameters")
		}

		params = append(params, parser.prev())
	}
	if !parser.matches(lexer.RIGHT_PAREN) {
		return nil, NewParsingError(parser.peek(), fmt.Sprintf("expected ')', got %s\n", parser.peek().TokenType()))
	}

	if !parser.matches(lexer.LEFT_BRACE) {
		return nil, NewParsingError(parser.peek(), fmt.Sprintf("expected '{', got %s\n", parser.peek().TokenType()))
	}
	body, err := parser.block()
	if err != nil {
		return nil, err
	}

	return ast.NewFunctionStatement(name, params, body), nil
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
	} else if parser.matches(lexer.IF) {
		return parser.conditionalStatement()
	} else if parser.matches(lexer.LEFT_BRACE) {
		return parser.blockStatement()
	} else if parser.matches(lexer.WHILE) {
		return parser.whileStatement()
	} else if parser.matches(lexer.FOR) {
		return parser.forStatement()
	}

	// --- parse regular statement
	return parser.expressionStatement()
}

func (parser *Parser) forStatement() (ast.Stmt, error) {
	if !parser.matches(lexer.LEFT_PAREN) {
		return nil, NewParsingError(parser.peek(), fmt.Sprintf("expected '(' but got %s", parser.peek().TokenType()))
	}

	// parse initializer
	var initializer ast.Stmt = nil
	var err error = nil
	if parser.matches(lexer.VAR) {
		initializer, err = parser.variableDeclaration()
	} else if parser.matches(lexer.SEMICOLON) {
		initializer, err = nil, nil
	} else {
		initializer, err = parser.expressionStatement()
	}
	if err != nil {
		return nil, NewParsingError(parser.peek(), fmt.Sprintf("invalid initializer in 'for' loop"))
	}

	// parse condition
	var condition ast.Expr = nil
	if parser.peek().TokenType() == lexer.SEMICOLON {
		condition, err = ast.NewLiteral(true), nil
	} else {
		condition, err = parser.expression()
	}
	if err != nil {
		return nil, err
	}
	if !parser.matches(lexer.SEMICOLON) {
		return nil, NewParsingError(parser.peek(), fmt.Sprintf("expected ';' but got %s", parser.peek().TokenType()))
	}

	// parse increment
	var increment ast.Expr = nil
	if parser.peek().TokenType() != lexer.RIGHT_PAREN {
		increment, err = parser.expression()
	}
	if err != nil {
		return nil, err
	}

	if !parser.matches(lexer.RIGHT_PAREN) {
		return nil, NewParsingError(parser.peek(), fmt.Sprintf("expected ')' but got %s", parser.peek().TokenType()))
	}

	// parse body
	body, err := parser.statement()
	if err != nil {
		return nil, err
	}

	return ast.NewForStatement(initializer, condition, increment, body), nil
}

func (parser *Parser) whileStatement() (ast.Stmt, error) {
	if !parser.matches(lexer.LEFT_PAREN) {
		return nil, NewParsingError(parser.peek(), fmt.Sprintf("expected '(' but got %s", parser.peek().TokenType()))
	}

	condition, err := parser.expression()
	if err != nil {
		return nil, err
	}

	if !parser.matches(lexer.RIGHT_PAREN) {
		return nil, NewParsingError(parser.peek(), fmt.Sprintf("expected ')' but got %s", parser.peek().TokenType()))
	}

	body, err := parser.statement()
	if err != nil {
		return nil, err
	}

	return ast.NewWhileStatement(condition, body), nil
}

func (parser *Parser) conditionalStatement() (ast.Stmt, error) {
	if !parser.matches(lexer.LEFT_PAREN) {
		return nil, NewParsingError(parser.peek(), fmt.Sprintf("expected '(' but got %s", parser.peek().TokenType()))
	}

	condition, err := parser.expression()
	if err != nil {
		return nil, err
	}

	if !parser.matches(lexer.RIGHT_PAREN) {
		return nil, NewParsingError(parser.peek(), fmt.Sprintf("expected ')' but got %s", parser.peek().TokenType()))
	}

	ifBranch, err := parser.statement()
	if err != nil {
		return nil, err
	}

	var elseBranch ast.Stmt = nil
	if parser.matches(lexer.ELSE) {
		var err error
		elseBranch, err = parser.statement()
		if err != nil {
			return nil, err
		}
	}

	return ast.NewConditionalStatement(condition, ifBranch, elseBranch), nil
}

func (parser *Parser) blockStatement() (ast.Stmt, error) {
	statements, err := parser.block()
	if err != nil {
		return nil, err
	}

	return ast.NewBlockStatement(statements), nil
}

func (parser *Parser) block() ([]ast.Stmt, error) {
	statements := make([]ast.Stmt, 0)

	for !parser.isAtEnd() && parser.peek().TokenType() != lexer.RIGHT_BRACE {
		stmt, err := parser.declaration()
		if err != nil {
			return nil, err
		}

		statements = append(statements, stmt)
	}

	// --- if the next token is not '}', error
	if !parser.matches(lexer.RIGHT_BRACE) {
		return nil, NewParsingError(parser.peek(), fmt.Sprintf("expected '}' but got %s", parser.peek().TokenType()))
	}

	return statements, nil
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
