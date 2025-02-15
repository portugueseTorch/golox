package ast

import "golox/src/lexer"

// marker interface for a node type
type Expr interface {
	marker()
}

// --- Binary expression: 1 + 2
type Binary struct {
	left     Expr
	operator lexer.Token
	right    Expr
}

func (t *Binary) marker() {}

// --- Grouping expression
type Grouping struct {
	expression Expr
}

func (t *Grouping) marker() {}

// --- Literal expression: "hello"
type Literal struct {
	value any
}

func (t *Literal) marker() {}

// --- Unary expression: !true | -1337
type Unary struct {
	operator   lexer.TokenType
	expression Expr
}

func (t *Unary) marker() {}
