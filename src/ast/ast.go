package ast

import "golox/src/lexer"

// marker interface for a node type
type Expr interface {
	marker()
}

// --- Binary expression: 1 + 2
type Binary struct {
	Left     Expr
	Operator lexer.Token
	Right    Expr
}

func (t *Binary) marker() {}

// --- Grouping expression
type Grouping struct {
	Expression Expr
}

func (t *Grouping) marker() {}

// --- Literal expression: "hello"
type Literal struct {
	Value any
}

func (t *Literal) marker() {}

// --- Unary expression: !true | -1337
type Unary struct {
	Operator   lexer.Token
	Expression Expr
}

func (t *Unary) marker() {}
