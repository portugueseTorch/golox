package parser

import (
	"golox/src/ast"
	"golox/src/lexer"
	"strconv"
)

/* Expression syntax in BNF
expression     → equality ;
equality       → comparison ( ( "!=" | "==" ) comparison )* ;
comparison     → term ( ( ">" | ">=" | "<" | "<=" ) term )* ;
term           → factor ( ( "-" | "+" ) factor )* ;
factor         → unary ( ( "/" | "*" ) unary )* ;
unary          → ( "!" | "-" ) unary
               | primary ;
primary        → NUMBER | STRING | "true" | "false" | "nil"
               | "(" expression ")" ;
*/

type Parser struct {
	tokens []lexer.Token
	cur    int
}

func NewParser(tokens []lexer.Token) Parser {
	return Parser{
		tokens: tokens,
		cur:    0,
	}
}

func (parser *Parser) expression() ast.Expr {
	return parser.equality()
}

func (parser *Parser) equality() ast.Expr {
	left := parser.comparison()

	// if the next token is a != or ==, evaluate next expression as a comparison
	for parser.matches(lexer.BANG_EQUAL, lexer.EQUAL_EQUAL) {
		operator := parser.prev()
		right := parser.comparison()

		left = ast.NewBinary(left, operator, right)
	}

	return left
}

func (parser *Parser) comparison() ast.Expr {
	left := parser.term()

	for parser.matches(lexer.LESS, lexer.LESS_EQUAL, lexer.GREATER, lexer.GREATER_EQUAL) {
		operator := parser.prev()
		right := parser.term()

		left = ast.NewBinary(left, operator, right)
	}

	return left
}

func (parser *Parser) term() ast.Expr {
	left := parser.factor()

	for parser.matches(lexer.MINUS, lexer.PLUS) {
		operator := parser.prev()
		right := parser.factor()

		left = ast.NewBinary(left, operator, right)
	}

	return left
}

func (parser *Parser) factor() ast.Expr {
	left := parser.unary()

	for parser.matches(lexer.SLASH, lexer.STAR) {
		operator := parser.prev()
		right := parser.unary()

		left = ast.NewBinary(left, operator, right)
	}

	return left
}

func (parser *Parser) unary() ast.Expr {
	// if the next token is a negation operator
	if parser.matches(lexer.BANG, lexer.MINUS) {
		operator := parser.prev()
		expr := parser.unary()

		return ast.NewUnary(operator, expr)
	}

	return parser.primary()
}

func (parser *Parser) primary() ast.Expr {
	if parser.matches(lexer.TRUE) {
		return ast.NewLiteral(true)
	} else if parser.matches(lexer.FALSE) {
		return ast.NewLiteral(false)
	} else if parser.matches(lexer.NIL) {
		return ast.NewLiteral(nil)
	} else if parser.matches(lexer.STRING) {
		return ast.NewLiteral(parser.prev().Literal())
	} else if parser.matches(lexer.NUMBER) {
		num, err := strconv.ParseFloat(*parser.prev().Literal(), 64)
		if err != nil {
			panic("Invalid float")
		}
		return ast.NewLiteral(num)
	}

	// check if it's a grouping expression
	if parser.matches(lexer.LEFT_PAREN) {
		expr := parser.expression()

		// if the next token is not a right paren, something is broken
		if !parser.matches(lexer.RIGHT_PAREN) {
			panic("Expected ')' missing")
		}

		return ast.NewGrouping(expr)
	}

	return nil
}
