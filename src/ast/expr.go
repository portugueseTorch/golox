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

func NewBinary(left Expr, op lexer.Token, right Expr) *Binary {
	return &Binary{
		Left:     left,
		Operator: op,
		Right:    right,
	}
}

func (t *Binary) marker() {}

// --- Grouping expression
type Grouping struct {
	Expression Expr
}

func NewGrouping(expr Expr) *Grouping {
	return &Grouping{
		Expression: expr,
	}
}

func (t *Grouping) marker() {}

// --- Literal expression: "hello"
type Literal struct {
	Value any
}

func NewLiteral(value any) *Literal {
	return &Literal{
		Value: value,
	}
}

func (t *Literal) marker() {}

// --- Unary expression: !true | -1337
type Unary struct {
	Operator   lexer.Token
	Expression Expr
}

func NewUnary(op lexer.Token, expr Expr) *Unary {
	return &Unary{
		Operator:   op,
		Expression: expr,
	}
}

func (t *Unary) marker() {}

// --- Variable node
type Variable struct {
	Name lexer.Token
}

func NewVariable(name lexer.Token) *Variable {
	return &Variable{
		Name: name,
	}
}

func (t *Variable) marker() {}

// --- Variable node
type Assignment struct {
	Name  lexer.Token
	Value Expr
}

func NewAssignment(name lexer.Token, val Expr) *Assignment {
	return &Assignment{
		Name:  name,
		Value: val,
	}
}

func (t *Assignment) marker() {}
